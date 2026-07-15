// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ConnectorEventService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorEventService] method instead.
type ConnectorEventService struct {
	Options []option.RequestOption
	Latest  *ConnectorEventLatestService
}

// NewConnectorEventService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorEventService(opts ...option.RequestOption) (r *ConnectorEventService) {
	r = &ConnectorEventService{}
	r.Options = opts
	r.Latest = NewConnectorEventLatestService(opts...)
	return
}

// Lists Magic WAN Connector Telemetry Events
func (r *ConnectorEventService) List(ctx context.Context, connectorID string, params ConnectorEventListParams, opts ...option.RequestOption) (res *ConnectorEventListResponse, err error) {
	var env ConnectorEventListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetches Magic WAN Connector Telemetry Event
func (r *ConnectorEventService) Get(ctx context.Context, connectorID string, eventT float64, eventN float64, query ConnectorEventGetParams, opts ...option.RequestOption) (res *ConnectorEventGetResponse, err error) {
	var env ConnectorEventGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events/%v.%v", query.AccountID, connectorID, eventT, eventN)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ConnectorEventListResponse struct {
	Count  float64                          `json:"count" api:"required"`
	Items  []ConnectorEventListResponseItem `json:"items" api:"required"`
	Cursor string                           `json:"cursor"`
	JSON   connectorEventListResponseJSON   `json:"-"`
}

// connectorEventListResponseJSON contains the JSON metadata for the struct
// [ConnectorEventListResponse]
type connectorEventListResponseJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	Cursor      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseItem struct {
	// Time the Event was collected (seconds since the Unix epoch)
	A float64 `json:"a" api:"required"`
	// Kind
	K string `json:"k" api:"required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n" api:"required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T    float64                            `json:"t" api:"required"`
	JSON connectorEventListResponseItemJSON `json:"-"`
}

// connectorEventListResponseItemJSON contains the JSON metadata for the struct
// [ConnectorEventListResponseItem]
type connectorEventListResponseItemJSON struct {
	A           apijson.Field
	K           apijson.Field
	N           apijson.Field
	T           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseItemJSON) RawJSON() string {
	return r.raw
}

// Recorded Event
type ConnectorEventGetResponse struct {
	// Event kind plus event-specific payload fields.
	//
	// Event kinds:
	//
	// - `Init`: Initialized process
	// - `Leave`: Stopped process
	// - `StartAttestation`: Started attestation
	// - `FinishAttestationSuccess`: Finished attestation
	// - `FinishAttestationFailure`: Failed attestation
	// - `StartRotateCryptKey`: Started crypt key rotation
	// - `FinishRotateCryptKeySuccess`: Finished crypt key rotation
	// - `FinishRotateCryptKeyFailure`: Failed crypt key rotation
	// - `StartRotatePki`: Started PKI rotation
	// - `FinishRotatePkiSuccess`: Finished PKI rotation
	// - `FinishRotatePkiFailure`: Failed PKI rotation
	// - `StartUpgrade`: Started upgrade
	// - `FinishUpgradeSuccess`: Finished upgrade
	// - `FinishUpgradeFailure`: Failed upgrade
	// - `Reconcile`: Reconciled
	// - `ConfigureCloudflaredTunnel`: Configured Cloudflared tunnel
	// - `RekeyInstallBoth`: Installed initial inbound and outbound keys
	// - `RekeyStart`: Installed new inbound key, kept old outbound
	// - `RekeyRestart`: Restarted in-progress rekey with newer key material
	// - `RekeyAdvance`: Confirmed traffic on new inbound key, swapped outbound to new
	// - `RekeyComplete`: Deleted old keys
	// - `RekeyReset`: Deleted all keys after receiving an unexpected key
	E ConnectorEventGetResponseE `json:"e" api:"required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n" api:"required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T float64 `json:"t" api:"required"`
	// Version
	V    string                        `json:"v"`
	JSON connectorEventGetResponseJSON `json:"-"`
}

// connectorEventGetResponseJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponse]
type connectorEventGetResponseJSON struct {
	E           apijson.Field
	N           apijson.Field
	T           apijson.Field
	V           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseJSON) RawJSON() string {
	return r.raw
}

// Event kind plus event-specific payload fields.
//
// Event kinds:
//
// - `Init`: Initialized process
// - `Leave`: Stopped process
// - `StartAttestation`: Started attestation
// - `FinishAttestationSuccess`: Finished attestation
// - `FinishAttestationFailure`: Failed attestation
// - `StartRotateCryptKey`: Started crypt key rotation
// - `FinishRotateCryptKeySuccess`: Finished crypt key rotation
// - `FinishRotateCryptKeyFailure`: Failed crypt key rotation
// - `StartRotatePki`: Started PKI rotation
// - `FinishRotatePkiSuccess`: Finished PKI rotation
// - `FinishRotatePkiFailure`: Failed PKI rotation
// - `StartUpgrade`: Started upgrade
// - `FinishUpgradeSuccess`: Finished upgrade
// - `FinishUpgradeFailure`: Failed upgrade
// - `Reconcile`: Reconciled
// - `ConfigureCloudflaredTunnel`: Configured Cloudflared tunnel
// - `RekeyInstallBoth`: Installed initial inbound and outbound keys
// - `RekeyStart`: Installed new inbound key, kept old outbound
// - `RekeyRestart`: Restarted in-progress rekey with newer key material
// - `RekeyAdvance`: Confirmed traffic on new inbound key, swapped outbound to new
// - `RekeyComplete`: Deleted old keys
// - `RekeyReset`: Deleted all keys after receiving an unexpected key
type ConnectorEventGetResponseE struct {
	// Event kind
	K           ConnectorEventGetResponseEK    `json:"k" api:"required"`
	ExtraFields map[string]interface{}         `json:"-" api:"extrafields"`
	JSON        connectorEventGetResponseEJSON `json:"-"`
}

// connectorEventGetResponseEJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseE]
type connectorEventGetResponseEJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEJSON) RawJSON() string {
	return r.raw
}

// Event kind
type ConnectorEventGetResponseEK string

const (
	ConnectorEventGetResponseEKInit                        ConnectorEventGetResponseEK = "Init"
	ConnectorEventGetResponseEKLeave                       ConnectorEventGetResponseEK = "Leave"
	ConnectorEventGetResponseEKStartAttestation            ConnectorEventGetResponseEK = "StartAttestation"
	ConnectorEventGetResponseEKFinishAttestationSuccess    ConnectorEventGetResponseEK = "FinishAttestationSuccess"
	ConnectorEventGetResponseEKFinishAttestationFailure    ConnectorEventGetResponseEK = "FinishAttestationFailure"
	ConnectorEventGetResponseEKStartRotateCryptKey         ConnectorEventGetResponseEK = "StartRotateCryptKey"
	ConnectorEventGetResponseEKFinishRotateCryptKeySuccess ConnectorEventGetResponseEK = "FinishRotateCryptKeySuccess"
	ConnectorEventGetResponseEKFinishRotateCryptKeyFailure ConnectorEventGetResponseEK = "FinishRotateCryptKeyFailure"
	ConnectorEventGetResponseEKStartRotatePki              ConnectorEventGetResponseEK = "StartRotatePki"
	ConnectorEventGetResponseEKFinishRotatePkiSuccess      ConnectorEventGetResponseEK = "FinishRotatePkiSuccess"
	ConnectorEventGetResponseEKFinishRotatePkiFailure      ConnectorEventGetResponseEK = "FinishRotatePkiFailure"
	ConnectorEventGetResponseEKStartUpgrade                ConnectorEventGetResponseEK = "StartUpgrade"
	ConnectorEventGetResponseEKFinishUpgradeSuccess        ConnectorEventGetResponseEK = "FinishUpgradeSuccess"
	ConnectorEventGetResponseEKFinishUpgradeFailure        ConnectorEventGetResponseEK = "FinishUpgradeFailure"
	ConnectorEventGetResponseEKReconcile                   ConnectorEventGetResponseEK = "Reconcile"
	ConnectorEventGetResponseEKConfigureCloudflaredTunnel  ConnectorEventGetResponseEK = "ConfigureCloudflaredTunnel"
	ConnectorEventGetResponseEKRekeyInstallBoth            ConnectorEventGetResponseEK = "RekeyInstallBoth"
	ConnectorEventGetResponseEKRekeyStart                  ConnectorEventGetResponseEK = "RekeyStart"
	ConnectorEventGetResponseEKRekeyRestart                ConnectorEventGetResponseEK = "RekeyRestart"
	ConnectorEventGetResponseEKRekeyAdvance                ConnectorEventGetResponseEK = "RekeyAdvance"
	ConnectorEventGetResponseEKRekeyComplete               ConnectorEventGetResponseEK = "RekeyComplete"
	ConnectorEventGetResponseEKRekeyReset                  ConnectorEventGetResponseEK = "RekeyReset"
)

func (r ConnectorEventGetResponseEK) IsKnown() bool {
	switch r {
	case ConnectorEventGetResponseEKInit, ConnectorEventGetResponseEKLeave, ConnectorEventGetResponseEKStartAttestation, ConnectorEventGetResponseEKFinishAttestationSuccess, ConnectorEventGetResponseEKFinishAttestationFailure, ConnectorEventGetResponseEKStartRotateCryptKey, ConnectorEventGetResponseEKFinishRotateCryptKeySuccess, ConnectorEventGetResponseEKFinishRotateCryptKeyFailure, ConnectorEventGetResponseEKStartRotatePki, ConnectorEventGetResponseEKFinishRotatePkiSuccess, ConnectorEventGetResponseEKFinishRotatePkiFailure, ConnectorEventGetResponseEKStartUpgrade, ConnectorEventGetResponseEKFinishUpgradeSuccess, ConnectorEventGetResponseEKFinishUpgradeFailure, ConnectorEventGetResponseEKReconcile, ConnectorEventGetResponseEKConfigureCloudflaredTunnel, ConnectorEventGetResponseEKRekeyInstallBoth, ConnectorEventGetResponseEKRekeyStart, ConnectorEventGetResponseEKRekeyRestart, ConnectorEventGetResponseEKRekeyAdvance, ConnectorEventGetResponseEKRekeyComplete, ConnectorEventGetResponseEKRekeyReset:
		return true
	}
	return false
}

type ConnectorEventListParams struct {
	// Account identifier
	AccountID param.Field[string]  `path:"account_id" api:"required"`
	From      param.Field[float64] `query:"from" api:"required"`
	To        param.Field[float64] `query:"to" api:"required"`
	Cursor    param.Field[string]  `query:"cursor"`
	// Filter by event kind
	K     param.Field[string]  `query:"k"`
	Limit param.Field[float64] `query:"limit"`
}

// URLQuery serializes [ConnectorEventListParams]'s query parameters as
// `url.Values`.
func (r ConnectorEventListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConnectorEventListResponseEnvelope struct {
	Result   ConnectorEventListResponse                   `json:"result" api:"required"`
	Success  bool                                         `json:"success" api:"required"`
	Errors   []ConnectorEventListResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventListResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventListResponseEnvelopeJSON       `json:"-"`
}

// connectorEventListResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEventListResponseEnvelope]
type connectorEventListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseEnvelopeErrors struct {
	Code    float64                                      `json:"code" api:"required"`
	Message string                                       `json:"message" api:"required"`
	JSON    connectorEventListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventListResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEventListResponseEnvelopeErrors]
type connectorEventListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventListResponseEnvelopeMessages struct {
	Code    float64                                        `json:"code" api:"required"`
	Message string                                         `json:"message" api:"required"`
	JSON    connectorEventListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventListResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ConnectorEventListResponseEnvelopeMessages]
type connectorEventListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConnectorEventGetResponseEnvelope struct {
	// Recorded Event
	Result   ConnectorEventGetResponse                   `json:"result" api:"required"`
	Success  bool                                        `json:"success" api:"required"`
	Errors   []ConnectorEventGetResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventGetResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventGetResponseEnvelopeJSON       `json:"-"`
}

// connectorEventGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEventGetResponseEnvelope]
type connectorEventGetResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseEnvelopeErrors struct {
	Code    float64                                     `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	JSON    connectorEventGetResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEnvelopeErrors]
type connectorEventGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventGetResponseEnvelopeMessages struct {
	Code    float64                                       `json:"code" api:"required"`
	Message string                                        `json:"message" api:"required"`
	JSON    connectorEventGetResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorEventGetResponseEnvelopeMessages]
type connectorEventGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
