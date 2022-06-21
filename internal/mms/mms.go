package mms

import (
	"encoding/json"
	"io/ioutil"
	"main/internal/additional"
	"net/http"
	"sort"

	log "github.com/sirupsen/logrus"
)

type MMSData struct {
	Country      string `json:""`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type MMSReport struct {
	sourceURL         string
	acceptedProviders []string
}

func New(sourceURL string, acceptedProviders []string) *MMSReport {
	return &MMSReport{
		sourceURL:         sourceURL,
		acceptedProviders: acceptedProviders,
	}
}

func (mr MMSReport) Make() ([]*MMSData, []*MMSData) {
	var sortByCountry []*MMSData

	resp, err := http.Get(mr.sourceURL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Incorrect status code, %v", resp.StatusCode)
	}

	resultJSON, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	err = json.Unmarshal(resultJSON, &sortByCountry)
	if err != nil {
		log.Printf("Failed to unmarshling JSON from body: %s", err)
	}

	sortByProvider := sortByCountry

	for i, sliceMMS := range sortByCountry {
		replace, err := additional.CountryCheck(sliceMMS.Country)
		if err != nil {
			log.Println(err)
			sortByCountry = append(sortByCountry[:i], sortByCountry[i+1:]...)
			continue
		}
		sortByCountry[i].Country = replace
		sortByProvider[i].Country = sortByCountry[i].Country

		err = additional.ProviderCheck(sliceMMS.Provider, mr.acceptedProviders)
		if err != nil {
			log.Printf("Incorrect data in Provider: %s\n", sliceMMS.Provider)
			sortByCountry = append(sortByCountry[:i], sortByCountry[i+1:]...)
			continue
		}
		sortByProvider[i].Provider = sortByCountry[i].Provider

		err = additional.BandwidthCheck(sliceMMS.Bandwidth)
		if err != nil {
			log.Printf("Incorrect data in Bandwidth: %s\n", sliceMMS.Bandwidth)
			sortByCountry = append(sortByCountry[:i], sortByCountry[i+1:]...)
			continue
		}
		sortByProvider[i].Bandwidth = sortByCountry[i].Bandwidth

		err = additional.RespTimeCheck(sliceMMS.ResponseTime)
		if err != nil {
			log.Printf("Incorrect data in ResponseTime: %s\n", sliceMMS.ResponseTime)
			sortByCountry = append(sortByCountry[:i], sortByCountry[i+1:]...)
			continue
		}
		sortByProvider[i].ResponseTime = sortByCountry[i].ResponseTime
	}

	sort.Slice(sortByProvider, func(i, j int) bool {
		return sortByProvider[i].Provider < sortByProvider[j].Provider
	})

	sort.Slice(sortByCountry, func(i, j int) bool {
		return sortByCountry[i].Country < sortByCountry[j].Country
	})

	return sortByProvider, sortByCountry
}
