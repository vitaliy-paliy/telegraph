package model

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

type User struct {
	ID              string             `json:"id" gorm:"primaryKey"`
	Name            string             `json:"name"`
	Username        string             `json:"username" gorm:"unique"`
	PhoneNumber     string             `json:"phone_number" gorm:"unique"`
	ProfileImageURL string             `json:"profile_image_url"`
	ActivityStatus  UserActivityStatus `json:"activity_status"`
	CreatedAt       time.Time          `json:"created_at"`
	UpdatedAt       time.Time          `json:"updated_at"`
	DeletedAt       *time.Time         `json:"deleted_at"`
}

type NewUserInput struct {
	Name            string `json:"name"`
	Username        string `json:"username"`
	PhoneNumber     string `json:"phone_number"`
	ProfileImageURL string `json:"profile_image_url"`
}

func (u *NewUserInput) Convert() *User {
	return &User{
		ID:              uuid.NewV4().String(),
		ActivityStatus:  UserActivityStatusEnum.OFFLINE,
		Name:            u.Name,
		Username:        u.Username,
		PhoneNumber:     u.PhoneNumber,
		ProfileImageURL: u.ProfileImageURL,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
}

// UserActivityStatus type.
type UserActivityStatus int

var UserActivityStatusEnum = struct {
	OFFLINE UserActivityStatus
	ONLINE  UserActivityStatus
}{
	OFFLINE: 0,
	ONLINE:  1,
}

func (u *UserActivityStatus) UnmarshalGQL(v interface{}) error {
	status, ok := v.(string)
	if !ok {
		return errors.New("UserActivityStatus must be a string.")
	}

	i, err := strconv.ParseInt(status, 10, 8)
	if err != nil {
		return err
	}
	*u = UserActivityStatus(i)

	return nil
}

func (u UserActivityStatus) MarshalGQL(w io.Writer) {
	if u == UserActivityStatusEnum.OFFLINE {
		w.Write([]byte(`"0"`))
	} else if u == UserActivityStatusEnum.ONLINE {
		w.Write([]byte(`"1"`))
	}
}
