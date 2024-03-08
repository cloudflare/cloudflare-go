// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-go/option"
)

// WorkersForPlatformDispatchNamespaceService contains methods and other services
// that help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewWorkersForPlatformDispatchNamespaceService] method instead.
type WorkersForPlatformDispatchNamespaceService struct {
	Options []option.RequestOption
	Scripts *WorkersForPlatformDispatchNamespaceScriptService
}

// NewWorkersForPlatformDispatchNamespaceService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewWorkersForPlatformDispatchNamespaceService(opts ...option.RequestOption) (r *WorkersForPlatformDispatchNamespaceService) {
	r = &WorkersForPlatformDispatchNamespaceService{}
	r.Options = opts
	r.Scripts = NewWorkersForPlatformDispatchNamespaceScriptService(opts...)
	return
}
