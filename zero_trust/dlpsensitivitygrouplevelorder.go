// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// DLPSensitivityGroupLevelOrderService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDLPSensitivityGroupLevelOrderService] method instead.
type DLPSensitivityGroupLevelOrderService struct {
	Options []option.RequestOption
}

// NewDLPSensitivityGroupLevelOrderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDLPSensitivityGroupLevelOrderService(opts ...option.RequestOption) (r *DLPSensitivityGroupLevelOrderService) {
	r = &DLPSensitivityGroupLevelOrderService{}
	r.Options = opts
	return
}

// Set the ordering of levels within a sensitivity group.
func (r *DLPSensitivityGroupLevelOrderService) Update(ctx context.Context, sensitivityGroupID string, params DLPSensitivityGroupLevelOrderUpdateParams, opts ...option.RequestOption) (res *DLPSensitivityGroupLevelOrderUpdateResponse, err error) {
	var env DLPSensitivityGroupLevelOrderUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s/level_order", params.AccountID, sensitivityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve the ordered list of level IDs for a sensitivity group.
func (r *DLPSensitivityGroupLevelOrderService) Get(ctx context.Context, sensitivityGroupID string, query DLPSensitivityGroupLevelOrderGetParams, opts ...option.RequestOption) (res *DLPSensitivityGroupLevelOrderGetResponse, err error) {
	var env DLPSensitivityGroupLevelOrderGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sensitivityGroupID == "" {
		err = errors.New("missing required sensitivity_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/dlp/sensitivity_groups/%s/level_order", query.AccountID, sensitivityGroupID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// DLPSensitivityGroupLevelOrderUpdateResponse is the ordered list of level IDs for a sensitivity group. Used to get and set the
// ordering of levels independently of level attributes.
type DLPSensitivityGroupLevelOrderUpdateResponse struct {
	LevelIDs []string                                        `json:"level_ids" api:"required" format:"uuid"`
	JSON     dlpSensitivityGroupLevelOrderUpdateResponseJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderUpdateResponseJSON contains the JSON metadata for
// the struct [DLPSensitivityGroupLevelOrderUpdateResponse]
type dlpSensitivityGroupLevelOrderUpdateResponseJSON struct {
	LevelIDs    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupLevelOrderGetResponse is the ordered list of level IDs for a sensitivity group. Used to get and set the
// ordering of levels independently of level attributes.
type DLPSensitivityGroupLevelOrderGetResponse struct {
	LevelIDs []string                                     `json:"level_ids" api:"required" format:"uuid"`
	JSON     dlpSensitivityGroupLevelOrderGetResponseJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderGetResponseJSON contains the JSON metadata for the
// struct [DLPSensitivityGroupLevelOrderGetResponse]
type dlpSensitivityGroupLevelOrderGetResponseJSON struct {
	LevelIDs    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderGetResponseJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderUpdateParams struct {
	AccountID param.Field[string]   `path:"account_id" api:"required"`
	LevelIDs  param.Field[[]string] `json:"level_ids" api:"required" format:"uuid"`
}

func (r DLPSensitivityGroupLevelOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DLPSensitivityGroupLevelOrderUpdateResponseEnvelope struct {
	Errors   []DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	// The ordered list of level IDs for a sensitivity group. Used to get and set the
	// ordering of levels independently of level attributes.
	Result DLPSensitivityGroupLevelOrderUpdateResponse             `json:"result"`
	JSON   dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelOrderUpdateResponseEnvelope]
type dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrors struct {
	Code             int64                                                           `json:"code" api:"required"`
	Message          string                                                          `json:"message" api:"required"`
	DocumentationURL string                                                          `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrors]
type dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                                              `json:"pointer"`
	JSON    dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSource]
type dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessages struct {
	Code             int64                                                             `json:"code" api:"required"`
	Message          string                                                            `json:"message" api:"required"`
	DocumentationURL string                                                            `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesJSON contains the
// JSON metadata for the struct
// [DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessages]
type dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                                                `json:"pointer"`
	JSON    dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSourceJSON contains
// the JSON metadata for the struct
// [DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSource]
type dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeSuccessTrue DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupLevelOrderUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DLPSensitivityGroupLevelOrderGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DLPSensitivityGroupLevelOrderGetResponseEnvelope struct {
	Errors   []DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success DLPSensitivityGroupLevelOrderGetResponseEnvelopeSuccess `json:"success" api:"required"`
	// The ordered list of level IDs for a sensitivity group. Used to get and set the
	// ordering of levels independently of level attributes.
	Result DLPSensitivityGroupLevelOrderGetResponse             `json:"result"`
	JSON   dlpSensitivityGroupLevelOrderGetResponseEnvelopeJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [DLPSensitivityGroupLevelOrderGetResponseEnvelope]
type dlpSensitivityGroupLevelOrderGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrors struct {
	Code             int64                                                        `json:"code" api:"required"`
	Message          string                                                       `json:"message" api:"required"`
	DocumentationURL string                                                       `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrors]
type dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSource struct {
	Pointer string                                                           `json:"pointer"`
	JSON    dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSourceJSON contains the
// JSON metadata for the struct
// [DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSource]
type dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessages struct {
	Code             int64                                                          `json:"code" api:"required"`
	Message          string                                                         `json:"message" api:"required"`
	DocumentationURL string                                                         `json:"documentation_url"`
	Source           DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessages]
type dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSource struct {
	Pointer string                                                             `json:"pointer"`
	JSON    dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSourceJSON contains the
// JSON metadata for the struct
// [DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSource]
type dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dlpSensitivityGroupLevelOrderGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// DLPSensitivityGroupLevelOrderGetResponseEnvelopeSuccess indicates whether the API call was successful.
type DLPSensitivityGroupLevelOrderGetResponseEnvelopeSuccess bool

const (
	DLPSensitivityGroupLevelOrderGetResponseEnvelopeSuccessTrue DLPSensitivityGroupLevelOrderGetResponseEnvelopeSuccess = true
)

func (r DLPSensitivityGroupLevelOrderGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DLPSensitivityGroupLevelOrderGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
