// Package fiserv provides a client for the North American API at https://docs.firstdata.com/org/gateway/docs/api.
package fiserv

//go:generate rm -f types.gen.go
//go:generate rm -f client.gen.go
//go:generate ./bin/oapi-codegen -templates=./templates -generate=types -package=fiserv -o types.gen.go YAMLSpec-6.14-v2-10-15-2020.yaml
//go:generate ./bin/oapi-codegen -templates=./templates -generate=client -package=fiserv -o client.gen.go YAMLSpec-6.14-v2-10-15-2020.yaml

const (
	ProductionServer = "https://prod.api.firstdata.com/gateway/v2"
	SandboxServer    = "https://cert.api.firstdata.com/gateway/v2"
)
