package entity

import "time"

type Message struct {
	CreateTime time.Time `db:"create_ts"`
	Content    string
}
