package simple

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/gofrs/uuid"
	"github.com/lufia/httpclientutil"
	"github.com/tooolbox/fiserv"
)

func TestApi(t *testing.T) {
	hc := http.Client{
		Transport: &httpclientutil.DumpTransport{},
	}

	if os.Getenv("FISERV_KEY") == "" {
		t.Fatal("Please specify FISERV_KEY, FISERV_SECRET, FISERV_STORE, etc.")
	}

	gw, err := NewGateway(
		Config{
			ApiKey:      os.Getenv("FISERV_KEY"),
			ApiSecret:   os.Getenv("FISERV_SECRET"),
			Environment: "sandbox",
		},
		fiserv.WithHTTPClient(&hc),
	)
	if err != nil {
		t.Fatal(err)
	}

	txnId, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	txnIdStr := txnId.String()

	storeId := os.Getenv("FISERV_STORE")
	origin := fiserv.TransactionOrigin_ECOM

	req := fiserv.SubmitPrimaryTransactionJSONRequestBody{
		MerchantTransactionId: &txnIdStr,
		Order:                 &fiserv.Order{},
		RequestType:           "requestType",
		StoreId:               &storeId,
		TransactionAmount: fiserv.Amount{
			Components: &fiserv.AmountComponents{},
			Currency:   "USD",
			Total:      132.45,
		},
		TransactionOrigin: &origin,
	}

	var body bytes.Buffer
	enc := json.NewEncoder(&body)
	if err := enc.Encode(req); err != nil {
		t.Fatal(err)
	}

	hdrs, err := gw.GenHeaders(body.Bytes())
	if err != nil {
		t.Fatal(err)
	}

	params := fiserv.SubmitPrimaryTransactionParams{
		// Content type.
		ContentType: hdrs.ContentType,

		// A client-generated ID for request tracking and signature creation, unique per request.  This is also used for idempotency control. We recommend 128-bit UUID format.
		ClientRequestId: hdrs.ClientRequestId,

		// Key given to merchant after boarding associating their requests with the appropriate app in Apigee.
		ApiKey: hdrs.ApiKey,

		// Epoch timestamp in milliseconds in the request from a client system. Used for Message Signature generation and time limit (5 mins).
		Timestamp: hdrs.Timestamp,

		// Used to ensure the request has not been tampered with during transmission. The Message-Signature is the Base64 encoded HMAC hash (SHA256 algorithm with the API Secret as the key.) For more information, refer to the supporting documentation on the Developer Portal.
		MessageSignature: &hdrs.MessageSignature,

		// Indicates the region where the client wants the transaction to be processed. This will override the default processing region identified for the client. Available options are argentina, brazil, germany, india and northamerica. Region specific store setup and APIGEE boarding is required in order to use an alternate region for processing.
		// Region *RegionParam,
	}

	resp, err := gw.SubmitPrimaryTransactionWithBody(context.Background(), &params, string(fiserv.ContentTypeParam_application_json), &body)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))
}
