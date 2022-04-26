package utils

import (
	"crypto/rand"
)

func GenerateOTP() (code string, err error) {
	nums := "0123456789"
	buff := make([]byte, 6)
	_, err = rand.Read(buff)
	if err != nil {
		return
	}

	for i := 0; i < 6; i++ {
		buff[i] = nums[int(buff[i])%10]
	}
	code = string(buff)

	return
}
