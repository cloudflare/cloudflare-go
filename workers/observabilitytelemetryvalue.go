// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
	"github.com/tidwall/gjson"
)

// ObservabilityTelemetryValueService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityTelemetryValueService] method instead.
type ObservabilityTelemetryValueService struct {
	Options []option.RequestOption
}

// NewObservabilityTelemetryValueService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilityTelemetryValueService(opts ...option.RequestOption) (r *ObservabilityTelemetryValueService) {
	r = &ObservabilityTelemetryValueService{}
	r.Options = opts
	return
}

// List unique values found in your events
func (r *ObservabilityTelemetryValueService) New(ctx context.Context, params ObservabilityTelemetryValueNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityTelemetryValueNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/values", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPost, path, params, &res, opts...)
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

// List unique values found in your events
func (r *ObservabilityTelemetryValueService) NewAutoPaging(ctx context.Context, params ObservabilityTelemetryValueNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityTelemetryValueNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, params, opts...))
}

type ObservabilityTelemetryValueNewResponse struct {
	Dataset string                                           `json:"dataset,required"`
	Key     string                                           `json:"key,required"`
	Type    ObservabilityTelemetryValueNewResponseType       `json:"type,required"`
	Value   ObservabilityTelemetryValueNewResponseValueUnion `json:"value,required"`
	JSON    observabilityTelemetryValueNewResponseJSON       `json:"-"`
}

// observabilityTelemetryValueNewResponseJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryValueNewResponse]
type observabilityTelemetryValueNewResponseJSON struct {
	Dataset     apijson.Field
	Key         apijson.Field
	Type        apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryValueNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryValueNewResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryValueNewResponseType string

const (
	ObservabilityTelemetryValueNewResponseTypeString  ObservabilityTelemetryValueNewResponseType = "string"
	ObservabilityTelemetryValueNewResponseTypeBoolean ObservabilityTelemetryValueNewResponseType = "boolean"
	ObservabilityTelemetryValueNewResponseTypeNumber  ObservabilityTelemetryValueNewResponseType = "number"
)

func (r ObservabilityTelemetryValueNewResponseType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValueNewResponseTypeString, ObservabilityTelemetryValueNewResponseTypeBoolean, ObservabilityTelemetryValueNewResponseTypeNumber:
		return true
	}
	return false
}

// Union satisfied by [shared.UnionString], [shared.UnionFloat] or
// [shared.UnionBool].
type ObservabilityTelemetryValueNewResponseValueUnion interface {
	ImplementsObservabilityTelemetryValueNewResponseValueUnion()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ObservabilityTelemetryValueNewResponseValueUnion)(nil)).Elem(),
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
	)
}

type ObservabilityTelemetryValueNewParams struct {
	AccountID param.Field[string]                                        `path:"account_id,required"`
	Datasets  param.Field[[]string]                                      `json:"datasets,required"`
	Key       param.Field[string]                                        `json:"key,required"`
	Timeframe param.Field[ObservabilityTelemetryValueNewParamsTimeframe] `json:"timeframe,required"`
	Type      param.Field[ObservabilityTelemetryValueNewParamsType]      `json:"type,required"`
	Filters   param.Field[[]ObservabilityTelemetryValueNewParamsFilter]  `json:"filters"`
	Limit     param.Field[float64]                                       `json:"limit"`
	// Search for a specific substring in the event.
	Needle param.Field[ObservabilityTelemetryValueNewParamsNeedle] `json:"needle"`
}

func (r ObservabilityTelemetryValueNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValueNewParamsTimeframe struct {
	From param.Field[float64] `json:"from,required"`
	To   param.Field[float64] `json:"to,required"`
}

func (r ObservabilityTelemetryValueNewParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValueNewParamsType string

const (
	ObservabilityTelemetryValueNewParamsTypeString  ObservabilityTelemetryValueNewParamsType = "string"
	ObservabilityTelemetryValueNewParamsTypeBoolean ObservabilityTelemetryValueNewParamsType = "boolean"
	ObservabilityTelemetryValueNewParamsTypeNumber  ObservabilityTelemetryValueNewParamsType = "number"
)

func (r ObservabilityTelemetryValueNewParamsType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValueNewParamsTypeString, ObservabilityTelemetryValueNewParamsTypeBoolean, ObservabilityTelemetryValueNewParamsTypeNumber:
		return true
	}
	return false
}

type ObservabilityTelemetryValueNewParamsFilter struct {
	Key       param.Field[string]                                                `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryValueNewParamsFiltersOperation]  `json:"operation,required"`
	Type      param.Field[ObservabilityTelemetryValueNewParamsFiltersType]       `json:"type,required"`
	Value     param.Field[ObservabilityTelemetryValueNewParamsFiltersValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryValueNewParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryValueNewParamsFiltersOperation string

const (
	ObservabilityTelemetryValueNewParamsFiltersOperationIncludes            ObservabilityTelemetryValueNewParamsFiltersOperation = "includes"
	ObservabilityTelemetryValueNewParamsFiltersOperationNotIncludes         ObservabilityTelemetryValueNewParamsFiltersOperation = "not_includes"
	ObservabilityTelemetryValueNewParamsFiltersOperationStartsWith          ObservabilityTelemetryValueNewParamsFiltersOperation = "starts_with"
	ObservabilityTelemetryValueNewParamsFiltersOperationRegex               ObservabilityTelemetryValueNewParamsFiltersOperation = "regex"
	ObservabilityTelemetryValueNewParamsFiltersOperationExists              ObservabilityTelemetryValueNewParamsFiltersOperation = "exists"
	ObservabilityTelemetryValueNewParamsFiltersOperationIsNull              ObservabilityTelemetryValueNewParamsFiltersOperation = "is_null"
	ObservabilityTelemetryValueNewParamsFiltersOperationIn                  ObservabilityTelemetryValueNewParamsFiltersOperation = "in"
	ObservabilityTelemetryValueNewParamsFiltersOperationNotIn               ObservabilityTelemetryValueNewParamsFiltersOperation = "not_in"
	ObservabilityTelemetryValueNewParamsFiltersOperationEq                  ObservabilityTelemetryValueNewParamsFiltersOperation = "eq"
	ObservabilityTelemetryValueNewParamsFiltersOperationNeq                 ObservabilityTelemetryValueNewParamsFiltersOperation = "neq"
	ObservabilityTelemetryValueNewParamsFiltersOperationGt                  ObservabilityTelemetryValueNewParamsFiltersOperation = "gt"
	ObservabilityTelemetryValueNewParamsFiltersOperationGte                 ObservabilityTelemetryValueNewParamsFiltersOperation = "gte"
	ObservabilityTelemetryValueNewParamsFiltersOperationLt                  ObservabilityTelemetryValueNewParamsFiltersOperation = "lt"
	ObservabilityTelemetryValueNewParamsFiltersOperationLte                 ObservabilityTelemetryValueNewParamsFiltersOperation = "lte"
	ObservabilityTelemetryValueNewParamsFiltersOperationUnknown7            ObservabilityTelemetryValueNewParamsFiltersOperation = "="
	ObservabilityTelemetryValueNewParamsFiltersOperationNotEquals           ObservabilityTelemetryValueNewParamsFiltersOperation = "!="
	ObservabilityTelemetryValueNewParamsFiltersOperationGreater             ObservabilityTelemetryValueNewParamsFiltersOperation = ">"
	ObservabilityTelemetryValueNewParamsFiltersOperationGreaterOrEquals     ObservabilityTelemetryValueNewParamsFiltersOperation = ">="
	ObservabilityTelemetryValueNewParamsFiltersOperationLess                ObservabilityTelemetryValueNewParamsFiltersOperation = "<"
	ObservabilityTelemetryValueNewParamsFiltersOperationLessOrEquals        ObservabilityTelemetryValueNewParamsFiltersOperation = "<="
	ObservabilityTelemetryValueNewParamsFiltersOperationIncludesUppercase   ObservabilityTelemetryValueNewParamsFiltersOperation = "INCLUDES"
	ObservabilityTelemetryValueNewParamsFiltersOperationDoesNotInclude      ObservabilityTelemetryValueNewParamsFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryValueNewParamsFiltersOperationMatchRegex          ObservabilityTelemetryValueNewParamsFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryValueNewParamsFiltersOperationExistsUppercase     ObservabilityTelemetryValueNewParamsFiltersOperation = "EXISTS"
	ObservabilityTelemetryValueNewParamsFiltersOperationDoesNotExist        ObservabilityTelemetryValueNewParamsFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryValueNewParamsFiltersOperationInUppercase         ObservabilityTelemetryValueNewParamsFiltersOperation = "IN"
	ObservabilityTelemetryValueNewParamsFiltersOperationNotInUppercase      ObservabilityTelemetryValueNewParamsFiltersOperation = "NOT_IN"
	ObservabilityTelemetryValueNewParamsFiltersOperationStartsWithUppercase ObservabilityTelemetryValueNewParamsFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryValueNewParamsFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValueNewParamsFiltersOperationIncludes, ObservabilityTelemetryValueNewParamsFiltersOperationNotIncludes, ObservabilityTelemetryValueNewParamsFiltersOperationStartsWith, ObservabilityTelemetryValueNewParamsFiltersOperationRegex, ObservabilityTelemetryValueNewParamsFiltersOperationExists, ObservabilityTelemetryValueNewParamsFiltersOperationIsNull, ObservabilityTelemetryValueNewParamsFiltersOperationIn, ObservabilityTelemetryValueNewParamsFiltersOperationNotIn, ObservabilityTelemetryValueNewParamsFiltersOperationEq, ObservabilityTelemetryValueNewParamsFiltersOperationNeq, ObservabilityTelemetryValueNewParamsFiltersOperationGt, ObservabilityTelemetryValueNewParamsFiltersOperationGte, ObservabilityTelemetryValueNewParamsFiltersOperationLt, ObservabilityTelemetryValueNewParamsFiltersOperationLte, ObservabilityTelemetryValueNewParamsFiltersOperationUnknown7, ObservabilityTelemetryValueNewParamsFiltersOperationNotEquals, ObservabilityTelemetryValueNewParamsFiltersOperationGreater, ObservabilityTelemetryValueNewParamsFiltersOperationGreaterOrEquals, ObservabilityTelemetryValueNewParamsFiltersOperationLess, ObservabilityTelemetryValueNewParamsFiltersOperationLessOrEquals, ObservabilityTelemetryValueNewParamsFiltersOperationIncludesUppercase, ObservabilityTelemetryValueNewParamsFiltersOperationDoesNotInclude, ObservabilityTelemetryValueNewParamsFiltersOperationMatchRegex, ObservabilityTelemetryValueNewParamsFiltersOperationExistsUppercase, ObservabilityTelemetryValueNewParamsFiltersOperationDoesNotExist, ObservabilityTelemetryValueNewParamsFiltersOperationInUppercase, ObservabilityTelemetryValueNewParamsFiltersOperationNotInUppercase, ObservabilityTelemetryValueNewParamsFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryValueNewParamsFiltersType string

const (
	ObservabilityTelemetryValueNewParamsFiltersTypeString  ObservabilityTelemetryValueNewParamsFiltersType = "string"
	ObservabilityTelemetryValueNewParamsFiltersTypeNumber  ObservabilityTelemetryValueNewParamsFiltersType = "number"
	ObservabilityTelemetryValueNewParamsFiltersTypeBoolean ObservabilityTelemetryValueNewParamsFiltersType = "boolean"
)

func (r ObservabilityTelemetryValueNewParamsFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryValueNewParamsFiltersTypeString, ObservabilityTelemetryValueNewParamsFiltersTypeNumber, ObservabilityTelemetryValueNewParamsFiltersTypeBoolean:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryValueNewParamsFiltersValueUnion interface {
	ImplementsObservabilityTelemetryValueNewParamsFiltersValueUnion()
}

// Search for a specific substring in the event.
type ObservabilityTelemetryValueNewParamsNeedle struct {
	Value     param.Field[ObservabilityTelemetryValueNewParamsNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                                 `json:"isRegex"`
	MatchCase param.Field[bool]                                                 `json:"matchCase"`
}

func (r ObservabilityTelemetryValueNewParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryValueNewParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryValueNewParamsNeedleValueUnion()
}
