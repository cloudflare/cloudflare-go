// File generated from our OpenAPI spec by Stainless.

package email_routing

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// RoutingAddressService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewRoutingAddressService] method
// instead.
type RoutingAddressService struct {
	Options []option.RequestOption
}

// NewRoutingAddressService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRoutingAddressService(opts ...option.RequestOption) (r *RoutingAddressService) {
	r = &RoutingAddressService{}
	r.Options = opts
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *RoutingAddressService) New(ctx context.Context, accountIdentifier string, body RoutingAddressNewParams, opts ...option.RequestOption) (res *RoutingAddressNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingAddressNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing destination addresses.
func (r *RoutingAddressService) List(ctx context.Context, accountIdentifier string, query RoutingAddressListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[RoutingAddressListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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
func (r *RoutingAddressService) ListAutoPaging(ctx context.Context, accountIdentifier string, query RoutingAddressListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[RoutingAddressListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountIdentifier, query, opts...))
}

// Deletes a specific destination address.
func (r *RoutingAddressService) Delete(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *RoutingAddressDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingAddressDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information for a specific destination email already created.
func (r *RoutingAddressService) Get(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *RoutingAddressGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env RoutingAddressGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type RoutingAddressNewResponse struct {
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
	Verified time.Time                     `json:"verified" format:"date-time"`
	JSON     routingAddressNewResponseJSON `json:"-"`
}

// routingAddressNewResponseJSON contains the JSON metadata for the struct
// [RoutingAddressNewResponse]
type routingAddressNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressNewResponseJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressListResponse struct {
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
	Verified time.Time                      `json:"verified" format:"date-time"`
	JSON     routingAddressListResponseJSON `json:"-"`
}

// routingAddressListResponseJSON contains the JSON metadata for the struct
// [RoutingAddressListResponse]
type routingAddressListResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressListResponseJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressDeleteResponse struct {
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
	Verified time.Time                        `json:"verified" format:"date-time"`
	JSON     routingAddressDeleteResponseJSON `json:"-"`
}

// routingAddressDeleteResponseJSON contains the JSON metadata for the struct
// [RoutingAddressDeleteResponse]
type routingAddressDeleteResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressGetResponse struct {
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
	Verified time.Time                     `json:"verified" format:"date-time"`
	JSON     routingAddressGetResponseJSON `json:"-"`
}

// routingAddressGetResponseJSON contains the JSON metadata for the struct
// [RoutingAddressGetResponse]
type routingAddressGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressGetResponseJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressNewParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r RoutingAddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type RoutingAddressNewResponseEnvelope struct {
	Errors   []RoutingAddressNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingAddressNewResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingAddressNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingAddressNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingAddressNewResponseEnvelopeJSON    `json:"-"`
}

// routingAddressNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingAddressNewResponseEnvelope]
type routingAddressNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressNewResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    routingAddressNewResponseEnvelopeErrorsJSON `json:"-"`
}

// routingAddressNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingAddressNewResponseEnvelopeErrors]
type routingAddressNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressNewResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    routingAddressNewResponseEnvelopeMessagesJSON `json:"-"`
}

// routingAddressNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingAddressNewResponseEnvelopeMessages]
type routingAddressNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingAddressNewResponseEnvelopeSuccess bool

const (
	RoutingAddressNewResponseEnvelopeSuccessTrue RoutingAddressNewResponseEnvelopeSuccess = true
)

type RoutingAddressListParams struct {
	// Sorts results in an ascending or descending order.
	Direction param.Field[RoutingAddressListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[RoutingAddressListParamsVerified] `query:"verified"`
}

// URLQuery serializes [RoutingAddressListParams]'s query parameters as
// `url.Values`.
func (r RoutingAddressListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts results in an ascending or descending order.
type RoutingAddressListParamsDirection string

const (
	RoutingAddressListParamsDirectionAsc  RoutingAddressListParamsDirection = "asc"
	RoutingAddressListParamsDirectionDesc RoutingAddressListParamsDirection = "desc"
)

// Filter by verified destination addresses.
type RoutingAddressListParamsVerified bool

const (
	RoutingAddressListParamsVerifiedTrue  RoutingAddressListParamsVerified = true
	RoutingAddressListParamsVerifiedFalse RoutingAddressListParamsVerified = false
)

type RoutingAddressDeleteResponseEnvelope struct {
	Errors   []RoutingAddressDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingAddressDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingAddressDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingAddressDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingAddressDeleteResponseEnvelopeJSON    `json:"-"`
}

// routingAddressDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [RoutingAddressDeleteResponseEnvelope]
type routingAddressDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressDeleteResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    routingAddressDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// routingAddressDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [RoutingAddressDeleteResponseEnvelopeErrors]
type routingAddressDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressDeleteResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    routingAddressDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// routingAddressDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [RoutingAddressDeleteResponseEnvelopeMessages]
type routingAddressDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingAddressDeleteResponseEnvelopeSuccess bool

const (
	RoutingAddressDeleteResponseEnvelopeSuccessTrue RoutingAddressDeleteResponseEnvelopeSuccess = true
)

type RoutingAddressGetResponseEnvelope struct {
	Errors   []RoutingAddressGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []RoutingAddressGetResponseEnvelopeMessages `json:"messages,required"`
	Result   RoutingAddressGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success RoutingAddressGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    routingAddressGetResponseEnvelopeJSON    `json:"-"`
}

// routingAddressGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [RoutingAddressGetResponseEnvelope]
type routingAddressGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressGetResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    routingAddressGetResponseEnvelopeErrorsJSON `json:"-"`
}

// routingAddressGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [RoutingAddressGetResponseEnvelopeErrors]
type routingAddressGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type RoutingAddressGetResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    routingAddressGetResponseEnvelopeMessagesJSON `json:"-"`
}

// routingAddressGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [RoutingAddressGetResponseEnvelopeMessages]
type routingAddressGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *RoutingAddressGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r routingAddressGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type RoutingAddressGetResponseEnvelopeSuccess bool

const (
	RoutingAddressGetResponseEnvelopeSuccessTrue RoutingAddressGetResponseEnvelopeSuccess = true
)
