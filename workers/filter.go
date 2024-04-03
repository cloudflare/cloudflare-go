// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// FilterService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewFilterService] method instead.
type FilterService struct {
	Options []option.RequestOption
}

// NewFilterService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFilterService(opts ...option.RequestOption) (r *FilterService) {
	r = &FilterService{}
	r.Options = opts
	return
}

// Create Filter
func (r *FilterService) New(ctx context.Context, params FilterNewParams, opts ...option.RequestOption) (res *FilterNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update Filter
func (r *FilterService) Update(ctx context.Context, filterID string, params FilterUpdateParams, opts ...option.RequestOption) (res *WorkersFilter, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterUpdateResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters/%s", params.ZoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List Filters
func (r *FilterService) List(ctx context.Context, query FilterListParams, opts ...option.RequestOption) (res *pagination.SinglePage[WorkersFilter], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("zones/%s/workers/filters", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

// List Filters
func (r *FilterService) ListAutoPaging(ctx context.Context, query FilterListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[WorkersFilter] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete Filter
func (r *FilterService) Delete(ctx context.Context, filterID string, params FilterDeleteParams, opts ...option.RequestOption) (res *FilterDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env FilterDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/workers/filters/%s", params.ZoneID, filterID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type WorkersFilter struct {
	// Identifier
	ID      string            `json:"id,required"`
	Enabled bool              `json:"enabled,required"`
	Pattern string            `json:"pattern,required"`
	JSON    workersFilterJSON `json:"-"`
}

// workersFilterJSON contains the JSON metadata for the struct [WorkersFilter]
type workersFilterJSON struct {
	ID          apijson.Field
	Enabled     apijson.Field
	Pattern     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *WorkersFilter) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r workersFilterJSON) RawJSON() string {
	return r.raw
}

type FilterNewResponse struct {
	// Identifier
	ID   string                `json:"id,required"`
	JSON filterNewResponseJSON `json:"-"`
}

// filterNewResponseJSON contains the JSON metadata for the struct
// [FilterNewResponse]
type filterNewResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterNewResponseJSON) RawJSON() string {
	return r.raw
}

type FilterDeleteResponse struct {
	// Identifier
	ID   string                   `json:"id,required"`
	JSON filterDeleteResponseJSON `json:"-"`
}

// filterDeleteResponseJSON contains the JSON metadata for the struct
// [FilterDeleteResponse]
type filterDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type FilterNewParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Enabled param.Field[bool]   `json:"enabled,required"`
	Pattern param.Field[string] `json:"pattern,required"`
}

func (r FilterNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FilterNewResponseEnvelope struct {
	Errors   []FilterNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterNewResponseEnvelopeMessages `json:"messages,required"`
	Result   FilterNewResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FilterNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterNewResponseEnvelopeJSON    `json:"-"`
}

// filterNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterNewResponseEnvelope]
type filterNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FilterNewResponseEnvelopeErrors struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    filterNewResponseEnvelopeErrorsJSON `json:"-"`
}

// filterNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FilterNewResponseEnvelopeErrors]
type filterNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FilterNewResponseEnvelopeMessages struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    filterNewResponseEnvelopeMessagesJSON `json:"-"`
}

// filterNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [FilterNewResponseEnvelopeMessages]
type filterNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FilterNewResponseEnvelopeSuccess bool

const (
	FilterNewResponseEnvelopeSuccessTrue FilterNewResponseEnvelopeSuccess = true
)

func (r FilterNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FilterNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FilterUpdateParams struct {
	// Identifier
	ZoneID  param.Field[string] `path:"zone_id,required"`
	Enabled param.Field[bool]   `json:"enabled,required"`
	Pattern param.Field[string] `json:"pattern,required"`
}

func (r FilterUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FilterUpdateResponseEnvelope struct {
	Errors   []FilterUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   WorkersFilter                          `json:"result,required"`
	// Whether the API call was successful
	Success FilterUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterUpdateResponseEnvelopeJSON    `json:"-"`
}

// filterUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterUpdateResponseEnvelope]
type filterUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FilterUpdateResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    filterUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// filterUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FilterUpdateResponseEnvelopeErrors]
type filterUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FilterUpdateResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    filterUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// filterUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FilterUpdateResponseEnvelopeMessages]
type filterUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FilterUpdateResponseEnvelopeSuccess bool

const (
	FilterUpdateResponseEnvelopeSuccessTrue FilterUpdateResponseEnvelopeSuccess = true
)

func (r FilterUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FilterUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FilterListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type FilterDeleteParams struct {
	// Identifier
	ZoneID param.Field[string]      `path:"zone_id,required"`
	Body   param.Field[interface{}] `json:"body,required"`
}

func (r FilterDeleteParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type FilterDeleteResponseEnvelope struct {
	Errors   []FilterDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []FilterDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   FilterDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful
	Success FilterDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    filterDeleteResponseEnvelopeJSON    `json:"-"`
}

// filterDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [FilterDeleteResponseEnvelope]
type filterDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FilterDeleteResponseEnvelopeErrors struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    filterDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// filterDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FilterDeleteResponseEnvelopeErrors]
type filterDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FilterDeleteResponseEnvelopeMessages struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    filterDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// filterDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FilterDeleteResponseEnvelopeMessages]
type filterDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FilterDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r filterDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type FilterDeleteResponseEnvelopeSuccess bool

const (
	FilterDeleteResponseEnvelopeSuccessTrue FilterDeleteResponseEnvelopeSuccess = true
)

func (r FilterDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FilterDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
