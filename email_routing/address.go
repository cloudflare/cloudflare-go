// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package email_routing

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// AddressService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAddressService] method instead.
type AddressService struct {
	Options []option.RequestOption
}

// NewAddressService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAddressService(opts ...option.RequestOption) (r *AddressService) {
	r = &AddressService{}
	r.Options = opts
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *AddressService) New(ctx context.Context, params AddressNewParams, opts ...option.RequestOption) (res *Address, err error) {
	var env AddressNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing destination addresses.
func (r *AddressService) List(ctx context.Context, params AddressListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[Address], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Lists existing destination addresses.
func (r *AddressService) ListAutoPaging(ctx context.Context, params AddressListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[Address] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

// Deletes a specific destination address.
func (r *AddressService) Delete(ctx context.Context, destinationAddressIdentifier string, body AddressDeleteParams, opts ...option.RequestOption) (res *Address, err error) {
	var env AddressDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if destinationAddressIdentifier == "" {
		err = errors.New("missing required destination_address_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", body.AccountID, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information for a specific destination email already created.
func (r *AddressService) Get(ctx context.Context, destinationAddressIdentifier string, query AddressGetParams, opts ...option.RequestOption) (res *Address, err error) {
	var env AddressGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if destinationAddressIdentifier == "" {
		err = errors.New("missing required destination_address_identifier parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", query.AccountID, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type Address struct {
	// Destination address identifier.
	ID string `json:"id"`
	// The date and time the destination address has been created.
	Created time.Time `json:"created" format:"date-time"`
	// The contact email address of the user.
	Email string `json:"email"`
	// The date and time the destination address was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// Destination address tag. (Deprecated, replaced by destination address
	// identifier)
	Tag string `json:"tag"`
	// The date and time the destination address has been verified. Null means not
	// verified yet.
	Verified time.Time   `json:"verified" format:"date-time"`
	JSON     addressJSON `json:"-"`
}

// addressJSON contains the JSON metadata for the struct [Address]
type addressJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *Address) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressJSON) RawJSON() string {
	return r.raw
}

type AddressNewParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r AddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AddressNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AddressNewResponseEnvelopeSuccess `json:"success,required"`
	Result  Address                           `json:"result"`
	JSON    addressNewResponseEnvelopeJSON    `json:"-"`
}

// addressNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressNewResponseEnvelope]
type addressNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressNewResponseEnvelopeSuccess bool

const (
	AddressNewResponseEnvelopeSuccessTrue AddressNewResponseEnvelopeSuccess = true
)

func (r AddressNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Sorts results in an ascending or descending order.
	Direction param.Field[AddressListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[AddressListParamsVerified] `query:"verified"`
}

// URLQuery serializes [AddressListParams]'s query parameters as `url.Values`.
func (r AddressListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Sorts results in an ascending or descending order.
type AddressListParamsDirection string

const (
	AddressListParamsDirectionAsc  AddressListParamsDirection = "asc"
	AddressListParamsDirectionDesc AddressListParamsDirection = "desc"
)

func (r AddressListParamsDirection) IsKnown() bool {
	switch r {
	case AddressListParamsDirectionAsc, AddressListParamsDirectionDesc:
		return true
	}
	return false
}

// Filter by verified destination addresses.
type AddressListParamsVerified bool

const (
	AddressListParamsVerifiedTrue  AddressListParamsVerified = true
	AddressListParamsVerifiedFalse AddressListParamsVerified = false
)

func (r AddressListParamsVerified) IsKnown() bool {
	switch r {
	case AddressListParamsVerifiedTrue, AddressListParamsVerifiedFalse:
		return true
	}
	return false
}

type AddressDeleteParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AddressDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  Address                              `json:"result"`
	JSON    addressDeleteResponseEnvelopeJSON    `json:"-"`
}

// addressDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressDeleteResponseEnvelope]
type addressDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressDeleteResponseEnvelopeSuccess bool

const (
	AddressDeleteResponseEnvelopeSuccessTrue AddressDeleteResponseEnvelopeSuccess = true
)

func (r AddressDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AddressGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AddressGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success AddressGetResponseEnvelopeSuccess `json:"success,required"`
	Result  Address                           `json:"result"`
	JSON    addressGetResponseEnvelopeJSON    `json:"-"`
}

// addressGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AddressGetResponseEnvelope]
type addressGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AddressGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r addressGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AddressGetResponseEnvelopeSuccess bool

const (
	AddressGetResponseEnvelopeSuccessTrue AddressGetResponseEnvelopeSuccess = true
)

func (r AddressGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AddressGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
