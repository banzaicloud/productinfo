// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/attributes"
	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/images"
	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/products"
	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/provider"
	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/providers"
	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/regions"
	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/service"
	"github.com/banzaicloud/productinfo/pkg/productinfo-client/client/services"
)

// Default productinfo HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new productinfo HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Productinfo {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new productinfo HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Productinfo {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new productinfo client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Productinfo {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Productinfo)
	cli.Transport = transport

	cli.Attributes = attributes.New(transport, formats)

	cli.Images = images.New(transport, formats)

	cli.Products = products.New(transport, formats)

	cli.Provider = provider.New(transport, formats)

	cli.Providers = providers.New(transport, formats)

	cli.Regions = regions.New(transport, formats)

	cli.Service = service.New(transport, formats)

	cli.Services = services.New(transport, formats)

	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Productinfo is a client for productinfo
type Productinfo struct {
	Attributes *attributes.Client

	Images *images.Client

	Products *products.Client

	Provider *provider.Client

	Providers *providers.Client

	Regions *regions.Client

	Service *service.Client

	Services *services.Client

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Productinfo) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport

	c.Attributes.SetTransport(transport)

	c.Images.SetTransport(transport)

	c.Products.SetTransport(transport)

	c.Provider.SetTransport(transport)

	c.Providers.SetTransport(transport)

	c.Regions.SetTransport(transport)

	c.Service.SetTransport(transport)

	c.Services.SetTransport(transport)

}
