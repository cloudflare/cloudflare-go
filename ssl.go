// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SSLService contains methods and other services that help with interacting with
// the cloudflare API. Note, unlike clients, this service does not read variables
// from the environment automatically. You should not instantiate this service
// directly, and instead use the [NewSSLService] method instead.
type SSLService struct {
	Options          []option.RequestOption
	Analyzes         *SSLAnalyzeService
	CertificatePacks *SSLCertificatePackService
	Recommendations  *SSLRecommendationService
	Universals       *SSLUniversalService
	Verifications    *SSLVerificationService
}

// NewSSLService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSSLService(opts ...option.RequestOption) (r *SSLService) {
	r = &SSLService{}
	r.Options = opts
	r.Analyzes = NewSSLAnalyzeService(opts...)
	r.CertificatePacks = NewSSLCertificatePackService(opts...)
	r.Recommendations = NewSSLRecommendationService(opts...)
	r.Universals = NewSSLUniversalService(opts...)
	r.Verifications = NewSSLVerificationService(opts...)
	return
}
