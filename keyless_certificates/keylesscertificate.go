// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package keyless_certificates

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// KeylessCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewKeylessCertificateService] method
// instead.
type KeylessCertificateService struct {
	Options []option.RequestOption
}

// NewKeylessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewKeylessCertificateService(opts ...option.RequestOption) (r *KeylessCertificateService) {
	r = &KeylessCertificateService{}
	r.Options = opts
	return
}

// Create Keyless SSL Configuration
func (r *KeylessCertificateService) New(ctx context.Context, params KeylessCertificateNewParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesBase, err error) {
	opts = append(r.Options[:], opts...)
	var env KeylessCertificateNewResponseEnvelope
	path := fmt.Sprintf("zones/%s/keyless_certificates", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List all Keyless SSL configurations for a given zone.
func (r *KeylessCertificateService) List(ctx context.Context, query KeylessCertificateListParams, opts ...option.RequestOption) (res *[]TLSCertificatesAndHostnamesBase, err error) {
	opts = append(r.Options[:], opts...)
	var env KeylessCertificateListResponseEnvelope
	path := fmt.Sprintf("zones/%s/keyless_certificates", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete Keyless SSL Configuration
func (r *KeylessCertificateService) Delete(ctx context.Context, keylessCertificateID string, body KeylessCertificateDeleteParams, opts ...option.RequestOption) (res *KeylessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env KeylessCertificateDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", body.ZoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// This will update attributes of a Keyless SSL. Consists of one or more of the
// following: host,name,port.
func (r *KeylessCertificateService) Edit(ctx context.Context, keylessCertificateID string, params KeylessCertificateEditParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesBase, err error) {
	opts = append(r.Options[:], opts...)
	var env KeylessCertificateEditResponseEnvelope
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", params.ZoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for one Keyless SSL configuration.
func (r *KeylessCertificateService) Get(ctx context.Context, keylessCertificateID string, query KeylessCertificateGetParams, opts ...option.RequestOption) (res *TLSCertificatesAndHostnamesBase, err error) {
	opts = append(r.Options[:], opts...)
	var env KeylessCertificateGetResponseEnvelope
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", query.ZoneID, keylessCertificateID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type TLSCertificatesAndHostnamesBase struct {
	// Keyless certificate identifier tag.
	ID string `json:"id,required"`
	// When the Keyless SSL was created.
	CreatedOn time.Time `json:"created_on,required" format:"date-time"`
	// Whether or not the Keyless SSL is on or off.
	Enabled bool `json:"enabled,required"`
	// The keyless SSL name.
	Host string `json:"host,required" format:"hostname"`
	// When the Keyless SSL was last modified.
	ModifiedOn time.Time `json:"modified_on,required" format:"date-time"`
	// The keyless SSL name.
	Name string `json:"name,required"`
	// Available permissions for the Keyless SSL for the current user requesting the
	// item.
	Permissions []interface{} `json:"permissions,required"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port float64 `json:"port,required"`
	// Status of the Keyless SSL.
	Status TLSCertificatesAndHostnamesBaseStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel TLSCertificatesAndHostnamesBaseTunnel `json:"tunnel"`
	JSON   tlsCertificatesAndHostnamesBaseJSON   `json:"-"`
}

// tlsCertificatesAndHostnamesBaseJSON contains the JSON metadata for the struct
// [TLSCertificatesAndHostnamesBase]
type tlsCertificatesAndHostnamesBaseJSON struct {
	ID          apijson.Field
	CreatedOn   apijson.Field
	Enabled     apijson.Field
	Host        apijson.Field
	ModifiedOn  apijson.Field
	Name        apijson.Field
	Permissions apijson.Field
	Port        apijson.Field
	Status      apijson.Field
	Tunnel      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesBase) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesBaseJSON) RawJSON() string {
	return r.raw
}

// Status of the Keyless SSL.
type TLSCertificatesAndHostnamesBaseStatus string

const (
	TLSCertificatesAndHostnamesBaseStatusActive  TLSCertificatesAndHostnamesBaseStatus = "active"
	TLSCertificatesAndHostnamesBaseStatusDeleted TLSCertificatesAndHostnamesBaseStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type TLSCertificatesAndHostnamesBaseTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                    `json:"vnet_id,required"`
	JSON   tlsCertificatesAndHostnamesBaseTunnelJSON `json:"-"`
}

// tlsCertificatesAndHostnamesBaseTunnelJSON contains the JSON metadata for the
// struct [TLSCertificatesAndHostnamesBaseTunnel]
type tlsCertificatesAndHostnamesBaseTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *TLSCertificatesAndHostnamesBaseTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r tlsCertificatesAndHostnamesBaseTunnelJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteResponse struct {
	// Identifier
	ID   string                               `json:"id"`
	JSON keylessCertificateDeleteResponseJSON `json:"-"`
}

// keylessCertificateDeleteResponseJSON contains the JSON metadata for the struct
// [KeylessCertificateDeleteResponse]
type keylessCertificateDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// The zone's SSL certificate or SSL certificate and intermediate(s).
	Certificate param.Field[string] `json:"certificate,required"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host,required" format:"hostname"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port,required"`
	// A ubiquitous bundle has the highest probability of being verified everywhere,
	// even by clients using outdated or unusual trust stores. An optimal bundle uses
	// the shortest chain and newest intermediates. And the force bundle verifies the
	// chain, but does not otherwise modify it.
	BundleMethod param.Field[KeylessCertificateNewParamsBundleMethod] `json:"bundle_method"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[KeylessCertificateNewParamsTunnel] `json:"tunnel"`
}

func (r KeylessCertificateNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type KeylessCertificateNewParamsBundleMethod string

const (
	KeylessCertificateNewParamsBundleMethodUbiquitous KeylessCertificateNewParamsBundleMethod = "ubiquitous"
	KeylessCertificateNewParamsBundleMethodOptimal    KeylessCertificateNewParamsBundleMethod = "optimal"
	KeylessCertificateNewParamsBundleMethodForce      KeylessCertificateNewParamsBundleMethod = "force"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type KeylessCertificateNewParamsTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r KeylessCertificateNewParamsTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeylessCertificateNewResponseEnvelope struct {
	Errors   []KeylessCertificateNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateNewResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesBase                 `json:"result,required"`
	// Whether the API call was successful
	Success KeylessCertificateNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    keylessCertificateNewResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateNewResponseEnvelope]
type keylessCertificateNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    keylessCertificateNewResponseEnvelopeErrorsJSON `json:"-"`
}

// keylessCertificateNewResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KeylessCertificateNewResponseEnvelopeErrors]
type keylessCertificateNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateNewResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    keylessCertificateNewResponseEnvelopeMessagesJSON `json:"-"`
}

// keylessCertificateNewResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [KeylessCertificateNewResponseEnvelopeMessages]
type keylessCertificateNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeylessCertificateNewResponseEnvelopeSuccess bool

const (
	KeylessCertificateNewResponseEnvelopeSuccessTrue KeylessCertificateNewResponseEnvelopeSuccess = true
)

type KeylessCertificateListParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type KeylessCertificateListResponseEnvelope struct {
	Errors   []KeylessCertificateListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateListResponseEnvelopeMessages `json:"messages,required"`
	Result   []TLSCertificatesAndHostnamesBase                `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    KeylessCertificateListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo KeylessCertificateListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       keylessCertificateListResponseEnvelopeJSON       `json:"-"`
}

// keylessCertificateListResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateListResponseEnvelope]
type keylessCertificateListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateListResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    keylessCertificateListResponseEnvelopeErrorsJSON `json:"-"`
}

// keylessCertificateListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KeylessCertificateListResponseEnvelopeErrors]
type keylessCertificateListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateListResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    keylessCertificateListResponseEnvelopeMessagesJSON `json:"-"`
}

// keylessCertificateListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KeylessCertificateListResponseEnvelopeMessages]
type keylessCertificateListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeylessCertificateListResponseEnvelopeSuccess bool

const (
	KeylessCertificateListResponseEnvelopeSuccessTrue KeylessCertificateListResponseEnvelopeSuccess = true
)

type KeylessCertificateListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                              `json:"total_count"`
	JSON       keylessCertificateListResponseEnvelopeResultInfoJSON `json:"-"`
}

// keylessCertificateListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [KeylessCertificateListResponseEnvelopeResultInfo]
type keylessCertificateListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type KeylessCertificateDeleteResponseEnvelope struct {
	Errors   []KeylessCertificateDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   KeylessCertificateDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success KeylessCertificateDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    keylessCertificateDeleteResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateDeleteResponseEnvelope]
type keylessCertificateDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteResponseEnvelopeErrors struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    keylessCertificateDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [KeylessCertificateDeleteResponseEnvelopeErrors]
type keylessCertificateDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateDeleteResponseEnvelopeMessages struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    keylessCertificateDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// keylessCertificateDeleteResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KeylessCertificateDeleteResponseEnvelopeMessages]
type keylessCertificateDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeylessCertificateDeleteResponseEnvelopeSuccess bool

const (
	KeylessCertificateDeleteResponseEnvelopeSuccessTrue KeylessCertificateDeleteResponseEnvelopeSuccess = true
)

type KeylessCertificateEditParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
	// Whether or not the Keyless SSL is on or off.
	Enabled param.Field[bool] `json:"enabled"`
	// The keyless SSL name.
	Host param.Field[string] `json:"host" format:"hostname"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// The keyless SSL port used to communicate between Cloudflare and the client's
	// Keyless SSL server.
	Port param.Field[float64] `json:"port"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[KeylessCertificateEditParamsTunnel] `json:"tunnel"`
}

func (r KeylessCertificateEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type KeylessCertificateEditParamsTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r KeylessCertificateEditParamsTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type KeylessCertificateEditResponseEnvelope struct {
	Errors   []KeylessCertificateEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateEditResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesBase                  `json:"result,required"`
	// Whether the API call was successful
	Success KeylessCertificateEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    keylessCertificateEditResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateEditResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateEditResponseEnvelope]
type keylessCertificateEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateEditResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    keylessCertificateEditResponseEnvelopeErrorsJSON `json:"-"`
}

// keylessCertificateEditResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KeylessCertificateEditResponseEnvelopeErrors]
type keylessCertificateEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateEditResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    keylessCertificateEditResponseEnvelopeMessagesJSON `json:"-"`
}

// keylessCertificateEditResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [KeylessCertificateEditResponseEnvelopeMessages]
type keylessCertificateEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeylessCertificateEditResponseEnvelopeSuccess bool

const (
	KeylessCertificateEditResponseEnvelopeSuccessTrue KeylessCertificateEditResponseEnvelopeSuccess = true
)

type KeylessCertificateGetParams struct {
	// Identifier
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type KeylessCertificateGetResponseEnvelope struct {
	Errors   []KeylessCertificateGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []KeylessCertificateGetResponseEnvelopeMessages `json:"messages,required"`
	Result   TLSCertificatesAndHostnamesBase                 `json:"result,required"`
	// Whether the API call was successful
	Success KeylessCertificateGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    keylessCertificateGetResponseEnvelopeJSON    `json:"-"`
}

// keylessCertificateGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [KeylessCertificateGetResponseEnvelope]
type keylessCertificateGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateGetResponseEnvelopeErrors struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    keylessCertificateGetResponseEnvelopeErrorsJSON `json:"-"`
}

// keylessCertificateGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [KeylessCertificateGetResponseEnvelopeErrors]
type keylessCertificateGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type KeylessCertificateGetResponseEnvelopeMessages struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    keylessCertificateGetResponseEnvelopeMessagesJSON `json:"-"`
}

// keylessCertificateGetResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [KeylessCertificateGetResponseEnvelopeMessages]
type keylessCertificateGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *KeylessCertificateGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r keylessCertificateGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type KeylessCertificateGetResponseEnvelopeSuccess bool

const (
	KeylessCertificateGetResponseEnvelopeSuccessTrue KeylessCertificateGetResponseEnvelopeSuccess = true
)
