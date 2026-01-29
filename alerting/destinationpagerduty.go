// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package alerting

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// DestinationPagerdutyService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDestinationPagerdutyService] method instead.
type DestinationPagerdutyService struct {
	Options []option.RequestOption
}

// NewDestinationPagerdutyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDestinationPagerdutyService(opts ...option.RequestOption) (r *DestinationPagerdutyService) {
	r = &DestinationPagerdutyService{}
	r.Options = opts
	return
}

// Creates a new token for integrating with PagerDuty.
func (r *DestinationPagerdutyService) New(ctx context.Context, body DestinationPagerdutyNewParams, opts ...option.RequestOption) (res *DestinationPagerdutyNewResponse, err error) {
	var env DestinationPagerdutyNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes all the PagerDuty Services connected to the account.
func (r *DestinationPagerdutyService) Delete(ctx context.Context, body DestinationPagerdutyDeleteParams, opts ...option.RequestOption) (res *DestinationPagerdutyDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", body.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get a list of all configured PagerDuty services.
func (r *DestinationPagerdutyService) Get(ctx context.Context, query DestinationPagerdutyGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[Pagerduty], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a list of all configured PagerDuty services.
func (r *DestinationPagerdutyService) GetAutoPaging(ctx context.Context, query DestinationPagerdutyGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[Pagerduty] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, query, opts...))
}

// Links PagerDuty with the account using the integration token.
func (r *DestinationPagerdutyService) Link(ctx context.Context, tokenID string, query DestinationPagerdutyLinkParams, opts ...option.RequestOption) (res *DestinationPagerdutyLinkResponse, err error) {
	var env DestinationPagerdutyLinkResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if tokenID == "" {
		err = errors.New("missing required token_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/destinations/pagerduty/connect/%s", query.AccountID, tokenID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Pagerduty struct {
	// UUID
	ID string `json:"id"`
	// The name of the pagerduty service.
	Name string        `json:"name"`
	JSON pagerdutyJSON `json:"-"`
}

// pagerdutyJSON contains the JSON metadata for the struct [Pagerduty]
type pagerdutyJSON struct {
	ID          apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Pagerduty) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r pagerdutyJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewResponse struct {
	// token in form of UUID
	ID   string                              `json:"id"`
	JSON destinationPagerdutyNewResponseJSON `json:"-"`
}

// destinationPagerdutyNewResponseJSON contains the JSON metadata for the struct
// [DestinationPagerdutyNewResponse]
type destinationPagerdutyNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyDeleteResponse struct {
	Errors   []DestinationPagerdutyDeleteResponseError   `json:"errors,required"`
	Messages []DestinationPagerdutyDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success DestinationPagerdutyDeleteResponseSuccess `json:"success,required"`
	JSON    destinationPagerdutyDeleteResponseJSON    `json:"-"`
}

// destinationPagerdutyDeleteResponseJSON contains the JSON metadata for the struct
// [DestinationPagerdutyDeleteResponse]
type destinationPagerdutyDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyDeleteResponseError struct {
	Message string                                      `json:"message,required"`
	Code    int64                                       `json:"code"`
	JSON    destinationPagerdutyDeleteResponseErrorJSON `json:"-"`
}

// destinationPagerdutyDeleteResponseErrorJSON contains the JSON metadata for the
// struct [DestinationPagerdutyDeleteResponseError]
type destinationPagerdutyDeleteResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyDeleteResponseMessage struct {
	Message string                                        `json:"message,required"`
	Code    int64                                         `json:"code"`
	JSON    destinationPagerdutyDeleteResponseMessageJSON `json:"-"`
}

// destinationPagerdutyDeleteResponseMessageJSON contains the JSON metadata for the
// struct [DestinationPagerdutyDeleteResponseMessage]
type destinationPagerdutyDeleteResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationPagerdutyDeleteResponseSuccess bool

const (
	DestinationPagerdutyDeleteResponseSuccessTrue DestinationPagerdutyDeleteResponseSuccess = true
)

func (r DestinationPagerdutyDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case DestinationPagerdutyDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type DestinationPagerdutyLinkResponse struct {
	// UUID
	ID   string                               `json:"id"`
	JSON destinationPagerdutyLinkResponseJSON `json:"-"`
}

// destinationPagerdutyLinkResponseJSON contains the JSON metadata for the struct
// [DestinationPagerdutyLinkResponse]
type destinationPagerdutyLinkResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyNewResponseEnvelope struct {
	Errors   []DestinationPagerdutyNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationPagerdutyNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success DestinationPagerdutyNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DestinationPagerdutyNewResponse                `json:"result"`
	JSON    destinationPagerdutyNewResponseEnvelopeJSON    `json:"-"`
}

// destinationPagerdutyNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationPagerdutyNewResponseEnvelope]
type destinationPagerdutyNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewResponseEnvelopeErrors struct {
	Message string                                            `json:"message,required"`
	Code    int64                                             `json:"code"`
	JSON    destinationPagerdutyNewResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationPagerdutyNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DestinationPagerdutyNewResponseEnvelopeErrors]
type destinationPagerdutyNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyNewResponseEnvelopeMessages struct {
	Message string                                              `json:"message,required"`
	Code    int64                                               `json:"code"`
	JSON    destinationPagerdutyNewResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationPagerdutyNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationPagerdutyNewResponseEnvelopeMessages]
type destinationPagerdutyNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationPagerdutyNewResponseEnvelopeSuccess bool

const (
	DestinationPagerdutyNewResponseEnvelopeSuccessTrue DestinationPagerdutyNewResponseEnvelopeSuccess = true
)

func (r DestinationPagerdutyNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationPagerdutyNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DestinationPagerdutyDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyLinkParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type DestinationPagerdutyLinkResponseEnvelope struct {
	Errors   []DestinationPagerdutyLinkResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DestinationPagerdutyLinkResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success DestinationPagerdutyLinkResponseEnvelopeSuccess `json:"success,required"`
	Result  DestinationPagerdutyLinkResponse                `json:"result"`
	JSON    destinationPagerdutyLinkResponseEnvelopeJSON    `json:"-"`
}

// destinationPagerdutyLinkResponseEnvelopeJSON contains the JSON metadata for the
// struct [DestinationPagerdutyLinkResponseEnvelope]
type destinationPagerdutyLinkResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyLinkResponseEnvelopeErrors struct {
	Message string                                             `json:"message,required"`
	Code    int64                                              `json:"code"`
	JSON    destinationPagerdutyLinkResponseEnvelopeErrorsJSON `json:"-"`
}

// destinationPagerdutyLinkResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DestinationPagerdutyLinkResponseEnvelopeErrors]
type destinationPagerdutyLinkResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DestinationPagerdutyLinkResponseEnvelopeMessages struct {
	Message string                                               `json:"message,required"`
	Code    int64                                                `json:"code"`
	JSON    destinationPagerdutyLinkResponseEnvelopeMessagesJSON `json:"-"`
}

// destinationPagerdutyLinkResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DestinationPagerdutyLinkResponseEnvelopeMessages]
type destinationPagerdutyLinkResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DestinationPagerdutyLinkResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r destinationPagerdutyLinkResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type DestinationPagerdutyLinkResponseEnvelopeSuccess bool

const (
	DestinationPagerdutyLinkResponseEnvelopeSuccessTrue DestinationPagerdutyLinkResponseEnvelopeSuccess = true
)

func (r DestinationPagerdutyLinkResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DestinationPagerdutyLinkResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
