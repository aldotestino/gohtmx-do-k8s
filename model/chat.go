package model

import (
	"time"
)

type Message struct {
	User     string `json:"user"`
	Content  string `json:"content"`
	Datetime string `json:"datetime"`
}

func NewMessage(user string, content string) *Message {
	return &Message{
		User:     user,
		Content:  content,
		Datetime: time.Now().Format(time.RFC822),
	}
}

type ChatStore interface {
	GetAll() []*Message
	Create(user *Message) (*Message, error)
}

type LocalChatStore struct {
	messages []*Message
}

func NewLocalChatStore() *LocalChatStore {
	return &LocalChatStore{
		messages: make([]*Message, 0),
	}
}

func (s *LocalChatStore) GetAll() []*Message {
	return s.messages
}

func (s *LocalChatStore) Create(message *Message) (*Message, error) {
	s.messages = append(s.messages, message)
	return message, nil
}
