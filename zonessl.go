// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZoneSslService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewZoneSslService] method instead.
type ZoneSslService struct {
	Options          []option.RequestOption
	Analyzes         *ZoneSslAnalyzeService
	CertificatePacks *ZoneSslCertificatePackService
	Recommendations  *ZoneSslRecommendationService
	Universals       *ZoneSslUniversalService
	Verifications    *ZoneSslVerificationService
}

// NewZoneSslService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewZoneSslService(opts ...option.RequestOption) (r *ZoneSslService) {
	r = &ZoneSslService{}
	r.Options = opts
	r.Analyzes = NewZoneSslAnalyzeService(opts...)
	r.CertificatePacks = NewZoneSslCertificatePackService(opts...)
	r.Recommendations = NewZoneSslRecommendationService(opts...)
	r.Universals = NewZoneSslUniversalService(opts...)
	r.Verifications = NewZoneSslVerificationService(opts...)
	return
}
