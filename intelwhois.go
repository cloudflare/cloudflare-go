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
func (r *IntelWhoisService) WhoisRecordGetWhoisRecord(ctx context.Context, accountID string, query IntelWhoisWhoisRecordGetWhoisRecordParams, opts ...option.RequestOption) (res *IntelWhoisWhoisRecordGetWhoisRecordResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/whois", accountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type IntelWhoisWhoisRecordGetWhoisRecordResponse struct {
	CreatedDate       time.Time                                       `json:"created_date" format:"date"`
	Domain            string                                          `json:"domain"`
	Nameservers       []string                                        `json:"nameservers"`
	Registrant        string                                          `json:"registrant"`
	RegistrantCountry string                                          `json:"registrant_country"`
	RegistrantEmail   string                                          `json:"registrant_email"`
	RegistrantOrg     string                                          `json:"registrant_org"`
	Registrar         string                                          `json:"registrar"`
	UpdatedDate       time.Time                                       `json:"updated_date" format:"date"`
	JSON              intelWhoisWhoisRecordGetWhoisRecordResponseJSON `json:"-"`
}

// intelWhoisWhoisRecordGetWhoisRecordResponseJSON contains the JSON metadata for
// the struct [IntelWhoisWhoisRecordGetWhoisRecordResponse]
type intelWhoisWhoisRecordGetWhoisRecordResponseJSON struct {
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

func (r *IntelWhoisWhoisRecordGetWhoisRecordResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelWhoisWhoisRecordGetWhoisRecordParams struct {
	Domain param.Field[string] `query:"domain"`
}

// URLQuery serializes [IntelWhoisWhoisRecordGetWhoisRecordParams]'s query
// parameters as `url.Values`.
func (r IntelWhoisWhoisRecordGetWhoisRecordParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelope struct {
	Errors   []IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeErrors   `json:"errors,required"`
	Messages []IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeMessages `json:"messages,required"`
	Result   IntelWhoisWhoisRecordGetWhoisRecordResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeSuccess `json:"success,required"`
	JSON    intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeJSON    `json:"-"`
}

// intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeJSON contains the JSON
// metadata for the struct [IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelope]
type intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeErrors struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeErrorsJSON `json:"-"`
}

// intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeErrors]
type intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeMessages struct {
	Code    int64                                                           `json:"code,required"`
	Message string                                                          `json:"message,required"`
	JSON    intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeMessagesJSON `json:"-"`
}

// intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeMessages]
type intelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeSuccess bool

const (
	IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeSuccessTrue IntelWhoisWhoisRecordGetWhoisRecordResponseEnvelopeSuccess = true
)
