package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Post struct {
	gorm.Model
	Title   string
	Content string
}

type Visitor struct {
	LastSeen time.Time
	Count    int
}
