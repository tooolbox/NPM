package simple

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofrs/uuid"
	"github.com/tooolbox/fiserv"
)

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

func (gw *Gateway) GenHeaders(payload []byte) (*HeaderData, error) {
	clientRequestId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	return gw.cfg.genHeaders(payload, clientRequestId.String(), time.Now())
}
