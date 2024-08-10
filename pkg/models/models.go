package models 


import (
	"errors"
	"time"
)

var ErrNoRecord = errors.New("models: no matching record found")


type Post struct {
	Id int
	Title string
	Content string
	Created time.Time
	Expires time.Time
}