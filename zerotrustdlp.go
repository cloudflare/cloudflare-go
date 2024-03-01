// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDLPService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDLPService] method
// instead.
type ZeroTrustDLPService struct {
	Options     []option.RequestOption
	Datasets    *ZeroTrustDLPDatasetService
	Patterns    *ZeroTrustDLPPatternService
	PayloadLogs *ZeroTrustDLPPayloadLogService
	Profiles    *ZeroTrustDLPProfileService
}

// NewZeroTrustDLPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewZeroTrustDLPService(opts ...option.RequestOption) (r *ZeroTrustDLPService) {
	r = &ZeroTrustDLPService{}
	r.Options = opts
	r.Datasets = NewZeroTrustDLPDatasetService(opts...)
	r.Patterns = NewZeroTrustDLPPatternService(opts...)
	r.PayloadLogs = NewZeroTrustDLPPayloadLogService(opts...)
	r.Profiles = NewZeroTrustDLPProfileService(opts...)
	return
}
