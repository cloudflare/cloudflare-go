// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ai_gateway

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
)

// LogResponseService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogResponseService] method instead.
type LogResponseService struct {
	Options []option.RequestOption
}

// NewLogResponseService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogResponseService(opts ...option.RequestOption) (r *LogResponseService) {
	r = &LogResponseService{}
	r.Options = opts
	return
}

// Get Gateway Log Response
func (r *LogResponseService) Get(ctx context.Context, id string, logID string, query LogResponseGetParams, opts ...option.RequestOption) (res *LogResponseGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	if logID == "" {
		err = errors.New("missing required logId parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s/response", query.AccountID, id, logID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogResponseGetResponse = interface{}

type LogResponseGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
