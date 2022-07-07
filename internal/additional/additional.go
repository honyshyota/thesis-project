package additional

import (
	"errors"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/pariz/gountries"
)

// This functions build strings path for sms, email, voice, billing methods
func GetFilePathByFileName(filename string) string {
	return os.Getenv("PATH_TO_READ") + filename
}

func GetFilePathByFileNameTest(filename string) string {
	return os.Getenv("PATH_TEST") + filename
}

// This functions for check or conversion values for sms, mms, voice and other
func VoiceParametersCheckAndConv(connStab, TTFB, voicePurity, medianOfCallsTime string) (float32, int, int, int, error) {
	var connFloat float64
	var TTFBtoInt, voicePurToInt, medianToInt int
	var err error
	switch connStab {
	case "":
		break
	default:
		connFloat, err = strconv.ParseFloat(connStab, 32)
		if err != nil {
			log.Printf("Incorrect data in ConnectionStability: %s, %s\n", connStab, err)
			return 0, 0, 0, 0, err
		}
	}

	switch TTFB {
	case "":
		break
	default:
		TTFBtoInt, err = strconv.Atoi(TTFB)
		if err != nil {
			log.Printf("Incorrect data in TTFB: %s, %s\n", TTFB, err)
			return 0, 0, 0, 0, err
		}
	}

	switch voicePurity {
	case "":
		break
	default:
		voicePurToInt, err = strconv.Atoi(voicePurity)
		if err != nil {
			log.Printf("Incorrect data in VoicePurity: %s, %s\n", voicePurity, err)
			return 0, 0, 0, 0, err
		}
	}

	switch medianOfCallsTime {
	case "":
		break
	default:
		medianToInt, err = strconv.Atoi(medianOfCallsTime)
		if err != nil {
			log.Printf("Incorrect data in medianOfCallsTime: %s, %s\n", medianOfCallsTime, err)
			return 0, 0, 0, 0, err
		}
	}

	return float32(connFloat), TTFBtoInt, voicePurToInt, medianToInt, nil
}

func CountryCheck(code string) (string, error) {
	query := gountries.New()                       //use library gountries to get countries collection
	country, err := query.FindCountryByAlpha(code) //using countries collection to check alpha2 code
	if err != nil {
		return "", err
	}

	return country.Name.Official, nil
}

func BandwidthCheck(band string) error {
	_, err := strconv.Atoi(band)
	if err != nil {
		return err
	}
	return nil
}

func RespTimeCheck(respTime string) error {
	_, err := strconv.Atoi(respTime)
	if err != nil {
		return err
	}
	return nil
}

func ProviderCheck(provider string, provName []string) error {
	err := errors.New("invalid syntax")
	for _, Name := range provName {
		if strings.Compare(provider, Name) == 0 {
			return nil
		}
	}
	return err
}

func DeliveryTimeCheckAndConv(deliveryTime string) (int, error) {
	num, err := strconv.Atoi(deliveryTime)
	if err != nil {
		return 0, err
	}
	return num, nil
}
