package store

import (
	"errors"
	"fmt"
	"os"
	"regexp"

	"telegraph/model"
	"telegraph/utils"

	"github.com/joho/godotenv"
	"github.com/satori/go.uuid"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

func (s *AuthStore) SendOTP(phoneNumber string) (code string, err error) {
	// Read .env file
	err = godotenv.Load()
	if err != nil {
		return
	}

	// Validate phone number. It should contain only 0-9 and be of length 11.
	rxpPhoneNumber := regexp.MustCompile("^[0-9]{11}$")
	if !rxpPhoneNumber.MatchString(phoneNumber) {
		err = errors.New("Invalid phone number format.")
		return
	}

	twilioPhoneNumber := os.Getenv("TWILIO_PHONE_NUMBER")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")

	twilioClient := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	code, err = utils.GenerateOTP()
	if err != nil {
		return
	}

	// Send OTP.
	params := &openapi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(twilioPhoneNumber)
	params.SetBody(fmt.Sprintf("Your OTP code: %s", code))
	_, err = twilioClient.ApiV2010.CreateMessage(params)

	return
}

func (s *AuthStore) SignUp(user *model.User) (*model.User, error) {
	// Validate username and phone number formats.
	if err := utils.ValidateNewUser(user); err != nil {
		return nil, err
	}

	user.ID = uuid.NewV4().String()
	user.ActivityStatus = model.UserActivityStatusEnum.OFFLINE

	err := s.db.Create(user).Error

	return user, err
}

func (s *AuthStore) SignIn(phoneNumber string) (user *model.User, err error) {
	user = &model.User{}

	if err := s.db.Where("phone_number = ?", phoneNumber).First(user).Error; err != nil {
		return nil, errors.New("User was not found.")
	}

	return
}
