package usps

import (
	"net/http"
	"net/url"
)

const (
	verifyAPI          = "Verify"
	zipLookupAPI       = "ZipCodeLookup"
	cityStateLookupAPI = "CityStateLookup"
)

var baseURL = url.URL{
	Scheme: "http",
	Host:   "production.shippingapis.com",
	Path:   "ShippingAPI.dll",
}

const xmlHeader = `<?xml version="1.0"?>`

type USPSClient struct {
	UserID     string // Your USPS user ID
	HTTPClient http.Client
}

func New(uid string) *USPSClient {
	return &USPSClient{UserID: uid}
}

func (c USPSClient) fmtURL(api, xml string) string {
	// Create the URL
	u := baseURL
	u.RawQuery = url.Values{
		"API": []string{api},
		"XML": []string{xml},
	}.Encode()

	// Return the URL
	return u.String()
}

type Address struct {
	FirmName     string
	Address1     string
	Address2     string
	City         string
	State        string
	Urbanization string
	Zip5         string
	Zip4         string
}

func (a Address) withID(id int) address {
	return address{
		ID:           id,
		FirmName:     a.FirmName,
		Address1:     a.Address1,
		Address2:     a.Address2,
		City:         a.City,
		State:        a.State,
		Urbanization: a.Urbanization,
		Zip5:         a.Zip5,
		Zip4:         a.Zip4,
	}
}

type address struct {
	ID           int `xml:",attr"`
	FirmName     string
	Address1     string
	Address2     string
	City         string
	State        string
	Urbanization string
	Zip5         string
	Zip4         string
}

type ZipCodeLookupRequest struct{}

func (r ZipCodeLookupRequest) Validate() error {
	return nil
}

type ZipCodeLookupResponse struct{}

func (c *USPSClient) ZipLookup(req *ZipCodeLookupRequest) (*ZipCodeLookupResponse, error) {
	return nil, nil
}

type CityStateLookupRequest struct{}

func (r CityStateLookupRequest) Validate() error {
	return nil
}

type CityStateLookupResponse struct{}

func (c *USPSClient) CityStateLookup(req *CityStateLookupRequest) (*CityStateLookupResponse, error) {
	return nil, nil
}
