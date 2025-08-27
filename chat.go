package main

import "time"

type User struct {
	name string
}

type Message struct {
	chat *Chat
	from *User
	sentAt time.Time
	text []byte // should be limited to maybe 4k bytes, enough for 1000 characters
}

func (msg *Message) Text() string {
	return string(msg.text)
}

type Chat struct {
	msgs []*Message
}