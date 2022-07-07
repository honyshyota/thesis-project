package billing_test

import (
	"log"
	"main/internal/additional"
	"main/internal/additional/billing"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestBilinnig_Make(t *testing.T) {
	if err := godotenv.Load("./testing/testing.env"); err != nil {
		log.Print("No .env file found")
	}

	billingReport := billing.New(additional.GetFilePathByFileNameTest(os.Getenv("BILLING_FILE_NAME")))
	billingData := billingReport.Make()
	assert.NotNil(t, billingData)
}
