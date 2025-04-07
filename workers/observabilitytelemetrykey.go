// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
)

// ObservabilityTelemetryKeyService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityTelemetryKeyService] method instead.
type ObservabilityTelemetryKeyService struct {
	Options []option.RequestOption
}

// NewObservabilityTelemetryKeyService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilityTelemetryKeyService(opts ...option.RequestOption) (r *ObservabilityTelemetryKeyService) {
	r = &ObservabilityTelemetryKeyService{}
	r.Options = opts
	return
}

// List all the keys in your telemetry events.
func (r *ObservabilityTelemetryKeyService) New(ctx context.Context, params ObservabilityTelemetryKeyNewParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityTelemetryKeyNewResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/telemetry/keys", params.AccountID)
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

// List all the keys in your telemetry events.
func (r *ObservabilityTelemetryKeyService) NewAutoPaging(ctx context.Context, params ObservabilityTelemetryKeyNewParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityTelemetryKeyNewResponse] {
	return pagination.NewSinglePageAutoPager(r.New(ctx, params, opts...))
}

type ObservabilityTelemetryKeyNewResponse struct {
	Key        string                                   `json:"key,required"`
	LastSeenAt float64                                  `json:"lastSeenAt,required"`
	Type       ObservabilityTelemetryKeyNewResponseType `json:"type,required"`
	JSON       observabilityTelemetryKeyNewResponseJSON `json:"-"`
}

// observabilityTelemetryKeyNewResponseJSON contains the JSON metadata for the
// struct [ObservabilityTelemetryKeyNewResponse]
type observabilityTelemetryKeyNewResponseJSON struct {
	Key         apijson.Field
	LastSeenAt  apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityTelemetryKeyNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityTelemetryKeyNewResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityTelemetryKeyNewResponseType string

const (
	ObservabilityTelemetryKeyNewResponseTypeString  ObservabilityTelemetryKeyNewResponseType = "string"
	ObservabilityTelemetryKeyNewResponseTypeBoolean ObservabilityTelemetryKeyNewResponseType = "boolean"
	ObservabilityTelemetryKeyNewResponseTypeNumber  ObservabilityTelemetryKeyNewResponseType = "number"
)

func (r ObservabilityTelemetryKeyNewResponseType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeyNewResponseTypeString, ObservabilityTelemetryKeyNewResponseTypeBoolean, ObservabilityTelemetryKeyNewResponseTypeNumber:
		return true
	}
	return false
}

type ObservabilityTelemetryKeyNewParams struct {
	AccountID param.Field[string]                                     `path:"account_id,required"`
	Datasets  param.Field[[]string]                                   `json:"datasets"`
	Filters   param.Field[[]ObservabilityTelemetryKeyNewParamsFilter] `json:"filters"`
	// Search for a specific substring in the keys.
	KeyNeedle param.Field[ObservabilityTelemetryKeyNewParamsKeyNeedle] `json:"keyNeedle"`
	Limit     param.Field[float64]                                     `json:"limit"`
	// Search for a specific substring in the event.
	Needle    param.Field[ObservabilityTelemetryKeyNewParamsNeedle]    `json:"needle"`
	Timeframe param.Field[ObservabilityTelemetryKeyNewParamsTimeframe] `json:"timeframe"`
}

func (r ObservabilityTelemetryKeyNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryKeyNewParamsFilter struct {
	Key       param.Field[string]                                              `json:"key,required"`
	Operation param.Field[ObservabilityTelemetryKeyNewParamsFiltersOperation]  `json:"operation,required"`
	Type      param.Field[ObservabilityTelemetryKeyNewParamsFiltersType]       `json:"type,required"`
	Value     param.Field[ObservabilityTelemetryKeyNewParamsFiltersValueUnion] `json:"value"`
}

func (r ObservabilityTelemetryKeyNewParamsFilter) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityTelemetryKeyNewParamsFiltersOperation string

const (
	ObservabilityTelemetryKeyNewParamsFiltersOperationIncludes            ObservabilityTelemetryKeyNewParamsFiltersOperation = "includes"
	ObservabilityTelemetryKeyNewParamsFiltersOperationNotIncludes         ObservabilityTelemetryKeyNewParamsFiltersOperation = "not_includes"
	ObservabilityTelemetryKeyNewParamsFiltersOperationStartsWith          ObservabilityTelemetryKeyNewParamsFiltersOperation = "starts_with"
	ObservabilityTelemetryKeyNewParamsFiltersOperationRegex               ObservabilityTelemetryKeyNewParamsFiltersOperation = "regex"
	ObservabilityTelemetryKeyNewParamsFiltersOperationExists              ObservabilityTelemetryKeyNewParamsFiltersOperation = "exists"
	ObservabilityTelemetryKeyNewParamsFiltersOperationIsNull              ObservabilityTelemetryKeyNewParamsFiltersOperation = "is_null"
	ObservabilityTelemetryKeyNewParamsFiltersOperationIn                  ObservabilityTelemetryKeyNewParamsFiltersOperation = "in"
	ObservabilityTelemetryKeyNewParamsFiltersOperationNotIn               ObservabilityTelemetryKeyNewParamsFiltersOperation = "not_in"
	ObservabilityTelemetryKeyNewParamsFiltersOperationEq                  ObservabilityTelemetryKeyNewParamsFiltersOperation = "eq"
	ObservabilityTelemetryKeyNewParamsFiltersOperationNeq                 ObservabilityTelemetryKeyNewParamsFiltersOperation = "neq"
	ObservabilityTelemetryKeyNewParamsFiltersOperationGt                  ObservabilityTelemetryKeyNewParamsFiltersOperation = "gt"
	ObservabilityTelemetryKeyNewParamsFiltersOperationGte                 ObservabilityTelemetryKeyNewParamsFiltersOperation = "gte"
	ObservabilityTelemetryKeyNewParamsFiltersOperationLt                  ObservabilityTelemetryKeyNewParamsFiltersOperation = "lt"
	ObservabilityTelemetryKeyNewParamsFiltersOperationLte                 ObservabilityTelemetryKeyNewParamsFiltersOperation = "lte"
	ObservabilityTelemetryKeyNewParamsFiltersOperationUnknown1            ObservabilityTelemetryKeyNewParamsFiltersOperation = "="
	ObservabilityTelemetryKeyNewParamsFiltersOperationNotEquals           ObservabilityTelemetryKeyNewParamsFiltersOperation = "!="
	ObservabilityTelemetryKeyNewParamsFiltersOperationGreater             ObservabilityTelemetryKeyNewParamsFiltersOperation = ">"
	ObservabilityTelemetryKeyNewParamsFiltersOperationGreaterOrEquals     ObservabilityTelemetryKeyNewParamsFiltersOperation = ">="
	ObservabilityTelemetryKeyNewParamsFiltersOperationLess                ObservabilityTelemetryKeyNewParamsFiltersOperation = "<"
	ObservabilityTelemetryKeyNewParamsFiltersOperationLessOrEquals        ObservabilityTelemetryKeyNewParamsFiltersOperation = "<="
	ObservabilityTelemetryKeyNewParamsFiltersOperationIncludesUppercase   ObservabilityTelemetryKeyNewParamsFiltersOperation = "INCLUDES"
	ObservabilityTelemetryKeyNewParamsFiltersOperationDoesNotInclude      ObservabilityTelemetryKeyNewParamsFiltersOperation = "DOES_NOT_INCLUDE"
	ObservabilityTelemetryKeyNewParamsFiltersOperationMatchRegex          ObservabilityTelemetryKeyNewParamsFiltersOperation = "MATCH_REGEX"
	ObservabilityTelemetryKeyNewParamsFiltersOperationExistsUppercase     ObservabilityTelemetryKeyNewParamsFiltersOperation = "EXISTS"
	ObservabilityTelemetryKeyNewParamsFiltersOperationDoesNotExist        ObservabilityTelemetryKeyNewParamsFiltersOperation = "DOES_NOT_EXIST"
	ObservabilityTelemetryKeyNewParamsFiltersOperationInUppercase         ObservabilityTelemetryKeyNewParamsFiltersOperation = "IN"
	ObservabilityTelemetryKeyNewParamsFiltersOperationNotInUppercase      ObservabilityTelemetryKeyNewParamsFiltersOperation = "NOT_IN"
	ObservabilityTelemetryKeyNewParamsFiltersOperationStartsWithUppercase ObservabilityTelemetryKeyNewParamsFiltersOperation = "STARTS_WITH"
)

func (r ObservabilityTelemetryKeyNewParamsFiltersOperation) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeyNewParamsFiltersOperationIncludes, ObservabilityTelemetryKeyNewParamsFiltersOperationNotIncludes, ObservabilityTelemetryKeyNewParamsFiltersOperationStartsWith, ObservabilityTelemetryKeyNewParamsFiltersOperationRegex, ObservabilityTelemetryKeyNewParamsFiltersOperationExists, ObservabilityTelemetryKeyNewParamsFiltersOperationIsNull, ObservabilityTelemetryKeyNewParamsFiltersOperationIn, ObservabilityTelemetryKeyNewParamsFiltersOperationNotIn, ObservabilityTelemetryKeyNewParamsFiltersOperationEq, ObservabilityTelemetryKeyNewParamsFiltersOperationNeq, ObservabilityTelemetryKeyNewParamsFiltersOperationGt, ObservabilityTelemetryKeyNewParamsFiltersOperationGte, ObservabilityTelemetryKeyNewParamsFiltersOperationLt, ObservabilityTelemetryKeyNewParamsFiltersOperationLte, ObservabilityTelemetryKeyNewParamsFiltersOperationUnknown1, ObservabilityTelemetryKeyNewParamsFiltersOperationNotEquals, ObservabilityTelemetryKeyNewParamsFiltersOperationGreater, ObservabilityTelemetryKeyNewParamsFiltersOperationGreaterOrEquals, ObservabilityTelemetryKeyNewParamsFiltersOperationLess, ObservabilityTelemetryKeyNewParamsFiltersOperationLessOrEquals, ObservabilityTelemetryKeyNewParamsFiltersOperationIncludesUppercase, ObservabilityTelemetryKeyNewParamsFiltersOperationDoesNotInclude, ObservabilityTelemetryKeyNewParamsFiltersOperationMatchRegex, ObservabilityTelemetryKeyNewParamsFiltersOperationExistsUppercase, ObservabilityTelemetryKeyNewParamsFiltersOperationDoesNotExist, ObservabilityTelemetryKeyNewParamsFiltersOperationInUppercase, ObservabilityTelemetryKeyNewParamsFiltersOperationNotInUppercase, ObservabilityTelemetryKeyNewParamsFiltersOperationStartsWithUppercase:
		return true
	}
	return false
}

type ObservabilityTelemetryKeyNewParamsFiltersType string

const (
	ObservabilityTelemetryKeyNewParamsFiltersTypeString  ObservabilityTelemetryKeyNewParamsFiltersType = "string"
	ObservabilityTelemetryKeyNewParamsFiltersTypeNumber  ObservabilityTelemetryKeyNewParamsFiltersType = "number"
	ObservabilityTelemetryKeyNewParamsFiltersTypeBoolean ObservabilityTelemetryKeyNewParamsFiltersType = "boolean"
)

func (r ObservabilityTelemetryKeyNewParamsFiltersType) IsKnown() bool {
	switch r {
	case ObservabilityTelemetryKeyNewParamsFiltersTypeString, ObservabilityTelemetryKeyNewParamsFiltersTypeNumber, ObservabilityTelemetryKeyNewParamsFiltersTypeBoolean:
		return true
	}
	return false
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeyNewParamsFiltersValueUnion interface {
	ImplementsObservabilityTelemetryKeyNewParamsFiltersValueUnion()
}

// Search for a specific substring in the keys.
type ObservabilityTelemetryKeyNewParamsKeyNeedle struct {
	Value     param.Field[ObservabilityTelemetryKeyNewParamsKeyNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                                  `json:"isRegex"`
	MatchCase param.Field[bool]                                                  `json:"matchCase"`
}

func (r ObservabilityTelemetryKeyNewParamsKeyNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeyNewParamsKeyNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeyNewParamsKeyNeedleValueUnion()
}

// Search for a specific substring in the event.
type ObservabilityTelemetryKeyNewParamsNeedle struct {
	Value     param.Field[ObservabilityTelemetryKeyNewParamsNeedleValueUnion] `json:"value,required"`
	IsRegex   param.Field[bool]                                               `json:"isRegex"`
	MatchCase param.Field[bool]                                               `json:"matchCase"`
}

func (r ObservabilityTelemetryKeyNewParamsNeedle) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Satisfied by [shared.UnionString], [shared.UnionFloat], [shared.UnionBool].
type ObservabilityTelemetryKeyNewParamsNeedleValueUnion interface {
	ImplementsObservabilityTelemetryKeyNewParamsNeedleValueUnion()
}

type ObservabilityTelemetryKeyNewParamsTimeframe struct {
	From param.Field[float64] `json:"from,required"`
	To   param.Field[float64] `json:"to,required"`
}

func (r ObservabilityTelemetryKeyNewParamsTimeframe) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
