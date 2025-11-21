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

// SilenceService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSilenceService] method instead.
type SilenceService struct {
	Options []option.RequestOption
}

// NewSilenceService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSilenceService(opts ...option.RequestOption) (r *SilenceService) {
	r = &SilenceService{}
	r.Options = opts
	return
}

// Creates a new silence for an account.
func (r *SilenceService) New(ctx context.Context, params SilenceNewParams, opts ...option.RequestOption) (res *SilenceNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/silences", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Updates existing silences for an account.
func (r *SilenceService) Update(ctx context.Context, params SilenceUpdateParams, opts ...option.RequestOption) (res *pagination.SinglePage[SilenceUpdateResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/silences", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodPut, path, params, &res, opts...)
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

// Updates existing silences for an account.
func (r *SilenceService) UpdateAutoPaging(ctx context.Context, params SilenceUpdateParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SilenceUpdateResponse] {
	return pagination.NewSinglePageAutoPager(r.Update(ctx, params, opts...))
}

// Gets a list of silences for an account.
func (r *SilenceService) List(ctx context.Context, query SilenceListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SilenceListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/silences", query.AccountID)
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

// Gets a list of silences for an account.
func (r *SilenceService) ListAutoPaging(ctx context.Context, query SilenceListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SilenceListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an existing silence for an account.
func (r *SilenceService) Delete(ctx context.Context, silenceID string, body SilenceDeleteParams, opts ...option.RequestOption) (res *SilenceDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/silences/%s", body.AccountID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Gets a specific silence for an account.
func (r *SilenceService) Get(ctx context.Context, silenceID string, query SilenceGetParams, opts ...option.RequestOption) (res *SilenceGetResponse, err error) {
	var env SilenceGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if silenceID == "" {
		err = errors.New("missing required silence_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/alerting/v3/silences/%s", query.AccountID, silenceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SilenceNewResponse struct {
	Errors   []SilenceNewResponseError   `json:"errors,required"`
	Messages []SilenceNewResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success SilenceNewResponseSuccess `json:"success,required"`
	JSON    silenceNewResponseJSON    `json:"-"`
}

// silenceNewResponseJSON contains the JSON metadata for the struct
// [SilenceNewResponse]
type silenceNewResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceNewResponseJSON) RawJSON() string {
	return r.raw
}

type SilenceNewResponseError struct {
	Message string                      `json:"message,required"`
	Code    int64                       `json:"code"`
	JSON    silenceNewResponseErrorJSON `json:"-"`
}

// silenceNewResponseErrorJSON contains the JSON metadata for the struct
// [SilenceNewResponseError]
type silenceNewResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceNewResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceNewResponseErrorJSON) RawJSON() string {
	return r.raw
}

type SilenceNewResponseMessage struct {
	Message string                        `json:"message,required"`
	Code    int64                         `json:"code"`
	JSON    silenceNewResponseMessageJSON `json:"-"`
}

// silenceNewResponseMessageJSON contains the JSON metadata for the struct
// [SilenceNewResponseMessage]
type silenceNewResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceNewResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceNewResponseMessageJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SilenceNewResponseSuccess bool

const (
	SilenceNewResponseSuccessTrue SilenceNewResponseSuccess = true
)

func (r SilenceNewResponseSuccess) IsKnown() bool {
	switch r {
	case SilenceNewResponseSuccessTrue:
		return true
	}
	return false
}

type SilenceUpdateResponse struct {
	// Silence ID
	ID string `json:"id"`
	// When the silence was created.
	CreatedAt string `json:"created_at"`
	// When the silence ends.
	EndTime string `json:"end_time"`
	// The unique identifier of a notification policy
	PolicyID string `json:"policy_id"`
	// When the silence starts.
	StartTime string `json:"start_time"`
	// When the silence was modified.
	UpdatedAt string                    `json:"updated_at"`
	JSON      silenceUpdateResponseJSON `json:"-"`
}

// silenceUpdateResponseJSON contains the JSON metadata for the struct
// [SilenceUpdateResponse]
type silenceUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	EndTime     apijson.Field
	PolicyID    apijson.Field
	StartTime   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SilenceListResponse struct {
	// Silence ID
	ID string `json:"id"`
	// When the silence was created.
	CreatedAt string `json:"created_at"`
	// When the silence ends.
	EndTime string `json:"end_time"`
	// The unique identifier of a notification policy
	PolicyID string `json:"policy_id"`
	// When the silence starts.
	StartTime string `json:"start_time"`
	// When the silence was modified.
	UpdatedAt string                  `json:"updated_at"`
	JSON      silenceListResponseJSON `json:"-"`
}

// silenceListResponseJSON contains the JSON metadata for the struct
// [SilenceListResponse]
type silenceListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	EndTime     apijson.Field
	PolicyID    apijson.Field
	StartTime   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceListResponseJSON) RawJSON() string {
	return r.raw
}

type SilenceDeleteResponse struct {
	Errors   []SilenceDeleteResponseError   `json:"errors,required"`
	Messages []SilenceDeleteResponseMessage `json:"messages,required"`
	// Whether the API call was successful
	Success SilenceDeleteResponseSuccess `json:"success,required"`
	JSON    silenceDeleteResponseJSON    `json:"-"`
}

// silenceDeleteResponseJSON contains the JSON metadata for the struct
// [SilenceDeleteResponse]
type silenceDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SilenceDeleteResponseError struct {
	Message string                         `json:"message,required"`
	Code    int64                          `json:"code"`
	JSON    silenceDeleteResponseErrorJSON `json:"-"`
}

// silenceDeleteResponseErrorJSON contains the JSON metadata for the struct
// [SilenceDeleteResponseError]
type silenceDeleteResponseErrorJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceDeleteResponseErrorJSON) RawJSON() string {
	return r.raw
}

type SilenceDeleteResponseMessage struct {
	Message string                           `json:"message,required"`
	Code    int64                            `json:"code"`
	JSON    silenceDeleteResponseMessageJSON `json:"-"`
}

// silenceDeleteResponseMessageJSON contains the JSON metadata for the struct
// [SilenceDeleteResponseMessage]
type silenceDeleteResponseMessageJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceDeleteResponseMessageJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SilenceDeleteResponseSuccess bool

const (
	SilenceDeleteResponseSuccessTrue SilenceDeleteResponseSuccess = true
)

func (r SilenceDeleteResponseSuccess) IsKnown() bool {
	switch r {
	case SilenceDeleteResponseSuccessTrue:
		return true
	}
	return false
}

type SilenceGetResponse struct {
	// Silence ID
	ID string `json:"id"`
	// When the silence was created.
	CreatedAt string `json:"created_at"`
	// When the silence ends.
	EndTime string `json:"end_time"`
	// The unique identifier of a notification policy
	PolicyID string `json:"policy_id"`
	// When the silence starts.
	StartTime string `json:"start_time"`
	// When the silence was modified.
	UpdatedAt string                 `json:"updated_at"`
	JSON      silenceGetResponseJSON `json:"-"`
}

// silenceGetResponseJSON contains the JSON metadata for the struct
// [SilenceGetResponse]
type silenceGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	EndTime     apijson.Field
	PolicyID    apijson.Field
	StartTime   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type SilenceNewParams struct {
	// The account id
	AccountID param.Field[string]    `path:"account_id,required"`
	Body      []SilenceNewParamsBody `json:"body,required"`
}

func (r SilenceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SilenceNewParamsBody struct {
	// When the silence ends.
	EndTime param.Field[string] `json:"end_time"`
	// The unique identifier of a notification policy
	PolicyID param.Field[string] `json:"policy_id"`
	// When the silence starts.
	StartTime param.Field[string] `json:"start_time"`
}

func (r SilenceNewParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SilenceUpdateParams struct {
	// The account id
	AccountID param.Field[string]       `path:"account_id,required"`
	Body      []SilenceUpdateParamsBody `json:"body,required"`
}

func (r SilenceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.Body)
}

type SilenceUpdateParamsBody struct {
	// Silence ID
	ID param.Field[string] `json:"id"`
	// When the silence ends.
	EndTime param.Field[string] `json:"end_time"`
	// When the silence starts.
	StartTime param.Field[string] `json:"start_time"`
}

func (r SilenceUpdateParamsBody) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SilenceListParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type SilenceDeleteParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type SilenceGetParams struct {
	// The account id
	AccountID param.Field[string] `path:"account_id,required"`
}

type SilenceGetResponseEnvelope struct {
	Errors   []SilenceGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SilenceGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful
	Success SilenceGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SilenceGetResponse                `json:"result"`
	JSON    silenceGetResponseEnvelopeJSON    `json:"-"`
}

// silenceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SilenceGetResponseEnvelope]
type silenceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type SilenceGetResponseEnvelopeErrors struct {
	Message string                               `json:"message,required"`
	Code    int64                                `json:"code"`
	JSON    silenceGetResponseEnvelopeErrorsJSON `json:"-"`
}

// silenceGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [SilenceGetResponseEnvelopeErrors]
type silenceGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type SilenceGetResponseEnvelopeMessages struct {
	Message string                                 `json:"message,required"`
	Code    int64                                  `json:"code"`
	JSON    silenceGetResponseEnvelopeMessagesJSON `json:"-"`
}

// silenceGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [SilenceGetResponseEnvelopeMessages]
type silenceGetResponseEnvelopeMessagesJSON struct {
	Message     apijson.Field
	Code        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SilenceGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r silenceGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SilenceGetResponseEnvelopeSuccess bool

const (
	SilenceGetResponseEnvelopeSuccessTrue SilenceGetResponseEnvelopeSuccess = true
)

func (r SilenceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SilenceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
