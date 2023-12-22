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

// ZoneCustomCertificatePrioritizeService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewZoneCustomCertificatePrioritizeService] method instead.
type ZoneCustomCertificatePrioritizeService struct {
	Options []option.RequestOption
}

// NewZoneCustomCertificatePrioritizeService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZoneCustomCertificatePrioritizeService(opts ...option.RequestOption) (r *ZoneCustomCertificatePrioritizeService) {
	r = &ZoneCustomCertificatePrioritizeService{}
	r.Options = opts
	return
}

// If a zone has multiple SSL certificates, you can set the order in which they
// should be used during a request. The higher priority will break ties across
// overlapping 'legacy_custom' certificates.
func (r *ZoneCustomCertificatePrioritizeService) CustomSslForAZoneRePrioritizeSslCertificates(ctx context.Context, zoneIdentifier string, body ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParams, opts ...option.RequestOption) (res *CertificateResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("zones/%s/custom_certificates/prioritize", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParams struct {
	// Array of ordered certificates.
	Certificates param.Field[[]ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParamsCertificate] `json:"certificates,required"`
}

func (r ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParamsCertificate struct {
	// The order/priority in which the certificate will be used in a request. The
	// higher priority will break ties across overlapping 'legacy_custom' certificates,
	// but 'legacy_custom' certificates will always supercede 'sni_custom'
	// certificates.
	Priority param.Field[float64] `json:"priority"`
}

func (r ZoneCustomCertificatePrioritizeCustomSslForAZoneRePrioritizeSslCertificatesParamsCertificate) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
