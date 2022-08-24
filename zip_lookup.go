package usps

// import (
// 	"context"
// 	"encoding/xml"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// type ZipCodeLookupRequest struct {
// 	Addresses []Address
// }

// type zipCodeLookupRequest struct {
// 	ZipCodeLookupRequest xml.Name
// 	UserID               string `xml:"USERID,attr"`
// 	Address              []struct {
// 		ID       int `xml:",attr"`
// 		FirmName string
// 		Address1 string
// 		Address2 string
// 		City     string
// 		State    string
// 		Zip5     string
// 		Zip4     string
// 	} `xml:"Address"`
// }

// type ZipCodeLookup struct {
// 	// (Optional) Firm name provided in request. Default is spaces.
// 	FirmName string

// 	// (Optional) Delivery Address in the destination address. May contain
// 	// secondary unit designator, such as APT or SUITE, for Accountable mail.)
// 	Address1 string

// 	// (Required) Delivery Address in the destination address.
// 	//
// 	// Required for all mail and packages, however 11-digit Destination Delivery Point
// 	// ZIP+4 Code can be provided as an alternative in
// 	Address2 string

// 	// (Optional) City name of the destination address.
// 	//
// 	// Field is required, unless a verified 11 digit DPV is provided
// 	// for the mailpiece.
// 	City string

// 	// (Optional) Two-character state code of the destination address.
// 	//
// 	// Default isspaces for International mail.
// 	State string

// 	// (Optional)
// 	Urbanization string

// 	// (Optional) Destination 5-digit ZIP Code.
// 	//
// 	// Must be 5-digits. Numeric values (0-9) only. If international, all zeroes.
// 	Zip5 string

// 	// (Optional) Destination ZIP+4.
// 	//
// 	// Numeric values (0-9) only. If International, all zeroes. Default to spaces if not available
// 	Zip4 string
// }

// type ZipCodeLookupResponse struct {
// 	Data  []ZipCodeLookup
// 	Error *XMLError
// }

// func (c *USPSClient) ZipLookup(ctx context.Context, req *ZipCodeLookupRequest) (*ZipCodeLookupResponse, error) {
// 	// Validate the request
// 	if len(req.Addresses) == 0 {
// 		return nil, fmt.Errorf("can't have empty Addresses field")
// 	}
// 	if len(req.Addresses) > 5 {
// 		return nil, fmt.Errorf("can't validate more than 5 addresses at a time")
// 	}

// 	// Build the request
// 	xreq := zipCodeLookupRequest{
// 		UserID: c.UserID,
// 	}
// 	for i, addr := range req.Addresses {
// 		xreq.Address = append(xreq.Address, struct {
// 			ID       int `xml:",attr"`
// 			FirmName string
// 			Address1 string
// 			Address2 string
// 			City     string
// 			State    string
// 			Zip5     string
// 			Zip4     string
// 		}{
// 			ID:       i,
// 			FirmName: addr.FirmName,
// 			Address1: addr.Address1,
// 			Address2: addr.Address2,
// 			City:     addr.City,
// 			State:    addr.State,
// 			Zip5:     addr.Zip5,
// 			Zip4:     addr.Zip4,
// 		})
// 	}

// 	// Marshal the request
// 	b, err := xml.Marshal(xreq)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal request as XML: %w", err)
// 	}

// 	// Add the header
// 	x := xmlHeader + string(b) // TODO – Is this necessary?

// 	// Get the URL
// 	u := c.fmtURL(zipLookupAPI, x)

// 	// Format the request
// 	r, err := http.NewRequestWithContext(ctx, http.MethodGet, u, nil)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to create http-request from data: %w", err)
// 	}

// 	// Send the request
// 	resp, err := c.HTTPClient.Do(r)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to send http-request: %w", err)
// 	}
// 	defer resp.Body.Close()

// 	// Read the body
// 	b, err = io.ReadAll(resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read response body: %w", err)
// 	}

// 	// Unmarshal the response
// 	type ZipLookupResponse struct {
// 		ZipCodeLookupResponse xml.Name
// 		Address               []ZipCodeLookup
// 	}
// 	var res struct {
// 		*ZipLookupResponse
// 		*XMLError
// 	}
// 	if err := xml.Unmarshal(b, &res); err != nil {
// 		err = fmt.Errorf("failed to unmarshal response as xml: %w", err)
// 		panic(err)
// 	}

// 	// Check for errors
// 	if res.XMLError != nil {
// 		return &ZipCodeLookupResponse{Error: res.XMLError}, nil
// 	}

// 	// Return success!
// 	return &ZipCodeLookupResponse{Data: res.ZipLookupResponse.Address}, nil
// }
