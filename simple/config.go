package simple

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/tooolbox/fiserv"
)

const (
	ProdHost    = "prod.api.firstdata.com/gateway/v2"
	SandboxHost = "cert.api.firstdata.com/gateway/v2"

	ProdEnvironment    = "production"
	SandboxEnvironment = "sandbox"
)

type Config struct {
	ApiKey      string
	ApiSecret   string
	Environment string
}

type HeaderData struct {
	ApiKey           fiserv.ApiKeyParam
	ContentType      fiserv.ContentTypeParam
	ClientRequestId  fiserv.ClientRequestIdParam
	Timestamp        fiserv.TimestampParam
	MessageSignature fiserv.MessageSignatureParam
}

func (cfg Config) genHeaders(payload []byte, clientRequestId string, now time.Time) (*HeaderData, error) {

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

	toSign := []byte(cfg.ApiKey + clientRequestId + fmt.Sprintf("%d", msSinceEpoch))
	if len(payload) > 0 {
		toSign = append(toSign, payload...)
	}

	sig, err := genMsgSignature(cfg.ApiSecret, toSign)
	if err != nil {
		return nil, err
	}

	return &HeaderData{
		ApiKey:           fiserv.ApiKeyParam(cfg.ApiKey),
		ContentType:      fiserv.ContentTypeParam_application_json,
		ClientRequestId:  fiserv.ClientRequestIdParam(clientRequestId),
		Timestamp:        fiserv.TimestampParam(msSinceEpoch),
		MessageSignature: fiserv.MessageSignatureParam(sig),
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
