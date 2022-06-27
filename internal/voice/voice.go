package voice

import (
	"io/ioutil"
	"main/internal/additional"
	"strings"

	log "github.com/sirupsen/logrus"
)

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}

type VoiceReport struct {
	sourcePath        string
	acceptedProviders []string
}

func New(sourcePath string, acceptedProviders []string) *VoiceReport {
	return &VoiceReport{
		sourcePath:        sourcePath,
		acceptedProviders: acceptedProviders,
	}
}

func (vr VoiceReport) Make() []*VoiceCallData {
	var resultVoiceData []*VoiceCallData

	voiceCollection, err := ioutil.ReadFile(vr.sourcePath)
	if err != nil {
		log.Fatalln("Failed to read data from file, ", err)
		return nil
	}
	voiceStringSlice := strings.Fields(string(voiceCollection))
	var voiceStringResult string

	for _, voiceString := range voiceStringSlice {
		sSlice := strings.Split(voiceString, ";")
		if len(sSlice) != 8 {
			log.Println("incorrect lenght data string in ", sSlice)
			continue
		}

		_, err := additional.CountryCheck(sSlice[0])
		if err != nil {
			log.Println(err)
			continue
		}

		err = additional.BandwidthCheck(sSlice[1])
		if err != nil {
			log.Printf("Incorrect data in Bandwidth: %s\n", sSlice[1])
			continue
		}

		err = additional.RespTimeCheck(sSlice[2])
		if err != nil {
			log.Printf("Incorrect data in ResponseTime: %s\n", sSlice[2])
			continue
		}

		err = additional.ProviderCheck(sSlice[3], vr.acceptedProviders)
		if err != nil {
			log.Printf("Incorrect data in Provider: %s\n", sSlice[3])
			continue
		}

		cStab, TTFB, vPurity, mTime, err := additional.VoiceParametersCheckAndConv(sSlice[4], sSlice[5], sSlice[6], sSlice[7])
		if err != nil {
			continue
		}

		var result = &VoiceCallData{
			Country:             sSlice[0],
			Bandwidth:           sSlice[1],
			ResponseTime:        sSlice[2],
			Provider:            sSlice[3],
			ConnectionStability: cStab,
			TTFB:                TTFB,
			VoicePurity:         vPurity,
			MedianOfCallsTime:   mTime,
		}

		resultVoiceData = append(resultVoiceData, result)
		voiceStringResult = voiceStringResult + voiceString + "\n"
	}

	return resultVoiceData
}
