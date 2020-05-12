package main

import (
	"fileget/util"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	allHeritageDetailList []HeritageDetail
	cnt int
)

func GetHeritageInfo(id int, heritageItem parseMsg) {
	startTime := time.Now()
	c := colly.NewCollector(
		colly.CacheDir("./whl"),
	)
	c.UserAgent = UserAgent

	c.OnRequest(func(req *colly.Request) {
		//lg.Debugf("(%v) on request: %v", id, req.URL)
	})

	c.OnHTML("div.content", func(e *colly.HTMLElement) {
		hd := new(HeritageDetail)
		hd.Name = heritageItem.Name

		fileName := strings.ReplaceAll(hd.Name, " ", "_")
		fileName  = strings.ReplaceAll(fileName, "/", "-")
		fileName  = strings.ReplaceAll(fileName, ":", "--")

		outputFile := fmt.Sprintf("%v%ctmp%c%v",  exeDirPath, os.PathSeparator, os.PathSeparator, fileName + ".json")
		if _, ok := allJson[fileName + ".json"]; ok {
			lg.Debugf("    json(%v) exists. do nothing!", fileName + ".json")
			return
		}

		brif, err := e.DOM.Find(".alternate").Html()
		if err != nil {
			lg.Errorf("get brief html failed: %v", err)
		} else {
			lg.Debugf("worker(%v) get brief", id)
			brif = strings.ReplaceAll(brif, "\"/", "\"" + urlPrefix + "/")
			hd.Brief = brif
		}

		e.DOM.Find(".alternate").ChildrenFiltered("div").Each(func(i int, s *goquery.Selection) {
			//lg.Debugf("parse list")
			if !strings.Contains(s.Text(), ",") && !strings.Contains(s.Text(), ":") && !strings.Contains(s.Text(), "(") && !strings.Contains(s.Text(), ")") {
				isCood := true
				desc := strings.TrimPrefix(s.Text(), " ")
				desc  = strings.TrimPrefix(desc, " ")
				desc  = strings.TrimSuffix(desc, " ")
				desc  = strings.ReplaceAll(s.Text(), "\n", "")
				desc  = strings.ReplaceAll(desc, "\t", "")
				lg.Debugf("parse desc: '%v'", desc)
				if strings.Contains(desc, " ") && desc != " " {
					a := int(desc[0])
					b := int(desc[1])

					if a >= 48 && a <= 57 {  // first char is number, next
						isCood = false
						return
					} else {
						if b < 48 || b > 57 { // second char is not number, next
							isCood = false
							return
						}
					}

					if isCood {
						if strings.Contains(desc, " ") {
							hd.Coordinate = desc
							//hd.Coordinate = strings.ReplaceAll(hd.Coordinate, "\t", "")
							lg.Debugf("parse coodinate: '%v'", desc)
							hd.CoordinateDigit[0], hd.CoordinateDigit[1] = coodinateConvert(hd.Coordinate)
						}
					}
				}

			}

			//switch i {
			//case 0:
			//	flagUrl, _ := s.ChildrenFiltered("img").Attr("src")
			//	hd.TheFlag = append(hd.TheFlag, urlPrefix + flagUrl)
			//	hd.Country = append(hd.Country, strings.Trim(s.ChildrenFiltered("a").Text(), "") )
			//	return
			//case 1:
			//	if len(s.ChildrenFiltered("img").Nodes) > 0 {  // still flag
			//		flagUrl, _ := s.ChildrenFiltered("img").Attr("src")
			//		hd.TheFlag = append(hd.TheFlag, urlPrefix + flagUrl)
			//		hd.Country = append(hd.Country, strings.Trim(s.ChildrenFiltered("a").Text(), ""))
			//		return
			//	}
			//	// not flag
			//	fallthrough
			//default:
			//	if len(hd.TheFlag) >= 2 { // multiple countries
			//		switch i {
			//		case 2:
			//			hd.Location = strings.ReplaceAll(s.Text(), "\n", "")
			//			hd.Location = strings.ReplaceAll(hd.Location, "\t", "")
			//		case 3:
			//			hd.Coordinate = strings.ReplaceAll(s.Text(), "\n", "")
			//			hd.Coordinate = strings.ReplaceAll(hd.Coordinate, "\t", "")
			//		case 4:
			//			hd.DateOfInscription = strings.ReplaceAll(s.Text(), "\n", "")
			//			hd.DateOfInscription = strings.ReplaceAll(hd.DateOfInscription, "\t", "")
			//		default:
			//		}
			//	} else {  // single country
			//		switch i {
			//		case 1:
			//			hd.Location = strings.Trim(s.Text(), " ")
			//			hd.Location = strings.ReplaceAll(hd.Location, "\n", "")
			//			hd.Location = strings.ReplaceAll(hd.Location, "\t", "")
			//		case 2:
			//			hd.Coordinate = strings.Trim(s.Text(), " ")
			//			hd.Coordinate = strings.ReplaceAll(hd.Coordinate, "\n", "")
			//			hd.Coordinate = strings.ReplaceAll(hd.Coordinate, "\t", "")
			//		case 3:
			//			hd.DateOfInscription = strings.Trim(s.Text(), " ")
			//			hd.DateOfInscription = strings.ReplaceAll(hd.DateOfInscription, "\n", "")
			//			hd.DateOfInscription = strings.ReplaceAll(hd.DateOfInscription, "\t", "")
			//		default:
			//		}
			//	}
			//}
		})

		zhDesc, err := e.DOM.Find("#contentdes_zh").Html()
		if err != nil {
			lg.Errorf("get zh html failed: %v", err)
		} else {
			lg.Debugf("worker(%v) use zh description", id)

			hd.Description = zhDesc
			hd.Description = strings.ReplaceAll(hd.Description, "\"/en", "\"" + urlPrefix + "/en")
			hd.Description = strings.ReplaceAll(hd.Description, "<h6>", "<a href=\"" + heritageItem.Url +"\"><h6><strong>")
			hd.Description = strings.ReplaceAll(hd.Description, "</h6>", "</h6></strong></a>")
		}

		e.DOM.Find("div.icaption.bordered").Find("img").Each(func(i int, s *goquery.Selection) {
			lg.Debugf("worker(%v) parse cover image", id)
			hd.CoverImageHref, _ = s.Attr("data-src")
			hd.CoverImageHref    = urlPrefix + strings.Trim(hd.CoverImageHref, " ")
		})

		if hd.Description == "" {
			// if no China contents, then use English instead
			var enDesc string
			var err    error
			enDesc, err = e.DOM.Find("#contentdes_en").Html()
			if err != nil {
				lg.Errorf("get english html failed: %v", err)
			} else {
				lg.Debugf("worker(%v) use english description", id)
				hd.Description = enDesc
				hd.Description = strings.ReplaceAll(hd.Description, "\"/en", "\"" + urlPrefix + "/en")
				hd.Description = strings.ReplaceAll(hd.Description, "<h6>", "<a href=\"" + heritageItem.Url +"\"><h6><strong>")
				hd.Description = strings.ReplaceAll(hd.Description, "</h6>", "</h6></strong></a>")
			}
		}
		//lg.Debugf("url: %v, detail; %v", c.String(), allHeritageDetailList)
		cnt++
		lg.Debugf("worker(%v) not json exists. write it(%v). count: %v", outputFile, outputFile, cnt)
		util.Write2JsonFile(hd, outputFile)

		lg.Debugf("worker(%v) parse %v spending %v seconds.", id, heritageItem.Url, time.Now().Sub(startTime).Seconds())
	})

	c.Visit(heritageItem.Url)
}

func str2Int(str string) (i int, err error) {
	i, err = strconv.Atoi(str)
	return
}

func str2Float(str string) (f float64, err error) {
	f, err = strconv.ParseFloat(str, 32)
	return
}

func coodinateConvert(co string) (latitude, longitude float64) {
	strs := strings.Split(co, " ")


	la_de, _  := str2Int(strs[0][1:])
	la_min, _ := str2Float(strs[1])
	laMin := la_min/60
	la_sec, _ := str2Float(strs[2])
	laSec := la_sec/3600

	latitude = float64(la_de) + laMin + laSec
	if strs[0][:1] == "S" {
		latitude *= -1
	}

	lo_de, _  := str2Int(strs[3][1:])
	lo_min, _ := str2Float(strs[4])
	loMin := lo_min/60
	lo_sec, _ := str2Float(strs[5])
	loSec := lo_sec/3600

	longitude = float64(lo_de) + loMin + loSec
	if strs[3][:1] == "W" {
		longitude *= -1
	}
	//
	return
}