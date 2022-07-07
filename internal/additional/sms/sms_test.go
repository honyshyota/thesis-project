package sms_test

import (
	"log"
	"main/internal/additional"
	"main/internal/additional/sms"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSMS_Make(t *testing.T) {
	if err := godotenv.Load("./testing/testing.env"); err != nil {
		log.Print("No .env file found")
	}

	smsReport := sms.New(additional.GetFilePathByFileNameTest(os.Getenv("SMS_FILE_NAME")), strings.Split(os.Getenv("SMS_MMS_PROV"), ", "))
	smsData1, smsData2 := smsReport.Make()
	assert.NotNil(t, smsData1)
	assert.NotNil(t, smsData2)
}
