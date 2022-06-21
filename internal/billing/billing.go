package billing

import (
	"io/ioutil"
	"strconv"

	log "github.com/sirupsen/logrus"
)

const (
	CreateCustomerMask = 1 << iota
	PurchaseMask
	PayoutMask
	ReccuringMask
	FraudControlMask
	CheckOutPageMask
)

type BillingData struct {
	CreateCustomer bool `json:"create_customer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"reccuring"`
	FraudControl   bool `json:"fraud_control"`
	CheckoutPage   bool `json:"checkout_page"`
}

type BillingReport struct {
	sourcePath string
}

func New(sourcePath string) *BillingReport {
	return &BillingReport{
		sourcePath: sourcePath,
	}
}

func (br BillingReport) Make() *BillingData {
	billingCollection, err := ioutil.ReadFile(br.sourcePath)
	if err != nil {
		log.Println(err)
	}

	mask, err := strconv.ParseInt(string(billingCollection), 2, 0)
	if err != nil {
		log.Println(err)
	}

	result := &BillingData{
		CreateCustomer: mask&CreateCustomerMask != 0,
		Purchase:       mask&PurchaseMask != 0,
		Payout:         mask&PayoutMask != 0,
		Recurring:      mask&ReccuringMask != 0,
		FraudControl:   mask&FraudControlMask != 0,
		CheckoutPage:   mask&CheckOutPageMask != 0,
	}

	return result
}
