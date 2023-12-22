// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountIntelASNService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelASNService] method
// instead.
type AccountIntelASNService struct {
	Options []option.RequestOption
	Subnets *AccountIntelASNSubnetService
}

// NewAccountIntelASNService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccountIntelASNService(opts ...option.RequestOption) (r *AccountIntelASNService) {
	r = &AccountIntelASNService{}
	r.Options = opts
	r.Subnets = NewAccountIntelASNSubnetService(opts...)
	return
}

// Get ASN Overview
func (r *AccountIntelASNService) Get(ctx context.Context, accountIdentifier string, asn interface{}, opts ...option.RequestOption) (res *ASNResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/asn/%v", accountIdentifier, asn)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ASNResponse struct {
	Errors   []ASNResponseError   `json:"errors"`
	Messages []ASNResponseMessage `json:"messages"`
	Result   ASNResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ASNResponseSuccess `json:"success"`
	JSON    asnResponseJSON    `json:"-"`
}

// asnResponseJSON contains the JSON metadata for the struct [ASNResponse]
type asnResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ASNResponseError struct {
	Code    int64                `json:"code,required"`
	Message string               `json:"message,required"`
	JSON    asnResponseErrorJSON `json:"-"`
}

// asnResponseErrorJSON contains the JSON metadata for the struct
// [ASNResponseError]
type asnResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ASNResponseMessage struct {
	Code    int64                  `json:"code,required"`
	Message string                 `json:"message,required"`
	JSON    asnResponseMessageJSON `json:"-"`
}

// asnResponseMessageJSON contains the JSON metadata for the struct
// [ASNResponseMessage]
type asnResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ASNResponseResult struct {
	ASN         int64    `json:"asn"`
	Country     string   `json:"country"`
	Description string   `json:"description"`
	DomainCount int64    `json:"domain_count"`
	TopDomains  []string `json:"top_domains"`
	// Infrastructure type of this ASN.
	Type ASNResponseResultType `json:"type"`
	JSON asnResponseResultJSON `json:"-"`
}

// asnResponseResultJSON contains the JSON metadata for the struct
// [ASNResponseResult]
type asnResponseResultJSON struct {
	ASN         apijson.Field
	Country     apijson.Field
	Description apijson.Field
	DomainCount apijson.Field
	TopDomains  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure type of this ASN.
type ASNResponseResultType string

const (
	ASNResponseResultTypeHostingProvider ASNResponseResultType = "hosting_provider"
	ASNResponseResultTypeIsp             ASNResponseResultType = "isp"
	ASNResponseResultTypeOrganization    ASNResponseResultType = "organization"
)

// Whether the API call was successful
type ASNResponseSuccess bool

const (
	ASNResponseSuccessTrue ASNResponseSuccess = true
)
