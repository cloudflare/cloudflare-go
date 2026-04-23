// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ObservabilityDestinationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewObservabilityDestinationService] method instead.
type ObservabilityDestinationService struct {
	Options []option.RequestOption
}

// NewObservabilityDestinationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewObservabilityDestinationService(opts ...option.RequestOption) (r *ObservabilityDestinationService) {
	r = &ObservabilityDestinationService{}
	r.Options = opts
	return
}

// Create a new Workers Observability Telemetry Destination.
func (r *ObservabilityDestinationService) New(ctx context.Context, params ObservabilityDestinationNewParams, opts ...option.RequestOption) (res *ObservabilityDestinationNewResponse, err error) {
	var env ObservabilityDestinationNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/destinations", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update an existing Workers Observability Telemetry Destination.
func (r *ObservabilityDestinationService) Update(ctx context.Context, slug string, params ObservabilityDestinationUpdateParams, opts ...option.RequestOption) (res *ObservabilityDestinationUpdateResponse, err error) {
	var env ObservabilityDestinationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/destinations/%s", params.AccountID, slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List your Workers Observability Telemetry Destinations.
func (r *ObservabilityDestinationService) List(ctx context.Context, params ObservabilityDestinationListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ObservabilityDestinationListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/destinations", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List your Workers Observability Telemetry Destinations.
func (r *ObservabilityDestinationService) ListAutoPaging(ctx context.Context, params ObservabilityDestinationListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ObservabilityDestinationListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Delete a Workers Observability Telemetry Destination.
func (r *ObservabilityDestinationService) Delete(ctx context.Context, slug string, body ObservabilityDestinationDeleteParams, opts ...option.RequestOption) (res *ObservabilityDestinationDeleteResponse, err error) {
	var env ObservabilityDestinationDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if slug == "" {
		err = errors.New("missing required slug parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/workers/observability/destinations/%s", body.AccountID, slug)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ObservabilityDestinationNewResponse struct {
	Configuration ObservabilityDestinationNewResponseConfiguration `json:"configuration" api:"required"`
	Enabled       bool                                             `json:"enabled" api:"required"`
	Name          string                                           `json:"name" api:"required"`
	Scripts       []string                                         `json:"scripts" api:"required"`
	Slug          string                                           `json:"slug" api:"required"`
	JSON          observabilityDestinationNewResponseJSON          `json:"-"`
}

// observabilityDestinationNewResponseJSON contains the JSON metadata for the
// struct [ObservabilityDestinationNewResponse]
type observabilityDestinationNewResponseJSON struct {
	Configuration apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Scripts       apijson.Field
	Slug          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityDestinationNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationNewResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationNewResponseConfiguration struct {
	DestinationConf string                                                         `json:"destination_conf" api:"required"`
	LogpushDataset  ObservabilityDestinationNewResponseConfigurationLogpushDataset `json:"logpushDataset" api:"required"`
	LogpushJob      float64                                                        `json:"logpushJob" api:"required"`
	Type            ObservabilityDestinationNewResponseConfigurationType           `json:"type" api:"required"`
	URL             string                                                         `json:"url" api:"required"`
	JSON            observabilityDestinationNewResponseConfigurationJSON           `json:"-"`
}

// observabilityDestinationNewResponseConfigurationJSON contains the JSON metadata
// for the struct [ObservabilityDestinationNewResponseConfiguration]
type observabilityDestinationNewResponseConfigurationJSON struct {
	DestinationConf apijson.Field
	LogpushDataset  apijson.Field
	LogpushJob      apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityDestinationNewResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationNewResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationNewResponseConfigurationLogpushDataset string

const (
	ObservabilityDestinationNewResponseConfigurationLogpushDatasetOpentelemetryTraces ObservabilityDestinationNewResponseConfigurationLogpushDataset = "opentelemetry-traces"
	ObservabilityDestinationNewResponseConfigurationLogpushDatasetOpentelemetryLogs   ObservabilityDestinationNewResponseConfigurationLogpushDataset = "opentelemetry-logs"
)

func (r ObservabilityDestinationNewResponseConfigurationLogpushDataset) IsKnown() bool {
	switch r {
	case ObservabilityDestinationNewResponseConfigurationLogpushDatasetOpentelemetryTraces, ObservabilityDestinationNewResponseConfigurationLogpushDatasetOpentelemetryLogs:
		return true
	}
	return false
}

type ObservabilityDestinationNewResponseConfigurationType string

const (
	ObservabilityDestinationNewResponseConfigurationTypeLogpush ObservabilityDestinationNewResponseConfigurationType = "logpush"
)

func (r ObservabilityDestinationNewResponseConfigurationType) IsKnown() bool {
	switch r {
	case ObservabilityDestinationNewResponseConfigurationTypeLogpush:
		return true
	}
	return false
}

type ObservabilityDestinationUpdateResponse struct {
	Configuration ObservabilityDestinationUpdateResponseConfiguration `json:"configuration" api:"required"`
	Enabled       bool                                                `json:"enabled" api:"required"`
	Name          string                                              `json:"name" api:"required"`
	Scripts       []string                                            `json:"scripts" api:"required"`
	Slug          string                                              `json:"slug" api:"required"`
	JSON          observabilityDestinationUpdateResponseJSON          `json:"-"`
}

// observabilityDestinationUpdateResponseJSON contains the JSON metadata for the
// struct [ObservabilityDestinationUpdateResponse]
type observabilityDestinationUpdateResponseJSON struct {
	Configuration apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Scripts       apijson.Field
	Slug          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityDestinationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationUpdateResponseConfiguration struct {
	DestinationConf string                                                            `json:"destination_conf" api:"required"`
	LogpushDataset  ObservabilityDestinationUpdateResponseConfigurationLogpushDataset `json:"logpushDataset" api:"required"`
	LogpushJob      float64                                                           `json:"logpushJob" api:"required"`
	Type            ObservabilityDestinationUpdateResponseConfigurationType           `json:"type" api:"required"`
	URL             string                                                            `json:"url" api:"required"`
	JSON            observabilityDestinationUpdateResponseConfigurationJSON           `json:"-"`
}

// observabilityDestinationUpdateResponseConfigurationJSON contains the JSON
// metadata for the struct [ObservabilityDestinationUpdateResponseConfiguration]
type observabilityDestinationUpdateResponseConfigurationJSON struct {
	DestinationConf apijson.Field
	LogpushDataset  apijson.Field
	LogpushJob      apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityDestinationUpdateResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationUpdateResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationUpdateResponseConfigurationLogpushDataset string

const (
	ObservabilityDestinationUpdateResponseConfigurationLogpushDatasetOpentelemetryTraces ObservabilityDestinationUpdateResponseConfigurationLogpushDataset = "opentelemetry-traces"
	ObservabilityDestinationUpdateResponseConfigurationLogpushDatasetOpentelemetryLogs   ObservabilityDestinationUpdateResponseConfigurationLogpushDataset = "opentelemetry-logs"
)

func (r ObservabilityDestinationUpdateResponseConfigurationLogpushDataset) IsKnown() bool {
	switch r {
	case ObservabilityDestinationUpdateResponseConfigurationLogpushDatasetOpentelemetryTraces, ObservabilityDestinationUpdateResponseConfigurationLogpushDatasetOpentelemetryLogs:
		return true
	}
	return false
}

type ObservabilityDestinationUpdateResponseConfigurationType string

const (
	ObservabilityDestinationUpdateResponseConfigurationTypeLogpush ObservabilityDestinationUpdateResponseConfigurationType = "logpush"
)

func (r ObservabilityDestinationUpdateResponseConfigurationType) IsKnown() bool {
	switch r {
	case ObservabilityDestinationUpdateResponseConfigurationTypeLogpush:
		return true
	}
	return false
}

type ObservabilityDestinationListResponse struct {
	Configuration ObservabilityDestinationListResponseConfiguration `json:"configuration" api:"required"`
	Enabled       bool                                              `json:"enabled" api:"required"`
	Name          string                                            `json:"name" api:"required"`
	Scripts       []string                                          `json:"scripts" api:"required"`
	Slug          string                                            `json:"slug" api:"required"`
	JSON          observabilityDestinationListResponseJSON          `json:"-"`
}

// observabilityDestinationListResponseJSON contains the JSON metadata for the
// struct [ObservabilityDestinationListResponse]
type observabilityDestinationListResponseJSON struct {
	Configuration apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Scripts       apijson.Field
	Slug          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityDestinationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationListResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationListResponseConfiguration struct {
	DestinationConf string                                                          `json:"destination_conf" api:"required"`
	Headers         map[string]string                                               `json:"headers" api:"required"`
	JobStatus       ObservabilityDestinationListResponseConfigurationJobStatus      `json:"jobStatus" api:"required"`
	LogpushDataset  ObservabilityDestinationListResponseConfigurationLogpushDataset `json:"logpushDataset" api:"required"`
	Type            ObservabilityDestinationListResponseConfigurationType           `json:"type" api:"required"`
	URL             string                                                          `json:"url" api:"required"`
	JSON            observabilityDestinationListResponseConfigurationJSON           `json:"-"`
}

// observabilityDestinationListResponseConfigurationJSON contains the JSON metadata
// for the struct [ObservabilityDestinationListResponseConfiguration]
type observabilityDestinationListResponseConfigurationJSON struct {
	DestinationConf apijson.Field
	Headers         apijson.Field
	JobStatus       apijson.Field
	LogpushDataset  apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityDestinationListResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationListResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationListResponseConfigurationJobStatus struct {
	ErrorMessage string                                                         `json:"error_message" api:"required"`
	LastComplete string                                                         `json:"last_complete" api:"required"`
	LastError    string                                                         `json:"last_error" api:"required"`
	JSON         observabilityDestinationListResponseConfigurationJobStatusJSON `json:"-"`
}

// observabilityDestinationListResponseConfigurationJobStatusJSON contains the JSON
// metadata for the struct
// [ObservabilityDestinationListResponseConfigurationJobStatus]
type observabilityDestinationListResponseConfigurationJobStatusJSON struct {
	ErrorMessage apijson.Field
	LastComplete apijson.Field
	LastError    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ObservabilityDestinationListResponseConfigurationJobStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationListResponseConfigurationJobStatusJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationListResponseConfigurationLogpushDataset string

const (
	ObservabilityDestinationListResponseConfigurationLogpushDatasetOpentelemetryTraces ObservabilityDestinationListResponseConfigurationLogpushDataset = "opentelemetry-traces"
	ObservabilityDestinationListResponseConfigurationLogpushDatasetOpentelemetryLogs   ObservabilityDestinationListResponseConfigurationLogpushDataset = "opentelemetry-logs"
)

func (r ObservabilityDestinationListResponseConfigurationLogpushDataset) IsKnown() bool {
	switch r {
	case ObservabilityDestinationListResponseConfigurationLogpushDatasetOpentelemetryTraces, ObservabilityDestinationListResponseConfigurationLogpushDatasetOpentelemetryLogs:
		return true
	}
	return false
}

type ObservabilityDestinationListResponseConfigurationType string

const (
	ObservabilityDestinationListResponseConfigurationTypeLogpush ObservabilityDestinationListResponseConfigurationType = "logpush"
)

func (r ObservabilityDestinationListResponseConfigurationType) IsKnown() bool {
	switch r {
	case ObservabilityDestinationListResponseConfigurationTypeLogpush:
		return true
	}
	return false
}

type ObservabilityDestinationDeleteResponse struct {
	Configuration ObservabilityDestinationDeleteResponseConfiguration `json:"configuration" api:"required"`
	Enabled       bool                                                `json:"enabled" api:"required"`
	Name          string                                              `json:"name" api:"required"`
	Scripts       []string                                            `json:"scripts" api:"required"`
	Slug          string                                              `json:"slug" api:"required"`
	JSON          observabilityDestinationDeleteResponseJSON          `json:"-"`
}

// observabilityDestinationDeleteResponseJSON contains the JSON metadata for the
// struct [ObservabilityDestinationDeleteResponse]
type observabilityDestinationDeleteResponseJSON struct {
	Configuration apijson.Field
	Enabled       apijson.Field
	Name          apijson.Field
	Scripts       apijson.Field
	Slug          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ObservabilityDestinationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationDeleteResponseConfiguration struct {
	DestinationConf string                                                            `json:"destination_conf" api:"required"`
	LogpushDataset  ObservabilityDestinationDeleteResponseConfigurationLogpushDataset `json:"logpushDataset" api:"required"`
	LogpushJob      float64                                                           `json:"logpushJob" api:"required"`
	Type            ObservabilityDestinationDeleteResponseConfigurationType           `json:"type" api:"required"`
	URL             string                                                            `json:"url" api:"required"`
	JSON            observabilityDestinationDeleteResponseConfigurationJSON           `json:"-"`
}

// observabilityDestinationDeleteResponseConfigurationJSON contains the JSON
// metadata for the struct [ObservabilityDestinationDeleteResponseConfiguration]
type observabilityDestinationDeleteResponseConfigurationJSON struct {
	DestinationConf apijson.Field
	LogpushDataset  apijson.Field
	LogpushJob      apijson.Field
	Type            apijson.Field
	URL             apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *ObservabilityDestinationDeleteResponseConfiguration) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationDeleteResponseConfigurationJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationDeleteResponseConfigurationLogpushDataset string

const (
	ObservabilityDestinationDeleteResponseConfigurationLogpushDatasetOpentelemetryTraces ObservabilityDestinationDeleteResponseConfigurationLogpushDataset = "opentelemetry-traces"
	ObservabilityDestinationDeleteResponseConfigurationLogpushDatasetOpentelemetryLogs   ObservabilityDestinationDeleteResponseConfigurationLogpushDataset = "opentelemetry-logs"
)

func (r ObservabilityDestinationDeleteResponseConfigurationLogpushDataset) IsKnown() bool {
	switch r {
	case ObservabilityDestinationDeleteResponseConfigurationLogpushDatasetOpentelemetryTraces, ObservabilityDestinationDeleteResponseConfigurationLogpushDatasetOpentelemetryLogs:
		return true
	}
	return false
}

type ObservabilityDestinationDeleteResponseConfigurationType string

const (
	ObservabilityDestinationDeleteResponseConfigurationTypeLogpush ObservabilityDestinationDeleteResponseConfigurationType = "logpush"
)

func (r ObservabilityDestinationDeleteResponseConfigurationType) IsKnown() bool {
	switch r {
	case ObservabilityDestinationDeleteResponseConfigurationTypeLogpush:
		return true
	}
	return false
}

type ObservabilityDestinationNewParams struct {
	AccountID          param.Field[string]                                         `path:"account_id" api:"required"`
	Configuration      param.Field[ObservabilityDestinationNewParamsConfiguration] `json:"configuration" api:"required"`
	Enabled            param.Field[bool]                                           `json:"enabled" api:"required"`
	Name               param.Field[string]                                         `json:"name" api:"required"`
	SkipPreflightCheck param.Field[bool]                                           `json:"skipPreflightCheck"`
}

func (r ObservabilityDestinationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityDestinationNewParamsConfiguration struct {
	Headers        param.Field[map[string]string]                                            `json:"headers" api:"required"`
	LogpushDataset param.Field[ObservabilityDestinationNewParamsConfigurationLogpushDataset] `json:"logpushDataset" api:"required"`
	Type           param.Field[ObservabilityDestinationNewParamsConfigurationType]           `json:"type" api:"required"`
	URL            param.Field[string]                                                       `json:"url" api:"required"`
}

func (r ObservabilityDestinationNewParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityDestinationNewParamsConfigurationLogpushDataset string

const (
	ObservabilityDestinationNewParamsConfigurationLogpushDatasetOpentelemetryTraces ObservabilityDestinationNewParamsConfigurationLogpushDataset = "opentelemetry-traces"
	ObservabilityDestinationNewParamsConfigurationLogpushDatasetOpentelemetryLogs   ObservabilityDestinationNewParamsConfigurationLogpushDataset = "opentelemetry-logs"
)

func (r ObservabilityDestinationNewParamsConfigurationLogpushDataset) IsKnown() bool {
	switch r {
	case ObservabilityDestinationNewParamsConfigurationLogpushDatasetOpentelemetryTraces, ObservabilityDestinationNewParamsConfigurationLogpushDatasetOpentelemetryLogs:
		return true
	}
	return false
}

type ObservabilityDestinationNewParamsConfigurationType string

const (
	ObservabilityDestinationNewParamsConfigurationTypeLogpush ObservabilityDestinationNewParamsConfigurationType = "logpush"
)

func (r ObservabilityDestinationNewParamsConfigurationType) IsKnown() bool {
	switch r {
	case ObservabilityDestinationNewParamsConfigurationTypeLogpush:
		return true
	}
	return false
}

type ObservabilityDestinationNewResponseEnvelope struct {
	Errors   []ObservabilityDestinationNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityDestinationNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityDestinationNewResponse                   `json:"result" api:"required"`
	Success  ObservabilityDestinationNewResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityDestinationNewResponseEnvelopeJSON       `json:"-"`
}

// observabilityDestinationNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ObservabilityDestinationNewResponseEnvelope]
type observabilityDestinationNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationNewResponseEnvelopeErrors struct {
	Message string                                                `json:"message" api:"required"`
	JSON    observabilityDestinationNewResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityDestinationNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ObservabilityDestinationNewResponseEnvelopeErrors]
type observabilityDestinationNewResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationNewResponseEnvelopeMessages struct {
	Message ObservabilityDestinationNewResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityDestinationNewResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityDestinationNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityDestinationNewResponseEnvelopeMessages]
type observabilityDestinationNewResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationNewResponseEnvelopeMessagesMessage string

const (
	ObservabilityDestinationNewResponseEnvelopeMessagesMessageResourceCreated ObservabilityDestinationNewResponseEnvelopeMessagesMessage = "Resource created"
)

func (r ObservabilityDestinationNewResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityDestinationNewResponseEnvelopeMessagesMessageResourceCreated:
		return true
	}
	return false
}

type ObservabilityDestinationNewResponseEnvelopeSuccess bool

const (
	ObservabilityDestinationNewResponseEnvelopeSuccessTrue ObservabilityDestinationNewResponseEnvelopeSuccess = true
)

func (r ObservabilityDestinationNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityDestinationNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityDestinationUpdateParams struct {
	AccountID     param.Field[string]                                            `path:"account_id" api:"required"`
	Configuration param.Field[ObservabilityDestinationUpdateParamsConfiguration] `json:"configuration" api:"required"`
	Enabled       param.Field[bool]                                              `json:"enabled" api:"required"`
}

func (r ObservabilityDestinationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityDestinationUpdateParamsConfiguration struct {
	Headers param.Field[map[string]string]                                     `json:"headers" api:"required"`
	Type    param.Field[ObservabilityDestinationUpdateParamsConfigurationType] `json:"type" api:"required"`
	URL     param.Field[string]                                                `json:"url" api:"required"`
}

func (r ObservabilityDestinationUpdateParamsConfiguration) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ObservabilityDestinationUpdateParamsConfigurationType string

const (
	ObservabilityDestinationUpdateParamsConfigurationTypeLogpush ObservabilityDestinationUpdateParamsConfigurationType = "logpush"
)

func (r ObservabilityDestinationUpdateParamsConfigurationType) IsKnown() bool {
	switch r {
	case ObservabilityDestinationUpdateParamsConfigurationTypeLogpush:
		return true
	}
	return false
}

type ObservabilityDestinationUpdateResponseEnvelope struct {
	Errors   []ObservabilityDestinationUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityDestinationUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ObservabilityDestinationUpdateResponse                   `json:"result" api:"required"`
	Success  ObservabilityDestinationUpdateResponseEnvelopeSuccess    `json:"success" api:"required"`
	JSON     observabilityDestinationUpdateResponseEnvelopeJSON       `json:"-"`
}

// observabilityDestinationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityDestinationUpdateResponseEnvelope]
type observabilityDestinationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationUpdateResponseEnvelopeErrors struct {
	Message string                                                   `json:"message" api:"required"`
	JSON    observabilityDestinationUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityDestinationUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityDestinationUpdateResponseEnvelopeErrors]
type observabilityDestinationUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationUpdateResponseEnvelopeMessages struct {
	Message ObservabilityDestinationUpdateResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityDestinationUpdateResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityDestinationUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityDestinationUpdateResponseEnvelopeMessages]
type observabilityDestinationUpdateResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationUpdateResponseEnvelopeMessagesMessage string

const (
	ObservabilityDestinationUpdateResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityDestinationUpdateResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityDestinationUpdateResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityDestinationUpdateResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityDestinationUpdateResponseEnvelopeSuccess bool

const (
	ObservabilityDestinationUpdateResponseEnvelopeSuccessTrue ObservabilityDestinationUpdateResponseEnvelopeSuccess = true
)

func (r ObservabilityDestinationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityDestinationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ObservabilityDestinationListParams struct {
	AccountID param.Field[string]                                    `path:"account_id" api:"required"`
	Order     param.Field[ObservabilityDestinationListParamsOrder]   `query:"order"`
	OrderBy   param.Field[ObservabilityDestinationListParamsOrderBy] `query:"orderBy"`
	Page      param.Field[float64]                                   `query:"page"`
	PerPage   param.Field[float64]                                   `query:"perPage"`
}

// URLQuery serializes [ObservabilityDestinationListParams]'s query parameters as
// `url.Values`.
func (r ObservabilityDestinationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ObservabilityDestinationListParamsOrder string

const (
	ObservabilityDestinationListParamsOrderAsc  ObservabilityDestinationListParamsOrder = "asc"
	ObservabilityDestinationListParamsOrderDesc ObservabilityDestinationListParamsOrder = "desc"
)

func (r ObservabilityDestinationListParamsOrder) IsKnown() bool {
	switch r {
	case ObservabilityDestinationListParamsOrderAsc, ObservabilityDestinationListParamsOrderDesc:
		return true
	}
	return false
}

type ObservabilityDestinationListParamsOrderBy string

const (
	ObservabilityDestinationListParamsOrderByCreated ObservabilityDestinationListParamsOrderBy = "created"
	ObservabilityDestinationListParamsOrderByUpdated ObservabilityDestinationListParamsOrderBy = "updated"
)

func (r ObservabilityDestinationListParamsOrderBy) IsKnown() bool {
	switch r {
	case ObservabilityDestinationListParamsOrderByCreated, ObservabilityDestinationListParamsOrderByUpdated:
		return true
	}
	return false
}

type ObservabilityDestinationDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ObservabilityDestinationDeleteResponseEnvelope struct {
	Errors   []ObservabilityDestinationDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ObservabilityDestinationDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Success  ObservabilityDestinationDeleteResponseEnvelopeSuccess    `json:"success" api:"required"`
	Result   ObservabilityDestinationDeleteResponse                   `json:"result"`
	JSON     observabilityDestinationDeleteResponseEnvelopeJSON       `json:"-"`
}

// observabilityDestinationDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ObservabilityDestinationDeleteResponseEnvelope]
type observabilityDestinationDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationDeleteResponseEnvelopeErrors struct {
	Message string                                                   `json:"message" api:"required"`
	JSON    observabilityDestinationDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// observabilityDestinationDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ObservabilityDestinationDeleteResponseEnvelopeErrors]
type observabilityDestinationDeleteResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationDeleteResponseEnvelopeMessages struct {
	Message ObservabilityDestinationDeleteResponseEnvelopeMessagesMessage `json:"message" api:"required"`
	JSON    observabilityDestinationDeleteResponseEnvelopeMessagesJSON    `json:"-"`
}

// observabilityDestinationDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ObservabilityDestinationDeleteResponseEnvelopeMessages]
type observabilityDestinationDeleteResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ObservabilityDestinationDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r observabilityDestinationDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ObservabilityDestinationDeleteResponseEnvelopeMessagesMessage string

const (
	ObservabilityDestinationDeleteResponseEnvelopeMessagesMessageSuccessfulRequest ObservabilityDestinationDeleteResponseEnvelopeMessagesMessage = "Successful request"
)

func (r ObservabilityDestinationDeleteResponseEnvelopeMessagesMessage) IsKnown() bool {
	switch r {
	case ObservabilityDestinationDeleteResponseEnvelopeMessagesMessageSuccessfulRequest:
		return true
	}
	return false
}

type ObservabilityDestinationDeleteResponseEnvelopeSuccess bool

const (
	ObservabilityDestinationDeleteResponseEnvelopeSuccessTrue ObservabilityDestinationDeleteResponseEnvelopeSuccess = true
)

func (r ObservabilityDestinationDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ObservabilityDestinationDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
