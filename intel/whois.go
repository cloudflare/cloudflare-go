// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// WhoisService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewWhoisService] method instead.
type WhoisService struct {
	Options []option.RequestOption
}

// NewWhoisService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWhoisService(opts ...option.RequestOption) (r *WhoisService) {
	r = &WhoisService{}
	r.Options = opts
	return
}

// Get WHOIS Record
func (r *WhoisService) Get(ctx context.Context, params WhoisGetParams, opts ...option.RequestOption) (res *IntelWhois, err error) {
	opts = append(r.Options[:], opts...)
	var env WhoisGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/whois", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelWhois struct {
	CreatedDate       time.Time      `json:"created_date" format:"date"`
	Domain            string         `json:"domain"`
	Nameservers       []string       `json:"nameservers"`
	Registrant        string         `json:"registrant"`
	RegistrantCountry string         `json:"registrant_country"`
	RegistrantEmail   string         `json:"registrant_email"`
	RegistrantOrg     string         `json:"registrant_org"`
	Registrar         string         `json:"registrar"`
	UpdatedDate       time.Time      `json:"updated_date" format:"date"`
	JSON              intelWhoisJSON `json:"-"`
}

// intelWhoisJSON contains the JSON metadata for the struct [IntelWhois]
type intelWhoisJSON struct {
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

func (r *IntelWhois) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelWhoisJSON) RawJSON() string {
	return r.raw
}

type WhoisGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Domain    param.Field[string] `query:"domain"`
}

// URLQuery serializes [WhoisGetParams]'s query parameters as `url.Values`.
func (r WhoisGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WhoisGetResponseEnvelope struct {
	Errors   []WhoisGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []WhoisGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelWhois                         `json:"result,required"`
	// Whether the API call was successful
	Success WhoisGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    whoisGetResponseEnvelopeJSON    `json:"-"`
}

// whoisGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [WhoisGetResponseEnvelope]
type whoisGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type WhoisGetResponseEnvelopeErrors struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    whoisGetResponseEnvelopeErrorsJSON `json:"-"`
}

// whoisGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [WhoisGetResponseEnvelopeErrors]
type whoisGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type WhoisGetResponseEnvelopeMessages struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    whoisGetResponseEnvelopeMessagesJSON `json:"-"`
}

// whoisGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [WhoisGetResponseEnvelopeMessages]
type whoisGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WhoisGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r whoisGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type WhoisGetResponseEnvelopeSuccess bool

const (
	WhoisGetResponseEnvelopeSuccessTrue WhoisGetResponseEnvelopeSuccess = true
)

func (r WhoisGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case WhoisGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
