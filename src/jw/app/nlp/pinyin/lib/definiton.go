package lib

type RspJsonDefault struct {
	Py []string  `json:"py"`
	Hz []string  `json:"hz"`
}

type RspJsonHeteronym struct {
	Py [][]string  `json:"py"`
	Hz []string    `json:"hz"`
}