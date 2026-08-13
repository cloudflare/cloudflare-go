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

// DLPDataClassService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPDataClassService] method instead.
type DLPDataClassService struct {
	Options []option.RequestOption
}

// NewDLPDataClassService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPDataClassService(opts ...option.RequestOption) (r *DLPDataClassService) {
	r = &DLPDataClassService{}
	r.Options = opts
	return
}

// Creates a new data class
func (r *DLPDataClassService) New(ctx context.Context, params DLPDataClassNewParams, opts ...option.RequestOption) (res *DLPDataClassNewResponse, err error) {
	var env DLPDataClassNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_classes", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Update the attributes of a single data class
func (r *DLPDataClassService) Update(ctx context.Context, dataClassID string, params DLPDataClassUpdateParams, opts ...option.RequestOption) (res *DLPDataClassUpdateResponse, err error) {
	var env DLPDataClassUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dataClassID == "" {
		err = errors.New("missing required data_class_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_classes/%s", params.AccountID, dataClassID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve all data classes in an account
func (r *DLPDataClassService) List(ctx context.Context, query DLPDataClassListParams, opts ...option.RequestOption) (res *pagination.SinglePage[DLPDataClassListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_classes", query.AccountID)
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

// Retrieve all data classes in an account
func (r *DLPDataClassService) ListAutoPaging(ctx context.Context, query DLPDataClassListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[DLPDataClassListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Delete a single data class
func (r *DLPDataClassService) Delete(ctx context.Context, dataClassID string, body DLPDataClassDeleteParams, opts ...option.RequestOption) (res *DLPDataClassDeleteResponse, err error) {
	var env DLPDataClassDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dataClassID == "" {
		err = errors.New("missing required data_class_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_classes/%s", body.AccountID, dataClassID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve a specific data class
func (r *DLPDataClassService) Get(ctx context.Context, dataClassID string, query DLPDataClassGetParams, opts ...option.RequestOption) (res *DLPDataClassGetResponse, err error) {
	var env DLPDataClassGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if dataClassID == "" {
		err = errors.New("missing required data_class_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/data_classes/%s", query.AccountID, dataClassID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type DLPDataClassNewResponse struct {
	ID                string                                    `json:"id" api:"required" format:"uuid"`
	CreatedAt         time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	DataTags          []string                                  `json:"data_tags" api:"required" format:"uuid"`
	Expression        string                                    `json:"expression" api:"required"`
	Name              string                                    `json:"name" api:"required"`
	SensitivityLevels []DLPDataClassNewResponseSensitivityLevel `json:"sensitivity_levels" api:"required"`
	UpdatedAt         time.Time                                 `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                    `json:"description" api:"nullable"`
	JSON              dlpDataClassNewResponseJSON               `json:"-"`
}

// dlpDataClassNewResponseJSON contains the JSON metadata for the struct
// [DLPDataClassNewResponse]
type dlpDataClassNewResponseJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	DataTags          apijson.Field
	Expression        apijson.Field
	Name              apijson.Field
	SensitivityLevels apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPDataClassNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassNewResponseJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassNewResponseSensitivityLevel is a reference pairing a sensitivity group with a specific level within that group.
type DLPDataClassNewResponseSensitivityLevel struct {
	GroupID string                                      `json:"group_id" api:"required" format:"uuid"`
	LevelID string                                      `json:"level_id" api:"required" format:"uuid"`
	JSON    dlpDataClassNewResponseSensitivityLevelJSON `json:"-"`
}

// dlpDataClassNewResponseSensitivityLevelJSON contains the JSON metadata for the
// struct [DLPDataClassNewResponseSensitivityLevel]
type dlpDataClassNewResponseSensitivityLevelJSON struct {
	GroupID     apijson.Field
	LevelID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassNewResponseSensitivityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassNewResponseSensitivityLevelJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassUpdateResponse struct {
	ID                string                                       `json:"id" api:"required" format:"uuid"`
	CreatedAt         time.Time                                    `json:"created_at" api:"required" format:"date-time"`
	DataTags          []string                                     `json:"data_tags" api:"required" format:"uuid"`
	Expression        string                                       `json:"expression" api:"required"`
	Name              string                                       `json:"name" api:"required"`
	SensitivityLevels []DLPDataClassUpdateResponseSensitivityLevel `json:"sensitivity_levels" api:"required"`
	UpdatedAt         time.Time                                    `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                       `json:"description" api:"nullable"`
	JSON              dlpDataClassUpdateResponseJSON               `json:"-"`
}

// dlpDataClassUpdateResponseJSON contains the JSON metadata for the struct
// [DLPDataClassUpdateResponse]
type dlpDataClassUpdateResponseJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	DataTags          apijson.Field
	Expression        apijson.Field
	Name              apijson.Field
	SensitivityLevels apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPDataClassUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassUpdateResponseSensitivityLevel is a reference pairing a sensitivity group with a specific level within that group.
type DLPDataClassUpdateResponseSensitivityLevel struct {
	GroupID string                                         `json:"group_id" api:"required" format:"uuid"`
	LevelID string                                         `json:"level_id" api:"required" format:"uuid"`
	JSON    dlpDataClassUpdateResponseSensitivityLevelJSON `json:"-"`
}

// dlpDataClassUpdateResponseSensitivityLevelJSON contains the JSON metadata for
// the struct [DLPDataClassUpdateResponseSensitivityLevel]
type dlpDataClassUpdateResponseSensitivityLevelJSON struct {
	GroupID     apijson.Field
	LevelID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassUpdateResponseSensitivityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassUpdateResponseSensitivityLevelJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassListResponse struct {
	ID                string                                     `json:"id" api:"required" format:"uuid"`
	CreatedAt         time.Time                                  `json:"created_at" api:"required" format:"date-time"`
	DataTags          []string                                   `json:"data_tags" api:"required" format:"uuid"`
	Expression        string                                     `json:"expression" api:"required"`
	Name              string                                     `json:"name" api:"required"`
	SensitivityLevels []DLPDataClassListResponseSensitivityLevel `json:"sensitivity_levels" api:"required"`
	UpdatedAt         time.Time                                  `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                     `json:"description" api:"nullable"`
	JSON              dlpDataClassListResponseJSON               `json:"-"`
}

// dlpDataClassListResponseJSON contains the JSON metadata for the struct
// [DLPDataClassListResponse]
type dlpDataClassListResponseJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	DataTags          apijson.Field
	Expression        apijson.Field
	Name              apijson.Field
	SensitivityLevels apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPDataClassListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassListResponseJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassListResponseSensitivityLevel is a reference pairing a sensitivity group with a specific level within that group.
type DLPDataClassListResponseSensitivityLevel struct {
	GroupID string                                       `json:"group_id" api:"required" format:"uuid"`
	LevelID string                                       `json:"level_id" api:"required" format:"uuid"`
	JSON    dlpDataClassListResponseSensitivityLevelJSON `json:"-"`
}

// dlpDataClassListResponseSensitivityLevelJSON contains the JSON metadata for the
// struct [DLPDataClassListResponseSensitivityLevel]
type dlpDataClassListResponseSensitivityLevelJSON struct {
	GroupID     apijson.Field
	LevelID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassListResponseSensitivityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassListResponseSensitivityLevelJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassDeleteResponse = interface{}

type DLPDataClassGetResponse struct {
	ID                string                                    `json:"id" api:"required" format:"uuid"`
	CreatedAt         time.Time                                 `json:"created_at" api:"required" format:"date-time"`
	DataTags          []string                                  `json:"data_tags" api:"required" format:"uuid"`
	Expression        string                                    `json:"expression" api:"required"`
	Name              string                                    `json:"name" api:"required"`
	SensitivityLevels []DLPDataClassGetResponseSensitivityLevel `json:"sensitivity_levels" api:"required"`
	UpdatedAt         time.Time                                 `json:"updated_at" api:"required" format:"date-time"`
	Description       string                                    `json:"description" api:"nullable"`
	JSON              dlpDataClassGetResponseJSON               `json:"-"`
}

// dlpDataClassGetResponseJSON contains the JSON metadata for the struct
// [DLPDataClassGetResponse]
type dlpDataClassGetResponseJSON struct {
	ID                apijson.Field
	CreatedAt         apijson.Field
	DataTags          apijson.Field
	Expression        apijson.Field
	Name              apijson.Field
	SensitivityLevels apijson.Field
	UpdatedAt         apijson.Field
	Description       apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *DLPDataClassGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassGetResponseJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassGetResponseSensitivityLevel is a reference pairing a sensitivity group with a specific level within that group.
type DLPDataClassGetResponseSensitivityLevel struct {
	GroupID string                                      `json:"group_id" api:"required" format:"uuid"`
	LevelID string                                      `json:"level_id" api:"required" format:"uuid"`
	JSON    dlpDataClassGetResponseSensitivityLevelJSON `json:"-"`
}

// dlpDataClassGetResponseSensitivityLevelJSON contains the JSON metadata for the
// struct [DLPDataClassGetResponseSensitivityLevel]
type dlpDataClassGetResponseSensitivityLevelJSON struct {
	GroupID     apijson.Field
	LevelID     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassGetResponseSensitivityLevel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassGetResponseSensitivityLevelJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassNewParams struct {
	AccountID         param.Field[string]                                  `path:"account_id" api:"required"`
	DataTags          param.Field[[]string]                                `json:"data_tags" api:"required" format:"uuid"`
	Expression        param.Field[string]                                  `json:"expression" api:"required"`
	Name              param.Field[string]                                  `json:"name" api:"required"`
	SensitivityLevels param.Field[[]DLPDataClassNewParamsSensitivityLevel] `json:"sensitivity_levels" api:"required"`
	Description       param.Field[string]                                  `json:"description"`
}

func (r DLPDataClassNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLPDataClassNewParamsSensitivityLevel is a reference pairing a sensitivity group with a specific level within that group.
type DLPDataClassNewParamsSensitivityLevel struct {
	GroupID param.Field[string] `json:"group_id" api:"required" format:"uuid"`
	LevelID param.Field[string] `json:"level_id" api:"required" format:"uuid"`
}

func (r DLPDataClassNewParamsSensitivityLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataClassNewResponseEnvelope struct {
	Errors   []DLPDataClassNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataClassNewResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataClassNewResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataClassNewResponse                `json:"result"`
	JSON    dlpDataClassNewResponseEnvelopeJSON    `json:"-"`
}

// dlpDataClassNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDataClassNewResponseEnvelope]
type dlpDataClassNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassNewResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPDataClassNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataClassNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataClassNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDataClassNewResponseEnvelopeErrors]
type dlpDataClassNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassNewResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpDataClassNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataClassNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPDataClassNewResponseEnvelopeErrorsSource]
type dlpDataClassNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassNewResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DLPDataClassNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataClassNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataClassNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDataClassNewResponseEnvelopeMessages]
type dlpDataClassNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassNewResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dlpDataClassNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataClassNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPDataClassNewResponseEnvelopeMessagesSource]
type dlpDataClassNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassNewResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataClassNewResponseEnvelopeSuccess bool

const (
	DLPDataClassNewResponseEnvelopeSuccessTrue DLPDataClassNewResponseEnvelopeSuccess = true
)

func (r DLPDataClassNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataClassNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataClassUpdateParams struct {
	AccountID         param.Field[string]                                     `path:"account_id" api:"required"`
	DataTags          param.Field[[]string]                                   `json:"data_tags" format:"uuid"`
	Description       param.Field[string]                                     `json:"description"`
	Expression        param.Field[string]                                     `json:"expression"`
	Name              param.Field[string]                                     `json:"name"`
	SensitivityLevels param.Field[[]DLPDataClassUpdateParamsSensitivityLevel] `json:"sensitivity_levels"`
}

func (r DLPDataClassUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// DLPDataClassUpdateParamsSensitivityLevel is a reference pairing a sensitivity group with a specific level within that group.
type DLPDataClassUpdateParamsSensitivityLevel struct {
	GroupID param.Field[string] `json:"group_id" api:"required" format:"uuid"`
	LevelID param.Field[string] `json:"level_id" api:"required" format:"uuid"`
}

func (r DLPDataClassUpdateParamsSensitivityLevel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPDataClassUpdateResponseEnvelope struct {
	Errors   []DLPDataClassUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataClassUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataClassUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataClassUpdateResponse                `json:"result"`
	JSON    dlpDataClassUpdateResponseEnvelopeJSON    `json:"-"`
}

// dlpDataClassUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDataClassUpdateResponseEnvelope]
type dlpDataClassUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassUpdateResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPDataClassUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataClassUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataClassUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDataClassUpdateResponseEnvelopeErrors]
type dlpDataClassUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpDataClassUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataClassUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDataClassUpdateResponseEnvelopeErrorsSource]
type dlpDataClassUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassUpdateResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPDataClassUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataClassUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataClassUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDataClassUpdateResponseEnvelopeMessages]
type dlpDataClassUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpDataClassUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataClassUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPDataClassUpdateResponseEnvelopeMessagesSource]
type dlpDataClassUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassUpdateResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataClassUpdateResponseEnvelopeSuccess bool

const (
	DLPDataClassUpdateResponseEnvelopeSuccessTrue DLPDataClassUpdateResponseEnvelopeSuccess = true
)

func (r DLPDataClassUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataClassUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataClassListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataClassDeleteParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataClassDeleteResponseEnvelope struct {
	Errors   []DLPDataClassDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataClassDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataClassDeleteResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataClassDeleteResponse                `json:"result" api:"nullable"`
	JSON    dlpDataClassDeleteResponseEnvelopeJSON    `json:"-"`
}

// dlpDataClassDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDataClassDeleteResponseEnvelope]
type dlpDataClassDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassDeleteResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code" api:"required"`
	Message          string                                         `json:"message" api:"required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           DLPDataClassDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataClassDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataClassDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDataClassDeleteResponseEnvelopeErrors]
type dlpDataClassDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    dlpDataClassDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataClassDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [DLPDataClassDeleteResponseEnvelopeErrorsSource]
type dlpDataClassDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassDeleteResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code" api:"required"`
	Message          string                                           `json:"message" api:"required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           DLPDataClassDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataClassDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataClassDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [DLPDataClassDeleteResponseEnvelopeMessages]
type dlpDataClassDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    dlpDataClassDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataClassDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [DLPDataClassDeleteResponseEnvelopeMessagesSource]
type dlpDataClassDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassDeleteResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataClassDeleteResponseEnvelopeSuccess bool

const (
	DLPDataClassDeleteResponseEnvelopeSuccessTrue DLPDataClassDeleteResponseEnvelopeSuccess = true
)

func (r DLPDataClassDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataClassDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPDataClassGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPDataClassGetResponseEnvelope struct {
	Errors   []DLPDataClassGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPDataClassGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPDataClassGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  DLPDataClassGetResponse                `json:"result"`
	JSON    dlpDataClassGetResponseEnvelopeJSON    `json:"-"`
}

// dlpDataClassGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DLPDataClassGetResponseEnvelope]
type dlpDataClassGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassGetResponseEnvelopeErrors struct {
	Code             int64                                       `json:"code" api:"required"`
	Message          string                                      `json:"message" api:"required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DLPDataClassGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpDataClassGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpDataClassGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DLPDataClassGetResponseEnvelopeErrors]
type dlpDataClassGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassGetResponseEnvelopeErrorsSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dlpDataClassGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpDataClassGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [DLPDataClassGetResponseEnvelopeErrorsSource]
type dlpDataClassGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassGetResponseEnvelopeMessages struct {
	Code             int64                                         `json:"code" api:"required"`
	Message          string                                        `json:"message" api:"required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           DLPDataClassGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpDataClassGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpDataClassGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DLPDataClassGetResponseEnvelopeMessages]
type dlpDataClassGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPDataClassGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPDataClassGetResponseEnvelopeMessagesSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    dlpDataClassGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpDataClassGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DLPDataClassGetResponseEnvelopeMessagesSource]
type dlpDataClassGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPDataClassGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpDataClassGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPDataClassGetResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPDataClassGetResponseEnvelopeSuccess bool

const (
	DLPDataClassGetResponseEnvelopeSuccessTrue DLPDataClassGetResponseEnvelopeSuccess = true
)

func (r DLPDataClassGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPDataClassGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
