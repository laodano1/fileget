package main



func LoadData() {
	subIList := make([]sbSubItem, 0)
	subIList = append(subIList, sbSubItem{Name: "mp3", Href: "/mp3"})
	subIList = append(subIList, sbSubItem{Name: "mp4", Href: "/mp4"})
	subIList = append(subIList, sbSubItem{Name: "mkv", Href: "/mkv"})

	pgItems1 := make([]pageItem, 0)
	pgItems1 = append(pgItems1, pageItem{})
	pgcnt1 := pageContent{PageObjs: pgItems1}

	pgItems2 := make([]pageItem, 0)
	pgItems2 = append(pgItems2, pageItem{Name: "tesla", Href: "/video/tesla.mp4"})
	pgcnt2 := pageContent{PageObjs: pgItems2}

	pgItems3 := make([]pageItem, 0)
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgItems3 = append(pgItems3, pageItem{Name: "ye-wen-4", Href: "/video/ye-wen-4.mkv"})
	pgcnt3 := pageContent{PageObjs: pgItems3}

	pgcList := make([]pageContent, 0)
	pgcList = append(pgcList, pgcnt1)
	pgcList = append(pgcList, pgcnt2)
	pgcList = append(pgcList, pgcnt3)

	sbmi := sidebarMainItem{
		Name:     "Media",
		SubItems: subIList,
		PageCnt:  pgcList,
	}

	sbmiList := make([]sidebarMainItem, 0)
	sbmiList = append(sbmiList, sbmi)
	sbObj := &sidebar{
		List: sbmiList,
	}

	sidebarData = *sbObj
}
