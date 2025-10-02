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
)

// AvailableAlertService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAvailableAlertService] method instead.
type AvailableAlertService struct {
	Options []option.RequestOption
}

// NewAvailableAlertService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAvailableAlertService(opts ...option.RequestOption) (r *AvailableAlertService) {
	r = &AvailableAlertService{}
	r.Options = opts
	return
}

// Gets a list of all alert types for which an account is eligible.
func (r *AvailableAlertService) List(ctx context.Context, query AvailableAlertListParams, opts ...option.RequestOption) (res *AvailableAlertListResponse, err error) {
	var env AvailableAlertListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/available_alerts", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AvailableAlertListResponse map[string][]AvailableAlertListResponseItem

type AvailableAlertListResponseItem struct {
	// Describes the alert type.
	Description string `json:"description"`
	// Alert type name.
	DisplayName string `json:"display_name"`
	// Format of additional configuration options (filters) for the alert type. Data
	// type of filters during policy creation: Array of strings.
	FilterOptions []interface{} `json:"filter_options"`
	// Use this value when creating and updating a notification policy.
	Type string                             `json:"type"`
	JSON availableAlertListResponseItemJSON `json:"-"`
}

// availableAlertListResponseItemJSON contains the JSON metadata for the struct
// [AvailableAlertListResponseItem]
type availableAlertListResponseItemJSON struct {
	Description   apijson.Field
	DisplayName   apijson.Field
	FilterOptions apijson.Field
	Type          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AvailableAlertListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableAlertListResponseItemJSON) RawJSON() string {
	return r.raw
}

type AvailableAlertListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type AvailableAlertListResponseEnvelope struct {
	Errors   []AvailableAlertListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AvailableAlertListResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success AvailableAlertListResponseEnvelopeSuccess `json:"success,required"`
	Result  AvailableAlertListResponse                `json:"result"`
	JSON    availableAlertListResponseEnvelopeJSON    `json:"-"`
}

// availableAlertListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AvailableAlertListResponseEnvelope]
type availableAlertListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableAlertListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableAlertListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AvailableAlertListResponseEnvelopeErrors struct {
	Message string                                       `json:"message,required"`
	Code    int64                                        `json:"code"`
	JSON    availableAlertListResponseEnvelopeErrorsJSON `json:"-"`
}

// availableAlertListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AvailableAlertListResponseEnvelopeErrors]
type availableAlertListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableAlertListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableAlertListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AvailableAlertListResponseEnvelopeMessages struct {
	Message string                                         `json:"message,required"`
	Code    int64                                          `json:"code"`
	JSON    availableAlertListResponseEnvelopeMessagesJSON `json:"-"`
}

// availableAlertListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AvailableAlertListResponseEnvelopeMessages]
type availableAlertListResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AvailableAlertListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r availableAlertListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AvailableAlertListResponseEnvelopeSuccess bool

const (
	AvailableAlertListResponseEnvelopeSuccessTrue AvailableAlertListResponseEnvelopeSuccess = true
)

func (r AvailableAlertListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AvailableAlertListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
