package simple

import (
	"fmt"
	"strings"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/tooolbox/firstvision/client"
)

const (
	ProdHost    = "prod.api.firstdata.com"
	SandboxHost = "cert.api.firstdata.com"

	ProdEnvironment    = "production"
	SandboxEnvironment = "sandbox"
)

type Config struct {
	ApiKey      string
	ApiSecret   string
	Environment string
}

type Gateway struct {
	cfg          Config
	debugLogging bool
}

func NewGateway(cfg Config) *Gateway {
	return &Gateway{
		cfg: cfg,
	}
}

func (gw *Gateway) WithDebugLogging(debug bool) *Gateway {
	gw.debugLogging = debug
	return gw
}

func (gw *Gateway) Transport() (*httptransport.Runtime, error) {

	var host string
	switch strings.ToLower(gw.cfg.Environment) {
	case ProdEnvironment:
		host = ProdHost
	case SandboxEnvironment:
		host = SandboxHost
	default:
		return nil, fmt.Errorf("unknown run environment '%s', please use one of: [%s, %s]", gw.cfg.Environment, ProdEnvironment, SandboxEnvironment)
	}

	tr := httptransport.New(
		host,
		client.DefaultBasePath,
		client.DefaultSchemes,
	)

	tr.SetDebug(gw.debugLogging)

	auth, err := gw.ClientAuthInfoWriter(host)
	if err != nil {
		return nil, err
	}

	tr.DefaultAuthentication = auth

	return tr, nil
}

func (gw *Gateway) ClientAuthInfoWriter(host string) (runtime.ClientAuthInfoWriter, error) {
	return gw.SignatureAuth(host), nil
}

func (gw *Gateway) NewHTTPClient(formats strfmt.Registry) (*client.FirstVision, error) {
	tr, err := gw.Transport()
	if err != nil {
		return nil, err
	}

	return client.New(tr, formats), nil
}
