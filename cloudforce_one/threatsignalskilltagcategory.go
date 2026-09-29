// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

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

// ThreatSignalSkillTagCategoryService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalSkillTagCategoryService] method instead.
type ThreatSignalSkillTagCategoryService struct {
	Options []option.RequestOption
}

// NewThreatSignalSkillTagCategoryService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewThreatSignalSkillTagCategoryService(opts ...option.RequestOption) (r *ThreatSignalSkillTagCategoryService) {
	r = &ThreatSignalSkillTagCategoryService{}
	r.Options = opts
	return
}

// Replaces the tag categories the default tagging skill may choose tags from.
func (r *ThreatSignalSkillTagCategoryService) Update(ctx context.Context, skillID ThreatSignalSkillTagCategoryUpdateParamsSkillID, params ThreatSignalSkillTagCategoryUpdateParams, opts ...option.RequestOption) (res *ThreatSignalSkillTagCategoryUpdateResponse, err error) {
	var env ThreatSignalSkillTagCategoryUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/skills/%v/tag-categories", params.AccountID, skillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieves the tag categories the default tagging skill may choose tags from.
func (r *ThreatSignalSkillTagCategoryService) Get(ctx context.Context, skillID ThreatSignalSkillTagCategoryGetParamsSkillID, query ThreatSignalSkillTagCategoryGetParams, opts ...option.RequestOption) (res *ThreatSignalSkillTagCategoryGetResponse, err error) {
	var env ThreatSignalSkillTagCategoryGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/skills/%v/tag-categories", query.AccountID, skillID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalSkillTagCategoryUpdateResponse struct {
	CategoryUUIDs []string                                          `json:"category_uuids" api:"required" format:"uuid"`
	SkillID       ThreatSignalSkillTagCategoryUpdateResponseSkillID `json:"skill_id" api:"required"`
	JSON          threatSignalSkillTagCategoryUpdateResponseJSON    `json:"-"`
}

// threatSignalSkillTagCategoryUpdateResponseJSON contains the JSON metadata for
// the struct [ThreatSignalSkillTagCategoryUpdateResponse]
type threatSignalSkillTagCategoryUpdateResponseJSON struct {
	CategoryUUIDs apijson.Field
	SkillID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatSignalSkillTagCategoryUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillTagCategoryUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillTagCategoryUpdateResponseSkillID string

const (
	ThreatSignalSkillTagCategoryUpdateResponseSkillIDDefaultTaggingSkill ThreatSignalSkillTagCategoryUpdateResponseSkillID = "default-tagging-skill"
)

func (r ThreatSignalSkillTagCategoryUpdateResponseSkillID) IsKnown() bool {
	switch r {
	case ThreatSignalSkillTagCategoryUpdateResponseSkillIDDefaultTaggingSkill:
		return true
	}
	return false
}

type ThreatSignalSkillTagCategoryGetResponse struct {
	CategoryUUIDs []string                                       `json:"category_uuids" api:"required" format:"uuid"`
	SkillID       ThreatSignalSkillTagCategoryGetResponseSkillID `json:"skill_id" api:"required"`
	JSON          threatSignalSkillTagCategoryGetResponseJSON    `json:"-"`
}

// threatSignalSkillTagCategoryGetResponseJSON contains the JSON metadata for the
// struct [ThreatSignalSkillTagCategoryGetResponse]
type threatSignalSkillTagCategoryGetResponseJSON struct {
	CategoryUUIDs apijson.Field
	SkillID       apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *ThreatSignalSkillTagCategoryGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillTagCategoryGetResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillTagCategoryGetResponseSkillID string

const (
	ThreatSignalSkillTagCategoryGetResponseSkillIDDefaultTaggingSkill ThreatSignalSkillTagCategoryGetResponseSkillID = "default-tagging-skill"
)

func (r ThreatSignalSkillTagCategoryGetResponseSkillID) IsKnown() bool {
	switch r {
	case ThreatSignalSkillTagCategoryGetResponseSkillIDDefaultTaggingSkill:
		return true
	}
	return false
}

type ThreatSignalSkillTagCategoryUpdateParams struct {
	AccountID     param.Field[string]   `path:"account_id" api:"required"`
	CategoryUUIDs param.Field[[]string] `json:"category_uuids" api:"required" format:"uuid"`
}

func (r ThreatSignalSkillTagCategoryUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatSignalSkillTagCategoryUpdateParamsSkillID string

const (
	ThreatSignalSkillTagCategoryUpdateParamsSkillIDDefaultTaggingSkill ThreatSignalSkillTagCategoryUpdateParamsSkillID = "default-tagging-skill"
)

func (r ThreatSignalSkillTagCategoryUpdateParamsSkillID) IsKnown() bool {
	switch r {
	case ThreatSignalSkillTagCategoryUpdateParamsSkillIDDefaultTaggingSkill:
		return true
	}
	return false
}

type ThreatSignalSkillTagCategoryUpdateResponseEnvelope struct {
	Errors  []ThreatSignalSkillTagCategoryUpdateResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalSkillTagCategoryUpdateResponse                 `json:"result" api:"required"`
	Success ThreatSignalSkillTagCategoryUpdateResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalSkillTagCategoryUpdateResponseEnvelopeJSON     `json:"-"`
}

// threatSignalSkillTagCategoryUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [ThreatSignalSkillTagCategoryUpdateResponseEnvelope]
type threatSignalSkillTagCategoryUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillTagCategoryUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillTagCategoryUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillTagCategoryUpdateResponseEnvelopeErrors struct {
	Message string                                                       `json:"message" api:"required"`
	JSON    threatSignalSkillTagCategoryUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalSkillTagCategoryUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ThreatSignalSkillTagCategoryUpdateResponseEnvelopeErrors]
type threatSignalSkillTagCategoryUpdateResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillTagCategoryUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillTagCategoryUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillTagCategoryUpdateResponseEnvelopeSuccess bool

const (
	ThreatSignalSkillTagCategoryUpdateResponseEnvelopeSuccessTrue ThreatSignalSkillTagCategoryUpdateResponseEnvelopeSuccess = true
)

func (r ThreatSignalSkillTagCategoryUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalSkillTagCategoryUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ThreatSignalSkillTagCategoryGetParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalSkillTagCategoryGetParamsSkillID string

const (
	ThreatSignalSkillTagCategoryGetParamsSkillIDDefaultTaggingSkill ThreatSignalSkillTagCategoryGetParamsSkillID = "default-tagging-skill"
)

func (r ThreatSignalSkillTagCategoryGetParamsSkillID) IsKnown() bool {
	switch r {
	case ThreatSignalSkillTagCategoryGetParamsSkillIDDefaultTaggingSkill:
		return true
	}
	return false
}

type ThreatSignalSkillTagCategoryGetResponseEnvelope struct {
	Errors  []ThreatSignalSkillTagCategoryGetResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalSkillTagCategoryGetResponse                 `json:"result" api:"required"`
	Success ThreatSignalSkillTagCategoryGetResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalSkillTagCategoryGetResponseEnvelopeJSON     `json:"-"`
}

// threatSignalSkillTagCategoryGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ThreatSignalSkillTagCategoryGetResponseEnvelope]
type threatSignalSkillTagCategoryGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillTagCategoryGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillTagCategoryGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillTagCategoryGetResponseEnvelopeErrors struct {
	Message string                                                    `json:"message" api:"required"`
	JSON    threatSignalSkillTagCategoryGetResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalSkillTagCategoryGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ThreatSignalSkillTagCategoryGetResponseEnvelopeErrors]
type threatSignalSkillTagCategoryGetResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalSkillTagCategoryGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalSkillTagCategoryGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalSkillTagCategoryGetResponseEnvelopeSuccess bool

const (
	ThreatSignalSkillTagCategoryGetResponseEnvelopeSuccessTrue ThreatSignalSkillTagCategoryGetResponseEnvelopeSuccess = true
)

func (r ThreatSignalSkillTagCategoryGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalSkillTagCategoryGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
