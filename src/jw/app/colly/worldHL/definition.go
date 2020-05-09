package main


type OneHeritage struct {
	Name string `json:"name"`
	Href string `json:"href"`
	BelongTo string `json:"belong_to"`
}

type HeritageItem struct {
	TypeOrder []string  `json:"type_order"`  // type order to manage types map key sequence
	Types     map[string][]OneHeritage `json:"types"`
	BelongTo  string `json:"belong_to"`
}

type CountryItem struct {
	Name string `json:"name"`
	Href string `json:"href"`
	Type string `json:"type,omitempty"`
	HeritageList []HeritageItem `json:"heritage_list"`
}

type WorldHeritageList struct {
	CountryList []CountryItem `json:"country_list"`
}


type msg struct {
	Status bool
	Url    string
}


