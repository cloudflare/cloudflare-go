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

// ZoneSslCertificatePackQuotaService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSslCertificatePackQuotaService] method instead.
type ZoneSslCertificatePackQuotaService struct {
	Options []option.RequestOption
}

// NewZoneSslCertificatePackQuotaService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSslCertificatePackQuotaService(opts ...option.RequestOption) (r *ZoneSslCertificatePackQuotaService) {
	r = &ZoneSslCertificatePackQuotaService{}
	r.Options = opts
	return
}

// For a given zone, list certificate pack quotas.
func (r *ZoneSslCertificatePackQuotaService) CertificatePacksGetCertificatePackQuotas(ctx context.Context, zoneIdentifier string, opts ...option.RequestOption) (res *ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/quota", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse struct {
	Errors   []ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseError   `json:"errors"`
	Messages []ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseMessage `json:"messages"`
	Result   ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResult    `json:"result"`
	// Whether the API call was successful
	Success ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseSuccess `json:"success"`
	JSON    zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseJSON    `json:"-"`
}

// zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse]
type zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseError struct {
	Code    int64                                                                                `json:"code,required"`
	Message string                                                                               `json:"message,required"`
	JSON    zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseErrorJSON `json:"-"`
}

// zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseErrorJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseError]
type zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseMessage struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseMessageJSON `json:"-"`
}

// zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseMessageJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseMessage]
type zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResult struct {
	Advanced ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultAdvanced `json:"advanced"`
	JSON     zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultJSON     `json:"-"`
}

// zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResult]
type zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultJSON struct {
	Advanced    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResult) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultAdvanced struct {
	// Quantity Allocated.
	Allocated int64 `json:"allocated"`
	// Quantity Used.
	Used int64                                                                                         `json:"used"`
	JSON zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultAdvancedJSON `json:"-"`
}

// zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultAdvancedJSON
// contains the JSON metadata for the struct
// [ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultAdvanced]
type zoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultAdvancedJSON struct {
	Allocated   apijson.Field
	Used        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseResultAdvanced) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseSuccess bool

const (
	ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseSuccessTrue ZoneSslCertificatePackQuotaCertificatePacksGetCertificatePackQuotasResponseSuccess = true
)
