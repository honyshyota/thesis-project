package incident

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
)

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}

type IncidentReport struct {
	sourceURL string
}

func New(sourceURL string) *IncidentReport {
	return &IncidentReport{
		sourceURL: sourceURL,
	}
}

func (ir IncidentReport) Make() []*IncidentData {
	var resultIncedent []*IncidentData

	resp, err := http.Get(ir.sourceURL)
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

	err = json.Unmarshal(resultJSON, &resultIncedent)
	if err != nil {
		log.Println(err)
	}

	sort.Slice(resultIncedent, func(i, j int) bool {
		return resultIncedent[i].Status < resultIncedent[j].Status
	})

	return resultIncedent
}
