// Code generated by go-swagger; DO NOT EDIT.

package region

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new region API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for region API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetRegion Provides the detailed info of a specific region of a cloud provider
*/
func (a *Client) GetRegion(params *GetRegionParams) (*GetRegionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRegionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRegion",
		Method:             "GET",
		PathPattern:        "/providers/{provider}/services/{service}/regions/{region}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRegionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRegionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
