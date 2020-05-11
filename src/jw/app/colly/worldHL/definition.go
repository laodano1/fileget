package main


type OneHeritage struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Country string `json:"country"`
}

type HeritageItem struct {
	TypeOrder []string  `json:"type_order"`  // type order to manage types map key sequence
	Types     map[string][]OneHeritage `json:"types"`
	Country   string `json:"country"`
}

type CountryItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Type string `json:"type,omitempty"`
	HeritageList []HeritageItem `json:"heritage_list"`
}

type WorldHeritageList struct {
	countryOrder []string               `json:"country_order"`
	CountryList  map[string]*CountryItem `json:"country_list"`
}


type parseMsg struct {
	Status bool
	Url    string
	Name   string
}

type HeritageDetail struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TheFlag     []string `json:"the_flag"`
	Country     []string  `json:"country"`
	Location    string  `json:"location"`
	Coordinate  string  `json:"coordinate"`
	DateOfInscription string `json:"date_of_inscription"`
	CoverImageHref    string  `json:"cover_image"`

}

