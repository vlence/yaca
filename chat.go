package main

import "time"

type User struct {
	name string
}

type Message struct {
	from *User
	to *User
	sentAt time.Time
	text string
}

func (msg *Message) Text() string {
	return msg.text
}