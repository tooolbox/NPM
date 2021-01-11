package fiserv

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

type Credentials struct {
	ApiKey    string
	ApiSecret string
}

func (crd Credentials) GenRequestHeaders(payload []byte) (*HeaderData, error) {
	clientRequestId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return crd.genHeaders(payload, clientRequestId.String(), time.Now())
}

func (crd Credentials) genHeaders(payload []byte, clientRequestId string, now time.Time) (*HeaderData, error) {

	// equivalent python code:
	//
	// self.api_key = params['api_key']
	// self.api_secret = params['api_secret']
	// self.client_request_id = str(uuid.uuid4())
	// self.timestamp = str(int(round(time.time()*1000)))
	//
	// data = self.api_key + self.client_request_id + self.timestamp
	// if payload:
	// 	data += json.dumps(self.sanitize_for_serialization(payload))

	msSinceEpoch := now.Unix() * 1000

	toSign := []byte(crd.ApiKey + clientRequestId + fmt.Sprintf("%d", msSinceEpoch))
	if len(payload) > 0 {
		toSign = append(toSign, payload...)
	}

	sig, err := genMsgSignature(crd.ApiSecret, toSign)
	if err != nil {
		return nil, err
	}

	return &HeaderData{
		ApiKey:           ApiKeyParam(crd.ApiKey),
		ContentType:      ContentTypeParam_application_json,
		ClientRequestId:  ClientRequestIdParam(clientRequestId),
		Timestamp:        TimestampParam(msSinceEpoch),
		MessageSignature: MessageSignatureParam(sig),
	}, nil
}

func genMsgSignature(apiSecret string, data []byte) (string, error) {

	// equivalent python code:
	//
	// h = hmac.new(self.api_secret, data, hashlib.sha256).hexdigest()
	// return base64.b64encode(h)

	// equivalent JS code:
	// var genMsgSignature = function (secret, data) {
	//     return Base64.stringify(crypto_js_1.HmacSHA256(data, secret));
	// };

	hasher := hmac.New(sha256.New, []byte(apiSecret))
	if _, err := hasher.Write(data); err != nil {
		return "", fmt.Errorf("unable to hash data: %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return b64, nil
}
