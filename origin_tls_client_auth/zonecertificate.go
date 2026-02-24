// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package origin_tls_client_auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ZoneCertificateService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneCertificateService] method instead.
type ZoneCertificateService struct {
	Options []option.RequestOption
}

// NewZoneCertificateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZoneCertificateService(opts ...option.RequestOption) (r *ZoneCertificateService) {
	r = &ZoneCertificateService{}
	r.Options = opts
	return
}

// Upload your own certificate you want Cloudflare to use for edge-to-origin
// communication to override the shared certificate. Please note that it is
// important to keep only one certificate active. Also, make sure to enable
// zone-level authenticated origin pulls by making a PUT call to settings endpoint
// to see the uploaded certificate in use.
func (r *ZoneCertificateService) New(ctx context.Context, params ZoneCertificateNewParams, opts ...option.RequestOption) (res *ZoneCertificateNewResponse, err error) {
	var env ZoneCertificateNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all client certificates configured for zone-level authenticated origin
// pulls.
func (r *ZoneCertificateService) List(ctx context.Context, query ZoneCertificateListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ZoneCertificateListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth", query.ZoneID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all client certificates configured for zone-level authenticated origin
// pulls.
func (r *ZoneCertificateService) ListAutoPaging(ctx context.Context, query ZoneCertificateListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ZoneCertificateListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Removes a client certificate used for zone-level authenticated origin pulls.
func (r *ZoneCertificateService) Delete(ctx context.Context, certificateID string, body ZoneCertificateDeleteParams, opts ...option.RequestOption) (res *ZoneCertificateDeleteResponse, err error) {
	var env ZoneCertificateDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", body.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retrieves details for a specific client certificate used in zone-level
// authenticated origin pulls.
func (r *ZoneCertificateService) Get(ctx context.Context, certificateID string, query ZoneCertificateGetParams, opts ...option.RequestOption) (res *ZoneCertificateGetResponse, err error) {
	var env ZoneCertificateGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	if certificateID == "" {
		err = errors.New("missing required certificate_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/origin_tls_client_auth/%s", query.ZoneID, certificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneAuthenticatedOriginPull struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// When the certificate from the authority expires.
	ExpiresOn time.Time `json:"expires_on" format:"date-time"`
	// The certificate authority that issued the certificate.
	Issuer string `json:"issuer"`
	// The type of hash used for the certificate.
	Signature string `json:"signature"`
	// Status of the certificate activation.
	Status ZoneAuthenticatedOriginPullStatus `json:"status"`
	// This is the time the certificate was uploaded.
	UploadedOn time.Time                       `json:"uploaded_on" format:"date-time"`
	JSON       zoneAuthenticatedOriginPullJSON `json:"-"`
}

// zoneAuthenticatedOriginPullJSON contains the JSON metadata for the struct
// [ZoneAuthenticatedOriginPull]
type zoneAuthenticatedOriginPullJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	ExpiresOn   apijson.Field
	Issuer      apijson.Field
	Signature   apijson.Field
	Status      apijson.Field
	UploadedOn  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneAuthenticatedOriginPull) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneAuthenticatedOriginPullJSON) RawJSON() string {
	return r.raw
}

// Status of the certificate activation.
type ZoneAuthenticatedOriginPullStatus string

const (
	ZoneAuthenticatedOriginPullStatusInitializing       ZoneAuthenticatedOriginPullStatus = "initializing"
	ZoneAuthenticatedOriginPullStatusPendingDeployment  ZoneAuthenticatedOriginPullStatus = "pending_deployment"
	ZoneAuthenticatedOriginPullStatusPendingDeletion    ZoneAuthenticatedOriginPullStatus = "pending_deletion"
	ZoneAuthenticatedOriginPullStatusActive             ZoneAuthenticatedOriginPullStatus = "active"
	ZoneAuthenticatedOriginPullStatusDeleted            ZoneAuthenticatedOriginPullStatus = "deleted"
	ZoneAuthenticatedOriginPullStatusDeploymentTimedOut ZoneAuthenticatedOriginPullStatus = "deployment_timed_out"
	ZoneAuthenticatedOriginPullStatusDeletionTimedOut   ZoneAuthenticatedOriginPullStatus = "deletion_timed_out"
)

func (r ZoneAuthenticatedOriginPullStatus) IsKnown() bool {
	switch r {
	case ZoneAuthenticatedOriginPullStatusInitializing, ZoneAuthenticatedOriginPullStatusPendingDeployment, ZoneAuthenticatedOriginPullStatusPendingDeletion, ZoneAuthenticatedOriginPullStatusActive, ZoneAuthenticatedOriginPullStatusDeleted, ZoneAuthenticatedOriginPullStatusDeploymentTimedOut, ZoneAuthenticatedOriginPullStatusDeletionTimedOut:
		return true
	}
	return false
}

type ZoneCertificateNewResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                         `json:"private_key"`
	JSON       zoneCertificateNewResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// zoneCertificateNewResponseJSON contains the JSON metadata for the struct
// [ZoneCertificateNewResponse]
type zoneCertificateNewResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateListResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                          `json:"private_key"`
	JSON       zoneCertificateListResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// zoneCertificateListResponseJSON contains the JSON metadata for the struct
// [ZoneCertificateListResponse]
type zoneCertificateListResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateListResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateDeleteResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                            `json:"private_key"`
	JSON       zoneCertificateDeleteResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// zoneCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneCertificateDeleteResponse]
type zoneCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateGetResponse struct {
	// Identifier.
	ID string `json:"id"`
	// The zone's leaf certificate.
	Certificate string `json:"certificate"`
	// Indicates whether zone-level authenticated origin pulls is enabled.
	Enabled bool `json:"enabled"`
	// The zone's private key.
	PrivateKey string                         `json:"private_key"`
	JSON       zoneCertificateGetResponseJSON `json:"-"`
	ZoneAuthenticatedOriginPull
}

// zoneCertificateGetResponseJSON contains the JSON metadata for the struct
// [ZoneCertificateGetResponse]
type zoneCertificateGetResponseJSON struct {
	ID          apijson.Field
	Certificate apijson.Field
	Enabled     apijson.Field
	PrivateKey  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateNewParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The zone's leaf certificate.
	Certificate param.Field[string] `json:"certificate,required"`
	// The zone's private key.
	PrivateKey param.Field[string] `json:"private_key,required"`
}

func (r ZoneCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCertificateNewResponseEnvelope struct {
	Errors   []ZoneCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneCertificateNewResponse                `json:"result"`
	JSON    zoneCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// zoneCertificateNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneCertificateNewResponseEnvelope]
type zoneCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateNewResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ZoneCertificateNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneCertificateNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneCertificateNewResponseEnvelopeErrors]
type zoneCertificateNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateNewResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    zoneCertificateNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneCertificateNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ZoneCertificateNewResponseEnvelopeErrorsSource]
type zoneCertificateNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateNewResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ZoneCertificateNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneCertificateNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneCertificateNewResponseEnvelopeMessages]
type zoneCertificateNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateNewResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    zoneCertificateNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneCertificateNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ZoneCertificateNewResponseEnvelopeMessagesSource]
type zoneCertificateNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneCertificateNewResponseEnvelopeSuccess bool

const (
	ZoneCertificateNewResponseEnvelopeSuccessTrue ZoneCertificateNewResponseEnvelopeSuccess = true
)

func (r ZoneCertificateNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneCertificateNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneCertificateListParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneCertificateDeleteParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneCertificateDeleteResponseEnvelope struct {
	Errors   []ZoneCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneCertificateDeleteResponse                `json:"result"`
	JSON    zoneCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// zoneCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneCertificateDeleteResponseEnvelope]
type zoneCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateDeleteResponseEnvelopeErrors struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ZoneCertificateDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneCertificateDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZoneCertificateDeleteResponseEnvelopeErrors]
type zoneCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    zoneCertificateDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneCertificateDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ZoneCertificateDeleteResponseEnvelopeErrorsSource]
type zoneCertificateDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateDeleteResponseEnvelopeMessages struct {
	Code             int64                                               `json:"code,required"`
	Message          string                                              `json:"message,required"`
	DocumentationURL string                                              `json:"documentation_url"`
	Source           ZoneCertificateDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneCertificateDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneCertificateDeleteResponseEnvelopeMessages]
type zoneCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                  `json:"pointer"`
	JSON    zoneCertificateDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneCertificateDeleteResponseEnvelopeMessagesSourceJSON contains the JSON
// metadata for the struct [ZoneCertificateDeleteResponseEnvelopeMessagesSource]
type zoneCertificateDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneCertificateDeleteResponseEnvelopeSuccess bool

const (
	ZoneCertificateDeleteResponseEnvelopeSuccessTrue ZoneCertificateDeleteResponseEnvelopeSuccess = true
)

func (r ZoneCertificateDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneCertificateDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneCertificateGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneCertificateGetResponseEnvelope struct {
	Errors   []ZoneCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZoneCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success ZoneCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneCertificateGetResponse                `json:"result"`
	JSON    zoneCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// zoneCertificateGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ZoneCertificateGetResponseEnvelope]
type zoneCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateGetResponseEnvelopeErrors struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ZoneCertificateGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             zoneCertificateGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// zoneCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ZoneCertificateGetResponseEnvelopeErrors]
type zoneCertificateGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateGetResponseEnvelopeErrorsSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    zoneCertificateGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// zoneCertificateGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ZoneCertificateGetResponseEnvelopeErrorsSource]
type zoneCertificateGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateGetResponseEnvelopeMessages struct {
	Code             int64                                            `json:"code,required"`
	Message          string                                           `json:"message,required"`
	DocumentationURL string                                           `json:"documentation_url"`
	Source           ZoneCertificateGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             zoneCertificateGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// zoneCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ZoneCertificateGetResponseEnvelopeMessages]
type zoneCertificateGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ZoneCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ZoneCertificateGetResponseEnvelopeMessagesSource struct {
	Pointer string                                               `json:"pointer"`
	JSON    zoneCertificateGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// zoneCertificateGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ZoneCertificateGetResponseEnvelopeMessagesSource]
type zoneCertificateGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneCertificateGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneCertificateGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ZoneCertificateGetResponseEnvelopeSuccess bool

const (
	ZoneCertificateGetResponseEnvelopeSuccessTrue ZoneCertificateGetResponseEnvelopeSuccess = true
)

func (r ZoneCertificateGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneCertificateGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
