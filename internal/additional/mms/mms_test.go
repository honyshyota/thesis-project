package mms_test

import (
	"log"
	"main/internal/additional/mms"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMms_Make(t *testing.T) {
	if err := godotenv.Load("./testing/testing.env"); err != nil {
		log.Print("No .env file found")
	}

	mmsReport := mms.New(os.Getenv("MMS_URL"), strings.Split(os.Getenv("SMS_MMS_PROV"), ", "))
	mmsData1, mmsData2 := mmsReport.Make()
	assert.NotNil(t, mmsData1)
	assert.NotNil(t, mmsData2)
}
