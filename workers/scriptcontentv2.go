// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// ScriptContentV2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewScriptContentV2Service] method
// instead.
type ScriptContentV2Service struct {
	Options []option.RequestOption
}

// NewScriptContentV2Service generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScriptContentV2Service(opts ...option.RequestOption) (r *ScriptContentV2Service) {
	r = &ScriptContentV2Service{}
	r.Options = opts
	return
}

// Fetch script content only
func (r *ScriptContentV2Service) Get(ctx context.Context, scriptName string, query ScriptContentV2GetParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content/v2", query.AccountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ScriptContentV2GetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
