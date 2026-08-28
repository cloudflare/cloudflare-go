// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// DevtoolBrowserLiveViewService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevtoolBrowserLiveViewService] method instead.
type DevtoolBrowserLiveViewService struct {
	Options []option.RequestOption
}

// NewDevtoolBrowserLiveViewService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevtoolBrowserLiveViewService(opts ...option.RequestOption) (r *DevtoolBrowserLiveViewService) {
	r = &DevtoolBrowserLiveViewService{}
	r.Options = opts
	return
}

// Generates time-limited URLs to view a remote browser session. Set
// `guardrails: { mode: 'readonly' }` to create a view-only link.
func (r *DevtoolBrowserLiveViewService) New(ctx context.Context, sessionID string, params DevtoolBrowserLiveViewNewParams, opts ...option.RequestOption) (res *DevtoolBrowserLiveViewNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/live_view", params.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type DevtoolBrowserLiveViewNewResponse struct {
	// Target ID
	ID string `json:"id" api:"required"`
	// URL to open the live view in a browser
	DevtoolsFrontendURL string                                   `json:"devtoolsFrontendUrl" api:"required" format:"uri"`
	Options             DevtoolBrowserLiveViewNewResponseOptions `json:"options" api:"required"`
	// WebSocket URL for CDP connection
	WebSocketDebuggerURL string                                `json:"webSocketDebuggerUrl" api:"required" format:"uri"`
	JSON                 devtoolBrowserLiveViewNewResponseJSON `json:"-"`
}

// devtoolBrowserLiveViewNewResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserLiveViewNewResponse]
type devtoolBrowserLiveViewNewResponseJSON struct {
	ID                   apijson.Field
	DevtoolsFrontendURL  apijson.Field
	Options              apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolBrowserLiveViewNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserLiveViewNewResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserLiveViewNewResponseOptions struct {
	// UI mode for the live view
	Mode DevtoolBrowserLiveViewNewResponseOptionsMode `json:"mode" api:"required"`
	// Connection guardrails applied to this link
	Guardrails DevtoolBrowserLiveViewNewResponseOptionsGuardrails `json:"guardrails"`
	JSON       devtoolBrowserLiveViewNewResponseOptionsJSON       `json:"-"`
}

// devtoolBrowserLiveViewNewResponseOptionsJSON contains the JSON metadata for the
// struct [DevtoolBrowserLiveViewNewResponseOptions]
type devtoolBrowserLiveViewNewResponseOptionsJSON struct {
	Mode        apijson.Field
	Guardrails  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevtoolBrowserLiveViewNewResponseOptions) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserLiveViewNewResponseOptionsJSON) RawJSON() string {
	return r.raw
}

// UI mode for the live view
type DevtoolBrowserLiveViewNewResponseOptionsMode string

const (
	DevtoolBrowserLiveViewNewResponseOptionsModeDevtools DevtoolBrowserLiveViewNewResponseOptionsMode = "devtools"
	DevtoolBrowserLiveViewNewResponseOptionsModeTab      DevtoolBrowserLiveViewNewResponseOptionsMode = "tab"
	DevtoolBrowserLiveViewNewResponseOptionsModeFull     DevtoolBrowserLiveViewNewResponseOptionsMode = "full"
)

func (r DevtoolBrowserLiveViewNewResponseOptionsMode) IsKnown() bool {
	switch r {
	case DevtoolBrowserLiveViewNewResponseOptionsModeDevtools, DevtoolBrowserLiveViewNewResponseOptionsModeTab, DevtoolBrowserLiveViewNewResponseOptionsModeFull:
		return true
	}
	return false
}

// Connection guardrails applied to this link
type DevtoolBrowserLiveViewNewResponseOptionsGuardrails struct {
	Mode DevtoolBrowserLiveViewNewResponseOptionsGuardrailsMode `json:"mode" api:"required"`
	JSON devtoolBrowserLiveViewNewResponseOptionsGuardrailsJSON `json:"-"`
}

// devtoolBrowserLiveViewNewResponseOptionsGuardrailsJSON contains the JSON
// metadata for the struct [DevtoolBrowserLiveViewNewResponseOptionsGuardrails]
type devtoolBrowserLiveViewNewResponseOptionsGuardrailsJSON struct {
	Mode        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevtoolBrowserLiveViewNewResponseOptionsGuardrails) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserLiveViewNewResponseOptionsGuardrailsJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserLiveViewNewResponseOptionsGuardrailsMode string

const (
	DevtoolBrowserLiveViewNewResponseOptionsGuardrailsModeReadonly DevtoolBrowserLiveViewNewResponseOptionsGuardrailsMode = "readonly"
)

func (r DevtoolBrowserLiveViewNewResponseOptionsGuardrailsMode) IsKnown() bool {
	switch r {
	case DevtoolBrowserLiveViewNewResponseOptionsGuardrailsModeReadonly:
		return true
	}
	return false
}

type DevtoolBrowserLiveViewNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// How long the live view URLs remain valid, in milliseconds. Default: 5 minutes.
	// Max: 60 minutes.
	ExpiresInMs param.Field[float64] `json:"expiresInMs"`
	// Connection guardrails. Use `{ mode: 'readonly' }` to generate a view-only link.
	Guardrails param.Field[DevtoolBrowserLiveViewNewParamsGuardrails] `json:"guardrails"`
	// UI mode: 'devtools' (Chrome DevTools), 'tab' (single tab view), 'full'
	// (multi-tab browser)
	Mode param.Field[DevtoolBrowserLiveViewNewParamsMode] `json:"mode"`
	// Target ID (page) to connect to. If omitted, auto-resolves to the first active
	// page.
	TargetID param.Field[string] `json:"targetId"`
}

func (r DevtoolBrowserLiveViewNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Connection guardrails. Use `{ mode: 'readonly' }` to generate a view-only link.
type DevtoolBrowserLiveViewNewParamsGuardrails struct {
	Mode param.Field[DevtoolBrowserLiveViewNewParamsGuardrailsMode] `json:"mode" api:"required"`
}

func (r DevtoolBrowserLiveViewNewParamsGuardrails) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DevtoolBrowserLiveViewNewParamsGuardrailsMode string

const (
	DevtoolBrowserLiveViewNewParamsGuardrailsModeReadonly DevtoolBrowserLiveViewNewParamsGuardrailsMode = "readonly"
)

func (r DevtoolBrowserLiveViewNewParamsGuardrailsMode) IsKnown() bool {
	switch r {
	case DevtoolBrowserLiveViewNewParamsGuardrailsModeReadonly:
		return true
	}
	return false
}

// UI mode: 'devtools' (Chrome DevTools), 'tab' (single tab view), 'full'
// (multi-tab browser)
type DevtoolBrowserLiveViewNewParamsMode string

const (
	DevtoolBrowserLiveViewNewParamsModeDevtools DevtoolBrowserLiveViewNewParamsMode = "devtools"
	DevtoolBrowserLiveViewNewParamsModeTab      DevtoolBrowserLiveViewNewParamsMode = "tab"
	DevtoolBrowserLiveViewNewParamsModeFull     DevtoolBrowserLiveViewNewParamsMode = "full"
)

func (r DevtoolBrowserLiveViewNewParamsMode) IsKnown() bool {
	switch r {
	case DevtoolBrowserLiveViewNewParamsModeDevtools, DevtoolBrowserLiveViewNewParamsModeTab, DevtoolBrowserLiveViewNewParamsModeFull:
		return true
	}
	return false
}
