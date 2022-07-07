package incident_test

import (
	"log"
	"main/internal/additional/incident"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestIncident_Make(t *testing.T) {
	if err := godotenv.Load("./testing/testing.env"); err != nil {
		log.Print("No .env file found")
	}

	incidentReport := incident.New(os.Getenv("INCIDENT_URL"))
	incidentData := incidentReport.Make()
	assert.NotNil(t, incidentData)
}
