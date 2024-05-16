package models

import "time"

type Person struct {
	ID       string    `uri:"id" `
	Name     string    `uri:"name" form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"5"`
}
