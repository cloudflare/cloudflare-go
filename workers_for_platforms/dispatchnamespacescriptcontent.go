// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// DispatchNamespaceScriptContentService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptContentService] method instead.
type DispatchNamespaceScriptContentService struct {
	Options  []option.RequestOption
	Scripts  *DispatchNamespaceScriptContentScriptService
	Settings *DispatchNamespaceScriptContentSettingService
	Bindings *DispatchNamespaceScriptContentBindingService
}

// NewDispatchNamespaceScriptContentService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptContentService(opts ...option.RequestOption) (r *DispatchNamespaceScriptContentService) {
	r = &DispatchNamespaceScriptContentService{}
	r.Options = opts
	r.Scripts = NewDispatchNamespaceScriptContentScriptService(opts...)
	r.Settings = NewDispatchNamespaceScriptContentSettingService(opts...)
	r.Bindings = NewDispatchNamespaceScriptContentBindingService(opts...)
	return
}
