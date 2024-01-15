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

// ZoneKeylessCertificateService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZoneKeylessCertificateService]
// method instead.
type ZoneKeylessCertificateService struct {
	Options []option.RequestOption
}

// NewZoneKeylessCertificateService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneKeylessCertificateService(opts ...option.RequestOption) (r *ZoneKeylessCertificateService) {
	r = &ZoneKeylessCertificateService{}
	r.Options = opts
	return
}

// Get details for one Keyless SSL configuration.
func (r *ZoneKeylessCertificateService) Get(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneKeylessCertificateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This will update attributes of a Keyless SSL. Consists of one or more of the
// following: host,name,port.
func (r *ZoneKeylessCertificateService) Update(ctx context.Context, zoneIdentifier string, identifier string, body ZoneKeylessCertificateUpdateParams, opts ...option.RequestOption) (res *ZoneKeylessCertificateUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Delete Keyless SSL Configuration
func (r *ZoneKeylessCertificateService) Delete(ctx context.Context, zoneIdentifier string, identifier string, opts ...option.RequestOption) (res *ZoneKeylessCertificateDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates/%s", zoneIdentifier, identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Create Keyless SSL Configuration
func (r *ZoneKeylessCertificateService) KeylessSslForAZoneNewKeylessSslConfiguration(ctx context.Context, zoneIdentifier string, body ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParams, opts ...option.RequestOption) (res *ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all Keyless SSL configurations for a given zone.
func (r *ZoneKeylessCertificateService) KeylessSslForAZoneListKeylessSslConfigurations(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/keyless_certificates", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneKeylessCertificateGetResponse struct {
	Errors   []ZoneKeylessCertificateGetResponseError   `json:"errors"`
	Messages []ZoneKeylessCertificateGetResponseMessage `json:"messages"`
	Result   ZoneKeylessCertificateGetResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneKeylessCertificateGetResponseSuccess `json:"success"`
	JSON    zoneKeylessCertificateGetResponseJSON    `json:"-"`
}

// zoneKeylessCertificateGetResponseJSON contains the JSON metadata for the struct
// [ZoneKeylessCertificateGetResponse]
type zoneKeylessCertificateGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateGetResponseError struct {
	Code    int64                                      `json:"code,required"`
	Message string                                     `json:"message,required"`
	JSON    zoneKeylessCertificateGetResponseErrorJSON `json:"-"`
}

// zoneKeylessCertificateGetResponseErrorJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateGetResponseError]
type zoneKeylessCertificateGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateGetResponseMessage struct {
	Code    int64                                        `json:"code,required"`
	Message string                                       `json:"message,required"`
	JSON    zoneKeylessCertificateGetResponseMessageJSON `json:"-"`
}

// zoneKeylessCertificateGetResponseMessageJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateGetResponseMessage]
type zoneKeylessCertificateGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateGetResponseResult struct {
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
	Status ZoneKeylessCertificateGetResponseResultStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel ZoneKeylessCertificateGetResponseResultTunnel `json:"tunnel"`
	JSON   zoneKeylessCertificateGetResponseResultJSON   `json:"-"`
}

// zoneKeylessCertificateGetResponseResultJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateGetResponseResult]
type zoneKeylessCertificateGetResponseResultJSON struct {
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

func (r *ZoneKeylessCertificateGetResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type ZoneKeylessCertificateGetResponseResultStatus string

const (
	ZoneKeylessCertificateGetResponseResultStatusActive  ZoneKeylessCertificateGetResponseResultStatus = "active"
	ZoneKeylessCertificateGetResponseResultStatusDeleted ZoneKeylessCertificateGetResponseResultStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateGetResponseResultTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                            `json:"vnet_id,required"`
	JSON   zoneKeylessCertificateGetResponseResultTunnelJSON `json:"-"`
}

// zoneKeylessCertificateGetResponseResultTunnelJSON contains the JSON metadata for
// the struct [ZoneKeylessCertificateGetResponseResultTunnel]
type zoneKeylessCertificateGetResponseResultTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateGetResponseResultTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneKeylessCertificateGetResponseSuccess bool

const (
	ZoneKeylessCertificateGetResponseSuccessTrue ZoneKeylessCertificateGetResponseSuccess = true
)

type ZoneKeylessCertificateUpdateResponse struct {
	Errors   []ZoneKeylessCertificateUpdateResponseError   `json:"errors"`
	Messages []ZoneKeylessCertificateUpdateResponseMessage `json:"messages"`
	Result   ZoneKeylessCertificateUpdateResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneKeylessCertificateUpdateResponseSuccess `json:"success"`
	JSON    zoneKeylessCertificateUpdateResponseJSON    `json:"-"`
}

// zoneKeylessCertificateUpdateResponseJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateUpdateResponse]
type zoneKeylessCertificateUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateUpdateResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneKeylessCertificateUpdateResponseErrorJSON `json:"-"`
}

// zoneKeylessCertificateUpdateResponseErrorJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateUpdateResponseError]
type zoneKeylessCertificateUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateUpdateResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneKeylessCertificateUpdateResponseMessageJSON `json:"-"`
}

// zoneKeylessCertificateUpdateResponseMessageJSON contains the JSON metadata for
// the struct [ZoneKeylessCertificateUpdateResponseMessage]
type zoneKeylessCertificateUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateUpdateResponseResult struct {
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
	Status ZoneKeylessCertificateUpdateResponseResultStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel ZoneKeylessCertificateUpdateResponseResultTunnel `json:"tunnel"`
	JSON   zoneKeylessCertificateUpdateResponseResultJSON   `json:"-"`
}

// zoneKeylessCertificateUpdateResponseResultJSON contains the JSON metadata for
// the struct [ZoneKeylessCertificateUpdateResponseResult]
type zoneKeylessCertificateUpdateResponseResultJSON struct {
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

func (r *ZoneKeylessCertificateUpdateResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type ZoneKeylessCertificateUpdateResponseResultStatus string

const (
	ZoneKeylessCertificateUpdateResponseResultStatusActive  ZoneKeylessCertificateUpdateResponseResultStatus = "active"
	ZoneKeylessCertificateUpdateResponseResultStatusDeleted ZoneKeylessCertificateUpdateResponseResultStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateUpdateResponseResultTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                               `json:"vnet_id,required"`
	JSON   zoneKeylessCertificateUpdateResponseResultTunnelJSON `json:"-"`
}

// zoneKeylessCertificateUpdateResponseResultTunnelJSON contains the JSON metadata
// for the struct [ZoneKeylessCertificateUpdateResponseResultTunnel]
type zoneKeylessCertificateUpdateResponseResultTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateUpdateResponseResultTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneKeylessCertificateUpdateResponseSuccess bool

const (
	ZoneKeylessCertificateUpdateResponseSuccessTrue ZoneKeylessCertificateUpdateResponseSuccess = true
)

type ZoneKeylessCertificateDeleteResponse struct {
	Errors   []ZoneKeylessCertificateDeleteResponseError   `json:"errors"`
	Messages []ZoneKeylessCertificateDeleteResponseMessage `json:"messages"`
	Result   ZoneKeylessCertificateDeleteResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneKeylessCertificateDeleteResponseSuccess `json:"success"`
	JSON    zoneKeylessCertificateDeleteResponseJSON    `json:"-"`
}

// zoneKeylessCertificateDeleteResponseJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateDeleteResponse]
type zoneKeylessCertificateDeleteResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateDeleteResponseError struct {
	Code    int64                                         `json:"code,required"`
	Message string                                        `json:"message,required"`
	JSON    zoneKeylessCertificateDeleteResponseErrorJSON `json:"-"`
}

// zoneKeylessCertificateDeleteResponseErrorJSON contains the JSON metadata for the
// struct [ZoneKeylessCertificateDeleteResponseError]
type zoneKeylessCertificateDeleteResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateDeleteResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateDeleteResponseMessage struct {
	Code    int64                                           `json:"code,required"`
	Message string                                          `json:"message,required"`
	JSON    zoneKeylessCertificateDeleteResponseMessageJSON `json:"-"`
}

// zoneKeylessCertificateDeleteResponseMessageJSON contains the JSON metadata for
// the struct [ZoneKeylessCertificateDeleteResponseMessage]
type zoneKeylessCertificateDeleteResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateDeleteResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateDeleteResponseResult struct {
	// Identifier
	ID   string                                         `json:"id"`
	JSON zoneKeylessCertificateDeleteResponseResultJSON `json:"-"`
}

// zoneKeylessCertificateDeleteResponseResultJSON contains the JSON metadata for
// the struct [ZoneKeylessCertificateDeleteResponseResult]
type zoneKeylessCertificateDeleteResponseResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateDeleteResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneKeylessCertificateDeleteResponseSuccess bool

const (
	ZoneKeylessCertificateDeleteResponseSuccessTrue ZoneKeylessCertificateDeleteResponseSuccess = true
)

type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponse struct {
	Errors   []ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseError   `json:"errors"`
	Messages []ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseMessage `json:"messages"`
	Result   ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseSuccess `json:"success"`
	JSON    zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseJSON    `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponse]
type zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseError struct {
	Code    int64                                                                               `json:"code,required"`
	Message string                                                                              `json:"message,required"`
	JSON    zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseErrorJSON `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseError]
type zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseMessage struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseMessageJSON `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseMessage]
type zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResult struct {
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
	Status ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultTunnel `json:"tunnel"`
	JSON   zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultJSON   `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResult]
type zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultJSON struct {
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

func (r *ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultStatus string

const (
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultStatusActive  ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultStatus = "active"
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultStatusDeleted ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                                                                     `json:"vnet_id,required"`
	JSON   zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultTunnelJSON `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultTunnelJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultTunnel]
type zoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseResultTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseSuccess bool

const (
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseSuccessTrue ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationResponseSuccess = true
)

type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponse struct {
	Errors     []ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseError    `json:"errors"`
	Messages   []ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseMessage  `json:"messages"`
	Result     []ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResult   `json:"result"`
	ResultInfo ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseSuccess `json:"success"`
	JSON    zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseJSON    `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponse]
type zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseError struct {
	Code    int64                                                                                 `json:"code,required"`
	Message string                                                                                `json:"message,required"`
	JSON    zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseErrorJSON `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseError]
type zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseMessage struct {
	Code    int64                                                                                   `json:"code,required"`
	Message string                                                                                  `json:"message,required"`
	JSON    zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseMessageJSON `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseMessage]
type zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResult struct {
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
	Status ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultStatus `json:"status,required"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultTunnel `json:"tunnel"`
	JSON   zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultJSON   `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResult]
type zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultJSON struct {
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

func (r *ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Status of the Keyless SSL.
type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultStatus string

const (
	ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultStatusActive  ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultStatus = "active"
	ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultStatusDeleted ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultStatus = "deleted"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP string `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID string                                                                                       `json:"vnet_id,required"`
	JSON   zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultTunnelJSON `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultTunnelJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultTunnel]
type zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultTunnelJSON struct {
	PrivateIP   apijson.Field
	VnetID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultTunnel) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                    `json:"total_count"`
	JSON       zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultInfoJSON `json:"-"`
}

// zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultInfoJSON
// contains the JSON metadata for the struct
// [ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultInfo]
type zoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseSuccess bool

const (
	ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseSuccessTrue ZoneKeylessCertificateKeylessSslForAZoneListKeylessSslConfigurationsResponseSuccess = true
)

type ZoneKeylessCertificateUpdateParams struct {
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
	Tunnel param.Field[ZoneKeylessCertificateUpdateParamsTunnel] `json:"tunnel"`
}

func (r ZoneKeylessCertificateUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateUpdateParamsTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r ZoneKeylessCertificateUpdateParamsTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParams struct {
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
	BundleMethod param.Field[ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod] `json:"bundle_method"`
	// The keyless SSL name.
	Name param.Field[string] `json:"name"`
	// Configuration for using Keyless SSL through a Cloudflare Tunnel
	Tunnel param.Field[ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsTunnel] `json:"tunnel"`
}

func (r ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// A ubiquitous bundle has the highest probability of being verified everywhere,
// even by clients using outdated or unusual trust stores. An optimal bundle uses
// the shortest chain and newest intermediates. And the force bundle verifies the
// chain, but does not otherwise modify it.
type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod string

const (
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethodUbiquitous ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod = "ubiquitous"
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethodOptimal    ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod = "optimal"
	ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethodForce      ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsBundleMethod = "force"
)

// Configuration for using Keyless SSL through a Cloudflare Tunnel
type ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsTunnel struct {
	// Private IP of the Key Server Host
	PrivateIP param.Field[string] `json:"private_ip,required"`
	// Cloudflare Tunnel Virtual Network ID
	VnetID param.Field[string] `json:"vnet_id,required"`
}

func (r ZoneKeylessCertificateKeylessSslForAZoneNewKeylessSslConfigurationParamsTunnel) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
