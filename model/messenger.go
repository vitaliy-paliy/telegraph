package model

import (
	"errors"
	"io"
	"strconv"
	"time"
)

type Message struct {
	ID             string        `json:"id"`
	ConversationID string        `json:"conversation_id"`
	Text           string        `json:"text"`
	Sender         string        `json:"sender"`
	Recipient      string        `json:"recipient"`
	Status         MessageStatus `json:"status"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
	DeletedAt      time.Time     `json:"deleted_at"`
}

type NewMessageInput struct {
	ConversationID string `json:"conversation_id"`
	Text           string `json:"text"`
	Sender         string `json:"sender"`
	Recipient      string `json:"recipient"`
}

func (i NewMessageInput) Convert() *Message {
	return &Message{
		ConversationID: i.ConversationID,
		Text:           i.Text,
		Sender:         i.Sender,
		Recipient:      i.Recipient,
	}
}

// Message status type
type MessageStatus int

var MessageStatusEnum = struct {
	SENT MessageStatus
	READ MessageStatus
}{
	SENT: 0,
	READ: 1,
}

func (s *MessageStatus) UnmarshalGQL(v interface{}) error {
	status, ok := v.(string)
	if !ok {
		return errors.New("An error occured.")
	}

	i, err := strconv.ParseInt(status, 10, 8)
	if err != nil {
		return err
	}
	*s = MessageStatus(i)

	return nil
}

func (s MessageStatus) MarshalGQL(w io.Writer) {
	switch s {
	case MessageStatusEnum.SENT:
		w.Write([]byte("0"))
	case MessageStatusEnum.READ:
		w.Write([]byte("1"))
	default:
		panic("Error while marshal scheme")
	}
}
