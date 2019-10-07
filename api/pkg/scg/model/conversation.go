package model

import (
	"time"
)

type Sender string

const (
	SystemSender = "system"
	UserSender   = "user"
)

type Message struct {
	Sender     Sender
	ReplyToken string
	Text       string
	CreateDate time.Time
}

func NewMessage(sender Sender, text string, time time.Time, replyToken string) *Message {
	return &Message{
		Sender:     sender,
		Text:       text,
		CreateDate: time,
		ReplyToken: replyToken,
	}
}

type Conversation struct {
	id       Id
	user     *User
	messages []*Message
}

func NewConversation(user *User) *Conversation {
	return &Conversation{
		id:       user.GetId(),
		user:     user,
		messages: make([]*Message, 0),
	}
}

func (c *Conversation) GetId() Id {
	return c.id
}

func (c *Conversation) GetUser() *User {
	return c.user
}

func (c *Conversation) GetMessages() []*Message {
	return c.messages
}

func (c *Conversation) AddMessage(sender Sender, text string, time time.Time, replyToken string) {
	c.messages = append(c.messages, NewMessage(sender, text, time, replyToken))
}

func (c *Conversation) GetLastestMessages() *Message {
	return c.messages[len(c.messages)-1]
}
