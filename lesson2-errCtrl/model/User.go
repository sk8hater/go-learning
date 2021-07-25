package model

import (
	"time"
)

type User struct {
	Id        uint32    `sql:"primaryKey;column:id"`
	Name      string    `sql:"column:name"`
	Age       int       `sql:"column:age"`
	CreatedBy string    `sql:"column:create_by"`
	CreatedAt time.Time `sql:"column:create_At"`
	IsDel     uint8     `sql:"column:is_del"`
	DeletedAt time.Time `sql:"column:deleted_At"`
}
