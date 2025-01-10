// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// PoolReferenceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPoolReferenceService] method instead.
type PoolReferenceService struct {
	Options []option.RequestOption
}

// NewPoolReferenceService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPoolReferenceService(opts ...option.RequestOption) (r *PoolReferenceService) {
	r = &PoolReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided pool.
func (r *PoolReferenceService) Get(ctx context.Context, poolID string, query PoolReferenceGetParams, opts ...option.RequestOption) (res *[]PoolReferenceGetResponse, err error) {
	var env PoolReferenceGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if poolID == "" {
		err = errors.New("missing required pool_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/pools/%s/references", query.AccountID, poolID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type PoolReferenceGetResponse struct {
	ReferenceType PoolReferenceGetResponseReferenceType `json:"reference_type"`
	ResourceID    string                                `json:"resource_id"`
	ResourceName  string                                `json:"resource_name"`
	ResourceType  string                                `json:"resource_type"`
	JSON          poolReferenceGetResponseJSON          `json:"-"`
}

// poolReferenceGetResponseJSON contains the JSON metadata for the struct
// [PoolReferenceGetResponse]
type poolReferenceGetResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *PoolReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type PoolReferenceGetResponseReferenceType string

const (
	PoolReferenceGetResponseReferenceTypeStar     PoolReferenceGetResponseReferenceType = "*"
	PoolReferenceGetResponseReferenceTypeReferral PoolReferenceGetResponseReferenceType = "referral"
	PoolReferenceGetResponseReferenceTypeReferrer PoolReferenceGetResponseReferenceType = "referrer"
)

func (r PoolReferenceGetResponseReferenceType) IsKnown() bool {
	switch r {
	case PoolReferenceGetResponseReferenceTypeStar, PoolReferenceGetResponseReferenceTypeReferral, PoolReferenceGetResponseReferenceTypeReferrer:
		return true
	}
	return false
}

type PoolReferenceGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type PoolReferenceGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// List of resources that reference a given pool.
	Result []PoolReferenceGetResponse `json:"result,required"`
	// Whether the API call was successful
	Success PoolReferenceGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    poolReferenceGetResponseEnvelopeJSON    `json:"-"`
}

// poolReferenceGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [PoolReferenceGetResponseEnvelope]
type poolReferenceGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoolReferenceGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r poolReferenceGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type PoolReferenceGetResponseEnvelopeSuccess bool

const (
	PoolReferenceGetResponseEnvelopeSuccessTrue PoolReferenceGetResponseEnvelopeSuccess = true
)

func (r PoolReferenceGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case PoolReferenceGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
