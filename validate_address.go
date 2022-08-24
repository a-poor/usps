package usps

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// Bool is a boolean value that can be unmarshalled from a string.
type Bool bool

func (b *Bool) UnmarshalText(text []byte) error {
	switch strings.ToUpper(string(text)) {
	case "Y", "TRUE":
		*b = true
	case "N", "FALSE":
		*b = false
	default:
		return fmt.Errorf("failed to unmarshal bool: %s", text)
	}
	return nil
}

type ValidateAddressRequest struct {
	Addresses []Address
}

func (r ValidateAddressRequest) Validate() error {
	if len(r.Addresses) == 0 {
		return fmt.Errorf("can't have empty Addresses field")
	}
	if len(r.Addresses) > 5 {
		return fmt.Errorf("can't validate more than 5 addresses at a time")
	}
	return nil
}

type validateAddressRequest struct {
	XMLName   xml.Name `xml:"AddressValidateRequest"`
	UserID    string   `xml:"USERID,attr"`
	Revision  int
	Addresses []address `xml:"Address"`
}

type ValidatedAddress struct {
	ID                   int `xml:",attr"` // Index of the address in the request
	FirmName             string
	Address1             string
	Address2             string
	Address2Abbreviation string
	City                 string
	CityAbbreviation     string
	State                string
	Urbanization         string
	Zip5                 string
	Zip4                 string
	DeliveryPoint        string
	ReturnText           string
	CarrierRoute         string
	Footnotes            Footnote
	DPVConfirmation      DPVConfirmation
	DPVCMRA              Bool
	DPVFootnotes         DVPFootnote
	Business             Bool
	CentralDeliveryPoint Bool
	Vacant               Bool
}

type ValidateAddressResponse struct {
	Addresses []ValidatedAddress
	Err       *XMLError
}

func ValidateAddressErrorRes(err *XMLError) *ValidateAddressResponse {
	return &ValidateAddressResponse{Err: err}
}

func ValidateAddressSuccessRes(addrs []ValidatedAddress) *ValidateAddressResponse {
	return &ValidateAddressResponse{Addresses: addrs}
}

func (res *ValidateAddressResponse) IsError() bool {
	return res.Err != nil
}

func (res *ValidateAddressResponse) GetError() error {
	return res.Err
}

func (c *USPSClient) ValidateAddress(ctx context.Context, req *ValidateAddressRequest) (*ValidateAddressResponse, error) {
	// Validate the request
	if req == nil {
		return nil, fmt.Errorf("can't validate nil request")
	}
	if err := req.Validate(); err != nil {
		return nil, fmt.Errorf("invalid request: %w", err)
	}

	// Format the addresses
	addrs := make([]address, len(req.Addresses))
	for i, a := range req.Addresses {
		addrs[i] = a.withID(i)
	}

	// Create the request with additional fields
	xreq := validateAddressRequest{
		UserID:    c.UserID,
		Revision:  1, // TODO – Make optional?
		Addresses: addrs,
	}

	// Marshal the request
	b, err := xml.Marshal(xreq)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request as XML: %w", err)
	}

	// Add the header
	x := xmlHeader + string(b) // TODO – Is this necessary?

	// Get the URL
	u := c.fmtURL(verifyAPI, x)

	// Format the request
	r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create http-request from data: %w", err)
	}

	// Send the request
	resp, err := c.HTTPClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("failed to send http-request: %w", err)
	}
	defer resp.Body.Close()

	// Read the body
	b, err = io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal the response
	type AddressValidateResponse struct {
		AddressValidateResponse xml.Name
		Address                 []ValidatedAddress
	}
	var res struct {
		*AddressValidateResponse
		*XMLError
	}
	if err := xml.Unmarshal(b, &res); err != nil {
		err = fmt.Errorf("failed to unmarshal response as xml: %w", err)
		panic(err)
	}

	// Check for errors
	if res.XMLError != nil {
		return ValidateAddressErrorRes(res.XMLError), nil
	}

	// Return success!
	return ValidateAddressSuccessRes(res.Address), nil
}
