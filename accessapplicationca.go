// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccessApplicationCaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessApplicationCaService]
// method instead.
type AccessApplicationCaService struct {
	Options []option.RequestOption
}

// NewAccessApplicationCaService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessApplicationCaService(opts ...option.RequestOption) (r *AccessApplicationCaService) {
	r = &AccessApplicationCaService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *AccessApplicationCaService) New(ctx context.Context, uuid string, body AccessApplicationCaNewParams, opts ...option.RequestOption) (res *AccessApplicationCaNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCaNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccessApplicationCaService) List(ctx context.Context, query AccessApplicationCaListParams, opts ...option.RequestOption) (res *[]AccessApplicationCaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCaListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/ca", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a short-lived certificate CA.
func (r *AccessApplicationCaService) Delete(ctx context.Context, uuid string, body AccessApplicationCaDeleteParams, opts ...option.RequestOption) (res *AccessApplicationCaDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCaDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccessApplicationCaService) Get(ctx context.Context, uuid string, query AccessApplicationCaGetParams, opts ...option.RequestOption) (res *AccessApplicationCaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCaGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessApplicationCaNewResponseUnknown] or
// [shared.UnionString].
type AccessApplicationCaNewResponse interface {
	ImplementsAccessApplicationCaNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationCaNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessApplicationCaListResponse struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	Aud string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string                              `json:"public_key"`
	JSON      accessApplicationCaListResponseJSON `json:"-"`
}

// accessApplicationCaListResponseJSON contains the JSON metadata for the struct
// [AccessApplicationCaListResponse]
type accessApplicationCaListResponseJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaDeleteResponse struct {
	// The ID of the CA.
	ID   string                                `json:"id"`
	JSON accessApplicationCaDeleteResponseJSON `json:"-"`
}

// accessApplicationCaDeleteResponseJSON contains the JSON metadata for the struct
// [AccessApplicationCaDeleteResponse]
type accessApplicationCaDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccessApplicationCaGetResponseUnknown] or
// [shared.UnionString].
type AccessApplicationCaGetResponse interface {
	ImplementsAccessApplicationCaGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationCaGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessApplicationCaNewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AccessApplicationCaNewResponseEnvelope struct {
	Errors   []AccessApplicationCaNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCaNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationCaNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationCaNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationCaNewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCaNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCaNewResponseEnvelope]
type accessApplicationCaNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaNewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessApplicationCaNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCaNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCaNewResponseEnvelopeErrors]
type accessApplicationCaNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaNewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessApplicationCaNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCaNewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCaNewResponseEnvelopeMessages]
type accessApplicationCaNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationCaNewResponseEnvelopeSuccess bool

const (
	AccessApplicationCaNewResponseEnvelopeSuccessTrue AccessApplicationCaNewResponseEnvelopeSuccess = true
)

type AccessApplicationCaListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AccessApplicationCaListResponseEnvelope struct {
	Errors   []AccessApplicationCaListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCaListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessApplicationCaListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessApplicationCaListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessApplicationCaListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessApplicationCaListResponseEnvelopeJSON       `json:"-"`
}

// accessApplicationCaListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCaListResponseEnvelope]
type accessApplicationCaListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessApplicationCaListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCaListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCaListResponseEnvelopeErrors]
type accessApplicationCaListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessApplicationCaListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCaListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCaListResponseEnvelopeMessages]
type accessApplicationCaListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationCaListResponseEnvelopeSuccess bool

const (
	AccessApplicationCaListResponseEnvelopeSuccessTrue AccessApplicationCaListResponseEnvelopeSuccess = true
)

type AccessApplicationCaListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       accessApplicationCaListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessApplicationCaListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AccessApplicationCaListResponseEnvelopeResultInfo]
type accessApplicationCaListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AccessApplicationCaDeleteResponseEnvelope struct {
	Errors   []AccessApplicationCaDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCaDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationCaDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationCaDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationCaDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCaDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCaDeleteResponseEnvelope]
type accessApplicationCaDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaDeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessApplicationCaDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCaDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationCaDeleteResponseEnvelopeErrors]
type accessApplicationCaDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaDeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessApplicationCaDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCaDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCaDeleteResponseEnvelopeMessages]
type accessApplicationCaDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationCaDeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationCaDeleteResponseEnvelopeSuccessTrue AccessApplicationCaDeleteResponseEnvelopeSuccess = true
)

type AccessApplicationCaGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id,required"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type AccessApplicationCaGetResponseEnvelope struct {
	Errors   []AccessApplicationCaGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCaGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationCaGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationCaGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationCaGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCaGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCaGetResponseEnvelope]
type accessApplicationCaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessApplicationCaGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCaGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCaGetResponseEnvelopeErrors]
type accessApplicationCaGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessApplicationCaGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessApplicationCaGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCaGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCaGetResponseEnvelopeMessages]
type accessApplicationCaGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCaGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessApplicationCaGetResponseEnvelopeSuccess bool

const (
	AccessApplicationCaGetResponseEnvelopeSuccessTrue AccessApplicationCaGetResponseEnvelopeSuccess = true
)
