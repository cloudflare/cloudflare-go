// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessCertificateService] method
// instead.
type AccessCertificateService struct {
	Options  []option.RequestOption
	Settings *AccessCertificateSettingService
}

// NewAccessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessCertificateService(opts ...option.RequestOption) (r *AccessCertificateService) {
	r = &AccessCertificateService{}
	r.Options = opts
	r.Settings = NewAccessCertificateSettingService(opts...)
	return
}

// Fetches a single mTLS certificate.
func (r *AccessCertificateService) Get(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateGetResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured mTLS certificate.
func (r *AccessCertificateService) Update(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, body AccessCertificateUpdateParams, opts ...option.RequestOption) (res *AccessCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateUpdateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an mTLS certificate.
func (r *AccessCertificateService) Delete(ctx context.Context, accountOrZone string, accountOrZoneID string, uuid string, opts ...option.RequestOption) (res *AccessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Adds a new mTLS root certificate to Access.
func (r *AccessCertificateService) AccessMTLSAuthenticationAddAnMTLSCertificate(ctx context.Context, accountOrZone string, accountOrZoneID string, body AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams, opts ...option.RequestOption) (res *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all mTLS root certificates.
func (r *AccessCertificateService) AccessMTLSAuthenticationListMTLSCertificates(ctx context.Context, accountOrZone string, accountOrZoneID string, opts ...option.RequestOption) (res *[]AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelope
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessCertificateGetResponse struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                           `json:"name"`
	UpdatedAt time.Time                        `json:"updated_at" format:"date-time"`
	JSON      accessCertificateGetResponseJSON `json:"-"`
}

// accessCertificateGetResponseJSON contains the JSON metadata for the struct
// [AccessCertificateGetResponse]
type accessCertificateGetResponseJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateUpdateResponse struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                              `json:"name"`
	UpdatedAt time.Time                           `json:"updated_at" format:"date-time"`
	JSON      accessCertificateUpdateResponseJSON `json:"-"`
}

// accessCertificateUpdateResponseJSON contains the JSON metadata for the struct
// [AccessCertificateUpdateResponse]
type accessCertificateUpdateResponseJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateDeleteResponse struct {
	// UUID
	ID   string                              `json:"id"`
	JSON accessCertificateDeleteResponseJSON `json:"-"`
}

// accessCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [AccessCertificateDeleteResponse]
type accessCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                                                    `json:"name"`
	UpdatedAt time.Time                                                                 `json:"updated_at" format:"date-time"`
	JSON      accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                                                                    `json:"name"`
	UpdatedAt time.Time                                                                 `json:"updated_at" format:"date-time"`
	JSON      accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseJSON struct {
	ID                  apijson.Field
	AssociatedHostnames apijson.Field
	CreatedAt           apijson.Field
	ExpiresOn           apijson.Field
	Fingerprint         apijson.Field
	Name                apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateGetResponseEnvelope struct {
	Errors   []AccessCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessCertificateGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateGetResponseEnvelope]
type accessCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessCertificateGetResponseEnvelopeErrors]
type accessCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessCertificateGetResponseEnvelopeMessages]
type accessCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateGetResponseEnvelopeSuccess bool

const (
	AccessCertificateGetResponseEnvelopeSuccessTrue AccessCertificateGetResponseEnvelopeSuccess = true
)

type AccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r AccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateUpdateResponseEnvelope struct {
	Errors   []AccessCertificateUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessCertificateUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessCertificateUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCertificateUpdateResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateUpdateResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateUpdateResponseEnvelope]
type accessCertificateUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateUpdateResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessCertificateUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessCertificateUpdateResponseEnvelopeErrors]
type accessCertificateUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateUpdateResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessCertificateUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateUpdateResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessCertificateUpdateResponseEnvelopeMessages]
type accessCertificateUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateUpdateResponseEnvelopeSuccess bool

const (
	AccessCertificateUpdateResponseEnvelopeSuccessTrue AccessCertificateUpdateResponseEnvelopeSuccess = true
)

type AccessCertificateDeleteResponseEnvelope struct {
	Errors   []AccessCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessCertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateDeleteResponseEnvelope]
type accessCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    accessCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessCertificateDeleteResponseEnvelopeErrors]
type accessCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    accessCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [AccessCertificateDeleteResponseEnvelopeMessages]
type accessCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateDeleteResponseEnvelopeSuccess bool

const (
	AccessCertificateDeleteResponseEnvelopeSuccessTrue AccessCertificateDeleteResponseEnvelopeSuccess = true
)

type AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r AccessCertificateAccessMTLSAuthenticationAddAnMTLSCertificateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelope struct {
	Errors   []AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelope]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeErrors struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeErrors]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeMessages struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeMessages]
type accessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeSuccess bool

const (
	AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeSuccessTrue AccessCertificateAccessMtlsAuthenticationAddAnMtlsCertificateResponseEnvelopeSuccess = true
)

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelope struct {
	Errors   []AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeResultInfo `json:"result_info"`
	JSON       accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeJSON       `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelope]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeErrors struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeErrors]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeMessages struct {
	Code    int64                                                                                     `json:"code,required"`
	Message string                                                                                    `json:"message,required"`
	JSON    accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeMessages]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeSuccess bool

const (
	AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeSuccessTrue AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeSuccess = true
)

type AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                     `json:"total_count"`
	JSON       accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeResultInfo]
type accessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateAccessMtlsAuthenticationListMtlsCertificatesResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
