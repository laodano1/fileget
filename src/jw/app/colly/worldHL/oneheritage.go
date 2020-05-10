package main

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
	"time"
)

var (
	allHeritageDetailList []HeritageDetail
)

func GetHeritageInfo(id int, heritageItem msg) {
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

		e.DOM.Find(".alternate").ChildrenFiltered("div").Each(func(i int, s *goquery.Selection) {
			lg.Debugf("parse list")
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
			lg.Debugf("use zh description")
			hd.Description = zhDesc
		}

		e.DOM.Find("div.icaption.bordered").Find("img").Each(func(i int, s *goquery.Selection) {
			lg.Debugf("parse cover image")
			hd.CoverImageHref, _ = s.Attr("data-src")
			hd.CoverImageHref    = urlPrefix + strings.Trim(hd.CoverImageHref, " ")
		})

		//lg.Debugf("url: %v, detail; %v", c.String(), allHeritageDetailList)
		fileName := strings.ReplaceAll(hd.Name, " ", "_")
		fileName  = strings.ReplaceAll(fileName, "/", "-")
		if hd.Description == "" {
			// if no China contents, then use English instead
			var enDesc string
			enDesc, err = e.DOM.Find("#contentdes_en").Html()
			if err != nil {
				lg.Errorf("get english html failed: %v", err)
			} else {
				lg.Debugf("use english description")
				hd.Description = enDesc
			}
		}
		//utils.Write2JsonFile(hd, fmt.Sprintf("%v%ctmp%c%v",  exeDirPath, os.PathSeparator, os.PathSeparator, fileName + ".json") )
		//lg.Debugf("detail; %v", allHeritageDetailList)
		allHeritageDetailList = append(allHeritageDetailList, *hd)

		lg.Debugf("parse %v spending %v seconds.", heritageItem.Url, time.Now().Sub(startTime).Seconds())
	})

	c.Visit(heritageItem.Url)
}
