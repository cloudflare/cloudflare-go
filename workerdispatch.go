// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerDispatchService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerDispatchService] method
// instead.
type WorkerDispatchService struct {
	Options    []option.RequestOption
	Namespaces *WorkerDispatchNamespaceService
}

// NewWorkerDispatchService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerDispatchService(opts ...option.RequestOption) (r *WorkerDispatchService) {
	r = &WorkerDispatchService{}
	r.Options = opts
	r.Namespaces = NewWorkerDispatchNamespaceService(opts...)
	return
}
