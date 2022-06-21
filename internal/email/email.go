package email

import (
	"io/ioutil"
	"main/internal/additional"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

type EmailReport struct {
	sourcePath        string
	acceptedProviders []string
}

func New(sourcePath string, acceptedProviders []string) *EmailReport {
	return &EmailReport{
		sourcePath:        sourcePath,
		acceptedProviders: acceptedProviders,
	}
}

func (er EmailReport) Make() map[string][][]*EmailData {
	var resultHighSpeed []*EmailData
	var resultLowSpeed []*EmailData
	var dataEmail [][]*EmailData

	emailCollection, err := ioutil.ReadFile(er.sourcePath)
	if err != nil {
		log.Println("Failed reading data from file. ", err)
	}

	emailStringSlice := strings.Fields(string(emailCollection))
	var emailResultString string

	for _, emailString := range emailStringSlice {
		splitSlice := strings.Split(emailString, ";")
		if len(splitSlice) != 3 {
			log.Println("incorrect lenght data in string: ", splitSlice)
			continue
		}

		_, err := additional.CountryCheck(splitSlice[0])
		if err != nil {
			log.Println(err)
			continue
		}

		err = additional.ProviderCheck(splitSlice[1], er.acceptedProviders)
		if err != nil {
			log.Printf("Incorrect data in Provider: %s\n", splitSlice[1])
			continue
		}

		delTime, err := additional.DeliveryTimeCheckAndConv(splitSlice[2])
		if err != nil {
			log.Printf("Incorrect data in DeliveryTime: %s, %s\n", splitSlice[2], err)
		}

		var highResult = &EmailData{
			Country:      splitSlice[0],
			Provider:     splitSlice[1],
			DeliveryTime: delTime,
		}

		var lowResult = &EmailData{
			Country:      splitSlice[0],
			Provider:     splitSlice[1],
			DeliveryTime: delTime,
		}

		resultHighSpeed = append(resultHighSpeed, highResult)
		resultLowSpeed = append(resultLowSpeed, lowResult)

		emailResultString = emailResultString + emailString + "\n"
	}

	sort.Slice(resultHighSpeed, func(i, j int) bool {
		return resultHighSpeed[i].DeliveryTime < resultHighSpeed[j].DeliveryTime
	})

	sort.Slice(resultLowSpeed, func(i, j int) bool {
		return resultLowSpeed[i].DeliveryTime > resultLowSpeed[j].DeliveryTime
	})

	dataEmail = append(dataEmail, resultHighSpeed, resultLowSpeed)
	resultSortedEmail := make(map[string][][]*EmailData)
	for i := range resultHighSpeed {
		resultSortedEmail[resultHighSpeed[i].Country] = nil
	}
	var con []*EmailData
	for i := range resultSortedEmail {
		for a, data := range dataEmail {
			for _, y := range data {
				if y.Country == i {
					if a == 0 {
						con = append(con, y)
						if len(con) > 2 {
							resultSortedEmail[i] = append(resultSortedEmail[i], con)
							con = nil
							break
						}
					} else {
						con = append(con, y)
						if len(con) > 2 {
							resultSortedEmail[i] = append(resultSortedEmail[i], con)
							con = nil
							break
						}
					}
				}
			}
		}
	}

	return resultSortedEmail
}
