package models

type Vacancy struct {
	Id       	uint   `json:"id"`
	Title     	string `json:"title"`
	Description string `json:"description"`
	Source 		string `json:"source"`
}
