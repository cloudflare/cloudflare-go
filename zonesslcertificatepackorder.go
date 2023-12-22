// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSslCertificatePackOrderService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZoneSslCertificatePackOrderService] method instead.
type ZoneSslCertificatePackOrderService struct {
	Options []option.RequestOption
}

// NewZoneSslCertificatePackOrderService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneSslCertificatePackOrderService(opts ...option.RequestOption) (r *ZoneSslCertificatePackOrderService) {
	r = &ZoneSslCertificatePackOrderService{}
	r.Options = opts
	return
}

// For a given zone, order an advanced certificate pack.
func (r *ZoneSslCertificatePackOrderService) CertificatePacksOrderAdvancedCertificateManagerCertificatePack(ctx context.Context, zoneIdentifier string, body ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams, opts ...option.RequestOption) (res *AdvancedCertificatePackResponseSingleQ8qWa5qi, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/ssl/certificate_packs/order", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams struct {
	// Certificate Authority selected for the order. Selecting Let's Encrypt will
	// reduce customization of other fields: validation_method must be 'txt',
	// validity_days must be 90, cloudflare_branding must be omitted, and hosts must
	// contain only 2 entries, one for the zone name and one for the subdomain wildcard
	// of the zone name (e.g. example.com, \*.example.com).
	CertificateAuthority param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority] `json:"certificate_authority,required"`
	// Comma separated list of valid host names for the certificate packs. Must contain
	// the zone apex, may not contain more than 50 hosts, and may not be empty.
	Hosts param.Field[[]string] `json:"hosts,required"`
	// Type of certificate pack.
	Type param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType] `json:"type,required"`
	// Validation Method selected for the order.
	ValidationMethod param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod] `json:"validation_method,required"`
	// Validity Days selected for the order.
	ValidityDays param.Field[ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays] `json:"validity_days,required"`
	// Whether or not to add Cloudflare Branding for the order. This will add
	// sni.cloudflaressl.com as the Common Name if set true.
	CloudflareBranding param.Field[bool] `json:"cloudflare_branding"`
}

func (r ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Certificate Authority selected for the order. Selecting Let's Encrypt will
// reduce customization of other fields: validation_method must be 'txt',
// validity_days must be 90, cloudflare_branding must be omitted, and hosts must
// contain only 2 entries, one for the zone name and one for the subdomain wildcard
// of the zone name (e.g. example.com, \*.example.com).
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthorityDigicert    ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority = "digicert"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthorityGoogle      ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority = "google"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthorityLetsEncrypt ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsCertificateAuthority = "lets_encrypt"
)

// Type of certificate pack.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsTypeAdvanced ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsType = "advanced"
)

// Validation Method selected for the order.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod string

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodTxt   ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "txt"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodHTTP  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "http"
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethodEmail ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidationMethod = "email"
)

// Validity Days selected for the order.
type ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays int64

const (
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays14  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 14
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays30  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 30
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays90  ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 90
	ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays365 ZoneSslCertificatePackOrderCertificatePacksOrderAdvancedCertificateManagerCertificatePackParamsValidityDays = 365
)
