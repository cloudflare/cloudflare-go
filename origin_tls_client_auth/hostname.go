// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// HostnameService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHostnameService] method instead.
type HostnameService struct {
	Options      []option.RequestOption
	Certificates *HostnameCertificateService
}

// NewHostnameService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHostnameService(opts ...option.RequestOption) (r *HostnameService) {
	r = &HostnameService{}
	r.Options = opts
	r.Certificates = NewHostnameCertificateService(opts...)
	return
}

// Associate a hostname to a certificate and enable, disable or invalidate the
// association. If disabled, client certificate will not be sent to the hostname
// even if activated at the zone level. 100 maximum associations on a single
// certificate are allowed. Note: Use a null value for parameter _enabled_ to
// invalidate the association.
func (r *HostnameService) Update(ctx context.Context, params HostnameUpdateParams, opts ...option.RequestOption) (res *[]AuthenticatedOriginPull, err error) {
	var env HostnameUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get the Hostname Status for Client Authentication
func (r *HostnameService) Get(ctx context.Context, hostname string, query HostnameGetParams, opts ...option.RequestOption) (res *AuthenticatedOriginPull, err error) {
	var env HostnameGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if hostname == "" {
		err = errors.New("missing required hostname parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/hostnames/%s", query.ZoneID, hostname)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AuthenticatedOriginPull struct {
	// Identifier
	CERTID string `json:"cert_id"`
	// Status of the certificate or the association.
	CERTStatus AuthenticatedOriginPullCERTStatus `json:"cert_status"`
	// The time when the certificate was updated.
	CERTUpdatedAt time.Time `json:"cert_updated_at" format:"date-time"`
	// The time when the certificate was uploaded.
	CERTUploadedOn time.Time `json:"cert_uploaded_on" format:"date-time"`
	// The hostname certificate.
	Certificate string `json:"certificate"`
	// The time when the certificate was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled bool `json:"enabled,nullable"`
	// The date when the certificate expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname string `json:"hostname"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The serial number on the uploaded certificate.
	SerialNumber string `json:"serial_number"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate or the association.
	Status AuthenticatedOriginPullStatus `json:"status"`
	// The time when the certificate was updated.
	UpdatedAt time.Time                   `json:"updated_at" format:"date-time"`
	JSON      authenticatedOriginPullJSON `json:"-"`
}

// authenticatedOriginPullJSON contains the JSON metadata for the struct
// [AuthenticatedOriginPull]
type authenticatedOriginPullJSON struct {
	CERTID         apijson.Field
	CERTStatus     apijson.Field
	CERTUpdatedAt  apijson.Field
	CERTUploadedOn apijson.Field
	Certificate    apijson.Field
	CreatedAt      apijson.Field
	Enabled        apijson.Field
	ExpiresOn      apijson.Field
	Hostname       apijson.Field
	Issuer         apijson.Field
	SerialNumber   apijson.Field
	Signature      apijson.Field
	Status         apijson.Field
	UpdatedAt      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AuthenticatedOriginPull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r authenticatedOriginPullJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate or the association.
type AuthenticatedOriginPullCERTStatus string

const (
	AuthenticatedOriginPullCERTStatusInitializing       AuthenticatedOriginPullCERTStatus = "initializing"
	AuthenticatedOriginPullCERTStatusPendingDeployment  AuthenticatedOriginPullCERTStatus = "pending_deployment"
	AuthenticatedOriginPullCERTStatusPendingDeletion    AuthenticatedOriginPullCERTStatus = "pending_deletion"
	AuthenticatedOriginPullCERTStatusActive             AuthenticatedOriginPullCERTStatus = "active"
	AuthenticatedOriginPullCERTStatusDeleted            AuthenticatedOriginPullCERTStatus = "deleted"
	AuthenticatedOriginPullCERTStatusDeploymentTimedOut AuthenticatedOriginPullCERTStatus = "deployment_timed_out"
	AuthenticatedOriginPullCERTStatusDeletionTimedOut   AuthenticatedOriginPullCERTStatus = "deletion_timed_out"
)

func (r AuthenticatedOriginPullCERTStatus) IsKnown() bool {
	switch r {
	case AuthenticatedOriginPullCERTStatusInitializing, AuthenticatedOriginPullCERTStatusPendingDeployment, AuthenticatedOriginPullCERTStatusPendingDeletion, AuthenticatedOriginPullCERTStatusActive, AuthenticatedOriginPullCERTStatusDeleted, AuthenticatedOriginPullCERTStatusDeploymentTimedOut, AuthenticatedOriginPullCERTStatusDeletionTimedOut:
		return true
	}
	return false
}

// Status of the certificate or the association.
type AuthenticatedOriginPullStatus string

const (
	AuthenticatedOriginPullStatusInitializing       AuthenticatedOriginPullStatus = "initializing"
	AuthenticatedOriginPullStatusPendingDeployment  AuthenticatedOriginPullStatus = "pending_deployment"
	AuthenticatedOriginPullStatusPendingDeletion    AuthenticatedOriginPullStatus = "pending_deletion"
	AuthenticatedOriginPullStatusActive             AuthenticatedOriginPullStatus = "active"
	AuthenticatedOriginPullStatusDeleted            AuthenticatedOriginPullStatus = "deleted"
	AuthenticatedOriginPullStatusDeploymentTimedOut AuthenticatedOriginPullStatus = "deployment_timed_out"
	AuthenticatedOriginPullStatusDeletionTimedOut   AuthenticatedOriginPullStatus = "deletion_timed_out"
)

func (r AuthenticatedOriginPullStatus) IsKnown() bool {
	switch r {
	case AuthenticatedOriginPullStatusInitializing, AuthenticatedOriginPullStatusPendingDeployment, AuthenticatedOriginPullStatusPendingDeletion, AuthenticatedOriginPullStatusActive, AuthenticatedOriginPullStatusDeleted, AuthenticatedOriginPullStatusDeploymentTimedOut, AuthenticatedOriginPullStatusDeletionTimedOut:
		return true
	}
	return false
}

type HostnameUpdateParams struct {
	// Identifier
	ZoneID param.Field[string]                       `path:"zone_id,required"`
	Config param.Field[[]HostnameUpdateParamsConfig] `json:"config,required"`
}

func (r HostnameUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameUpdateParamsConfig struct {
	// Certificate identifier tag.
	CERTID param.Field[string] `json:"cert_id"`
	// Indicates whether hostname-level authenticated origin pulls is enabled. A null
	// value voids the association.
	Enabled param.Field[bool] `json:"enabled"`
	// The hostname on the origin for which the client certificate uploaded will be
	// used.
	Hostname param.Field[string] `json:"hostname"`
}

func (r HostnameUpdateParamsConfig) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type HostnameUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success    HostnameUpdateResponseEnvelopeSuccess    `json:"success,required"`
	Result     []AuthenticatedOriginPull                `json:"result"`
	ResultInfo HostnameUpdateResponseEnvelopeResultInfo `json:"result_info"`
	JSON       hostnameUpdateResponseEnvelopeJSON       `json:"-"`
}

// hostnameUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameUpdateResponseEnvelope]
type hostnameUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameUpdateResponseEnvelopeSuccess bool

const (
	HostnameUpdateResponseEnvelopeSuccessTrue HostnameUpdateResponseEnvelopeSuccess = true
)

func (r HostnameUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type HostnameUpdateResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                      `json:"total_count"`
	JSON       hostnameUpdateResponseEnvelopeResultInfoJSON `json:"-"`
}

// hostnameUpdateResponseEnvelopeResultInfoJSON contains the JSON metadata for the
// struct [HostnameUpdateResponseEnvelopeResultInfo]
type hostnameUpdateResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameUpdateResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameUpdateResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type HostnameGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type HostnameGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success HostnameGetResponseEnvelopeSuccess `json:"success,required"`
	Result  AuthenticatedOriginPull            `json:"result"`
	JSON    hostnameGetResponseEnvelopeJSON    `json:"-"`
}

// hostnameGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [HostnameGetResponseEnvelope]
type hostnameGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HostnameGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r hostnameGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type HostnameGetResponseEnvelopeSuccess bool

const (
	HostnameGetResponseEnvelopeSuccessTrue HostnameGetResponseEnvelopeSuccess = true
)

func (r HostnameGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case HostnameGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
