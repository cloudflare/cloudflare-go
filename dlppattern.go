// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// DLPPatternService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewDLPPatternService] method instead.
type DLPPatternService struct {
	Options   []option.RequestOption
	Validates *DLPPatternValidateService
}

// NewDLPPatternService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDLPPatternService(opts ...option.RequestOption) (r *DLPPatternService) {
	r = &DLPPatternService{}
	r.Options = opts
	r.Validates = NewDLPPatternValidateService(opts...)
	return
}
