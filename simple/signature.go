package simple

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/tooolbox/fiserv"
)

func (gw *Gateway) signature(clientRequestId uuid.UUID, msSinceEpoch int64, body []byte) (fiserv.MessageSignatureParam, error) {

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

	toSign := []byte(gw.cfg.ApiKey + clientRequestId.String() + fmt.Sprintf("%d", msSinceEpoch))
	if len(body) > 0 {
		toSign = append(toSign, body...)
	}

	sig, err := generateSignature(gw.cfg.ApiSecret, toSign)
	if err != nil {
		return "", err
	}

	return fiserv.MessageSignatureParam(sig), nil
}

func generateSignature(apiSecret string, data []byte) (string, error) {

	// equivalent python code:
	//
	// h = hmac.new(self.api_secret, data, hashlib.sha256).hexdigest()
	// return base64.b64encode(h)

	decodedKey, err := base64.StdEncoding.DecodeString(apiSecret)
	if err != nil {
		return "", fmt.Errorf("failed to decode the secret key: %v", err)
	}

	hasher := hmac.New(sha256.New, decodedKey)
	if _, err := hasher.Write(data); err != nil {
		return "", fmt.Errorf("unable to hash data: %v", err)
	}

	b64 := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	return b64, nil
}
