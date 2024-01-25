// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountWorkerDispatchNamespaceService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewAccountWorkerDispatchNamespaceService] method instead.
type AccountWorkerDispatchNamespaceService struct {
	Options []option.RequestOption
	Scripts *AccountWorkerDispatchNamespaceScriptService
}

// NewAccountWorkerDispatchNamespaceService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAccountWorkerDispatchNamespaceService(opts ...option.RequestOption) (r *AccountWorkerDispatchNamespaceService) {
	r = &AccountWorkerDispatchNamespaceService{}
	r.Options = opts
	r.Scripts = NewAccountWorkerDispatchNamespaceScriptService(opts...)
	return
}
