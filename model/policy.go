package model

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Friendship Policy
var FriendshipPolicyEnum = struct {
	FRIEND    string
	SENDER    string
	RECIPIENT string
}{
	FRIEND:    "friend",
	SENDER:    "sender",
	RECIPIENT: "recipient",
}

// Action type
type Action string

const (
	ActionCreate  Action = "Create"
	ActionGetOne  Action = "GetOne"
	ActionGetMany Action = "GetMany"
	ActionUpdate  Action = "Update"
	ActionDelete  Action = "Delete"
)

func (a Action) String() string {
	return string(a)
}

func (a Action) IsValid() bool {
	switch a {
	case ActionCreate, ActionGetOne, ActionGetMany, ActionUpdate, ActionDelete:
		return true
	}

	return false
}

func (a *Action) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return errors.New("FriendshipStatus must be a string.")
	}

	*a = Action(str)
	if !a.IsValid() {
		return fmt.Errorf("%s is not a valid Action", str)
	}

	return nil
}

func (a Action) MarshalGQL(w io.Writer) {
	w.Write([]byte(strconv.Quote(a.String())))
}
