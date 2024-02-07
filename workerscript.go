// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerScriptService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptService] method
// instead.
type WorkerScriptService struct {
	Options   []option.RequestOption
	Content   *WorkerScriptContentService
	ContentV2 *WorkerScriptContentV2Service
	Settings  *WorkerScriptSettingService
}

// NewWorkerScriptService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkerScriptService(opts ...option.RequestOption) (r *WorkerScriptService) {
	r = &WorkerScriptService{}
	r.Options = opts
	r.Content = NewWorkerScriptContentService(opts...)
	r.ContentV2 = NewWorkerScriptContentV2Service(opts...)
	r.Settings = NewWorkerScriptSettingService(opts...)
	return
}
