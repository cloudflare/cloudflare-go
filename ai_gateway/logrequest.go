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

// LogRequestService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLogRequestService] method instead.
type LogRequestService struct {
	Options []option.RequestOption
}

// NewLogRequestService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLogRequestService(opts ...option.RequestOption) (r *LogRequestService) {
	r = &LogRequestService{}
	r.Options = opts
	return
}

// Get Gateway Log Request
func (r *LogRequestService) Get(ctx context.Context, id string, logID string, query LogRequestGetParams, opts ...option.RequestOption) (res *LogRequestGetResponse, err error) {
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
	path := fmt.Sprintf("accounts/%s/ai-gateway/gateways/%s/logs/%s/request", query.AccountID, id, logID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type LogRequestGetResponse = interface{}

type LogRequestGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
