// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package api_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// LabelUserResourceOperationService contains methods and other services that help
// with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLabelUserResourceOperationService] method instead.
type LabelUserResourceOperationService struct {
	Options []option.RequestOption
}

// NewLabelUserResourceOperationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLabelUserResourceOperationService(opts ...option.RequestOption) (r *LabelUserResourceOperationService) {
	r = &LabelUserResourceOperationService{}
	r.Options = opts
	return
}

// Replace all operations(s) attached to a user label
func (r *LabelUserResourceOperationService) Update(ctx context.Context, name string, params LabelUserResourceOperationUpdateParams, opts ...option.RequestOption) (res *LabelUserResourceOperationUpdateResponse, err error) {
	var env LabelUserResourceOperationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/user/%s/resources/operation", params.ZoneID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type LabelUserResourceOperationUpdateResponse struct {
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
	Source LabelUserResourceOperationUpdateResponseSource `json:"source" api:"required"`
	// Provides counts of what resources are linked to this label
	MappedResources interface{}                                  `json:"mapped_resources"`
	JSON            labelUserResourceOperationUpdateResponseJSON `json:"-"`
}

// labelUserResourceOperationUpdateResponseJSON contains the JSON metadata for the
// struct [LabelUserResourceOperationUpdateResponse]
type labelUserResourceOperationUpdateResponseJSON struct {
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

func (r *LabelUserResourceOperationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserResourceOperationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelUserResourceOperationUpdateResponseSource string

const (
	LabelUserResourceOperationUpdateResponseSourceUser    LabelUserResourceOperationUpdateResponseSource = "user"
	LabelUserResourceOperationUpdateResponseSourceManaged LabelUserResourceOperationUpdateResponseSource = "managed"
)

func (r LabelUserResourceOperationUpdateResponseSource) IsKnown() bool {
	switch r {
	case LabelUserResourceOperationUpdateResponseSourceUser, LabelUserResourceOperationUpdateResponseSourceManaged:
		return true
	}
	return false
}

type LabelUserResourceOperationUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
	// Operation IDs selector
	Selector param.Field[LabelUserResourceOperationUpdateParamsSelector] `json:"selector" api:"required"`
}

func (r LabelUserResourceOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Operation IDs selector
type LabelUserResourceOperationUpdateParamsSelector struct {
	Include param.Field[LabelUserResourceOperationUpdateParamsSelectorInclude] `json:"include" api:"required"`
}

func (r LabelUserResourceOperationUpdateParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LabelUserResourceOperationUpdateParamsSelectorInclude struct {
	OperationIDs param.Field[[]string] `json:"operation_ids" api:"required"`
}

func (r LabelUserResourceOperationUpdateParamsSelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LabelUserResourceOperationUpdateResponseEnvelope struct {
	Errors   Message                                  `json:"errors" api:"required"`
	Messages Message                                  `json:"messages" api:"required"`
	Result   LabelUserResourceOperationUpdateResponse `json:"result" api:"required"`
	// Whether the API call was successful.
	Success LabelUserResourceOperationUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	JSON    labelUserResourceOperationUpdateResponseEnvelopeJSON    `json:"-"`
}

// labelUserResourceOperationUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [LabelUserResourceOperationUpdateResponseEnvelope]
type labelUserResourceOperationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelUserResourceOperationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelUserResourceOperationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LabelUserResourceOperationUpdateResponseEnvelopeSuccess bool

const (
	LabelUserResourceOperationUpdateResponseEnvelopeSuccessTrue LabelUserResourceOperationUpdateResponseEnvelopeSuccess = true
)

func (r LabelUserResourceOperationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LabelUserResourceOperationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
