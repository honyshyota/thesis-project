package email_test

import (
	"log"
	"main/internal/additional"
	"main/internal/additional/email"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestEmail_Make(t *testing.T) {
	if err := godotenv.Load("./testing/testing.env"); err != nil {
		log.Print("No .env file found")
	}

	emailReport := email.New(additional.GetFilePathByFileNameTest(os.Getenv("EMAIL_FILE_NAME")), strings.Split(os.Getenv("EMAIL_PROV"), ", "))
	emailData := emailReport.Make()
	assert.NotNil(t, emailData)
}
