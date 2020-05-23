package model

import "time"

type TodoList struct {
	ID        string    `bson:"_id" json:"id"`
	Text      string    `bson:"text" json:"text"`
	CreatedAt time.Time `bson:"createdAt" json:"createdAt"`
}

//SetID(string)
//	SetCreated(time.Time)

func (t *TodoList) SetID(ID string) {
	t.ID = ID
}

func (t *TodoList) SetCreated(created time.Time) {
	t.CreatedAt = created
}
