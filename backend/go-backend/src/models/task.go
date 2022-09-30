package models

type Task struct {
	ID 			string `json:"id"`
	Label 		string `json:"label"`
	Done		bool `json:"done"`
	Date		string `json:"date"`
}
