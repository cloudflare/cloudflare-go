// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/internal/shared"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/tidwall/gjson"
)

// AccessApplicationCAService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessApplicationCAService]
// method instead.
type AccessApplicationCAService struct {
	Options []option.RequestOption
}

// NewAccessApplicationCAService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessApplicationCAService(opts ...option.RequestOption) (r *AccessApplicationCAService) {
	r = &AccessApplicationCAService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *AccessApplicationCAService) New(ctx context.Context, uuid string, body AccessApplicationCANewParams, opts ...option.RequestOption) (res *AccessApplicationCANewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCANewResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccessApplicationCAService) List(ctx context.Context, query AccessApplicationCAListParams, opts ...option.RequestOption) (res *[]ZeroTrustCA, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCAListResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a short-lived certificate CA.
func (r *AccessApplicationCAService) Delete(ctx context.Context, uuid string, body AccessApplicationCADeleteParams, opts ...option.RequestOption) (res *AccessApplicationCADeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCADeleteResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccessApplicationCAService) Get(ctx context.Context, uuid string, query AccessApplicationCAGetParams, opts ...option.RequestOption) (res *AccessApplicationCAGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessApplicationCAGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustCA struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	Aud string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string          `json:"public_key"`
	JSON      zeroTrustCAJSON `json:"-"`
}

// zeroTrustCAJSON contains the JSON metadata for the struct [ZeroTrustCA]
type zeroTrustCAJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustCA) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustCAJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.AccessApplicationCANewResponseUnknown] or
// [shared.UnionString].
type AccessApplicationCANewResponse interface {
	ImplementsZeroTrustAccessApplicationCANewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationCANewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessApplicationCADeleteResponse struct {
	// The ID of the CA.
	ID   string                                `json:"id"`
	JSON accessApplicationCADeleteResponseJSON `json:"-"`
}

// accessApplicationCADeleteResponseJSON contains the JSON metadata for the struct
// [AccessApplicationCADeleteResponse]
type accessApplicationCADeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseJSON) RawJSON() string {
	return r.raw
}

// Union satisfied by [zero_trust.AccessApplicationCAGetResponseUnknown] or
// [shared.UnionString].
type AccessApplicationCAGetResponse interface {
	ImplementsZeroTrustAccessApplicationCAGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessApplicationCAGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessApplicationCANewParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationCANewResponseEnvelope struct {
	Errors   []AccessApplicationCANewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCANewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationCANewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationCANewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationCANewResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCANewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCANewResponseEnvelope]
type accessApplicationCANewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCANewResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessApplicationCANewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCANewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCANewResponseEnvelopeErrors]
type accessApplicationCANewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCANewResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessApplicationCANewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCANewResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCANewResponseEnvelopeMessages]
type accessApplicationCANewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCANewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCANewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationCANewResponseEnvelopeSuccess bool

const (
	AccessApplicationCANewResponseEnvelopeSuccessTrue AccessApplicationCANewResponseEnvelopeSuccess = true
)

func (r AccessApplicationCANewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCANewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationCAListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationCAListResponseEnvelope struct {
	Errors   []AccessApplicationCAListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCAListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustCA                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessApplicationCAListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessApplicationCAListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessApplicationCAListResponseEnvelopeJSON       `json:"-"`
}

// accessApplicationCAListResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCAListResponseEnvelope]
type accessApplicationCAListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessApplicationCAListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCAListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCAListResponseEnvelopeErrors]
type accessApplicationCAListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessApplicationCAListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCAListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCAListResponseEnvelopeMessages]
type accessApplicationCAListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationCAListResponseEnvelopeSuccess bool

const (
	AccessApplicationCAListResponseEnvelopeSuccessTrue AccessApplicationCAListResponseEnvelopeSuccess = true
)

func (r AccessApplicationCAListResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCAListResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationCAListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       accessApplicationCAListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessApplicationCAListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [AccessApplicationCAListResponseEnvelopeResultInfo]
type accessApplicationCAListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationCADeleteResponseEnvelope struct {
	Errors   []AccessApplicationCADeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCADeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationCADeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationCADeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationCADeleteResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCADeleteResponseEnvelope]
type accessApplicationCADeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteResponseEnvelopeErrors struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessApplicationCADeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [AccessApplicationCADeleteResponseEnvelopeErrors]
type accessApplicationCADeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCADeleteResponseEnvelopeMessages struct {
	Code    int64                                                 `json:"code,required"`
	Message string                                                `json:"message,required"`
	JSON    accessApplicationCADeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCADeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCADeleteResponseEnvelopeMessages]
type accessApplicationCADeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCADeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCADeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationCADeleteResponseEnvelopeSuccess bool

const (
	AccessApplicationCADeleteResponseEnvelopeSuccessTrue AccessApplicationCADeleteResponseEnvelopeSuccess = true
)

func (r AccessApplicationCADeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCADeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessApplicationCAGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessApplicationCAGetResponseEnvelope struct {
	Errors   []AccessApplicationCAGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessApplicationCAGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessApplicationCAGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessApplicationCAGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessApplicationCAGetResponseEnvelopeJSON    `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessApplicationCAGetResponseEnvelope]
type accessApplicationCAGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessApplicationCAGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessApplicationCAGetResponseEnvelopeErrors]
type accessApplicationCAGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessApplicationCAGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    accessApplicationCAGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessApplicationCAGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessApplicationCAGetResponseEnvelopeMessages]
type accessApplicationCAGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessApplicationCAGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessApplicationCAGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessApplicationCAGetResponseEnvelopeSuccess bool

const (
	AccessApplicationCAGetResponseEnvelopeSuccessTrue AccessApplicationCAGetResponseEnvelopeSuccess = true
)

func (r AccessApplicationCAGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessApplicationCAGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
