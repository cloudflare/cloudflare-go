// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package workers_for_platforms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/workers"
)

// DispatchNamespaceScriptBindingService contains methods and other services that
// help with interacting with the cloudflare API. Note, unlike clients, this
// service does not read variables from the environment automatically. You should
// not instantiate this service directly, and instead use the
// [NewDispatchNamespaceScriptBindingService] method instead.
type DispatchNamespaceScriptBindingService struct {
	Options []option.RequestOption
}

// NewDispatchNamespaceScriptBindingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDispatchNamespaceScriptBindingService(opts ...option.RequestOption) (r *DispatchNamespaceScriptBindingService) {
	r = &DispatchNamespaceScriptBindingService{}
	r.Options = opts
	return
}

// Fetch script bindings from a script uploaded to a Workers for Platforms
// namespace.
func (r *DispatchNamespaceScriptBindingService) Get(ctx context.Context, dispatchNamespace string, scriptName string, query DispatchNamespaceScriptBindingGetParams, opts ...option.RequestOption) (res *workers.Binding, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/workers/dispatch/namespaces/%s/scripts/%s/bindings", query.AccountID, dispatchNamespace, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type DispatchNamespaceScriptBindingGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}
