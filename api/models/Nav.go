package models

type Nav struct {
	Id       int    `json:"Id"`
	Name     string `json:"Name"`
	ParentId int    `json:"ParentId"`
	SortId   int    `json:"SortId"`
	NavIcon  string `json:"NavIcon"`
	NavPath  string `json:"NavPath"`
	NavAcc   string `json:"NavAcc"`
}

type Navs []Nav

var menu = Navs{
	Nav{
		1,
		"Home",
		0,
		1,
		"<DashboardIcon/>",
		"/",
		"public",
	},
	Nav{
		2,
		"Table",
		0,
		2,
		"<TableIcon/>",
		"/#",
		"private",
	},
}
