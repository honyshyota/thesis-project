package resultSet

import (
	configuration "main/configs"
	"main/internal/additional"
	"main/internal/billing"
	"main/internal/email"
	"main/internal/incident"
	"main/internal/mms"
	"main/internal/sms"
	"main/internal/support"
	"main/internal/voice"
	"sync"
)

// ResultSet collection all data and represents one structure

type ResultSetT struct {
	SMS       [][]*sms.SMSData                `json:"sms"`
	MMS       [][]*mms.MMSData                `json:"mms"`
	VoiceCall []*voice.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]*email.EmailData `json:"email"`
	Billing   *billing.BillingData            `json:"billing"`
	Support   []int                           `json:"support"`
	Incidents []*incident.IncidentData        `json:"incident"`
}

type ResultReport struct {
	cfg *configuration.Configuration
}

func New(cfg *configuration.Configuration) *ResultReport {
	return &ResultReport{
		cfg: cfg,
	}
}

func (rr ResultReport) GetResultData() *ResultSetT {
	var wg sync.WaitGroup

	var result *ResultSetT
	var smsData [][]*sms.SMSData
	var mmsData [][]*mms.MMSData
	var voiceData []*voice.VoiceCallData
	var emailData map[string][][]*email.EmailData
	var billingData *billing.BillingData
	var supportData []int
	var incidentData []*incident.IncidentData

	wg.Add(7)
	// Implementation of concurency
	go func() {
		defer wg.Done()

		sms := sms.New(additional.GetFilePathByFileName(rr.cfg.SmsFileName), rr.cfg.SmsMmsProv)
		smsAlpha2, smsCountry := sms.Make()
		smsData = append(smsData, smsAlpha2, smsCountry)
	}()
	go func() {
		defer wg.Done()

		mms := mms.New(rr.cfg.MmsURL, rr.cfg.SmsMmsProv)
		mmsAlpha2, mmsCountry := mms.Make()
		mmsData = append(mmsData, mmsAlpha2, mmsCountry)
	}()
	go func() {
		defer wg.Done()
		voice := voice.New(additional.GetFilePathByFileName(rr.cfg.VoiceFileName), rr.cfg.VoiceProv)
		voiceData = append(voiceData, voice.Make()...)
	}()
	go func() {
		defer wg.Done()
		email := email.New(additional.GetFilePathByFileName(rr.cfg.EmailFileName), rr.cfg.EmailProv)
		emailSlice := email.Make()
		emailData = emailSlice
	}()
	go func() {
		defer wg.Done()
		billing := billing.New(additional.GetFilePathByFileName(rr.cfg.BillingFileName))
		billingResult := billing.Make()
		billingData = billingResult
	}()
	go func() {
		defer wg.Done()
		support := support.New(rr.cfg.SupportURL)
		supportData = append(supportData, support.Make()...)
	}()
	go func() {
		defer wg.Done()
		incident := incident.New(rr.cfg.IncidentURL)
		incidentData = append(incidentData, incident.Make()...)
	}()

	wg.Wait()

	result = &ResultSetT{
		SMS:       smsData,
		MMS:       mmsData,
		VoiceCall: voiceData,
		Email:     emailData,
		Billing:   billingData,
		Support:   supportData,
		Incidents: incidentData,
	}

	return result
}
