// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
	"github.com/cloudflare/cloudflare-sdk-go/option"
	"github.com/tidwall/gjson"
)

// AccessAppCaService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessAppCaService] method
// instead.
type AccessAppCaService struct {
	Options []option.RequestOption
}

// NewAccessAppCaService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessAppCaService(opts ...option.RequestOption) (r *AccessAppCaService) {
	r = &AccessAppCaService{}
	r.Options = opts
	return
}

// Generates a new short-lived certificate CA and public key.
func (r *AccessAppCaService) New(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppCaNewResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccessAppCaService) List(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessAppCaListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppCaListResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/ca", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a short-lived certificate CA.
func (r *AccessAppCaService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppCaDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccessAppCaService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppCaGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [AccessAppCaNewResponseUnknown] or [shared.UnionString].
type AccessAppCaNewResponse interface {
	ImplementsAccessAppCaNewResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessAppCaNewResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessAppCaListResponse struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	Aud string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string                      `json:"public_key"`
	JSON      accessAppCaListResponseJSON `json:"-"`
}

// accessAppCaListResponseJSON contains the JSON metadata for the struct
// [AccessAppCaListResponse]
type accessAppCaListResponseJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponse struct {
	// The ID of the CA.
	ID   string                        `json:"id"`
	JSON accessAppCaDeleteResponseJSON `json:"-"`
}

// accessAppCaDeleteResponseJSON contains the JSON metadata for the struct
// [AccessAppCaDeleteResponse]
type accessAppCaDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [AccessAppCaGetResponseUnknown] or [shared.UnionString].
type AccessAppCaGetResponse interface {
	ImplementsAccessAppCaGetResponse()
}

func init() {
	apijson.RegisterUnion(
		reflect.TypeOf((*AccessAppCaGetResponse)(nil)).Elem(),
		"",
		apijson.UnionVariant{
			TypeFilter: gjson.String,
			Type:       reflect.TypeOf(shared.UnionString("")),
		},
	)
}

type AccessAppCaNewResponseEnvelope struct {
	Errors   []AccessAppCaNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppCaNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppCaNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppCaNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppCaNewResponseEnvelopeJSON    `json:"-"`
}

// accessAppCaNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppCaNewResponseEnvelope]
type accessAppCaNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaNewResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessAppCaNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppCaNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppCaNewResponseEnvelopeErrors]
type accessAppCaNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaNewResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessAppCaNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppCaNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppCaNewResponseEnvelopeMessages]
type accessAppCaNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaNewResponseEnvelopeSuccess bool

const (
	AccessAppCaNewResponseEnvelopeSuccessTrue AccessAppCaNewResponseEnvelopeSuccess = true
)

type AccessAppCaListResponseEnvelope struct {
	Errors   []AccessAppCaListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppCaListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessAppCaListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessAppCaListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessAppCaListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessAppCaListResponseEnvelopeJSON       `json:"-"`
}

// accessAppCaListResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppCaListResponseEnvelope]
type accessAppCaListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaListResponseEnvelopeErrors struct {
	Code    int64                                     `json:"code,required"`
	Message string                                    `json:"message,required"`
	JSON    accessAppCaListResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppCaListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppCaListResponseEnvelopeErrors]
type accessAppCaListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaListResponseEnvelopeMessages struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessAppCaListResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppCaListResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppCaListResponseEnvelopeMessages]
type accessAppCaListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaListResponseEnvelopeSuccess bool

const (
	AccessAppCaListResponseEnvelopeSuccessTrue AccessAppCaListResponseEnvelopeSuccess = true
)

type AccessAppCaListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                       `json:"total_count"`
	JSON       accessAppCaListResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessAppCaListResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [AccessAppCaListResponseEnvelopeResultInfo]
type accessAppCaListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponseEnvelope struct {
	Errors   []AccessAppCaDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppCaDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppCaDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppCaDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppCaDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessAppCaDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppCaDeleteResponseEnvelope]
type accessAppCaDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponseEnvelopeErrors struct {
	Code    int64                                       `json:"code,required"`
	Message string                                      `json:"message,required"`
	JSON    accessAppCaDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppCaDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppCaDeleteResponseEnvelopeErrors]
type accessAppCaDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponseEnvelopeMessages struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    accessAppCaDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppCaDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppCaDeleteResponseEnvelopeMessages]
type accessAppCaDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaDeleteResponseEnvelopeSuccess bool

const (
	AccessAppCaDeleteResponseEnvelopeSuccessTrue AccessAppCaDeleteResponseEnvelopeSuccess = true
)

type AccessAppCaGetResponseEnvelope struct {
	Errors   []AccessAppCaGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessAppCaGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessAppCaGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessAppCaGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessAppCaGetResponseEnvelopeJSON    `json:"-"`
}

// accessAppCaGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [AccessAppCaGetResponseEnvelope]
type accessAppCaGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaGetResponseEnvelopeErrors struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    accessAppCaGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppCaGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [AccessAppCaGetResponseEnvelopeErrors]
type accessAppCaGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaGetResponseEnvelopeMessages struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    accessAppCaGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppCaGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [AccessAppCaGetResponseEnvelopeMessages]
type accessAppCaGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaGetResponseEnvelopeSuccess bool

const (
	AccessAppCaGetResponseEnvelopeSuccessTrue AccessAppCaGetResponseEnvelopeSuccess = true
)
