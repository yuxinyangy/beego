package models

type Personto struct {
	//Name string `json:"name"`
	Nick string `json:"nick"`
	Birthday string `json:"birthday"`
	Address string `json:"address"`
	Password string `json:"password"`
}
/*
type Personto struct {
	Name string `form:"name"`
	Birthday string  `form:"birthday"`
	Address string  `form:"address"`
	Nick string `form:"nick"`
}*/
