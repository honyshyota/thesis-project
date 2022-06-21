package check

import (
	"encoding/json"
	configuration "main/configs"
	resultSet "main/internal/result_set"

	log "github.com/sirupsen/logrus"
)

// CheckResult use result_set.go and check data on error and conversion in json return into handler from router.go
type ResultT struct {
	Status bool                  `json:"status"`
	Data   *resultSet.ResultSetT `json:"data"`
	Error  string                `json:"error"`
}

type ConvertJson struct {
	cfg *configuration.Configuration
}

func New(cfg *configuration.Configuration) *ConvertJson {
	return &ConvertJson{
		cfg: cfg,
	}
}

func (cj *ConvertJson) CheckResult() []byte {
	var result *ResultT
	r := resultSet.New(cj.cfg).GetResultData()

	if len(r.SMS[0]) == 0 || len(r.MMS[0]) == 0 || len(r.VoiceCall) == 0 ||
		len(r.Email) == 0 || r.Support[1] == 0 || len(r.Incidents) == 0 {

		result = &ResultT{
			Status: false,
			Error:  "Error on collect data",
		}

		jsonResult, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Println(err)
		}

		return jsonResult
	} else {
		result = &ResultT{
			Status: true,
			Data:   r,
		}

		jsonResult, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			log.Println(err)
		}

		return jsonResult
	}

}
