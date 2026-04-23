// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// DevtoolBrowserPageService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevtoolBrowserPageService] method instead.
type DevtoolBrowserPageService struct {
	Options []option.RequestOption
}

// NewDevtoolBrowserPageService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevtoolBrowserPageService(opts ...option.RequestOption) (r *DevtoolBrowserPageService) {
	r = &DevtoolBrowserPageService{}
	r.Options = opts
	return
}

// Establishes a WebSocket connection to a specific Chrome DevTools target or page.
func (r *DevtoolBrowserPageService) Get(ctx context.Context, sessionID string, targetID string, query DevtoolBrowserPageGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return err
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/page/%s", query.AccountID, sessionID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return err
}

type DevtoolBrowserPageGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
