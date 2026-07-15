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
)

// ConnectorEventLatestService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorEventLatestService] method instead.
type ConnectorEventLatestService struct {
	Options []option.RequestOption
}

// NewConnectorEventLatestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectorEventLatestService(opts ...option.RequestOption) (r *ConnectorEventLatestService) {
	r = &ConnectorEventLatestService{}
	r.Options = opts
	return
}

// Fetches latest Magic WAN Connector Telemetry Events
func (r *ConnectorEventLatestService) List(ctx context.Context, connectorID string, query ConnectorEventLatestListParams, opts ...option.RequestOption) (res *ConnectorEventLatestListResponse, err error) {
	var env ConnectorEventLatestListResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s/telemetry/events/latest", query.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ConnectorEventLatestListResponse struct {
	Count float64                                `json:"count" api:"required"`
	Items []ConnectorEventLatestListResponseItem `json:"items" api:"required"`
	JSON  connectorEventLatestListResponseJSON   `json:"-"`
}

// connectorEventLatestListResponseJSON contains the JSON metadata for the struct
// [ConnectorEventLatestListResponse]
type connectorEventLatestListResponseJSON struct {
	Count       apijson.Field
	Items       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseJSON) RawJSON() string {
	return r.raw
}

// Recorded Event
type ConnectorEventLatestListResponseItem struct {
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
	E ConnectorEventLatestListResponseItemsE `json:"e" api:"required"`
	// Sequence number, used to order events with the same timestamp
	N float64 `json:"n" api:"required"`
	// Time the Event was recorded (seconds since the Unix epoch)
	T float64 `json:"t" api:"required"`
	// Version
	V    string                                   `json:"v"`
	JSON connectorEventLatestListResponseItemJSON `json:"-"`
}

// connectorEventLatestListResponseItemJSON contains the JSON metadata for the
// struct [ConnectorEventLatestListResponseItem]
type connectorEventLatestListResponseItemJSON struct {
	E           apijson.Field
	N           apijson.Field
	T           apijson.Field
	V           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItem) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemJSON) RawJSON() string {
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
type ConnectorEventLatestListResponseItemsE struct {
	// Event kind
	K           ConnectorEventLatestListResponseItemsEK    `json:"k" api:"required"`
	ExtraFields map[string]interface{}                     `json:"-" api:"extrafields"`
	JSON        connectorEventLatestListResponseItemsEJSON `json:"-"`
}

// connectorEventLatestListResponseItemsEJSON contains the JSON metadata for the
// struct [ConnectorEventLatestListResponseItemsE]
type connectorEventLatestListResponseItemsEJSON struct {
	K           apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseItemsE) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseItemsEJSON) RawJSON() string {
	return r.raw
}

// Event kind
type ConnectorEventLatestListResponseItemsEK string

const (
	ConnectorEventLatestListResponseItemsEKInit                        ConnectorEventLatestListResponseItemsEK = "Init"
	ConnectorEventLatestListResponseItemsEKLeave                       ConnectorEventLatestListResponseItemsEK = "Leave"
	ConnectorEventLatestListResponseItemsEKStartAttestation            ConnectorEventLatestListResponseItemsEK = "StartAttestation"
	ConnectorEventLatestListResponseItemsEKFinishAttestationSuccess    ConnectorEventLatestListResponseItemsEK = "FinishAttestationSuccess"
	ConnectorEventLatestListResponseItemsEKFinishAttestationFailure    ConnectorEventLatestListResponseItemsEK = "FinishAttestationFailure"
	ConnectorEventLatestListResponseItemsEKStartRotateCryptKey         ConnectorEventLatestListResponseItemsEK = "StartRotateCryptKey"
	ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeySuccess ConnectorEventLatestListResponseItemsEK = "FinishRotateCryptKeySuccess"
	ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeyFailure ConnectorEventLatestListResponseItemsEK = "FinishRotateCryptKeyFailure"
	ConnectorEventLatestListResponseItemsEKStartRotatePki              ConnectorEventLatestListResponseItemsEK = "StartRotatePki"
	ConnectorEventLatestListResponseItemsEKFinishRotatePkiSuccess      ConnectorEventLatestListResponseItemsEK = "FinishRotatePkiSuccess"
	ConnectorEventLatestListResponseItemsEKFinishRotatePkiFailure      ConnectorEventLatestListResponseItemsEK = "FinishRotatePkiFailure"
	ConnectorEventLatestListResponseItemsEKStartUpgrade                ConnectorEventLatestListResponseItemsEK = "StartUpgrade"
	ConnectorEventLatestListResponseItemsEKFinishUpgradeSuccess        ConnectorEventLatestListResponseItemsEK = "FinishUpgradeSuccess"
	ConnectorEventLatestListResponseItemsEKFinishUpgradeFailure        ConnectorEventLatestListResponseItemsEK = "FinishUpgradeFailure"
	ConnectorEventLatestListResponseItemsEKReconcile                   ConnectorEventLatestListResponseItemsEK = "Reconcile"
	ConnectorEventLatestListResponseItemsEKConfigureCloudflaredTunnel  ConnectorEventLatestListResponseItemsEK = "ConfigureCloudflaredTunnel"
	ConnectorEventLatestListResponseItemsEKRekeyInstallBoth            ConnectorEventLatestListResponseItemsEK = "RekeyInstallBoth"
	ConnectorEventLatestListResponseItemsEKRekeyStart                  ConnectorEventLatestListResponseItemsEK = "RekeyStart"
	ConnectorEventLatestListResponseItemsEKRekeyRestart                ConnectorEventLatestListResponseItemsEK = "RekeyRestart"
	ConnectorEventLatestListResponseItemsEKRekeyAdvance                ConnectorEventLatestListResponseItemsEK = "RekeyAdvance"
	ConnectorEventLatestListResponseItemsEKRekeyComplete               ConnectorEventLatestListResponseItemsEK = "RekeyComplete"
	ConnectorEventLatestListResponseItemsEKRekeyReset                  ConnectorEventLatestListResponseItemsEK = "RekeyReset"
)

func (r ConnectorEventLatestListResponseItemsEK) IsKnown() bool {
	switch r {
	case ConnectorEventLatestListResponseItemsEKInit, ConnectorEventLatestListResponseItemsEKLeave, ConnectorEventLatestListResponseItemsEKStartAttestation, ConnectorEventLatestListResponseItemsEKFinishAttestationSuccess, ConnectorEventLatestListResponseItemsEKFinishAttestationFailure, ConnectorEventLatestListResponseItemsEKStartRotateCryptKey, ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeySuccess, ConnectorEventLatestListResponseItemsEKFinishRotateCryptKeyFailure, ConnectorEventLatestListResponseItemsEKStartRotatePki, ConnectorEventLatestListResponseItemsEKFinishRotatePkiSuccess, ConnectorEventLatestListResponseItemsEKFinishRotatePkiFailure, ConnectorEventLatestListResponseItemsEKStartUpgrade, ConnectorEventLatestListResponseItemsEKFinishUpgradeSuccess, ConnectorEventLatestListResponseItemsEKFinishUpgradeFailure, ConnectorEventLatestListResponseItemsEKReconcile, ConnectorEventLatestListResponseItemsEKConfigureCloudflaredTunnel, ConnectorEventLatestListResponseItemsEKRekeyInstallBoth, ConnectorEventLatestListResponseItemsEKRekeyStart, ConnectorEventLatestListResponseItemsEKRekeyRestart, ConnectorEventLatestListResponseItemsEKRekeyAdvance, ConnectorEventLatestListResponseItemsEKRekeyComplete, ConnectorEventLatestListResponseItemsEKRekeyReset:
		return true
	}
	return false
}

type ConnectorEventLatestListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConnectorEventLatestListResponseEnvelope struct {
	Result   ConnectorEventLatestListResponse                   `json:"result" api:"required"`
	Success  bool                                               `json:"success" api:"required"`
	Errors   []ConnectorEventLatestListResponseEnvelopeErrors   `json:"errors"`
	Messages []ConnectorEventLatestListResponseEnvelopeMessages `json:"messages"`
	JSON     connectorEventLatestListResponseEnvelopeJSON       `json:"-"`
}

// connectorEventLatestListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ConnectorEventLatestListResponseEnvelope]
type connectorEventLatestListResponseEnvelopeJSON struct {
	Result      apijson.Field
	Success     apijson.Field
	Errors      apijson.Field
	Messages    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventLatestListResponseEnvelopeErrors struct {
	Code    float64                                            `json:"code" api:"required"`
	Message string                                             `json:"message" api:"required"`
	JSON    connectorEventLatestListResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEventLatestListResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [ConnectorEventLatestListResponseEnvelopeErrors]
type connectorEventLatestListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEventLatestListResponseEnvelopeMessages struct {
	Code    float64                                              `json:"code" api:"required"`
	Message string                                               `json:"message" api:"required"`
	JSON    connectorEventLatestListResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEventLatestListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ConnectorEventLatestListResponseEnvelopeMessages]
type connectorEventLatestListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEventLatestListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEventLatestListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
