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

// DevtoolBrowserService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevtoolBrowserService] method instead.
type DevtoolBrowserService struct {
	Options []option.RequestOption
	Page    *DevtoolBrowserPageService
	Targets *DevtoolBrowserTargetService
}

// NewDevtoolBrowserService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevtoolBrowserService(opts ...option.RequestOption) (r *DevtoolBrowserService) {
	r = &DevtoolBrowserService{}
	r.Options = opts
	r.Page = NewDevtoolBrowserPageService(opts...)
	r.Targets = NewDevtoolBrowserTargetService(opts...)
	return
}

// Get a browser session ID.
func (r *DevtoolBrowserService) New(ctx context.Context, params DevtoolBrowserNewParams, opts ...option.RequestOption) (res *DevtoolBrowserNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Closes an existing browser session.
func (r *DevtoolBrowserService) Delete(ctx context.Context, sessionID string, body DevtoolBrowserDeleteParams, opts ...option.RequestOption) (res *DevtoolBrowserDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s", body.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Establishes a WebSocket connection to an existing browser session.
func (r *DevtoolBrowserService) Connect(ctx context.Context, sessionID string, params DevtoolBrowserConnectParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s", params.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return err
}

// Acquires and establishes a WebSocket connection to a browser session.
func (r *DevtoolBrowserService) Launch(ctx context.Context, params DevtoolBrowserLaunchParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return err
}

// Returns the complete Chrome DevTools Protocol schema including all domains,
// commands, events, and types. This schema describes the entire CDP API surface.
func (r *DevtoolBrowserService) Protocol(ctx context.Context, sessionID string, query DevtoolBrowserProtocolParams, opts ...option.RequestOption) (res *DevtoolBrowserProtocolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/json/protocol", query.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Get browser version metadata.
func (r *DevtoolBrowserService) Version(ctx context.Context, sessionID string, query DevtoolBrowserVersionParams, opts ...option.RequestOption) (res *DevtoolBrowserVersionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/browser/%s/json/version", query.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DevtoolBrowserNewResponse struct {
	// Browser session ID.
	SessionID string `json:"sessionId" api:"required"`
	// WebSocket URL for the session.
	WebSocketDebuggerURL string                        `json:"webSocketDebuggerUrl"`
	JSON                 devtoolBrowserNewResponseJSON `json:"-"`
}

// devtoolBrowserNewResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserNewResponse]
type devtoolBrowserNewResponseJSON struct {
	SessionID            apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolBrowserNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserNewResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserDeleteResponse struct {
	Status DevtoolBrowserDeleteResponseStatus `json:"status" api:"required"`
	JSON   devtoolBrowserDeleteResponseJSON   `json:"-"`
}

// devtoolBrowserDeleteResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserDeleteResponse]
type devtoolBrowserDeleteResponseJSON struct {
	Status      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevtoolBrowserDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserDeleteResponseStatus string

const (
	DevtoolBrowserDeleteResponseStatusClosing DevtoolBrowserDeleteResponseStatus = "closing"
	DevtoolBrowserDeleteResponseStatusClosed  DevtoolBrowserDeleteResponseStatus = "closed"
)

func (r DevtoolBrowserDeleteResponseStatus) IsKnown() bool {
	switch r {
	case DevtoolBrowserDeleteResponseStatusClosing, DevtoolBrowserDeleteResponseStatusClosed:
		return true
	}
	return false
}

type DevtoolBrowserProtocolResponse struct {
	// List of protocol domains.
	Domains []DevtoolBrowserProtocolResponseDomain `json:"domains" api:"required"`
	// Protocol version.
	Version DevtoolBrowserProtocolResponseVersion `json:"version"`
	JSON    devtoolBrowserProtocolResponseJSON    `json:"-"`
}

// devtoolBrowserProtocolResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserProtocolResponse]
type devtoolBrowserProtocolResponseJSON struct {
	Domains     apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevtoolBrowserProtocolResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserProtocolResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserProtocolResponseDomain struct {
	// Domain name.
	Domain string `json:"domain" api:"required"`
	// Available commands.
	Commands []map[string]interface{} `json:"commands"`
	// Domain dependencies.
	Dependencies []string `json:"dependencies"`
	// Available events.
	Events []map[string]interface{} `json:"events"`
	// Whether this domain is experimental.
	Experimental bool `json:"experimental"`
	// Type definitions.
	Types []map[string]interface{}                 `json:"types"`
	JSON  devtoolBrowserProtocolResponseDomainJSON `json:"-"`
}

// devtoolBrowserProtocolResponseDomainJSON contains the JSON metadata for the
// struct [DevtoolBrowserProtocolResponseDomain]
type devtoolBrowserProtocolResponseDomainJSON struct {
	Domain       apijson.Field
	Commands     apijson.Field
	Dependencies apijson.Field
	Events       apijson.Field
	Experimental apijson.Field
	Types        apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *DevtoolBrowserProtocolResponseDomain) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserProtocolResponseDomainJSON) RawJSON() string {
	return r.raw
}

// Protocol version.
type DevtoolBrowserProtocolResponseVersion struct {
	// Major version.
	Major string `json:"major" api:"required"`
	// Minor version.
	Minor string                                    `json:"minor" api:"required"`
	JSON  devtoolBrowserProtocolResponseVersionJSON `json:"-"`
}

// devtoolBrowserProtocolResponseVersionJSON contains the JSON metadata for the
// struct [DevtoolBrowserProtocolResponseVersion]
type devtoolBrowserProtocolResponseVersionJSON struct {
	Major       apijson.Field
	Minor       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DevtoolBrowserProtocolResponseVersion) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserProtocolResponseVersionJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserVersionResponse struct {
	// Browser name and version.
	Browser string `json:"Browser" api:"required"`
	// Chrome DevTools Protocol version.
	ProtocolVersion string `json:"Protocol-Version" api:"required"`
	// User agent string.
	UserAgent string `json:"User-Agent" api:"required"`
	// V8 JavaScript engine version.
	V8Version string `json:"V8-Version" api:"required"`
	// WebKit version.
	WebKitVersion string `json:"WebKit-Version" api:"required"`
	// WebSocket URL for debugging the browser.
	WebSocketDebuggerURL string                            `json:"webSocketDebuggerUrl" api:"required"`
	JSON                 devtoolBrowserVersionResponseJSON `json:"-"`
}

// devtoolBrowserVersionResponseJSON contains the JSON metadata for the struct
// [DevtoolBrowserVersionResponse]
type devtoolBrowserVersionResponseJSON struct {
	Browser              apijson.Field
	ProtocolVersion      apijson.Field
	UserAgent            apijson.Field
	V8Version            apijson.Field
	WebKitVersion        apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolBrowserVersionResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolBrowserVersionResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolBrowserNewParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Keep-alive time in milliseconds.
	KeepAlive param.Field[float64] `query:"keep_alive"`
	// Use experimental browser.
	Lab       param.Field[bool] `query:"lab"`
	Recording param.Field[bool] `query:"recording"`
	// Include browser targets in response.
	Targets param.Field[bool] `query:"targets"`
}

// URLQuery serializes [DevtoolBrowserNewParams]'s query parameters as
// `url.Values`.
func (r DevtoolBrowserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DevtoolBrowserDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DevtoolBrowserConnectParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Keep-alive time in ms (only valid when acquiring new session).
	KeepAlive param.Field[float64] `query:"keep_alive"`
	// Use experimental browser.
	Lab       param.Field[bool] `query:"lab"`
	Recording param.Field[bool] `query:"recording"`
}

// URLQuery serializes [DevtoolBrowserConnectParams]'s query parameters as
// `url.Values`.
func (r DevtoolBrowserConnectParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DevtoolBrowserLaunchParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Keep-alive time in ms (only valid when acquiring new session).
	KeepAlive param.Field[float64] `query:"keep_alive"`
	// Use experimental browser.
	Lab       param.Field[bool] `query:"lab"`
	Recording param.Field[bool] `query:"recording"`
}

// URLQuery serializes [DevtoolBrowserLaunchParams]'s query parameters as
// `url.Values`.
func (r DevtoolBrowserLaunchParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DevtoolBrowserProtocolParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type DevtoolBrowserVersionParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
