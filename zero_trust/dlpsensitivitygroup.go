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

// DLPSensitivityGroupService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPSensitivityGroupService] method instead.
type DLPSensitivityGroupService struct {
	Options []option.RequestOption
	Levels  *DLPSensitivityGroupLevelService
}

// NewDLPSensitivityGroupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDLPSensitivityGroupService(opts ...option.RequestOption) (r *DLPSensitivityGroupService) {
	r = &DLPSensitivityGroupService{}
	r.Options = opts
	r.Levels = NewDLPSensitivityGroupLevelService(opts...)
	return
}

// Creates a new sensitivity group.
func (r *DLPSensitivityGroupService) New(ctx context.Context, params DLPSensitivityGroupNewParams, opts ...option.RequestOption) (res *DLPSensitivityGroupNewResponse, err error) {
	var env DLPSensitivityGroupNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the attributes of a single sensitivity group.
func (r *DLPSensitivityGroupService) Update(ctx context.Context, sensitivityGroupID string, params DLPSensitivityGroupUpdateParams, opts ...option.RequestOption) (res *DLPSensitivityGroupUpdateResponse, err error) {
	var env DLPSensitivityGroupUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s", params.AccountID, sensitivityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve all sensitivity groups in an account
func (r *DLPSensitivityGroupService) List(ctx context.Context, query DLPSensitivityGroupListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPSensitivityGroupListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups", query.AccountID)
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

// Retrieve all sensitivity groups in an account
func (r *DLPSensitivityGroupService) ListAutoPaging(ctx context.Context, query DLPSensitivityGroupListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPSensitivityGroupListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a single sensitivity group.
func (r *DLPSensitivityGroupService) Delete(ctx context.Context, sensitivityGroupID string, body DLPSensitivityGroupDeleteParams, opts ...option.RequestOption) (res *DLPSensitivityGroupDeleteResponse, err error) {
	var env DLPSensitivityGroupDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s", body.AccountID, sensitivityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a specific sensitivity group.
func (r *DLPSensitivityGroupService) Get(ctx context.Context, sensitivityGroupID string, query DLPSensitivityGroupGetParams, opts ...option.RequestOption) (res *DLPSensitivityGroupGetResponse, err error) {
	var env DLPSensitivityGroupGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s", query.AccountID, sensitivityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DLPSensitivityGroupNewResponse struct {
	ID          string                                `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                             `json:"created_at" api:"required" format:"date-time"`
	Levels      []DLPSensitivityGroupNewResponseLevel `json:"levels" api:"required"`
	Name        string                                `json:"name" api:"required"`
	UpdatedAt   time.Time                             `json:"updated_at" api:"required" format:"date-time"`
	Description string                                `json:"description" api:"nullable"`
	TemplateID  string                                `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpSensitivityGroupNewResponseJSON    `json:"-"`
}

// dlpSensitivityGroupNewResponseJSON contains the JSON metadata for the struct
// [DLPSensitivityGroupNewResponse]
type dlpSensitivityGroupNewResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Levels      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupNewResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupNewResponseLevel struct {
	ID          string                                  `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Name        string                                  `json:"name" api:"required"`
	UpdatedAt   time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Description string                                  `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupNewResponseLevelJSON `json:"-"`
}

// dlpSensitivityGroupNewResponseLevelJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupNewResponseLevel]
type dlpSensitivityGroupNewResponseLevelJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupNewResponseLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupNewResponseLevelJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupUpdateResponse struct {
	ID          string                                   `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Levels      []DLPSensitivityGroupUpdateResponseLevel `json:"levels" api:"required"`
	Name        string                                   `json:"name" api:"required"`
	UpdatedAt   time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	Description string                                   `json:"description" api:"nullable"`
	TemplateID  string                                   `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpSensitivityGroupUpdateResponseJSON    `json:"-"`
}

// dlpSensitivityGroupUpdateResponseJSON contains the JSON metadata for the struct
// [DLPSensitivityGroupUpdateResponse]
type dlpSensitivityGroupUpdateResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Levels      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupUpdateResponseLevel struct {
	ID          string                                     `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	Name        string                                     `json:"name" api:"required"`
	UpdatedAt   time.Time                                  `json:"updated_at" api:"required" format:"date-time"`
	Description string                                     `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupUpdateResponseLevelJSON `json:"-"`
}

// dlpSensitivityGroupUpdateResponseLevelJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupUpdateResponseLevel]
type dlpSensitivityGroupUpdateResponseLevelJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupUpdateResponseLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupUpdateResponseLevelJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupListResponse struct {
	ID          string                                 `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                              `json:"created_at" api:"required" format:"date-time"`
	Levels      []DLPSensitivityGroupListResponseLevel `json:"levels" api:"required"`
	Name        string                                 `json:"name" api:"required"`
	UpdatedAt   time.Time                              `json:"updated_at" api:"required" format:"date-time"`
	Description string                                 `json:"description" api:"nullable"`
	TemplateID  string                                 `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpSensitivityGroupListResponseJSON    `json:"-"`
}

// dlpSensitivityGroupListResponseJSON contains the JSON metadata for the struct
// [DLPSensitivityGroupListResponse]
type dlpSensitivityGroupListResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Levels      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupListResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupListResponseLevel struct {
	ID          string                                   `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                                `json:"created_at" api:"required" format:"date-time"`
	Name        string                                   `json:"name" api:"required"`
	UpdatedAt   time.Time                                `json:"updated_at" api:"required" format:"date-time"`
	Description string                                   `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupListResponseLevelJSON `json:"-"`
}

// dlpSensitivityGroupListResponseLevelJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupListResponseLevel]
type dlpSensitivityGroupListResponseLevelJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupListResponseLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupListResponseLevelJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupDeleteResponse = interface{}

type DLPSensitivityGroupGetResponse struct {
	ID          string                                `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                             `json:"created_at" api:"required" format:"date-time"`
	Levels      []DLPSensitivityGroupGetResponseLevel `json:"levels" api:"required"`
	Name        string                                `json:"name" api:"required"`
	UpdatedAt   time.Time                             `json:"updated_at" api:"required" format:"date-time"`
	Description string                                `json:"description" api:"nullable"`
	TemplateID  string                                `json:"template_id" api:"nullable" format:"uuid"`
	JSON        dlpSensitivityGroupGetResponseJSON    `json:"-"`
}

// dlpSensitivityGroupGetResponseJSON contains the JSON metadata for the struct
// [DLPSensitivityGroupGetResponse]
type dlpSensitivityGroupGetResponseJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Levels      apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	TemplateID  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupGetResponseLevel struct {
	ID          string                                  `json:"id" api:"required" format:"uuid"`
	CreatedAt   time.Time                               `json:"created_at" api:"required" format:"date-time"`
	Name        string                                  `json:"name" api:"required"`
	UpdatedAt   time.Time                               `json:"updated_at" api:"required" format:"date-time"`
	Description string                                  `json:"description" api:"nullable"`
	JSON        dlpSensitivityGroupGetResponseLevelJSON `json:"-"`
}

// dlpSensitivityGroupGetResponseLevelJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupGetResponseLevel]
type dlpSensitivityGroupGetResponseLevelJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	Description apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupGetResponseLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupGetResponseLevelJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupNewParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
	// Levels to create with the group. Mutually exclusive with `template_id`.
	Levels     param.Field[[]DLPSensitivityGroupNewParamsLevel] `json:"levels"`
	TemplateID param.Field[string]                              `json:"template_id" format:"uuid"`
}

func (r DLPSensitivityGroupNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPSensitivityGroupNewParamsLevel struct {
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
}

func (r DLPSensitivityGroupNewParamsLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPSensitivityGroupNewResponseEnvelope struct {
	Errors   []DLPSensitivityGroupNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupNewResponse                `json:"result"`
	JSON    dlpSensitivityGroupNewResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupNewResponseEnvelope]
type dlpSensitivityGroupNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupNewResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPSensitivityGroupNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPSensitivityGroupNewResponseEnvelopeErrors]
type dlpSensitivityGroupNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupNewResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpSensitivityGroupNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupNewResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupNewResponseEnvelopeErrorsSource]
type dlpSensitivityGroupNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupNewResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPSensitivityGroupNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupNewResponseEnvelopeMessages]
type dlpSensitivityGroupNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupNewResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpSensitivityGroupNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupNewResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupNewResponseEnvelopeMessagesSource]
type dlpSensitivityGroupNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupNewResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupNewResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupNewResponseEnvelopeSuccessTrue DLPSensitivityGroupNewResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSensitivityGroupUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description"`
	// The desired final state of levels.
	//
	// - `None` (omitted): no level changes.
	// - `Some([])`: delete all levels.
	// - `Some([...])`: desired final set + order.
	Levels param.Field[[]DLPSensitivityGroupUpdateParamsLevel] `json:"levels"`
	Name   param.Field[string]                                 `json:"name"`
}

func (r DLPSensitivityGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPSensitivityGroupUpdateParamsLevel struct {
	// If `None` (omitted), a new level will be created. Otherwise, an existing level
	// will be updated.
	ID          param.Field[string] `json:"id" format:"uuid"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
}

func (r DLPSensitivityGroupUpdateParamsLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPSensitivityGroupUpdateResponseEnvelope struct {
	Errors   []DLPSensitivityGroupUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupUpdateResponse                `json:"result"`
	JSON    dlpSensitivityGroupUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupUpdateResponseEnvelope]
type dlpSensitivityGroupUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupUpdateResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPSensitivityGroupUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupUpdateResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupUpdateResponseEnvelopeErrors]
type dlpSensitivityGroupUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpSensitivityGroupUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupUpdateResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupUpdateResponseEnvelopeErrorsSource]
type dlpSensitivityGroupUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupUpdateResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPSensitivityGroupUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupUpdateResponseEnvelopeMessages]
type dlpSensitivityGroupUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpSensitivityGroupUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupUpdateResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupUpdateResponseEnvelopeMessagesSource]
type dlpSensitivityGroupUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupUpdateResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupUpdateResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupUpdateResponseEnvelopeSuccessTrue DLPSensitivityGroupUpdateResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSensitivityGroupListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSensitivityGroupDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSensitivityGroupDeleteResponseEnvelope struct {
	Errors   []DLPSensitivityGroupDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupDeleteResponse                `json:"result" api:"nullable"`
	JSON    dlpSensitivityGroupDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupDeleteResponseEnvelope]
type dlpSensitivityGroupDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupDeleteResponseEnvelopeErrors struct {
	Code             int64                                                 `json:"code" api:"required"`
	Message          string                                                `json:"message" api:"required"`
	DocumentationURL string                                                `json:"documentation_url"`
	Source           DLPSensitivityGroupDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupDeleteResponseEnvelopeErrors]
type dlpSensitivityGroupDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                    `json:"pointer"`
	JSON    dlpSensitivityGroupDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupDeleteResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupDeleteResponseEnvelopeErrorsSource]
type dlpSensitivityGroupDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupDeleteResponseEnvelopeMessages struct {
	Code             int64                                                   `json:"code" api:"required"`
	Message          string                                                  `json:"message" api:"required"`
	DocumentationURL string                                                  `json:"documentation_url"`
	Source           DLPSensitivityGroupDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupDeleteResponseEnvelopeMessages]
type dlpSensitivityGroupDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                      `json:"pointer"`
	JSON    dlpSensitivityGroupDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupDeleteResponseEnvelopeMessagesSource]
type dlpSensitivityGroupDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupDeleteResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupDeleteResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupDeleteResponseEnvelopeSuccessTrue DLPSensitivityGroupDeleteResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSensitivityGroupGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSensitivityGroupGetResponseEnvelope struct {
	Errors   []DLPSensitivityGroupGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPSensitivityGroupGetResponse                `json:"result"`
	JSON    dlpSensitivityGroupGetResponseEnvelopeJSON    `json:"-"`
}

// dlpSensitivityGroupGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupGetResponseEnvelope]
type dlpSensitivityGroupGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupGetResponseEnvelopeErrors struct {
	Code             int64                                              `json:"code" api:"required"`
	Message          string                                             `json:"message" api:"required"`
	DocumentationURL string                                             `json:"documentation_url"`
	Source           DLPSensitivityGroupGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [DLPSensitivityGroupGetResponseEnvelopeErrors]
type dlpSensitivityGroupGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                 `json:"pointer"`
	JSON    dlpSensitivityGroupGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupGetResponseEnvelopeErrorsSourceJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupGetResponseEnvelopeErrorsSource]
type dlpSensitivityGroupGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupGetResponseEnvelopeMessages struct {
	Code             int64                                                `json:"code" api:"required"`
	Message          string                                               `json:"message" api:"required"`
	DocumentationURL string                                               `json:"documentation_url"`
	Source           DLPSensitivityGroupGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupGetResponseEnvelopeMessages]
type dlpSensitivityGroupGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                   `json:"pointer"`
	JSON    dlpSensitivityGroupGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupGetResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupGetResponseEnvelopeMessagesSource]
type dlpSensitivityGroupGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupGetResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupGetResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupGetResponseEnvelopeSuccessTrue DLPSensitivityGroupGetResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
