package simple

import (
	"fmt"
	"strings"

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

type Gateway struct {
	cfg Config
	fiserv.ClientInterface
}

func NewGateway(cfg Config, opts ...fiserv.ClientOption) (*Gateway, error) {

	var endpoint string
	switch strings.ToLower(cfg.Environment) {
	case ProdEnvironment:
		endpoint = ProdHost
	case SandboxEnvironment:
		endpoint = SandboxHost
	default:
		return nil, fmt.Errorf("unknown run environment '%s', please use one of: [%s, %s]", cfg.Environment, ProdEnvironment, SandboxEnvironment)
	}

	opts = append(opts)

	clnt, err := fiserv.NewClient("https://"+endpoint, opts...)
	if err != nil {
		return nil, err
	}

	return &Gateway{
		cfg:             cfg,
		ClientInterface: clnt,
	}, nil
}

// func (gw *Gateway) Auth(ctx context.Context, req *http.Request) error

// auth, err := gw.ClientAuthInfoWriter(host)
// if err != nil {
// 	return nil, err
// }

// tr.DefaultAuthentication = auth

// func (gw *Gateway) ClientAuthInfoWriter(host string) (runtime.ClientAuthInfoWriter, error) {
// 	return gw.SignatureAuth(host), nil
// }

// func (gw *Gateway) NewHTTPClient(formats strfmt.Registry) (*client.FirstVision, error) {
// 	tr, err := gw.Transport()
// 	if err != nil {
// 		return nil, err
// 	}

// 	return client.New(tr, formats), nil
// }
