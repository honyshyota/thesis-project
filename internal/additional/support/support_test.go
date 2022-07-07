package support_test

import (
	"log"
	"main/internal/additional/support"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestSupport_Make(t *testing.T) {
	if err := godotenv.Load("./testing/testing.env"); err != nil {
		log.Print("No .env file found")
	}

	supportReport := support.New(os.Getenv("SUPPORT_URL"))
	supportData := supportReport.Make()
	assert.NotNil(t, supportData)
}
