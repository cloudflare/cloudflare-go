// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// WorkerScriptContentV2Service contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewWorkerScriptContentV2Service]
// method instead.
type WorkerScriptContentV2Service struct {
	Options []option.RequestOption
}

// NewWorkerScriptContentV2Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWorkerScriptContentV2Service(opts ...option.RequestOption) (r *WorkerScriptContentV2Service) {
	r = &WorkerScriptContentV2Service{}
	r.Options = opts
	return
}

// Fetch script content only
func (r *WorkerScriptContentV2Service) Get(ctx context.Context, accountID string, scriptName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "string")}, opts...)
	path := fmt.Sprintf("accounts/%s/workers/scripts/%s/content/v2", accountID, scriptName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}
