// Package fiserv provides a client for the North American API at https://docs.firstdata.com/org/gateway/docs/api.
package fiserv

//go:generate echo rm -f types.gen.go
//go:generate rm -f types.gen.go
//go:generate echo rm -f client.gen.go
//go:generate rm -f client.gen.go
//go:generate echo oapi-codegen -generate=types -package=fiserv YAMLSpec-6.14-v2-10-15-2020.yaml
//go:generate oapi-codegen -generate=types -package=fiserv -o types.gen.go YAMLSpec-6.14-v2-10-15-2020.yaml
//go:generate echo oapi-codegen -generate=client -package=fiserv YAMLSpec-6.14-v2-10-15-2020.yaml
//go:generate oapi-codegen -generate=client -package=fiserv -o client.gen.go YAMLSpec-6.14-v2-10-15-2020.yaml
