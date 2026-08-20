// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package flagship

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/shared"
	"github.com/tidwall/gjson"
)

// AppEvaluateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAppEvaluateService] method instead.
type AppEvaluateService struct {
	Options []option.RequestOption
}

// NewAppEvaluateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAppEvaluateService(opts ...option.RequestOption) (r *AppEvaluateService) {
	r = &AppEvaluateService{}
	r.Options = opts
	return
}

// Evaluates a flag against the provided context. Pass context attributes as query
// parameters; boolean and numeric strings are coerced automatically. For
// low-latency in-Worker evaluation, prefer the Flagship binding over this
// endpoint.
func (r *AppEvaluateService) Get(ctx context.Context, appID string, params AppEvaluateGetParams, opts ...option.RequestOption) (res *AppEvaluateGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/flagship/apps/%s/evaluate", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type AppEvaluateGetResponse struct {
	FlagKey string                           `json:"flagKey" api:"required"`
	Reason  AppEvaluateGetResponseReason     `json:"reason" api:"required"`
	Variant string                           `json:"variant" api:"required"`
	Value   AppEvaluateGetResponseValueUnion `json:"value"`
	JSON    appEvaluateGetResponseJSON       `json:"-"`
}

// appEvaluateGetResponseJSON contains the JSON metadata for the struct
// [AppEvaluateGetResponse]
type appEvaluateGetResponseJSON struct {
	FlagKey     apijson.Field
	Reason      apijson.Field
	Variant     apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AppEvaluateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r appEvaluateGetResponseJSON) RawJSON() string {
	return r.raw
}

type AppEvaluateGetResponseReason string

const (
	AppEvaluateGetResponseReasonTargetingMatch AppEvaluateGetResponseReason = "TARGETING_MATCH"
	AppEvaluateGetResponseReasonDefault        AppEvaluateGetResponseReason = "DEFAULT"
	AppEvaluateGetResponseReasonDisabled       AppEvaluateGetResponseReason = "DISABLED"
	AppEvaluateGetResponseReasonSplit          AppEvaluateGetResponseReason = "SPLIT"
)

func (r AppEvaluateGetResponseReason) IsKnown() bool {
	switch r {
	case AppEvaluateGetResponseReasonTargetingMatch, AppEvaluateGetResponseReasonDefault, AppEvaluateGetResponseReasonDisabled, AppEvaluateGetResponseReasonSplit:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat],
// [shared.UnionBool], [AppEvaluateGetResponseValueMap] or
// [AppEvaluateGetResponseValueArray].
type AppEvaluateGetResponseValueUnion interface {
	ImplementsAppEvaluateGetResponseValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AppEvaluateGetResponseValueUnion)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.Number,
			Type:       reflect.TypeOf(shared.UnionFloat(0)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.True,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.False,
			Type:       reflect.TypeOf(shared.UnionBool(false)),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppEvaluateGetResponseValueMap{}),
		},
		apijson.UnionVariant{
			TypeFilter: gjson.JSON,
			Type:       reflect.TypeOf(AppEvaluateGetResponseValueArray{}),
		},
	)
}

type AppEvaluateGetResponseValueMap map[string]interface{}

func (r AppEvaluateGetResponseValueMap) ImplementsAppEvaluateGetResponseValueUnion() {}

type AppEvaluateGetResponseValueArray []interface{}

func (r AppEvaluateGetResponseValueArray) ImplementsAppEvaluateGetResponseValueUnion() {}

type AppEvaluateGetParams struct {
	// Cloudflare account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// The flag key to evaluate.
	FlagKey param.Field[string] `query:"flagKey" api:"required"`
	// Context targeting key (per OpenFeature spec); used for percentage rollout
	// bucketing.
	TargetingKey param.Field[string] `query:"targetingKey"`
}

// URLQuery serializes [AppEvaluateGetParams]'s query parameters as `url.Values`.
func (r AppEvaluateGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
