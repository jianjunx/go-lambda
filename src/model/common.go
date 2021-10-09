package model

type PageListStruct struct {
	Page  int         `json:"page"`
	Size  int         `json:"size"`
	Total int         `json:"total"`
	List  interface{} `json:"list"`
}
