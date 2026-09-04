// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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
	"github.com/cloudflare/cloudflare-go/v7/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v7/shared"
)

// ConnectorInterruptService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorInterruptService] method instead.
type ConnectorInterruptService struct {
	Options []option.RequestOption
}

// NewConnectorInterruptService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorInterruptService(opts ...option.RequestOption) (r *ConnectorInterruptService) {
	r = &ConnectorInterruptService{}
	r.Options = opts
	return
}

// Creates an interrupt for a Magic WAN Connector.
func (r *ConnectorInterruptService) New(ctx context.Context, connectorID string, params ConnectorInterruptNewParams, opts ...option.RequestOption) (res *ConnectorInterruptNewResponse, err error) {
	var env ConnectorInterruptNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/interrupts", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Lists interrupts for a Magic WAN Connector.
func (r *ConnectorInterruptService) List(ctx context.Context, connectorID string, query ConnectorInterruptListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ConnectorInterruptListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/interrupts", query.AccountID, connectorID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Lists interrupts for a Magic WAN Connector.
func (r *ConnectorInterruptService) ListAutoPaging(ctx context.Context, connectorID string, query ConnectorInterruptListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ConnectorInterruptListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, connectorID, query, opts...))
}

// Interrupt action for a connector.
type ConnectorInterruptNewResponse struct {
	SubmittedAt string                                `json:"submitted_at" api:"required"`
	Reboot      ConnectorInterruptNewResponseReboot   `json:"reboot"`
	Restart     ConnectorInterruptNewResponseRestart  `json:"restart"`
	Shutdown    ConnectorInterruptNewResponseShutdown `json:"shutdown"`
	TriggeredAt string                                `json:"triggered_at"`
	JSON        connectorInterruptNewResponseJSON     `json:"-"`
}

// connectorInterruptNewResponseJSON contains the JSON metadata for the struct
// [ConnectorInterruptNewResponse]
type connectorInterruptNewResponseJSON struct {
	SubmittedAt apijson.Field
	Reboot      apijson.Field
	Restart     apijson.Field
	Shutdown    apijson.Field
	TriggeredAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptNewResponseReboot struct {
	// Purge connector state.
	Purge bool                                    `json:"purge"`
	JSON  connectorInterruptNewResponseRebootJSON `json:"-"`
}

// connectorInterruptNewResponseRebootJSON contains the JSON metadata for the
// struct [ConnectorInterruptNewResponseReboot]
type connectorInterruptNewResponseRebootJSON struct {
	Purge       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptNewResponseReboot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptNewResponseRebootJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptNewResponseRestart struct {
	// Purge connector state.
	Purge bool                                     `json:"purge"`
	JSON  connectorInterruptNewResponseRestartJSON `json:"-"`
}

// connectorInterruptNewResponseRestartJSON contains the JSON metadata for the
// struct [ConnectorInterruptNewResponseRestart]
type connectorInterruptNewResponseRestartJSON struct {
	Purge       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptNewResponseRestart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptNewResponseRestartJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptNewResponseShutdown struct {
	// Purge connector state.
	Purge bool                                      `json:"purge"`
	JSON  connectorInterruptNewResponseShutdownJSON `json:"-"`
}

// connectorInterruptNewResponseShutdownJSON contains the JSON metadata for the
// struct [ConnectorInterruptNewResponseShutdown]
type connectorInterruptNewResponseShutdownJSON struct {
	Purge       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptNewResponseShutdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptNewResponseShutdownJSON) RawJSON() string {
	return r.raw
}

// Interrupt action for a connector.
type ConnectorInterruptListResponse struct {
	SubmittedAt string                                 `json:"submitted_at" api:"required"`
	Reboot      ConnectorInterruptListResponseReboot   `json:"reboot"`
	Restart     ConnectorInterruptListResponseRestart  `json:"restart"`
	Shutdown    ConnectorInterruptListResponseShutdown `json:"shutdown"`
	TriggeredAt string                                 `json:"triggered_at"`
	JSON        connectorInterruptListResponseJSON     `json:"-"`
}

// connectorInterruptListResponseJSON contains the JSON metadata for the struct
// [ConnectorInterruptListResponse]
type connectorInterruptListResponseJSON struct {
	SubmittedAt apijson.Field
	Reboot      apijson.Field
	Restart     apijson.Field
	Shutdown    apijson.Field
	TriggeredAt apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptListResponseReboot struct {
	// Purge connector state.
	Purge bool                                     `json:"purge"`
	JSON  connectorInterruptListResponseRebootJSON `json:"-"`
}

// connectorInterruptListResponseRebootJSON contains the JSON metadata for the
// struct [ConnectorInterruptListResponseReboot]
type connectorInterruptListResponseRebootJSON struct {
	Purge       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptListResponseReboot) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptListResponseRebootJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptListResponseRestart struct {
	// Purge connector state.
	Purge bool                                      `json:"purge"`
	JSON  connectorInterruptListResponseRestartJSON `json:"-"`
}

// connectorInterruptListResponseRestartJSON contains the JSON metadata for the
// struct [ConnectorInterruptListResponseRestart]
type connectorInterruptListResponseRestartJSON struct {
	Purge       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptListResponseRestart) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptListResponseRestartJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptListResponseShutdown struct {
	// Purge connector state.
	Purge bool                                       `json:"purge"`
	JSON  connectorInterruptListResponseShutdownJSON `json:"-"`
}

// connectorInterruptListResponseShutdownJSON contains the JSON metadata for the
// struct [ConnectorInterruptListResponseShutdown]
type connectorInterruptListResponseShutdownJSON struct {
	Purge       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptListResponseShutdown) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptListResponseShutdownJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptNewParams struct {
	AccountID param.Field[string]                              `path:"account_id" api:"required"`
	Reboot    param.Field[ConnectorInterruptNewParamsReboot]   `json:"reboot"`
	Restart   param.Field[ConnectorInterruptNewParamsRestart]  `json:"restart"`
	Shutdown  param.Field[ConnectorInterruptNewParamsShutdown] `json:"shutdown"`
}

func (r ConnectorInterruptNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorInterruptNewParamsReboot struct {
	// Purge connector state.
	Purge param.Field[bool] `json:"purge"`
}

func (r ConnectorInterruptNewParamsReboot) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorInterruptNewParamsRestart struct {
	// Purge connector state.
	Purge param.Field[bool] `json:"purge"`
}

func (r ConnectorInterruptNewParamsRestart) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorInterruptNewParamsShutdown struct {
	// Purge connector state.
	Purge param.Field[bool] `json:"purge"`
}

func (r ConnectorInterruptNewParamsShutdown) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorInterruptNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors" api:"required"`
	Messages []shared.ResponseInfo `json:"messages" api:"required"`
	// Interrupt action for a connector.
	Result  ConnectorInterruptNewResponse             `json:"result" api:"required"`
	Success bool                                      `json:"success" api:"required"`
	JSON    connectorInterruptNewResponseEnvelopeJSON `json:"-"`
}

// connectorInterruptNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConnectorInterruptNewResponseEnvelope]
type connectorInterruptNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorInterruptNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorInterruptNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorInterruptListParams struct {
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
