package models

import (
	"github.com/jinzhu/gorm"
	"members"
)

type Group struct {
	ID uint `json: "id" gorm: primary_key`
	name string
	member []Members `gorm: "Foreign_Key: "`
}