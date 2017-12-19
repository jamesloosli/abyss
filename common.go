package main

import (
	crypto/md5
	time
)

const BlockSize = 64
const Size = 32

// Timestamps should be the time the message posts.
type Timestamp struct {
	time.Time
}

// Common format for chats, e.g. Slack - JobMachine - Abyss
type Chat struct {
	Client string
	Server string
	Room string
}

func (c *Chat) Hash() string {
	data := []byte(c.Client + c.Server + c.Room)
	return md5.Sum(data)

// Common message object
type Message struct {
	Stamp Timestamp
	Source ChatSource
	Author string
	Message string
}

// Interface for taking in messages from a chatsource
type Reciever interface {
	ingest() string
}

// Interface for writing messages to a chat
type Emitter interface {
	write() string
}
