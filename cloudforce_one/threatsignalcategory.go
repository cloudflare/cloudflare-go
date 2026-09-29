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

// ThreatSignalCategoryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatSignalCategoryService] method instead.
type ThreatSignalCategoryService struct {
	Options []option.RequestOption
}

// NewThreatSignalCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatSignalCategoryService(opts ...option.RequestOption) (r *ThreatSignalCategoryService) {
	r = &ThreatSignalCategoryService{}
	r.Options = opts
	return
}

// Lists the predefined categories that can be assigned to feeds.
func (r *ThreatSignalCategoryService) List(ctx context.Context, query ThreatSignalCategoryListParams, opts ...option.RequestOption) (res *ThreatSignalCategoryListResponse, err error) {
	var env ThreatSignalCategoryListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/v2/threat-signals/categories", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ThreatSignalCategoryListResponse struct {
	Categories []ThreatSignalCategoryListResponseCategory `json:"categories" api:"required"`
	JSON       threatSignalCategoryListResponseJSON       `json:"-"`
}

// threatSignalCategoryListResponseJSON contains the JSON metadata for the struct
// [ThreatSignalCategoryListResponse]
type threatSignalCategoryListResponseJSON struct {
	Categories  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalCategoryListResponseCategory struct {
	// Wire value accepted by the feed `category_id` field.
	ID string `json:"id" api:"required" format:"uuid"`
	// Plain-language description of the category.
	Description string `json:"description" api:"required"`
	// Human-readable display label.
	Name string                                       `json:"name" api:"required"`
	JSON threatSignalCategoryListResponseCategoryJSON `json:"-"`
}

// threatSignalCategoryListResponseCategoryJSON contains the JSON metadata for the
// struct [ThreatSignalCategoryListResponseCategory]
type threatSignalCategoryListResponseCategoryJSON struct {
	ID          apijson.Field
	Description apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalCategoryListResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalCategoryListResponseCategoryJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalCategoryListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatSignalCategoryListResponseEnvelope struct {
	Errors  []ThreatSignalCategoryListResponseEnvelopeErrors `json:"errors" api:"required"`
	Result  ThreatSignalCategoryListResponse                 `json:"result" api:"required"`
	Success ThreatSignalCategoryListResponseEnvelopeSuccess  `json:"success" api:"required"`
	JSON    threatSignalCategoryListResponseEnvelopeJSON     `json:"-"`
}

// threatSignalCategoryListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ThreatSignalCategoryListResponseEnvelope]
type threatSignalCategoryListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalCategoryListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalCategoryListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalCategoryListResponseEnvelopeErrors struct {
	Message string                                             `json:"message" api:"required"`
	JSON    threatSignalCategoryListResponseEnvelopeErrorsJSON `json:"-"`
}

// threatSignalCategoryListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ThreatSignalCategoryListResponseEnvelopeErrors]
type threatSignalCategoryListResponseEnvelopeErrorsJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatSignalCategoryListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatSignalCategoryListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ThreatSignalCategoryListResponseEnvelopeSuccess bool

const (
	ThreatSignalCategoryListResponseEnvelopeSuccessTrue ThreatSignalCategoryListResponseEnvelopeSuccess = true
)

func (r ThreatSignalCategoryListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ThreatSignalCategoryListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
