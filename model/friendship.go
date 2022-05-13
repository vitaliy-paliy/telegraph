package model

import (
	"errors"
	"io"
	"strconv"
	"time"
)

type Friendship struct {
	ID        string           `json:"id"`
	Sender    string           `json:"sender"`
	Recipient string           `json:"recipient"`
	Status    FriendshipStatus `json:"status"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt time.Time        `json:"deleted_at"`
}

func (f *Friendship) Invert() *Friendship {
	return &Friendship{Sender: f.Recipient, Recipient: f.Sender}
}

type NewFriendshipInput struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
}

func (f *NewFriendshipInput) Convert() *Friendship {
	return &Friendship{
		Sender:    f.Sender,
		Recipient: f.Recipient,
	}
}

// FriendshipStatus type
type FriendshipStatus int

var FriendshipStatusEnum = struct {
	PENDING  FriendshipStatus
	ACCEPTED FriendshipStatus
	BLOCKED  FriendshipStatus
}{
	PENDING:  0,
	ACCEPTED: 1,
	BLOCKED:  2,
}

func (f *FriendshipStatus) UnmarshalGQL(v interface{}) error {
	status, ok := v.(string)
	if !ok {
		return errors.New("FriendshipStatus must be a string.")
	}

	i, err := strconv.ParseInt(status, 10, 8)
	if err != nil {
		return err
	}
	*f = FriendshipStatus(i)

	return nil
}

func (f FriendshipStatus) MarshalGQL(w io.Writer) {
	if f == FriendshipStatusEnum.PENDING {
		w.Write([]byte(`"0"`))
	} else if f == FriendshipStatusEnum.ACCEPTED {
		w.Write([]byte(`"1"`))
	} else if f == FriendshipStatusEnum.BLOCKED {
		w.Write([]byte(`"2"`))
	}
}
