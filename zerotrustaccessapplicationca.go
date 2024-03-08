// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/internal/apijson"
	"github.com/cloudflare/cloudflare-go/internal/param"
	"github.com/cloudflare/cloudflare-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/internal/shared"
	"github.com/cloudflare/cloudflare-go/option"
	"github.com/tidwall/gjson"
)

// ZeroTrustAccessApplicationCAService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZeroTrustAccessApplicationCAService] method instead.
type ZeroTrustAccessApplicationCAService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessApplicationCAService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessApplicationCAService(opts ...option.RequestOption) (r *ZeroTrustAccessApplicationCAService) {
	r = &ZeroTrustAccessApplicationCAService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *ZeroTrustAccessApplicationCAService) New(ctx context.Context, uuid string, body ZeroTrustAccessApplicationCANewParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationCANewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationCANewResponseEnvelope
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
func (r *ZeroTrustAccessApplicationCAService) List(ctx context.Context, query ZeroTrustAccessApplicationCAListParams, opts ...option.RequestOption) (res *[]ZeroTrustAccessApplicationCAListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationCAListResponseEnvelope
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
func (r *ZeroTrustAccessApplicationCAService) Delete(ctx context.Context, uuid string, body ZeroTrustAccessApplicationCADeleteParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationCADeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationCADeleteResponseEnvelope
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
func (r *ZeroTrustAccessApplicationCAService) Get(ctx context.Context, uuid string, query ZeroTrustAccessApplicationCAGetParams, opts ...option.RequestOption) (res *ZeroTrustAccessApplicationCAGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessApplicationCAGetResponseEnvelope
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

// Union satisfied by [ZeroTrustAccessApplicationCANewResponseUnknown] or
// [shared.UnionString].
type ZeroTrustAccessApplicationCANewResponse interface {
	ImplementsZeroTrustAccessApplicationCANewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationCANewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustAccessApplicationCAListResponse struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	Aud string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string                                       `json:"public_key"`
	JSON      zeroTrustAccessApplicationCAListResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationCAListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessApplicationCAListResponse]
type zeroTrustAccessApplicationCAListResponseJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCADeleteResponse struct {
	// The ID of the CA.
	ID   string                                         `json:"id"`
	JSON zeroTrustAccessApplicationCADeleteResponseJSON `json:"-"`
}

// zeroTrustAccessApplicationCADeleteResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessApplicationCADeleteResponse]
type zeroTrustAccessApplicationCADeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCADeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCADeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [ZeroTrustAccessApplicationCAGetResponseUnknown] or
// [shared.UnionString].
type ZeroTrustAccessApplicationCAGetResponse interface {
	ImplementsZeroTrustAccessApplicationCAGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*ZeroTrustAccessApplicationCAGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type ZeroTrustAccessApplicationCANewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationCANewResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationCANewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationCANewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessApplicationCANewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationCANewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationCANewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationCANewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationCANewResponseEnvelope]
type zeroTrustAccessApplicationCANewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCANewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCANewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCANewResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessApplicationCANewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationCANewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationCANewResponseEnvelopeErrors]
type zeroTrustAccessApplicationCANewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCANewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCANewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCANewResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessApplicationCANewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationCANewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationCANewResponseEnvelopeMessages]
type zeroTrustAccessApplicationCANewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCANewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCANewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessApplicationCANewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationCANewResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationCANewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationCAListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationCAListResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationCAListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationCAListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessApplicationCAListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessApplicationCAListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessApplicationCAListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessApplicationCAListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessApplicationCAListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationCAListResponseEnvelope]
type zeroTrustAccessApplicationCAListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCAListResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessApplicationCAListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationCAListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationCAListResponseEnvelopeErrors]
type zeroTrustAccessApplicationCAListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCAListResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessApplicationCAListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationCAListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationCAListResponseEnvelopeMessages]
type zeroTrustAccessApplicationCAListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessApplicationCAListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationCAListResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationCAListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationCAListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                        `json:"total_count"`
	JSON       zeroTrustAccessApplicationCAListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessApplicationCAListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationCAListResponseEnvelopeResultInfo]
type zeroTrustAccessApplicationCAListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCADeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationCADeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationCADeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationCADeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessApplicationCADeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationCADeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationCADeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationCADeleteResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationCADeleteResponseEnvelope]
type zeroTrustAccessApplicationCADeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCADeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCADeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCADeleteResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessApplicationCADeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationCADeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationCADeleteResponseEnvelopeErrors]
type zeroTrustAccessApplicationCADeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCADeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCADeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCADeleteResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustAccessApplicationCADeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationCADeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationCADeleteResponseEnvelopeMessages]
type zeroTrustAccessApplicationCADeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCADeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCADeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessApplicationCADeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationCADeleteResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationCADeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessApplicationCAGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessApplicationCAGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessApplicationCAGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessApplicationCAGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessApplicationCAGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessApplicationCAGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessApplicationCAGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessApplicationCAGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessApplicationCAGetResponseEnvelope]
type zeroTrustAccessApplicationCAGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCAGetResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessApplicationCAGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessApplicationCAGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessApplicationCAGetResponseEnvelopeErrors]
type zeroTrustAccessApplicationCAGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessApplicationCAGetResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessApplicationCAGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessApplicationCAGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessApplicationCAGetResponseEnvelopeMessages]
type zeroTrustAccessApplicationCAGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessApplicationCAGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessApplicationCAGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessApplicationCAGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessApplicationCAGetResponseEnvelopeSuccessTrue ZeroTrustAccessApplicationCAGetResponseEnvelopeSuccess = true
)
