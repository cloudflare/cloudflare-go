// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountIntelWhoisService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountIntelWhoisService] method
// instead.
type AccountIntelWhoisService struct {
	Options []option.RequestOption
}

// NewAccountIntelWhoisService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountIntelWhoisService(opts ...option.RequestOption) (r *AccountIntelWhoisService) {
	r = &AccountIntelWhoisService{}
	r.Options = opts
	return
}

// Get WHOIS Record
func (r *AccountIntelWhoisService) WhoisRecordGetWhoisRecord(ctx context.Context, accountIdentifier string, query AccountIntelWhoisWhoisRecordGetWhoisRecordParams, opts ...option.RequestOption) (res *WhoisSingleResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/intel/whois", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type WhoisSingleResponse struct {
	Errors   []WhoisSingleResponseError   `json:"errors"`
	Messages []WhoisSingleResponseMessage `json:"messages"`
	Result   WhoisSingleResponseResult    `json:"result"`
	// Whether the API call was successful
	Success WhoisSingleResponseSuccess `json:"success"`
	JSON    whoisSingleResponseJSON    `json:"-"`
}

// whoisSingleResponseJSON contains the JSON metadata for the struct
// [WhoisSingleResponse]
type whoisSingleResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisSingleResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WhoisSingleResponseError struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    whoisSingleResponseErrorJSON `json:"-"`
}

// whoisSingleResponseErrorJSON contains the JSON metadata for the struct
// [WhoisSingleResponseError]
type whoisSingleResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisSingleResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WhoisSingleResponseMessage struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    whoisSingleResponseMessageJSON `json:"-"`
}

// whoisSingleResponseMessageJSON contains the JSON metadata for the struct
// [WhoisSingleResponseMessage]
type whoisSingleResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisSingleResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type WhoisSingleResponseResult struct {
	CreatedDate       time.Time                     `json:"created_date" format:"date"`
	Domain            string                        `json:"domain"`
	Nameservers       []string                      `json:"nameservers"`
	Registrant        string                        `json:"registrant"`
	RegistrantCountry string                        `json:"registrant_country"`
	RegistrantEmail   string                        `json:"registrant_email"`
	RegistrantOrg     string                        `json:"registrant_org"`
	Registrar         string                        `json:"registrar"`
	UpdatedDate       time.Time                     `json:"updated_date" format:"date"`
	JSON              whoisSingleResponseResultJSON `json:"-"`
}

// whoisSingleResponseResultJSON contains the JSON metadata for the struct
// [WhoisSingleResponseResult]
type whoisSingleResponseResultJSON struct {
	CreatedDate       apijson.Field
	Domain            apijson.Field
	Nameservers       apijson.Field
	Registrant        apijson.Field
	RegistrantCountry apijson.Field
	RegistrantEmail   apijson.Field
	RegistrantOrg     apijson.Field
	Registrar         apijson.Field
	UpdatedDate       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *WhoisSingleResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type WhoisSingleResponseSuccess bool

const (
	WhoisSingleResponseSuccessTrue WhoisSingleResponseSuccess = true
)

type AccountIntelWhoisWhoisRecordGetWhoisRecordParams struct {
	Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes [AccountIntelWhoisWhoisRecordGetWhoisRecordParams]'s query
// parameters as `url.Values`.
func (r AccountIntelWhoisWhoisRecordGetWhoisRecordParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
