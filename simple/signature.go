package simple

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"
)

func (gw *Gateway) SignatureAuth(host string) (f AuthenticateRequestFunc) {
	f = func(req runtime.ClientRequest, reg strfmt.Registry) error {

		// self.api_key = params['api_key']
		// self.api_secret = params['api_secret']
		// self.client_request_id = str(uuid.uuid4())
		// self.timestamp = str(int(round(time.time()*1000)))

		clientRequestId, err := uuid.NewV4()
		if err != nil {
			return err
		}

		timestamp := fmt.Sprintf("%d", time.Now().Unix())

		if gw.debugLogging {
			log.Println("Performing Signature Auth...")
			log.Println("ApiKey: " + gw.cfg.ApiKey)
		}

		hdrs := req.GetHeaderParams()

		hdrs["Client-Request-Id"] = []string{clientRequestId.String()}
		hdrs["Api-Key"] = []string{gw.cfg.ApiKey}
		hdrs["Timestamp"] = []string{timestamp}
		hdrs["Region"] = []string{}

		body := req.GetBody()

		if gw.debugLogging {
			log.Printf("Body: `%s`", string(body))
		}

		// equivalent python code:
		//
		// data = self.api_key + self.client_request_id + self.timestamp
		// if payload:
		// 	data += json.dumps(self.sanitize_for_serialization(payload))

		toSign := []byte(gw.cfg.ApiKey + clientRequestId.String() + timestamp)
		if len(body) > 0 {
			toSign = append(toSign, body...)
		}

		signature, err := generateSignature(gw.cfg.ApiSecret, toSign)
		if err != nil {
			return err
		}
		hdrs["Message-Signature"] = []string{signature}

		return nil
	}
	return
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
