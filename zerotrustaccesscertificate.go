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

// ZeroTrustAccessCertificateService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustAccessCertificateService] method instead.
type ZeroTrustAccessCertificateService struct {
	Options  []option.RequestOption
	Settings *ZeroTrustAccessCertificateSettingService
}

// NewZeroTrustAccessCertificateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessCertificateService(opts ...option.RequestOption) (r *ZeroTrustAccessCertificateService) {
	r = &ZeroTrustAccessCertificateService{}
	r.Options = opts
	r.Settings = NewZeroTrustAccessCertificateSettingService(opts...)
	return
}

// Adds a new mTLS root certificate to Access.
func (r *ZeroTrustAccessCertificateService) New(ctx context.Context, params ZeroTrustAccessCertificateNewParams, opts ...option.RequestOption) (res *AccessCertificates, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCertificateNewResponseEnvelope
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
func (r *ZeroTrustAccessCertificateService) Update(ctx context.Context, uuid string, params ZeroTrustAccessCertificateUpdateParams, opts ...option.RequestOption) (res *AccessCertificates, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCertificateUpdateResponseEnvelope
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
func (r *ZeroTrustAccessCertificateService) List(ctx context.Context, query ZeroTrustAccessCertificateListParams, opts ...option.RequestOption) (res *[]AccessCertificates, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCertificateListResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes an mTLS certificate.
func (r *ZeroTrustAccessCertificateService) Delete(ctx context.Context, uuid string, body ZeroTrustAccessCertificateDeleteParams, opts ...option.RequestOption) (res *ZeroTrustAccessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCertificateDeleteResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a single mTLS certificate.
func (r *ZeroTrustAccessCertificateService) Get(ctx context.Context, uuid string, query ZeroTrustAccessCertificateGetParams, opts ...option.RequestOption) (res *AccessCertificates, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessCertificateGetResponseEnvelope
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
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessCertificates struct {
	// The ID of the application that will use this certificate.
	ID interface{} `json:"id"`
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames []string  `json:"associated_hostnames"`
	CreatedAt           time.Time `json:"created_at" format:"date-time"`
	ExpiresOn           time.Time `json:"expires_on" format:"date-time"`
	// The MD5 fingerprint of the certificate.
	Fingerprint string `json:"fingerprint"`
	// The name of the certificate.
	Name      string                 `json:"name"`
	UpdatedAt time.Time              `json:"updated_at" format:"date-time"`
	JSON      accessCertificatesJSON `json:"-"`
}

// accessCertificatesJSON contains the JSON metadata for the struct
// [AccessCertificates]
type accessCertificatesJSON struct {
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

func (r *AccessCertificates) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateDeleteResponse struct {
	// UUID
	ID   string                                       `json:"id"`
	JSON zeroTrustAccessCertificateDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessCertificateDeleteResponse]
type zeroTrustAccessCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateNewParams struct {
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

func (r ZeroTrustAccessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessCertificateNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessCertificates                                      `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCertificateNewResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessCertificateNewResponseEnvelope]
type zeroTrustAccessCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCertificateNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateNewResponseEnvelopeErrors]
type zeroTrustAccessCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCertificateNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateNewResponseEnvelopeMessages]
type zeroTrustAccessCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessCertificateNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCertificateNewResponseEnvelopeSuccessTrue ZeroTrustAccessCertificateNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCertificateUpdateParams struct {
	// The hostnames of the applications that will use this certificate.
	AssociatedHostnames param.Field[[]string] `json:"associated_hostnames,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The name of the certificate.
	Name param.Field[string] `json:"name"`
}

func (r ZeroTrustAccessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessCertificateUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessCertificateUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCertificateUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessCertificates                                         `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCertificateUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCertificateUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCertificateUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessCertificateUpdateResponseEnvelope]
type zeroTrustAccessCertificateUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateUpdateResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessCertificateUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCertificateUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateUpdateResponseEnvelopeErrors]
type zeroTrustAccessCertificateUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateUpdateResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessCertificateUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCertificateUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessCertificateUpdateResponseEnvelopeMessages]
type zeroTrustAccessCertificateUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessCertificateUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCertificateUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessCertificateUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCertificateListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessCertificateListResponseEnvelope struct {
	Errors   []ZeroTrustAccessCertificateListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCertificateListResponseEnvelopeMessages `json:"messages,required"`
	Result   []AccessCertificates                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessCertificateListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessCertificateListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessCertificateListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessCertificateListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessCertificateListResponseEnvelope]
type zeroTrustAccessCertificateListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateListResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustAccessCertificateListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCertificateListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateListResponseEnvelopeErrors]
type zeroTrustAccessCertificateListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateListResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessCertificateListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCertificateListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateListResponseEnvelopeMessages]
type zeroTrustAccessCertificateListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessCertificateListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCertificateListResponseEnvelopeSuccessTrue ZeroTrustAccessCertificateListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCertificateListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                      `json:"total_count"`
	JSON       zeroTrustAccessCertificateListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessCertificateListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessCertificateListResponseEnvelopeResultInfo]
type zeroTrustAccessCertificateListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessCertificateDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessCertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCertificateDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessCertificateDeleteResponseEnvelope]
type zeroTrustAccessCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateDeleteResponseEnvelopeErrors]
type zeroTrustAccessCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessCertificateDeleteResponseEnvelopeMessages]
type zeroTrustAccessCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessCertificateDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCertificateDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessCertificateDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessCertificateGetParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessCertificateGetResponseEnvelope struct {
	Errors   []ZeroTrustAccessCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   AccessCertificates                                      `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessCertificateGetResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZeroTrustAccessCertificateGetResponseEnvelope]
type zeroTrustAccessCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                                   `json:"code,required"`
	Message string                                                  `json:"message,required"`
	JSON    zeroTrustAccessCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessCertificateGetResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateGetResponseEnvelopeErrors]
type zeroTrustAccessCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustAccessCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessCertificateGetResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessCertificateGetResponseEnvelopeMessages]
type zeroTrustAccessCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustAccessCertificateGetResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessCertificateGetResponseEnvelopeSuccessTrue ZeroTrustAccessCertificateGetResponseEnvelopeSuccess = true
)
