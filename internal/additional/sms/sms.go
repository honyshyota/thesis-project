package sms

import (
	"io/ioutil"
	"main/internal/additional"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type SMSReport struct {
	sourcePath        string
	acceptedProviders []string
}

func New(sourcePath string, acceptedProviders []string) *SMSReport {
	return &SMSReport{
		sourcePath:        sourcePath,
		acceptedProviders: acceptedProviders,
	}
}

func (sr SMSReport) Make() ([]*SMSData, []*SMSData) {
	var replaceCountry []*SMSData
	var resultSMS []*SMSData
	smsCollection, err := ioutil.ReadFile(sr.sourcePath)
	if err != nil {
		log.Println("Failed reading sms.data", err)
		return nil, nil
	}

	smsStringSlice := strings.Fields(string(smsCollection))
	var smsResultString string

	for _, smsString := range smsStringSlice {
		splitSlice := strings.Split(smsString, ";")

		if len(splitSlice) != 4 {
			log.Printf("Incorrect lenght data string, %s\n", splitSlice)
			continue
		}

		countryName, err := additional.CountryCheck(splitSlice[0])
		if err != nil {
			log.Println(err)
			continue
		}

		err = additional.BandwidthCheck(splitSlice[1])
		if err != nil {
			log.Printf("Incorrect data in Bandwidth: %s, %s\n", splitSlice[1], err)
			continue
		}

		err = additional.RespTimeCheck(splitSlice[2])
		if err != nil {
			log.Printf("Incorrect data in ResponseTime: %s, %s\n", splitSlice[2], err)
			continue
		}

		err = additional.ProviderCheck(splitSlice[3], sr.acceptedProviders)
		if err != nil {
			log.Printf("Incorrect data in Provider: %s, %s\n", splitSlice[3], err)
			continue
		}

		var result = &SMSData{
			Country:      countryName,
			Bandwidth:    splitSlice[1],
			ResponseTime: splitSlice[2],
			Provider:     splitSlice[3],
		}

		var replace = &SMSData{
			Country:      countryName,
			Bandwidth:    splitSlice[1],
			ResponseTime: splitSlice[2],
			Provider:     splitSlice[3],
		}
		resultSMS = append(resultSMS, result)
		replaceCountry = append(replaceCountry, replace)
		smsResultString = smsResultString + smsString + "\n"
	}

	sort.Slice(resultSMS, func(i, j int) bool {
		return resultSMS[i].Provider < resultSMS[j].Provider
	})

	sort.Slice(replaceCountry, func(i, j int) bool {
		return replaceCountry[i].Country < replaceCountry[j].Country
	})

	return resultSMS, replaceCountry
}
