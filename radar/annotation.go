// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package radar

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// AnnotationService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAnnotationService] method instead.
type AnnotationService struct {
	Options []option.RequestOption
	Outages *AnnotationOutageService
}

// NewAnnotationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAnnotationService(opts ...option.RequestOption) (r *AnnotationService) {
	r = &AnnotationService{}
	r.Options = opts
	r.Outages = NewAnnotationOutageService(opts...)
	return
}
