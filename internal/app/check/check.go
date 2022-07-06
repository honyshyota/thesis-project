package check

import (
	"encoding/json"
	resultSet "main/internal/app/resultSet"
)

// CheckResult use result_set.go and check data on error and conversion in json return into handler from router.go
type ResultT struct {
	Status bool                  `json:"status"`
	Data   *resultSet.ResultSetT `json:"data"`
	Error  string                `json:"error"`
}

type ConvertJson struct{}

func New() *ConvertJson {
	return &ConvertJson{}
}

func (cj *ConvertJson) CheckResult() ([]byte, error) {
	var result *ResultT
	r := resultSet.New().GetResultData()

	if len(r.SMS[0]) == 0 || len(r.MMS[0]) == 0 || len(r.VoiceCall) == 0 ||
		len(r.Email) == 0 || r.Support[1] == 0 || len(r.Incidents) == 0 {

		result = &ResultT{
			Status: false,
			Error:  "Error on collect data",
		}

		jsonResult, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return nil, err
		}

		return jsonResult, nil
	} else {
		result = &ResultT{
			Status: true,
			Data:   r,
		}

		jsonResult, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return nil, err
		}

		return jsonResult, nil
	}

}
