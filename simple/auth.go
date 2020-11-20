package simple

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

type AuthenticateRequestFunc func(req runtime.ClientRequest, reg strfmt.Registry) error

func (f AuthenticateRequestFunc) AuthenticateRequest(req runtime.ClientRequest, reg strfmt.Registry) error {
	return f(req, reg)
}