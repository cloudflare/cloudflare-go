// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
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

// Generates a new short-lived certificate CA and public key.
func (r *AccessAppCaService) AccessShortLivedCertificateCAsNewAShortLivedCertificateCa(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a short-lived certificate CA and its public key.
func (r *AccessAppCaService) AccessShortLivedCertificateCAsGetAShortLivedCertificateCa(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/%s/ca", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists short-lived certificate CAs and their public keys.
func (r *AccessAppCaService) AccessShortLivedCertificateCAsListShortLivedCertificateCAs(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/apps/ca", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
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

type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse = interface{}

type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse = interface{}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse struct {
	// The ID of the CA.
	ID string `json:"id"`
	// The Application Audience (AUD) tag. Identifies the application associated with
	// the CA.
	Aud string `json:"aud"`
	// The public key to add to your SSH server configuration.
	PublicKey string                                                                            `json:"public_key"`
	JSON      accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseJSON struct {
	ID          apijson.Field
	Aud         apijson.Field
	PublicKey   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaDeleteResponseEnvelope struct {
	Errors   []AccessAppCaDeleteResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppCaDeleteResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppCaDeleteResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppCaDeleteResponseEnvelopeSuccess `json:"success"`
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

type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelope struct {
	Errors   []AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeSuccess `json:"success"`
	JSON    accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeJSON    `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelope]
type accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeErrors struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeErrors]
type accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeMessages struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeMessages]
type accessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeSuccess bool

const (
	AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeSuccessTrue AccessAppCaAccessShortLivedCertificateCAsNewAShortLivedCertificateCaResponseEnvelopeSuccess = true
)

type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelope struct {
	Errors   []AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeErrors   `json:"errors"`
	Messages []AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeMessages `json:"messages"`
	Result   AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponse                   `json:"result"`
	// Whether the API call was successful
	Success AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeSuccess `json:"success"`
	JSON    accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeJSON    `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelope]
type accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeErrors struct {
	Code    int64                                                                                          `json:"code,required"`
	Message string                                                                                         `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeErrors]
type accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeMessages struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeMessages]
type accessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeSuccess bool

const (
	AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeSuccessTrue AccessAppCaAccessShortLivedCertificateCAsGetAShortLivedCertificateCaResponseEnvelopeSuccess = true
)

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelope struct {
	Errors     []AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeMessages `json:"messages"`
	Result     []AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponse                 `json:"result"`
	ResultInfo AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeSuccess `json:"success"`
	JSON    accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeJSON    `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelope]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeErrors struct {
	Code    int64                                                                                           `json:"code,required"`
	Message string                                                                                          `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeErrorsJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeErrors]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeMessages struct {
	Code    int64                                                                                             `json:"code,required"`
	Message string                                                                                            `json:"message,required"`
	JSON    accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeMessagesJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeMessages]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                             `json:"total_count"`
	JSON       accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeResultInfo]
type accessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeSuccess bool

const (
	AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeSuccessTrue AccessAppCaAccessShortLivedCertificateCAsListShortLivedCertificateCAsResponseEnvelopeSuccess = true
)
