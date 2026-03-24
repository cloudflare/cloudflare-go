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

// LabelManagedResourceOperationService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLabelManagedResourceOperationService] method instead.
type LabelManagedResourceOperationService struct {
	Options []option.RequestOption
}

// NewLabelManagedResourceOperationService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewLabelManagedResourceOperationService(opts ...option.RequestOption) (r *LabelManagedResourceOperationService) {
	r = &LabelManagedResourceOperationService{}
	r.Options = opts
	return
}

// Replace all operations(s) attached to a managed label
func (r *LabelManagedResourceOperationService) Update(ctx context.Context, name string, params LabelManagedResourceOperationUpdateParams, opts ...option.RequestOption) (res *LabelManagedResourceOperationUpdateResponse, err error) {
	var env LabelManagedResourceOperationUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/api_gateway/labels/managed/%s/resources/operation", params.ZoneID, name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type LabelManagedResourceOperationUpdateResponse struct {
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// The description of the label
	Description string    `json:"description,required"`
	LastUpdated time.Time `json:"last_updated,required" format:"date-time"`
	// Metadata for the label
	Metadata interface{} `json:"metadata,required"`
	// The name of the label
	Name string `json:"name,required"`
	// - `user` - label is owned by the user
	// - `managed` - label is owned by cloudflare
	Source LabelManagedResourceOperationUpdateResponseSource `json:"source,required"`
	// Provides counts of what resources are linked to this label
	MappedResources interface{}                                     `json:"mapped_resources"`
	JSON            labelManagedResourceOperationUpdateResponseJSON `json:"-"`
}

// labelManagedResourceOperationUpdateResponseJSON contains the JSON metadata for
// the struct [LabelManagedResourceOperationUpdateResponse]
type labelManagedResourceOperationUpdateResponseJSON struct {
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

func (r *LabelManagedResourceOperationUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelManagedResourceOperationUpdateResponseJSON) RawJSON() string {
	return r.raw
}

// - `user` - label is owned by the user
// - `managed` - label is owned by cloudflare
type LabelManagedResourceOperationUpdateResponseSource string

const (
	LabelManagedResourceOperationUpdateResponseSourceUser    LabelManagedResourceOperationUpdateResponseSource = "user"
	LabelManagedResourceOperationUpdateResponseSourceManaged LabelManagedResourceOperationUpdateResponseSource = "managed"
)

func (r LabelManagedResourceOperationUpdateResponseSource) IsKnown() bool {
	switch r {
	case LabelManagedResourceOperationUpdateResponseSourceUser, LabelManagedResourceOperationUpdateResponseSourceManaged:
		return true
	}
	return false
}

type LabelManagedResourceOperationUpdateParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Operation IDs selector
	Selector param.Field[LabelManagedResourceOperationUpdateParamsSelector] `json:"selector,required"`
}

func (r LabelManagedResourceOperationUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Operation IDs selector
type LabelManagedResourceOperationUpdateParamsSelector struct {
	Include param.Field[LabelManagedResourceOperationUpdateParamsSelectorInclude] `json:"include,required"`
}

func (r LabelManagedResourceOperationUpdateParamsSelector) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LabelManagedResourceOperationUpdateParamsSelectorInclude struct {
	OperationIDs param.Field[[]string] `json:"operation_ids,required"`
}

func (r LabelManagedResourceOperationUpdateParamsSelectorInclude) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type LabelManagedResourceOperationUpdateResponseEnvelope struct {
	Errors   Message                                     `json:"errors,required"`
	Messages Message                                     `json:"messages,required"`
	Result   LabelManagedResourceOperationUpdateResponse `json:"result,required"`
	// Whether the API call was successful.
	Success LabelManagedResourceOperationUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    labelManagedResourceOperationUpdateResponseEnvelopeJSON    `json:"-"`
}

// labelManagedResourceOperationUpdateResponseEnvelopeJSON contains the JSON
// metadata for the struct [LabelManagedResourceOperationUpdateResponseEnvelope]
type labelManagedResourceOperationUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *LabelManagedResourceOperationUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r labelManagedResourceOperationUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type LabelManagedResourceOperationUpdateResponseEnvelopeSuccess bool

const (
	LabelManagedResourceOperationUpdateResponseEnvelopeSuccessTrue LabelManagedResourceOperationUpdateResponseEnvelopeSuccess = true
)

func (r LabelManagedResourceOperationUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case LabelManagedResourceOperationUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
