// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package browser_rendering

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// DevtoolBrowserTargetService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevtoolBrowserTargetService] method instead.
type DevtoolBrowserTargetService struct {
	Options []option.RequestOption
}

// NewDevtoolBrowserTargetService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDevtoolBrowserTargetService(opts ...option.RequestOption) (r *DevtoolBrowserTargetService) {
	r = &DevtoolBrowserTargetService{}
	r.Options = opts
	return
}

// Opens a new tab in the browser. Optionally specify a URL to navigate to.
func (r *DevtoolBrowserTargetService) New(ctx context.Context, sessionID string, params DevtoolBrowserTargetNewParams, opts ...option.RequestOption) (res *DevtoolBrowserTargetNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/json/new", params.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns a list of all debuggable targets including tabs, pages, service workers,
// and other browser contexts.
func (r *DevtoolBrowserTargetService) List(ctx context.Context, sessionID string, query DevtoolBrowserTargetListParams, opts ...option.RequestOption) (res *[]DevtoolBrowserTargetListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/json/list", query.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Activates (brings to front) a specific browser target by its ID.
func (r *DevtoolBrowserTargetService) Activate(ctx context.Context, sessionID string, targetID string, query DevtoolBrowserTargetActivateParams, opts ...option.RequestOption) (res *DevtoolBrowserTargetActivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/json/activate/%s", query.AccountID, sessionID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Closes a specific browser target (tab, page, etc.) by its ID. Returns 'Target is
// closing' on success or an error if the target is not found.
func (r *DevtoolBrowserTargetService) Close(ctx context.Context, sessionID string, targetID string, query DevtoolBrowserTargetCloseParams, opts ...option.RequestOption) (res *DevtoolBrowserTargetCloseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/json/close/%s", query.AccountID, sessionID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the debuggable target with the given ID.
func (r *DevtoolBrowserTargetService) Get(ctx context.Context, sessionID string, targetID string, query DevtoolBrowserTargetGetParams, opts ...option.RequestOption) (res *DevtoolBrowserTargetGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	if targetID == "" {
		err = errors.New("missing required target_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/json/list/%s", query.AccountID, sessionID, targetID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DevtoolBrowserTargetNewResponse struct {
	// Target ID.
	ID string `json:"id" api:"required"`
	// Target type (page, background_page, worker, etc.).
	Type string `json:"type" api:"required"`
	// URL of the target.
	URL string `json:"url" api:"required"`
	// Target description.
	Description string `json:"description"`
	// DevTools frontend URL.
	DevtoolsFrontendURL string `json:"devtoolsFrontendUrl"`
	// Title of the target.
	Title string `json:"title"`
	// WebSocket URL for debugging this target.
	WebSocketDebuggerURL string                              `json:"webSocketDebuggerUrl"`
	JSON                 devtoolBrowserTargetNewResponseJSON `json:"-"`
}

// devtoolBrowserTargetNewResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserTargetNewResponse]
type devtoolBrowserTargetNewResponseJSON struct {
	ID                   apijson.Field
	Type                 apijson.Field
	URL                  apijson.Field
	Description          apijson.Field
	DevtoolsFrontendURL  apijson.Field
	Title                apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolBrowserTargetNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserTargetNewResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserTargetListResponse struct {
	// Target ID.
	ID string `json:"id" api:"required"`
	// Target type (page, background_page, worker, etc.).
	Type string `json:"type" api:"required"`
	// URL of the target.
	URL string `json:"url" api:"required"`
	// Target description.
	Description string `json:"description"`
	// DevTools frontend URL.
	DevtoolsFrontendURL string `json:"devtoolsFrontendUrl"`
	// Title of the target.
	Title string `json:"title"`
	// WebSocket URL for debugging this target.
	WebSocketDebuggerURL string                               `json:"webSocketDebuggerUrl"`
	JSON                 devtoolBrowserTargetListResponseJSON `json:"-"`
}

// devtoolBrowserTargetListResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserTargetListResponse]
type devtoolBrowserTargetListResponseJSON struct {
	ID                   apijson.Field
	Type                 apijson.Field
	URL                  apijson.Field
	Description          apijson.Field
	DevtoolsFrontendURL  apijson.Field
	Title                apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolBrowserTargetListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserTargetListResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserTargetActivateResponse struct {
	// Target activated.
	Message string                                   `json:"message" api:"required"`
	JSON    devtoolBrowserTargetActivateResponseJSON `json:"-"`
}

// devtoolBrowserTargetActivateResponseJSON contains the JSON metadata for the
// struct [DevtoolBrowserTargetActivateResponse]
type devtoolBrowserTargetActivateResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevtoolBrowserTargetActivateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserTargetActivateResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserTargetCloseResponse struct {
	// Target is closing.
	Message string                                `json:"message" api:"required"`
	JSON    devtoolBrowserTargetCloseResponseJSON `json:"-"`
}

// devtoolBrowserTargetCloseResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserTargetCloseResponse]
type devtoolBrowserTargetCloseResponseJSON struct {
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevtoolBrowserTargetCloseResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserTargetCloseResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserTargetGetResponse struct {
	// Target ID.
	ID string `json:"id" api:"required"`
	// Target type (page, background_page, worker, etc.).
	Type string `json:"type" api:"required"`
	// URL of the target.
	URL string `json:"url" api:"required"`
	// Target description.
	Description string `json:"description"`
	// DevTools frontend URL.
	DevtoolsFrontendURL string `json:"devtoolsFrontendUrl"`
	// Title of the target.
	Title string `json:"title"`
	// WebSocket URL for debugging this target.
	WebSocketDebuggerURL string                              `json:"webSocketDebuggerUrl"`
	JSON                 devtoolBrowserTargetGetResponseJSON `json:"-"`
}

// devtoolBrowserTargetGetResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserTargetGetResponse]
type devtoolBrowserTargetGetResponseJSON struct {
	ID                   apijson.Field
	Type                 apijson.Field
	URL                  apijson.Field
	Description          apijson.Field
	DevtoolsFrontendURL  apijson.Field
	Title                apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolBrowserTargetGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserTargetGetResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserTargetNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	URL       param.Field[string] `query:"url" format:"uri"`
}

// URLQuery serializes [DevtoolBrowserTargetNewParams]'s query parameters as
// `url.Values`.
func (r DevtoolBrowserTargetNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DevtoolBrowserTargetListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DevtoolBrowserTargetActivateParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DevtoolBrowserTargetCloseParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DevtoolBrowserTargetGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
