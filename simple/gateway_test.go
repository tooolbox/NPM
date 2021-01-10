package simple

import (
	"context"
	"net/http"
	"strings"
	"testing"

	"github.com/lufia/httpclientutil"
	"github.com/tooolbox/fiserv"
)

func TestGateway(t *testing.T) {
	cfg := Config{
		ApiKey:      "csn5gVMfGgXh1cnFtimlHQOH1zNERw7Q",
		ApiSecret:   "9JiLZgNhQPTA1KIt",
		Environment: "sandbox",
	}

	client := http.Client{
		Transport: &httpclientutil.DumpTransport{},
	}

	gw, err := NewGateway(cfg, fiserv.WithHTTPClient(&client))
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

	hdrs, err := gw.GenHeaders([]byte(payload))
	if err != nil {
		t.Fatal(err)
	}

	params := fiserv.SubmitPrimaryTransactionParams{
		ApiKey:           hdrs.ApiKey,
		Timestamp:        hdrs.Timestamp,
		MessageSignature: &hdrs.MessageSignature,
		ContentType:      hdrs.ContentType,
		ClientRequestId:  hdrs.ClientRequestId,
	}

	if _, err := gw.SubmitPrimaryTransactionWithBody(
		context.Background(),
		&params,
		string(hdrs.ContentType),
		strings.NewReader(payload),
	); err != nil {
		t.Fatal(err)
	}

}
