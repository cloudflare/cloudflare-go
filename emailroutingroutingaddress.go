// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// EmailRoutingRoutingAddressService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewEmailRoutingRoutingAddressService] method instead.
type EmailRoutingRoutingAddressService struct {
	Options []option.RequestOption
}

// NewEmailRoutingRoutingAddressService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEmailRoutingRoutingAddressService(opts ...option.RequestOption) (r *EmailRoutingRoutingAddressService) {
	r = &EmailRoutingRoutingAddressService{}
	r.Options = opts
	return
}

// Create a destination address to forward your emails to. Destination addresses
// need to be verified before they can be used.
func (r *EmailRoutingRoutingAddressService) New(ctx context.Context, accountIdentifier string, body EmailRoutingRoutingAddressNewParams, opts ...option.RequestOption) (res *EmailRoutingRoutingAddressNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingAddressNewResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists existing destination addresses.
func (r *EmailRoutingRoutingAddressService) List(ctx context.Context, accountIdentifier string, query EmailRoutingRoutingAddressListParams, opts ...option.RequestOption) (res *shared.V4PagePaginationArray[EmailRoutingRoutingAddressListResponse], err error) {
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
func (r *EmailRoutingRoutingAddressService) ListAutoPaging(ctx context.Context, accountIdentifier string, query EmailRoutingRoutingAddressListParams, opts ...option.RequestOption) *shared.V4PagePaginationArrayAutoPager[EmailRoutingRoutingAddressListResponse] {
	return shared.NewV4PagePaginationArrayAutoPager(r.List(ctx, accountIdentifier, query, opts...))
}

// Deletes a specific destination address.
func (r *EmailRoutingRoutingAddressService) Delete(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingAddressDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingAddressDeleteResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Gets information for a specific destination email already created.
func (r *EmailRoutingRoutingAddressService) Get(ctx context.Context, accountIdentifier string, destinationAddressIdentifier string, opts ...option.RequestOption) (res *EmailRoutingRoutingAddressGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env EmailRoutingRoutingAddressGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/email/routing/addresses/%s", accountIdentifier, destinationAddressIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type EmailRoutingRoutingAddressNewResponse struct {
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
	Verified time.Time                                 `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressNewResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressNewResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingAddressNewResponse]
type emailRoutingRoutingAddressNewResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressNewResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressListResponse struct {
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
	Verified time.Time                                  `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressListResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressListResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingAddressListResponse]
type emailRoutingRoutingAddressListResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressListResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressDeleteResponse struct {
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
	Verified time.Time                                    `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressDeleteResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingAddressDeleteResponse]
type emailRoutingRoutingAddressDeleteResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressGetResponse struct {
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
	Verified time.Time                                 `json:"verified" format:"date-time"`
	JSON     emailRoutingRoutingAddressGetResponseJSON `json:"-"`
}

// emailRoutingRoutingAddressGetResponseJSON contains the JSON metadata for the
// struct [EmailRoutingRoutingAddressGetResponse]
type emailRoutingRoutingAddressGetResponseJSON struct {
	ID          apijson.Field
	Created     apijson.Field
	Email       apijson.Field
	Modified    apijson.Field
	Tag         apijson.Field
	Verified    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressGetResponseJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressNewParams struct {
	// The contact email address of the user.
	Email param.Field[string] `json:"email,required"`
}

func (r EmailRoutingRoutingAddressNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type EmailRoutingRoutingAddressNewResponseEnvelope struct {
	Errors   []EmailRoutingRoutingAddressNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingAddressNewResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingAddressNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingAddressNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingAddressNewResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingAddressNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingAddressNewResponseEnvelope]
type emailRoutingRoutingAddressNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRoutingAddressNewResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingAddressNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressNewResponseEnvelopeErrors]
type emailRoutingRoutingAddressNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRoutingAddressNewResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingAddressNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressNewResponseEnvelopeMessages]
type emailRoutingRoutingAddressNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingAddressNewResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingAddressNewResponseEnvelopeSuccessTrue EmailRoutingRoutingAddressNewResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingAddressListParams struct {
	// Sorts results in an ascending or descending order.
	Direction param.Field[EmailRoutingRoutingAddressListParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Filter by verified destination addresses.
	Verified param.Field[EmailRoutingRoutingAddressListParamsVerified] `query:"verified"`
}

// URLQuery serializes [EmailRoutingRoutingAddressListParams]'s query parameters as
// `url.Values`.
func (r EmailRoutingRoutingAddressListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sorts results in an ascending or descending order.
type EmailRoutingRoutingAddressListParamsDirection string

const (
	EmailRoutingRoutingAddressListParamsDirectionAsc  EmailRoutingRoutingAddressListParamsDirection = "asc"
	EmailRoutingRoutingAddressListParamsDirectionDesc EmailRoutingRoutingAddressListParamsDirection = "desc"
)

// Filter by verified destination addresses.
type EmailRoutingRoutingAddressListParamsVerified bool

const (
	EmailRoutingRoutingAddressListParamsVerifiedTrue  EmailRoutingRoutingAddressListParamsVerified = true
	EmailRoutingRoutingAddressListParamsVerifiedFalse EmailRoutingRoutingAddressListParamsVerified = false
)

type EmailRoutingRoutingAddressDeleteResponseEnvelope struct {
	Errors   []EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingAddressDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingAddressDeleteResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [EmailRoutingRoutingAddressDeleteResponseEnvelope]
type emailRoutingRoutingAddressDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    emailRoutingRoutingAddressDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors]
type emailRoutingRoutingAddressDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    emailRoutingRoutingAddressDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingAddressDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages]
type emailRoutingRoutingAddressDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccessTrue EmailRoutingRoutingAddressDeleteResponseEnvelopeSuccess = true
)

type EmailRoutingRoutingAddressGetResponseEnvelope struct {
	Errors   []EmailRoutingRoutingAddressGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []EmailRoutingRoutingAddressGetResponseEnvelopeMessages `json:"messages,required"`
	Result   EmailRoutingRoutingAddressGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success EmailRoutingRoutingAddressGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    emailRoutingRoutingAddressGetResponseEnvelopeJSON    `json:"-"`
}

// emailRoutingRoutingAddressGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [EmailRoutingRoutingAddressGetResponseEnvelope]
type emailRoutingRoutingAddressGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    emailRoutingRoutingAddressGetResponseEnvelopeErrorsJSON `json:"-"`
}

// emailRoutingRoutingAddressGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressGetResponseEnvelopeErrors]
type emailRoutingRoutingAddressGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type EmailRoutingRoutingAddressGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    emailRoutingRoutingAddressGetResponseEnvelopeMessagesJSON `json:"-"`
}

// emailRoutingRoutingAddressGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [EmailRoutingRoutingAddressGetResponseEnvelopeMessages]
type emailRoutingRoutingAddressGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EmailRoutingRoutingAddressGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r emailRoutingRoutingAddressGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type EmailRoutingRoutingAddressGetResponseEnvelopeSuccess bool

const (
	EmailRoutingRoutingAddressGetResponseEnvelopeSuccessTrue EmailRoutingRoutingAddressGetResponseEnvelopeSuccess = true
)
