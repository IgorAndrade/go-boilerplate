package model

import "time"

type TodoList struct {
	ID         string    `bson:"_id" json:"id"`
	Text       string    `bson:"text" json:"text"`
	Created_at time.Time `bson:"created_at" json:"created_at"`
}

//SetID(string)
//	SetCreated(time.Time)

func (t *TodoList) SetID(ID string) {
	t.ID = ID
}

func (t *TodoList) SetCreated(created time.Time) {
	t.Created_at = created
}
