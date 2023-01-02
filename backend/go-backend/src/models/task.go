package models

type Task struct {
	ID 			int64 `json:"id" bun:"id,pk,autoincrement"`
	Label 		string `json:"label,omitempty" bun:"label"`
	Done		bool `json:"done" bun:"done"`
	Date		string `json:"date,omitempty" bun:"date"`
}
