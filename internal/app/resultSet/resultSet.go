package resultSet

import (
	"main/internal/additional"
	"main/internal/additional/billing"
	"main/internal/additional/email"
	"main/internal/additional/incident"
	"main/internal/additional/mms"
	"main/internal/additional/sms"
	"main/internal/additional/support"
	"main/internal/additional/voice"
	"os"
	"strings"
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

type ResultReport struct{}

func New() *ResultReport {
	return &ResultReport{}
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

		sms := sms.New(additional.GetFilePathByFileName(os.Getenv("SMS_FILE_NAME")), strings.Split(os.Getenv("SMS_MMS_PROV"), ", "))
		smsAlpha2, smsCountry := sms.Make()
		smsData = append(smsData, smsAlpha2, smsCountry)
	}()
	go func() {
		defer wg.Done()

		mms := mms.New(os.Getenv("MMS_URL"), strings.Split(os.Getenv("SMS_MMS_PROV"), ", "))
		mmsAlpha2, mmsCountry := mms.Make()

		mmsData = append(mmsData, mmsAlpha2, mmsCountry)
	}()
	go func() {
		defer wg.Done()
		voice := voice.New(additional.GetFilePathByFileName(os.Getenv("VOICE_FILE_NAME")), strings.Split(os.Getenv("VOICE_PROV"), ", "))
		voiceData = append(voiceData, voice.Make()...)
	}()
	go func() {
		defer wg.Done()
		email := email.New(additional.GetFilePathByFileName(os.Getenv("EMAIL_FILE_NAME")), strings.Split(os.Getenv("EMAIL_PROV"), ", "))
		emailSlice := email.Make()
		emailData = emailSlice
	}()
	go func() {
		defer wg.Done()
		billing := billing.New(additional.GetFilePathByFileName(os.Getenv("BILLING_FILE_NAME")))
		billingResult := billing.Make()
		billingData = billingResult
	}()
	go func() {
		defer wg.Done()
		support := support.New(os.Getenv("SUPPORT_URL"))
		supportData = append(supportData, support.Make()...)
	}()
	go func() {
		defer wg.Done()
		incident := incident.New(os.Getenv("INCIDENT_URL"))
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
