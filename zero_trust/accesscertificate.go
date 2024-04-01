// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/pagination"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
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

// Adds a new mTLS root certificate to Access.
func (r *AccessCertificateService) New(ctx context.Context, params AccessCertificateNewParams, opts ...option.RequestOption) (res *ZeroTrustCertificates, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured mTLS certificate.
func (r *AccessCertificateService) Update(ctx context.Context, uuid string, params AccessCertificateUpdateParams, opts ...option.RequestOption) (res *ZeroTrustCertificates, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all mTLS root certificates.
func (r *AccessCertificateService) List(ctx context.Context, query AccessCertificateListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZeroTrustCertificates], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates", accountOrZone, accountOrZoneID)
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

// Lists all mTLS root certificates.
func (r *AccessCertificateService) ListAutoPaging(ctx context.Context, query AccessCertificateListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZeroTrustCertificates] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an mTLS certificate.
func (r *AccessCertificateService) Delete(ctx context.Context, uuid string, body AccessCertificateDeleteParams, opts ...option.RequestOption) (res *AccessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single mTLS certificate.
func (r *AccessCertificateService) Get(ctx context.Context, uuid string, query AccessCertificateGetParams, opts ...option.RequestOption) (res *ZeroTrustCertificates, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessCertificateGetResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/certificates/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustCertificates struct {
	// The ID of the application that will use this certificate.
	ID string `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                    `json:"name"`
	UpdatedAt time.Time                 `json:"updated_at" format:"date-time"`
	JSON      zeroTrustCertificatesJSON `json:"-"`
}

// zeroTrustCertificatesJSON contains the JSON metadata for the struct
// [ZeroTrustCertificates]
type zeroTrustCertificatesJSON struct {
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

func (r *ZeroTrustCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustCertificatesJSON) RawJSON() string {
	return r.raw
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

func (r accessCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateNewParams struct {
	// The certificate content.
	Certificate param.Field[string] `json:"certificate,required"`
	// The name of the certificate.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames"`
}

func (r AccessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateNewResponseEnvelope struct {
	Errors   []AccessCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustCertificates                          `json:"result,required"`
	// Whether the API call was successful
	Success AccessCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    accessCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// accessCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [AccessCertificateNewResponseEnvelope]
type accessCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                          `json:"code,required"`
	Message string                                         `json:"message,required"`
	JSON    accessCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// accessCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [AccessCertificateNewResponseEnvelopeErrors]
type accessCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type AccessCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    accessCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// accessCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [AccessCertificateNewResponseEnvelopeMessages]
type accessCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateNewResponseEnvelopeSuccess bool

const (
	AccessCertificateNewResponseEnvelopeSuccessTrue AccessCertificateNewResponseEnvelopeSuccess = true
)

func (r AccessCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r AccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccessCertificateUpdateResponseEnvelope struct {
	Errors   []AccessCertificateUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustCertificates                             `json:"result,required"`
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

func (r accessCertificateUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessCertificateUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessCertificateUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateUpdateResponseEnvelopeSuccess bool

const (
	AccessCertificateUpdateResponseEnvelopeSuccessTrue AccessCertificateUpdateResponseEnvelopeSuccess = true
)

func (r AccessCertificateUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessCertificateDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

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

func (r accessCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateDeleteResponseEnvelopeSuccess bool

const (
	AccessCertificateDeleteResponseEnvelopeSuccessTrue AccessCertificateDeleteResponseEnvelopeSuccess = true
)

func (r AccessCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type AccessCertificateGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type AccessCertificateGetResponseEnvelope struct {
	Errors   []AccessCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []AccessCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustCertificates                          `json:"result,required"`
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

func (r accessCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
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

func (r accessCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
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

func (r accessCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AccessCertificateGetResponseEnvelopeSuccess bool

const (
	AccessCertificateGetResponseEnvelopeSuccessTrue AccessCertificateGetResponseEnvelopeSuccess = true
)

func (r AccessCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AccessCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
