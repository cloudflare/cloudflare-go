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

// IntelWhoisService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewIntelWhoisService] method instead.
type IntelWhoisService struct {
	Options []option.RequestOption
}

// NewIntelWhoisService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntelWhoisService(opts ...option.RequestOption) (r *IntelWhoisService) {
	r = &IntelWhoisService{}
	r.Options = opts
	return
}

// Get WHOIS Record
func (r *IntelWhoisService) Get(ctx context.Context, params IntelWhoisGetParams, opts ...option.RequestOption) (res *IntelWhoisGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelWhoisGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/whois", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelWhoisGetResponse struct {
	CreatedDate       time.Time                 `json:"created_date" format:"date"`
	Domain            string                    `json:"domain"`
	Nameservers       []string                  `json:"nameservers"`
	Registrant        string                    `json:"registrant"`
	RegistrantCountry string                    `json:"registrant_country"`
	RegistrantEmail   string                    `json:"registrant_email"`
	RegistrantOrg     string                    `json:"registrant_org"`
	Registrar         string                    `json:"registrar"`
	UpdatedDate       time.Time                 `json:"updated_date" format:"date"`
	JSON              intelWhoisGetResponseJSON `json:"-"`
}

// intelWhoisGetResponseJSON contains the JSON metadata for the struct
// [IntelWhoisGetResponse]
type intelWhoisGetResponseJSON struct {
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

func (r *IntelWhoisGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelWhoisGetResponseJSON) RawJSON() string {
	return r.raw
}

type IntelWhoisGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	Domain    param.Field[string] `query:"domain"`
}

// URLQuery serializes [IntelWhoisGetParams]'s query parameters as `url.Values`.
func (r IntelWhoisGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelWhoisGetResponseEnvelope struct {
	Errors   []IntelWhoisGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelWhoisGetResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelWhoisGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelWhoisGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelWhoisGetResponseEnvelopeJSON    `json:"-"`
}

// intelWhoisGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [IntelWhoisGetResponseEnvelope]
type intelWhoisGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelWhoisGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelWhoisGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type IntelWhoisGetResponseEnvelopeErrors struct {
	Code    int64                                   `json:"code,required"`
	Message string                                  `json:"message,required"`
	JSON    intelWhoisGetResponseEnvelopeErrorsJSON `json:"-"`
}

// intelWhoisGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [IntelWhoisGetResponseEnvelopeErrors]
type intelWhoisGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelWhoisGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelWhoisGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type IntelWhoisGetResponseEnvelopeMessages struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    intelWhoisGetResponseEnvelopeMessagesJSON `json:"-"`
}

// intelWhoisGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [IntelWhoisGetResponseEnvelopeMessages]
type intelWhoisGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelWhoisGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r intelWhoisGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type IntelWhoisGetResponseEnvelopeSuccess bool

const (
	IntelWhoisGetResponseEnvelopeSuccessTrue IntelWhoisGetResponseEnvelopeSuccess = true
)
