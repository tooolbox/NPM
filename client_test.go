package fiserv

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/lufia/httpclientutil"
)

func TestClient(t *testing.T) {

	// Credentials are from https://docs.firstdata.com/org/gateway/docs/api
	// and they may be subject to change.
	creds := Credentials{
		ApiKey:    "csn5gVMfGgXh1cnFtimlHQOH1zNERw7Q",
		ApiSecret: "9JiLZgNhQPTA1KIt",
	}

	httper := http.Client{
		Transport: &httpclientutil.DumpTransport{},
	}

	clnt, err := NewClientWithResponses(SandboxServer, creds, WithHTTPClient(&httper))
	if err != nil {
		t.Fatal(err)
	}

	payload := `
{
  "requestType": "PaymentCardSaleTransaction",
  "transactionAmount": {
    "total": "12.04",
    "currency": "USD"
  },
  "paymentMethod": {
    "paymentCard": {
      "number": "5424180279791732",
      "securityCode": "977",
      "expiryDate": {
        "month": "12",
        "year": "24"
      }
    }
  }
}`

	resp, err := clnt.SubmitPrimaryTransactionWithBodyWithResponse(
		context.Background(),
		&SubmitPrimaryTransactionParams{},
		"application/json",
		strings.NewReader(payload),
	)
	if err != nil {
		t.Fatal(err)
	}

	if resp.JSON200 == nil {
		t.Fatalf("Unexpected status code: %d", resp.StatusCode())
	}

	t.Logf("Approved amount: %f", resp.JSON200.ApprovedAmount.Total)

}
