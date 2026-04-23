// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// LabelManagedService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLabelManagedService] method instead.
type LabelManagedService struct {
	Options   []option.RequestOption
	Resources *LabelManagedResourceService
}

// NewLabelManagedService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLabelManagedService(opts ...option.RequestOption) (r *LabelManagedService) {
	r = &LabelManagedService{}
	r.Options = opts
	r.Resources = NewLabelManagedResourceService(opts...)
	return
}

// Retrieve managed label
func (r *LabelManagedService) Get(ctx context.Context, name string, params LabelManagedGetParams, opts ...option.RequestOption) (res *LabelManagedGetResponse, err error) {
	var env LabelManagedGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/managed/%s", params.ZoneID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type LabelManagedGetResponse struct {
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// The description of the label
	Description string    `json:"description" api:"required"`
	LastUpdated time.Time `json:"last_updated" api:"required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata" api:"required"`
	// The name of the label
	Name string `json:"name" api:"required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelManagedGetResponseSource `json:"source" api:"required"`
	// Provides counts of what resources are linked to this label
	MappedResources interface{}                 `json:"mapped_resources"`
	JSON            labelManagedGetResponseJSON `json:"-"`
}

// labelManagedGetResponseJSON contains the JSON metadata for the struct
// [LabelManagedGetResponse]
type labelManagedGetResponseJSON struct {
	CreatedAt       apijson.Field
	Description     apijson.Field
	LastUpdated     apijson.Field
	Metadata        apijson.Field
	Name            apijson.Field
	Source          apijson.Field
	MappedResources apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *LabelManagedGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelManagedGetResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelManagedGetResponseSource string

const (
	LabelManagedGetResponseSourceUser    LabelManagedGetResponseSource = "user"
	LabelManagedGetResponseSourceManaged LabelManagedGetResponseSource = "managed"
)

func (r LabelManagedGetResponseSource) IsKnown() bool {
	switch r {
	case LabelManagedGetResponseSourceUser, LabelManagedGetResponseSourceManaged:
		return true
	}
	return false
}

type LabelManagedGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Include `mapped_resources` for each label
	WithMappedResourceCounts param.Field[bool] `query:"with_mapped_resource_counts"`
}

// URLQuery serializes [LabelManagedGetParams]'s query parameters as `url.Values`.
func (r LabelManagedGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type LabelManagedGetResponseEnvelope struct {
	Errors   Message                 `json:"errors" api:"required"`
	Messages Message                 `json:"messages" api:"required"`
	Result   LabelManagedGetResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success LabelManagedGetResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    labelManagedGetResponseEnvelopeJSON    `json:"-"`
}

// labelManagedGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [LabelManagedGetResponseEnvelope]
type labelManagedGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelManagedGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelManagedGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LabelManagedGetResponseEnvelopeSuccess bool

const (
	LabelManagedGetResponseEnvelopeSuccessTrue LabelManagedGetResponseEnvelopeSuccess = true
)

func (r LabelManagedGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LabelManagedGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
