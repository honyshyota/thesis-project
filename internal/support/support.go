package support

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type SupportReport struct {
	sourceURL string
}

func New(sourceURL string) *SupportReport {
	return &SupportReport{
		sourceURL: sourceURL,
	}
}

func (supr SupportReport) Make() []int {
	var resultSupport []*SupportData
	resp, err := http.Get(supr.sourceURL)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Incorrect status code, %v", resp.StatusCode)
	}

	jsnUnmrshl, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	dec := json.NewDecoder(strings.NewReader(string(jsnUnmrshl)))

	err = dec.Decode(&resultSupport)
	if err == io.EOF {
		log.Println(err)
	} else if err != nil {
		log.Println(err)
	}

	sumTickets := 0
	load := 0

	for _, data := range resultSupport {
		sumTickets += data.ActiveTickets
	}

	if sumTickets <= 9 {
		load = 1
	} else if sumTickets > 9 && sumTickets <= 16 {
		load = 2
	} else {
		load = 3
	}

	responseTime := sumTickets * 60 / 18
	result := make([]int, 2)
	result[0] = load
	result[1] = responseTime

	return result
}
