// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
)

// DLPSensitivityGroupLevelService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPSensitivityGroupLevelService] method instead.
type DLPSensitivityGroupLevelService struct {
	Options []option.RequestOption
	Order   *DLPSensitivityGroupLevelOrderService
}

// NewDLPSensitivityGroupLevelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDLPSensitivityGroupLevelService(opts ...option.RequestOption) (r *DLPSensitivityGroupLevelService) {
	r = &DLPSensitivityGroupLevelService{}
	r.Options = opts
	r.Order = NewDLPSensitivityGroupLevelOrderService(opts...)
	return
}

// Creates a new sensitivity level.
func (r *DLPSensitivityGroupLevelService) New(ctx context.Context, sensitivityGroupID string, params DLPSensitivityGroupLevelNewParams, opts ...option.RequestOption) (res *DLPSensitivityGroupLevelNewResponse, err error) {
	var env DLPSensitivityGroupLevelNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s/levels", params.AccountID, sensitivityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the attributes of a single sensitivity level.
func (r *DLPSensitivityGroupLevelService) Update(ctx context.Context, sensitivityGroupID string, sensitivityLevelID string, params DLPSensitivityGroupLevelUpdateParams, opts ...option.RequestOption) (res *DLPSensitivityGroupLevelUpdateResponse, err error) {
	var env DLPSensitivityGroupLevelUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	if sensitivityLevelID == "" {
		err = errors.New("missing required sensitivity_level_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s/levels/%s", params.AccountID, sensitivityGroupID, sensitivityLevelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve all sensitivity levels in a sensitivity group
func (r *DLPSensitivityGroupLevelService) List(ctx context.Context, sensitivityGroupID string, query DLPSensitivityGroupLevelListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPSensitivityGroupLevelListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s/levels", query.AccountID, sensitivityGroupID)
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

// Retrieve all sensitivity levels in a sensitivity group
func (r *DLPSensitivityGroupLevelService) ListAutoPaging(ctx context.Context, sensitivityGroupID string, query DLPSensitivityGroupLevelListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPSensitivityGroupLevelListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, sensitivityGroupID, query, opts...))
}

// Delete a single sensitivity level.
func (r *DLPSensitivityGroupLevelService) Delete(ctx context.Context, sensitivityGroupID string, sensitivityLevelID string, body DLPSensitivityGroupLevelDeleteParams, opts ...option.RequestOption) (res *DLPSensitivityGroupLevelDeleteResponse, err error) {
	var env DLPSensitivityGroupLevelDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	if sensitivityLevelID == "" {
		err = errors.New("missing required sensitivity_level_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s/levels/%s", body.AccountID, sensitivityGroupID, sensitivityLevelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a specific sensitivity level.
func (r *DLPSensitivityGroupLevelService) Get(ctx context.Context, sensitivityGroupID string, sensitivityLevelID string, query DLPSensitivityGroupLevelGetParams, opts ...option.RequestOption) (res *DLPSensitivityGroupLevelGetResponse, err error) {
	var env DLPSensitivityGroupLevelGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	if sensitivityLevelID == "" {
		err = errors.New("missing required sensitivity_level_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s/levels/%s", query.AccountID, sensitivityGroupID, sensitivityLevelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DLPSensitivityGroupLevelNewResponse struct {
	ID          string                                  `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Name        string                                  `json:"name" api:"required"`
	UpdatedAt   time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Description string                                  `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupLevelNewResponseJSON `json:"-"`
}

// dlpSensitivityGroupLevelNewResponseJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupLevelNewResponse]
type dlpSensitivityGroupLevelNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelUpdateResponse struct {
	ID          string                                     `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	Name        string                                     `json:"name" api:"required"`
	UpdatedAt   time.Time                                  `json:"updated_at" api:"required" format:"date-time"`
	Description string                                     `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupLevelUpdateResponseJSON `json:"-"`
}

// dlpSensitivityGroupLevelUpdateResponseJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupLevelUpdateResponse]
type dlpSensitivityGroupLevelUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelListResponse struct {
	ID          string                                   `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Name        string                                   `json:"name" api:"required"`
	UpdatedAt   time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	Description string                                   `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupLevelListResponseJSON `json:"-"`
}

// dlpSensitivityGroupLevelListResponseJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupLevelListResponse]
type dlpSensitivityGroupLevelListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelListResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelDeleteResponse = interface{}

type DLPSensitivityGroupLevelGetResponse struct {
	ID          string                                  `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Name        string                                  `json:"name" api:"required"`
	UpdatedAt   time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Description string                                  `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupLevelGetResponseJSON `json:"-"`
}

// dlpSensitivityGroupLevelGetResponseJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupLevelGetResponse]
type dlpSensitivityGroupLevelGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r DLPSensitivityGroupLevelNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPSensitivityGroupLevelNewResponseEnvelope struct {
	Errors   []DLPSensitivityGroupLevelNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupLevelNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupLevelNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupLevelNewResponse                `json:"result"`
	JSON    dlpSensitivityGroupLevelNewResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupLevelNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPSensitivityGroupLevelNewResponseEnvelope]
type dlpSensitivityGroupLevelNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelNewResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupLevelNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupLevelNewResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupLevelNewResponseEnvelopeErrors]
type dlpSensitivityGroupLevelNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpSensitivityGroupLevelNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelNewResponseEnvelopeErrorsSource]
type dlpSensitivityGroupLevelNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelNewResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupLevelNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupLevelNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelNewResponseEnvelopeMessages]
type dlpSensitivityGroupLevelNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dlpSensitivityGroupLevelNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelNewResponseEnvelopeMessagesSource]
type dlpSensitivityGroupLevelNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupLevelNewResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupLevelNewResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupLevelNewResponseEnvelopeSuccessTrue DLPSensitivityGroupLevelNewResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupLevelNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupLevelNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSensitivityGroupLevelUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DLPSensitivityGroupLevelUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPSensitivityGroupLevelUpdateResponseEnvelope struct {
	Errors   []DLPSensitivityGroupLevelUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupLevelUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupLevelUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupLevelUpdateResponse                `json:"result"`
	JSON    dlpSensitivityGroupLevelUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupLevelUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupLevelUpdateResponseEnvelope]
type dlpSensitivityGroupLevelUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelUpdateResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelUpdateResponseEnvelopeErrors]
type dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelUpdateResponseEnvelopeErrorsSource]
type dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelUpdateResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelUpdateResponseEnvelopeMessages]
type dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DLPSensitivityGroupLevelUpdateResponseEnvelopeMessagesSource]
type dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupLevelUpdateResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupLevelUpdateResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupLevelUpdateResponseEnvelopeSuccessTrue DLPSensitivityGroupLevelUpdateResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupLevelUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupLevelUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSensitivityGroupLevelListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSensitivityGroupLevelDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSensitivityGroupLevelDeleteResponseEnvelope struct {
	Errors   []DLPSensitivityGroupLevelDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupLevelDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupLevelDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupLevelDeleteResponse                `json:"result" api:"nullable"`
	JSON    dlpSensitivityGroupLevelDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupLevelDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupLevelDeleteResponseEnvelope]
type dlpSensitivityGroupLevelDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelDeleteResponseEnvelopeErrors struct {
	Code             int64                                                      `json:"code" api:"required"`
	Message          string                                                     `json:"message" api:"required"`
	DocumentationURL string                                                     `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelDeleteResponseEnvelopeErrors]
type dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                         `json:"pointer"`
	JSON    dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelDeleteResponseEnvelopeErrorsSource]
type dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelDeleteResponseEnvelopeMessages struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelDeleteResponseEnvelopeMessages]
type dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DLPSensitivityGroupLevelDeleteResponseEnvelopeMessagesSource]
type dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupLevelDeleteResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupLevelDeleteResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupLevelDeleteResponseEnvelopeSuccessTrue DLPSensitivityGroupLevelDeleteResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupLevelDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupLevelDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSensitivityGroupLevelGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSensitivityGroupLevelGetResponseEnvelope struct {
	Errors   []DLPSensitivityGroupLevelGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupLevelGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupLevelGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupLevelGetResponse                `json:"result"`
	JSON    dlpSensitivityGroupLevelGetResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupLevelGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [DLPSensitivityGroupLevelGetResponseEnvelope]
type dlpSensitivityGroupLevelGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelGetResponseEnvelopeErrors struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupLevelGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupLevelGetResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupLevelGetResponseEnvelopeErrors]
type dlpSensitivityGroupLevelGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpSensitivityGroupLevelGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelGetResponseEnvelopeErrorsSource]
type dlpSensitivityGroupLevelGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelGetResponseEnvelopeMessages struct {
	Code             int64                                                     `json:"code" api:"required"`
	Message          string                                                    `json:"message" api:"required"`
	DocumentationURL string                                                    `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupLevelGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupLevelGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelGetResponseEnvelopeMessages]
type dlpSensitivityGroupLevelGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                        `json:"pointer"`
	JSON    dlpSensitivityGroupLevelGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelGetResponseEnvelopeMessagesSource]
type dlpSensitivityGroupLevelGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupLevelGetResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupLevelGetResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupLevelGetResponseEnvelopeSuccessTrue DLPSensitivityGroupLevelGetResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupLevelGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupLevelGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
