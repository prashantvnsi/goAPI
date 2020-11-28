package models

import (
	//"github.com/jinzhu/gorm"
)

type Members struct{
	ID uint `json: "id" gorm: primary_key`
	name string

}