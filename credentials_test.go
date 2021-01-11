package fiserv

import (
	"testing"
	"time"
)

func TestGenMsgSignature(t *testing.T) {
	secret := "5bAXQnRHT9bS8kdz"
	data := `H1r5Zqnfg1Yh32FmvUU7rabZuapaw7B1e3413f75-140a-40da-a0ae-5fdc295ad73a1610253841000{"merchantTransactionId":"410cb84c-d6ca-418b-acbd-bfaab9fcce08","order":{},"requestType":"requestType","storeId":"987654321987","transactionAmount":{"components":{},"currency":"USD","total":132.45},"transactionOrigin":"ECOM"}
`

	sig, err := genMsgSignature(secret, []byte(data))
	if err != nil {
		t.Fatal(err)
	}

	if expect := `pqPHd0HbMO/Ummoh+XiAmi6pOiLS47+VSiH9d2/dZqQ=`; sig != expect {
		t.Fatalf("Expected %s but found %s", expect, sig)
	}
}

func TestGenHeaders(t *testing.T) {
	creds := Credentials{
		ApiKey:    "H1r5Zqnfg1Yh32FmvUU7rabZuapaw7B1",
		ApiSecret: "5bAXQnRHT9bS8kdz",
	}

	hdrs, err := creds.genHeaders(
		[]byte(`{"merchantTransactionId":"410cb84c-d6ca-418b-acbd-bfaab9fcce08","order":{},"requestType":"requestType","storeId":"987654321987","transactionAmount":{"components":{},"currency":"USD","total":132.45},"transactionOrigin":"ECOM"}
`),
		"e3413f75-140a-40da-a0ae-5fdc295ad73a",
		time.Unix(1610253841, 0),
	)
	if err != nil {
		t.Fatal(err)
	}

	if string(hdrs.ApiKey) != creds.ApiKey {
		t.Fatalf("Unexpected ApiKey value: %s", hdrs.ApiKey)
	}

	if hdrs.ContentType != "application/json" {
		t.Fatalf("Unexpected ContentType value: %s", hdrs.ContentType)
	}

	if hdrs.Timestamp != 1610253841000 {
		t.Fatalf("Unexpected Timestamp value: %d", hdrs.Timestamp)
	}

	if hdrs.ClientRequestId != "e3413f75-140a-40da-a0ae-5fdc295ad73a" {
		t.Fatalf("Unexpected ClientRequestId value: %s", hdrs.ClientRequestId)
	}

	if hdrs.MessageSignature != `pqPHd0HbMO/Ummoh+XiAmi6pOiLS47+VSiH9d2/dZqQ=` {
		t.Fatalf("Unexpected MessageSignature value: %s", hdrs.MessageSignature)
	}

}
