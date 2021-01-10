package simple

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"time"

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

	requestId, err := uuid.NewV4()
	if err != nil {
		t.Fatal(err)
	}

	params := fiserv.SubmitPrimaryTransactionParams{
		ContentType: fiserv.ContentTypeParam_application_json,

		// A client-generated ID for request tracking and signature creation, unique per request.  This is also used for idempotency control. We recommend 128-bit UUID format.
		ClientRequestId: fiserv.ClientRequestIdParam(requestId.String()),

		// Key given to merchant after boarding associating their requests with the appropriate app in Apigee.
		ApiKey: fiserv.ApiKeyParam(os.Getenv("FISERV_KEY")),

		// Epoch timestamp in milliseconds in the request from a client system. Used for Message Signature generation and time limit (5 mins).
		Timestamp: fiserv.TimestampParam(time.Now().Unix()),

		// Used to ensure the request has not been tampered with during transmission. The Message-Signature is the Base64 encoded HMAC hash (SHA256 algorithm with the API Secret as the key.) For more information, refer to the supporting documentation on the Developer Portal.
		// MessageSignature MessageSignatureParam,

		// Indicates the region where the client wants the transaction to be processed. This will override the default processing region identified for the client. Available options are argentina, brazil, germany, india and northamerica. Region specific store setup and APIGEE boarding is required in order to use an alternate region for processing.
		// Region *RegionParam,
	}

	txnId := "bbq"
	storeId := os.Getenv("FISERV_STORE")
	origin := fiserv.TransactionOrigin_ECOM

	req := fiserv.SubmitPrimaryTransactionJSONRequestBody{
		MerchantTransactionId: &txnId,
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

	resp, err := gw.SubmitPrimaryTransaction(context.Background(), &params, req)
	if err != nil {
		t.Fatal(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))
}
