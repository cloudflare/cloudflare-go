// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// ScriptAssetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScriptAssetService] method instead.
type ScriptAssetService struct {
	Options []option.RequestOption
	Upload  *ScriptAssetUploadService
}

// NewScriptAssetService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptAssetService(opts ...option.RequestOption) (r *ScriptAssetService) {
	r = &ScriptAssetService{}
	r.Options = opts
	r.Upload = NewScriptAssetUploadService(opts...)
	return
}