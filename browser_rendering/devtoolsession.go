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

// DevtoolSessionService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevtoolSessionService] method instead.
type DevtoolSessionService struct {
	Options []option.RequestOption
}

// NewDevtoolSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevtoolSessionService(opts ...option.RequestOption) (r *DevtoolSessionService) {
	r = &DevtoolSessionService{}
	r.Options = opts
	return
}

// List active browser sessions.
func (r *DevtoolSessionService) List(ctx context.Context, params DevtoolSessionListParams, opts ...option.RequestOption) (res *[]DevtoolSessionListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/session", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Get details for a specific browser session.
func (r *DevtoolSessionService) Get(ctx context.Context, sessionID string, query DevtoolSessionGetParams, opts ...option.RequestOption) (res *DevtoolSessionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if sessionID == "" {
		err = errors.New("missing required session_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/browser-rendering/devtools/session/%s", query.AccountID, sessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type DevtoolSessionListResponse struct {
	// Session ID.
	SessionID string `json:"sessionId" api:"required" format:"uuid"`
	// Reason for session closure.
	CloseReason string `json:"closeReason"`
	// Human-readable close reason.
	CloseReasonText string `json:"closeReasonText"`
	// Connection end time.
	ConnectionEndTime float64 `json:"connectionEndTime"`
	// Connection ID.
	ConnectionID string `json:"connectionId"`
	// Connection start time.
	ConnectionStartTime float64 `json:"connectionStartTime"`
	// DevTools frontend URL.
	DevtoolsFrontendURL string `json:"devtoolsFrontendUrl"`
	// Session end time.
	EndTime float64 `json:"endTime"`
	// Last updated timestamp.
	LastUpdated float64 `json:"lastUpdated"`
	// Session start time.
	StartTime float64 `json:"startTime"`
	// WebSocket URL for debugging this target.
	WebSocketDebuggerURL string                         `json:"webSocketDebuggerUrl"`
	JSON                 devtoolSessionListResponseJSON `json:"-"`
}

// devtoolSessionListResponseJSON contains the JSON metadata for the struct
// [DevtoolSessionListResponse]
type devtoolSessionListResponseJSON struct {
	SessionID            apijson.Field
	CloseReason          apijson.Field
	CloseReasonText      apijson.Field
	ConnectionEndTime    apijson.Field
	ConnectionID         apijson.Field
	ConnectionStartTime  apijson.Field
	DevtoolsFrontendURL  apijson.Field
	EndTime              apijson.Field
	LastUpdated          apijson.Field
	StartTime            apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolSessionListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolSessionListResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolSessionGetResponse struct {
	// Session ID.
	SessionID string `json:"sessionId" api:"required" format:"uuid"`
	// Reason for session closure.
	CloseReason string `json:"closeReason"`
	// Human-readable close reason.
	CloseReasonText string `json:"closeReasonText"`
	// Connection end time.
	ConnectionEndTime float64 `json:"connectionEndTime"`
	// Connection ID.
	ConnectionID string `json:"connectionId"`
	// Connection start time.
	ConnectionStartTime float64 `json:"connectionStartTime"`
	// DevTools frontend URL.
	DevtoolsFrontendURL string `json:"devtoolsFrontendUrl"`
	// Session end time.
	EndTime float64 `json:"endTime"`
	// Last updated timestamp.
	LastUpdated float64 `json:"lastUpdated"`
	// Session start time.
	StartTime float64 `json:"startTime"`
	// WebSocket URL for debugging this target.
	WebSocketDebuggerURL string                        `json:"webSocketDebuggerUrl"`
	JSON                 devtoolSessionGetResponseJSON `json:"-"`
}

// devtoolSessionGetResponseJSON contains the JSON metadata for the struct
// [DevtoolSessionGetResponse]
type devtoolSessionGetResponseJSON struct {
	SessionID            apijson.Field
	CloseReason          apijson.Field
	CloseReasonText      apijson.Field
	ConnectionEndTime    apijson.Field
	ConnectionID         apijson.Field
	ConnectionStartTime  apijson.Field
	DevtoolsFrontendURL  apijson.Field
	EndTime              apijson.Field
	LastUpdated          apijson.Field
	StartTime            apijson.Field
	WebSocketDebuggerURL apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *DevtoolSessionGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r devtoolSessionGetResponseJSON) RawJSON() string {
	return r.raw
}

type DevtoolSessionListParams struct {
	// Account ID.
	AccountID param.Field[string]  `path:"account_id" api:"required"`
	Limit     param.Field[float64] `query:"limit"`
	Offset    param.Field[float64] `query:"offset"`
}

// URLQuery serializes [DevtoolSessionListParams]'s query parameters as
// `url.Values`.
func (r DevtoolSessionListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DevtoolSessionGetParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
