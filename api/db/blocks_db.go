package db

import (
	"gorm.io/gorm"
)

type Block struct {
	gorm.Model
	Number     int64
	Hash       string
	Time       string
	ParentHash string
}
