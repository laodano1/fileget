package main

import (
	"fileget/src/jw/app/colly/worldHL/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"os"
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
			lg.Debugf("json(%v) exists. do nothing!", fileName + ".json")
			return
		}

		e.DOM.Find(".alternate").ChildrenFiltered("div").Each(func(i int, s *goquery.Selection) {
			//lg.Debugf("parse list")
			switch i {
			case 0:
				flagUrl, _ := s.ChildrenFiltered("img").Attr("src")
				hd.TheFlag = append(hd.TheFlag, urlPrefix + flagUrl)
				hd.Country = append(hd.Country, strings.Trim(s.ChildrenFiltered("a").Text(), "") )
				return
			case 1:
				if len(s.ChildrenFiltered("img").Nodes) > 0 {  // still flag
					flagUrl, _ := s.ChildrenFiltered("img").Attr("src")
					hd.TheFlag = append(hd.TheFlag, urlPrefix + flagUrl)
					hd.Country = append(hd.Country, strings.Trim(s.ChildrenFiltered("a").Text(), ""))
					return
				}
				// not flag
				fallthrough
			default:
				if len(hd.TheFlag) >= 2 { // multiple countries
					switch i {
					case 2:
						hd.Location = strings.ReplaceAll(s.Text(), "\n", "")
						hd.Location = strings.ReplaceAll(hd.Location, "\t", "")
					case 3:
						hd.Coordinate = strings.ReplaceAll(s.Text(), "\n", "")
						hd.Coordinate = strings.ReplaceAll(hd.Coordinate, "\t", "")
					case 4:
						hd.DateOfInscription = strings.ReplaceAll(s.Text(), "\n", "")
						hd.DateOfInscription = strings.ReplaceAll(hd.DateOfInscription, "\t", "")
					default:
					}
				} else {  // single country
					switch i {
					case 1:
						hd.Location = strings.Trim(s.Text(), " ")
						hd.Location = strings.ReplaceAll(hd.Location, "\n", "")
						hd.Location = strings.ReplaceAll(hd.Location, "\t", "")
					case 2:
						hd.Coordinate = strings.Trim(s.Text(), " ")
						hd.Coordinate = strings.ReplaceAll(hd.Coordinate, "\n", "")
						hd.Coordinate = strings.ReplaceAll(hd.Coordinate, "\t", "")
					case 3:
						hd.DateOfInscription = strings.Trim(s.Text(), " ")
						hd.DateOfInscription = strings.ReplaceAll(hd.DateOfInscription, "\n", "")
						hd.DateOfInscription = strings.ReplaceAll(hd.DateOfInscription, "\t", "")
					default:
					}
				}
			}
		})

		zhDesc, err := e.DOM.Find("#contentdes_zh").Html()
		if err != nil {
			lg.Errorf("get zh html failed: %v", err)
		} else {
			lg.Debugf("worker(%v) use zh description", id)
			hd.Description = zhDesc
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
			}
		}
		//lg.Debugf("url: %v, detail; %v", c.String(), allHeritageDetailList)
		cnt++
		lg.Debugf("worker(%v) not json exists. write it(%v). count: %v", outputFile, outputFile, cnt)
		utils.Write2JsonFile(hd, outputFile)

		lg.Debugf("worker(%v) parse %v spending %v seconds.", id, heritageItem.Url, time.Now().Sub(startTime).Seconds())
	})

	c.Visit(heritageItem.Url)
}
