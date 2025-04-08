// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package botnet_feed

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ASNService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewASNService] method instead.
type ASNService struct {
	Options []option.RequestOption
}

// NewASNService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewASNService(opts ...option.RequestOption) (r *ASNService) {
	r = &ASNService{}
	r.Options = opts
	return
}

// Gets all the data the botnet tracking database has for a given ASN registered to
// user account for given date. If no date is given, it will return results for the
// previous day.
func (r *ASNService) DayReport(ctx context.Context, asnID int64, params ASNDayReportParams, opts ...option.RequestOption) (res *ASNDayReportResponse, err error) {
	var env ASNDayReportResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/asn/%v/day_report", params.AccountID, asnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets all the data the botnet threat feed tracking database has for a given ASN
// registered to user account.
func (r *ASNService) FullReport(ctx context.Context, asnID int64, query ASNFullReportParams, opts ...option.RequestOption) (res *ASNFullReportResponse, err error) {
	var env ASNFullReportResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/botnet_feed/asn/%v/full_report", query.AccountID, asnID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ASNDayReportResponse struct {
	CIDR         string                   `json:"cidr"`
	Date         time.Time                `json:"date" format:"date-time"`
	OffenseCount int64                    `json:"offense_count"`
	JSON         asnDayReportResponseJSON `json:"-"`
}

// asnDayReportResponseJSON contains the JSON metadata for the struct
// [ASNDayReportResponse]
type asnDayReportResponseJSON struct {
	CIDR         apijson.Field
	Date         apijson.Field
	OffenseCount apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ASNDayReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnDayReportResponseJSON) RawJSON() string {
	return r.raw
}

type ASNFullReportResponse struct {
	CIDR         string                    `json:"cidr"`
	Date         time.Time                 `json:"date" format:"date-time"`
	OffenseCount int64                     `json:"offense_count"`
	JSON         asnFullReportResponseJSON `json:"-"`
}

// asnFullReportResponseJSON contains the JSON metadata for the struct
// [ASNFullReportResponse]
type asnFullReportResponseJSON struct {
	CIDR         apijson.Field
	Date         apijson.Field
	OffenseCount apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ASNFullReportResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnFullReportResponseJSON) RawJSON() string {
	return r.raw
}

type ASNDayReportParams struct {
	// Identifier
	AccountID param.Field[string]    `path:"account_id,required"`
	Date      param.Field[time.Time] `query:"date" format:"date-time"`
}

// URLQuery serializes [ASNDayReportParams]'s query parameters as `url.Values`.
func (r ASNDayReportParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ASNDayReportResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ASNDayReportResponseEnvelopeSuccess `json:"success,required"`
	Result  ASNDayReportResponse                `json:"result"`
	JSON    asnDayReportResponseEnvelopeJSON    `json:"-"`
}

// asnDayReportResponseEnvelopeJSON contains the JSON metadata for the struct
// [ASNDayReportResponseEnvelope]
type asnDayReportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNDayReportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnDayReportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ASNDayReportResponseEnvelopeSuccess bool

const (
	ASNDayReportResponseEnvelopeSuccessTrue ASNDayReportResponseEnvelopeSuccess = true
)

func (r ASNDayReportResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ASNDayReportResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ASNFullReportParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ASNFullReportResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ASNFullReportResponseEnvelopeSuccess `json:"success,required"`
	Result  ASNFullReportResponse                `json:"result"`
	JSON    asnFullReportResponseEnvelopeJSON    `json:"-"`
}

// asnFullReportResponseEnvelopeJSON contains the JSON metadata for the struct
// [ASNFullReportResponseEnvelope]
type asnFullReportResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ASNFullReportResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r asnFullReportResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ASNFullReportResponseEnvelopeSuccess bool

const (
	ASNFullReportResponseEnvelopeSuccessTrue ASNFullReportResponseEnvelopeSuccess = true
)

func (r ASNFullReportResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ASNFullReportResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
