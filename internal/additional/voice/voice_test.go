package voice_test

import (
	"log"
	"main/internal/additional"
	"main/internal/additional/voice"
	"os"
	"strings"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestVoice_Make(t *testing.T) {
	if err := godotenv.Load("./testing/testing.env"); err != nil {
		log.Print("No .env file found")
	}

	voiceReport := voice.New(additional.GetFilePathByFileNameTest(os.Getenv("VOICE_FILE_NAME")), strings.Split(os.Getenv("VOICE_PROV"), ", "))
	voiceData := voiceReport.Make()
	assert.NotNil(t, voiceData)
}
