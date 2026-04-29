// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package magic_transit

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
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ConnectorService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectorService] method instead.
type ConnectorService struct {
	Options   []option.RequestOption
	Events    *ConnectorEventService
	Snapshots *ConnectorSnapshotService
}

// NewConnectorService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConnectorService(opts ...option.RequestOption) (r *ConnectorService) {
	r = &ConnectorService{}
	r.Options = opts
	r.Events = NewConnectorEventService(opts...)
	r.Snapshots = NewConnectorSnapshotService(opts...)
	return
}

// Add a connector to your account
func (r *ConnectorService) New(ctx context.Context, params ConnectorNewParams, opts ...option.RequestOption) (res *ConnectorNewResponse, err error) {
	var env ConnectorNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Replace Connector or Re-provision License Key
func (r *ConnectorService) Update(ctx context.Context, connectorID string, params ConnectorUpdateParams, opts ...option.RequestOption) (res *ConnectorUpdateResponse, err error) {
	var env ConnectorUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// List Connectors
func (r *ConnectorService) List(ctx context.Context, params ConnectorListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ConnectorListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// List Connectors
func (r *ConnectorService) ListAutoPaging(ctx context.Context, params ConnectorListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ConnectorListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

// Remove a connector from your account
func (r *ConnectorService) Delete(ctx context.Context, connectorID string, body ConnectorDeleteParams, opts ...option.RequestOption) (res *ConnectorDeleteResponse, err error) {
	var env ConnectorDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", body.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Edit Connector to update specific properties or Re-provision License Key
func (r *ConnectorService) Edit(ctx context.Context, connectorID string, params ConnectorEditParams, opts ...option.RequestOption) (res *ConnectorEditResponse, err error) {
	var env ConnectorEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", params.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Fetch Connector
func (r *ConnectorService) Get(ctx context.Context, connectorID string, query ConnectorGetParams, opts ...option.RequestOption) (res *ConnectorGetResponse, err error) {
	var env ConnectorGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if connectorID == "" {
		err = errors.New("missing required connector_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/magic/connectors/%s", query.AccountID, connectorID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type ConnectorNewResponse struct {
	ID        string `json:"id" api:"required"`
	Activated bool   `json:"activated" api:"required"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    []ConnectorNewResponseInterruptWindowDaysOfWeek `json:"interrupt_window_days_of_week" api:"required"`
	InterruptWindowDurationHours float64                                         `json:"interrupt_window_duration_hours" api:"required"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates []string                   `json:"interrupt_window_embargo_dates" api:"required"`
	InterruptWindowHourOfDay    float64                    `json:"interrupt_window_hour_of_day" api:"required"`
	LastUpdated                 string                     `json:"last_updated" api:"required"`
	Notes                       string                     `json:"notes" api:"required"`
	Timezone                    string                     `json:"timezone" api:"required"`
	Device                      ConnectorNewResponseDevice `json:"device"`
	LastHeartbeat               string                     `json:"last_heartbeat"`
	LastSeenVersion             string                     `json:"last_seen_version"`
	LicenseKey                  string                     `json:"license_key"`
	JSON                        connectorNewResponseJSON   `json:"-"`
}

// connectorNewResponseJSON contains the JSON metadata for the struct
// [ConnectorNewResponse]
type connectorNewResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDaysOfWeek    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowEmbargoDates  apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	LicenseKey                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewResponseInterruptWindowDaysOfWeek string

const (
	ConnectorNewResponseInterruptWindowDaysOfWeekSunday    ConnectorNewResponseInterruptWindowDaysOfWeek = "Sunday"
	ConnectorNewResponseInterruptWindowDaysOfWeekMonday    ConnectorNewResponseInterruptWindowDaysOfWeek = "Monday"
	ConnectorNewResponseInterruptWindowDaysOfWeekTuesday   ConnectorNewResponseInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorNewResponseInterruptWindowDaysOfWeekWednesday ConnectorNewResponseInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorNewResponseInterruptWindowDaysOfWeekThursday  ConnectorNewResponseInterruptWindowDaysOfWeek = "Thursday"
	ConnectorNewResponseInterruptWindowDaysOfWeekFriday    ConnectorNewResponseInterruptWindowDaysOfWeek = "Friday"
	ConnectorNewResponseInterruptWindowDaysOfWeekSaturday  ConnectorNewResponseInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorNewResponseInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorNewResponseInterruptWindowDaysOfWeekSunday, ConnectorNewResponseInterruptWindowDaysOfWeekMonday, ConnectorNewResponseInterruptWindowDaysOfWeekTuesday, ConnectorNewResponseInterruptWindowDaysOfWeekWednesday, ConnectorNewResponseInterruptWindowDaysOfWeekThursday, ConnectorNewResponseInterruptWindowDaysOfWeekFriday, ConnectorNewResponseInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorNewResponseDevice struct {
	ID           string                         `json:"id" api:"required"`
	SerialNumber string                         `json:"serial_number"`
	Type         ConnectorNewResponseDeviceType `json:"type"`
	JSON         connectorNewResponseDeviceJSON `json:"-"`
}

// connectorNewResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorNewResponseDevice]
type connectorNewResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorNewResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewResponseDeviceType string

const (
	ConnectorNewResponseDeviceTypeManaged  ConnectorNewResponseDeviceType = "MANAGED"
	ConnectorNewResponseDeviceTypeLicensed ConnectorNewResponseDeviceType = "LICENSED"
)

func (r ConnectorNewResponseDeviceType) IsKnown() bool {
	switch r {
	case ConnectorNewResponseDeviceTypeManaged, ConnectorNewResponseDeviceTypeLicensed:
		return true
	}
	return false
}

type ConnectorUpdateResponse struct {
	ID        string `json:"id" api:"required"`
	Activated bool   `json:"activated" api:"required"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    []ConnectorUpdateResponseInterruptWindowDaysOfWeek `json:"interrupt_window_days_of_week" api:"required"`
	InterruptWindowDurationHours float64                                            `json:"interrupt_window_duration_hours" api:"required"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates []string                      `json:"interrupt_window_embargo_dates" api:"required"`
	InterruptWindowHourOfDay    float64                       `json:"interrupt_window_hour_of_day" api:"required"`
	LastUpdated                 string                        `json:"last_updated" api:"required"`
	Notes                       string                        `json:"notes" api:"required"`
	Timezone                    string                        `json:"timezone" api:"required"`
	Device                      ConnectorUpdateResponseDevice `json:"device"`
	LastHeartbeat               string                        `json:"last_heartbeat"`
	LastSeenVersion             string                        `json:"last_seen_version"`
	LicenseKey                  string                        `json:"license_key"`
	JSON                        connectorUpdateResponseJSON   `json:"-"`
}

// connectorUpdateResponseJSON contains the JSON metadata for the struct
// [ConnectorUpdateResponse]
type connectorUpdateResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDaysOfWeek    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowEmbargoDates  apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	LicenseKey                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponseInterruptWindowDaysOfWeek string

const (
	ConnectorUpdateResponseInterruptWindowDaysOfWeekSunday    ConnectorUpdateResponseInterruptWindowDaysOfWeek = "Sunday"
	ConnectorUpdateResponseInterruptWindowDaysOfWeekMonday    ConnectorUpdateResponseInterruptWindowDaysOfWeek = "Monday"
	ConnectorUpdateResponseInterruptWindowDaysOfWeekTuesday   ConnectorUpdateResponseInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorUpdateResponseInterruptWindowDaysOfWeekWednesday ConnectorUpdateResponseInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorUpdateResponseInterruptWindowDaysOfWeekThursday  ConnectorUpdateResponseInterruptWindowDaysOfWeek = "Thursday"
	ConnectorUpdateResponseInterruptWindowDaysOfWeekFriday    ConnectorUpdateResponseInterruptWindowDaysOfWeek = "Friday"
	ConnectorUpdateResponseInterruptWindowDaysOfWeekSaturday  ConnectorUpdateResponseInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorUpdateResponseInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorUpdateResponseInterruptWindowDaysOfWeekSunday, ConnectorUpdateResponseInterruptWindowDaysOfWeekMonday, ConnectorUpdateResponseInterruptWindowDaysOfWeekTuesday, ConnectorUpdateResponseInterruptWindowDaysOfWeekWednesday, ConnectorUpdateResponseInterruptWindowDaysOfWeekThursday, ConnectorUpdateResponseInterruptWindowDaysOfWeekFriday, ConnectorUpdateResponseInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorUpdateResponseDevice struct {
	ID           string                            `json:"id" api:"required"`
	SerialNumber string                            `json:"serial_number"`
	Type         ConnectorUpdateResponseDeviceType `json:"type"`
	JSON         connectorUpdateResponseDeviceJSON `json:"-"`
}

// connectorUpdateResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorUpdateResponseDevice]
type connectorUpdateResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorUpdateResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponseDeviceType string

const (
	ConnectorUpdateResponseDeviceTypeManaged  ConnectorUpdateResponseDeviceType = "MANAGED"
	ConnectorUpdateResponseDeviceTypeLicensed ConnectorUpdateResponseDeviceType = "LICENSED"
)

func (r ConnectorUpdateResponseDeviceType) IsKnown() bool {
	switch r {
	case ConnectorUpdateResponseDeviceTypeManaged, ConnectorUpdateResponseDeviceTypeLicensed:
		return true
	}
	return false
}

type ConnectorListResponse struct {
	ID        string `json:"id" api:"required"`
	Activated bool   `json:"activated" api:"required"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    []ConnectorListResponseInterruptWindowDaysOfWeek `json:"interrupt_window_days_of_week" api:"required"`
	InterruptWindowDurationHours float64                                          `json:"interrupt_window_duration_hours" api:"required"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates []string                    `json:"interrupt_window_embargo_dates" api:"required"`
	InterruptWindowHourOfDay    float64                     `json:"interrupt_window_hour_of_day" api:"required"`
	LastUpdated                 string                      `json:"last_updated" api:"required"`
	Notes                       string                      `json:"notes" api:"required"`
	Timezone                    string                      `json:"timezone" api:"required"`
	Device                      ConnectorListResponseDevice `json:"device"`
	LastHeartbeat               string                      `json:"last_heartbeat"`
	LastSeenVersion             string                      `json:"last_seen_version"`
	LicenseKey                  string                      `json:"license_key"`
	JSON                        connectorListResponseJSON   `json:"-"`
}

// connectorListResponseJSON contains the JSON metadata for the struct
// [ConnectorListResponse]
type connectorListResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDaysOfWeek    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowEmbargoDates  apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	LicenseKey                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorListResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorListResponseInterruptWindowDaysOfWeek string

const (
	ConnectorListResponseInterruptWindowDaysOfWeekSunday    ConnectorListResponseInterruptWindowDaysOfWeek = "Sunday"
	ConnectorListResponseInterruptWindowDaysOfWeekMonday    ConnectorListResponseInterruptWindowDaysOfWeek = "Monday"
	ConnectorListResponseInterruptWindowDaysOfWeekTuesday   ConnectorListResponseInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorListResponseInterruptWindowDaysOfWeekWednesday ConnectorListResponseInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorListResponseInterruptWindowDaysOfWeekThursday  ConnectorListResponseInterruptWindowDaysOfWeek = "Thursday"
	ConnectorListResponseInterruptWindowDaysOfWeekFriday    ConnectorListResponseInterruptWindowDaysOfWeek = "Friday"
	ConnectorListResponseInterruptWindowDaysOfWeekSaturday  ConnectorListResponseInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorListResponseInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorListResponseInterruptWindowDaysOfWeekSunday, ConnectorListResponseInterruptWindowDaysOfWeekMonday, ConnectorListResponseInterruptWindowDaysOfWeekTuesday, ConnectorListResponseInterruptWindowDaysOfWeekWednesday, ConnectorListResponseInterruptWindowDaysOfWeekThursday, ConnectorListResponseInterruptWindowDaysOfWeekFriday, ConnectorListResponseInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorListResponseDevice struct {
	ID           string                          `json:"id" api:"required"`
	SerialNumber string                          `json:"serial_number"`
	Type         ConnectorListResponseDeviceType `json:"type"`
	JSON         connectorListResponseDeviceJSON `json:"-"`
}

// connectorListResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorListResponseDevice]
type connectorListResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorListResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorListResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorListResponseDeviceType string

const (
	ConnectorListResponseDeviceTypeManaged  ConnectorListResponseDeviceType = "MANAGED"
	ConnectorListResponseDeviceTypeLicensed ConnectorListResponseDeviceType = "LICENSED"
)

func (r ConnectorListResponseDeviceType) IsKnown() bool {
	switch r {
	case ConnectorListResponseDeviceTypeManaged, ConnectorListResponseDeviceTypeLicensed:
		return true
	}
	return false
}

type ConnectorDeleteResponse struct {
	ID        string `json:"id" api:"required"`
	Activated bool   `json:"activated" api:"required"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    []ConnectorDeleteResponseInterruptWindowDaysOfWeek `json:"interrupt_window_days_of_week" api:"required"`
	InterruptWindowDurationHours float64                                            `json:"interrupt_window_duration_hours" api:"required"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates []string                      `json:"interrupt_window_embargo_dates" api:"required"`
	InterruptWindowHourOfDay    float64                       `json:"interrupt_window_hour_of_day" api:"required"`
	LastUpdated                 string                        `json:"last_updated" api:"required"`
	Notes                       string                        `json:"notes" api:"required"`
	Timezone                    string                        `json:"timezone" api:"required"`
	Device                      ConnectorDeleteResponseDevice `json:"device"`
	LastHeartbeat               string                        `json:"last_heartbeat"`
	LastSeenVersion             string                        `json:"last_seen_version"`
	LicenseKey                  string                        `json:"license_key"`
	JSON                        connectorDeleteResponseJSON   `json:"-"`
}

// connectorDeleteResponseJSON contains the JSON metadata for the struct
// [ConnectorDeleteResponse]
type connectorDeleteResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDaysOfWeek    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowEmbargoDates  apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	LicenseKey                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponseInterruptWindowDaysOfWeek string

const (
	ConnectorDeleteResponseInterruptWindowDaysOfWeekSunday    ConnectorDeleteResponseInterruptWindowDaysOfWeek = "Sunday"
	ConnectorDeleteResponseInterruptWindowDaysOfWeekMonday    ConnectorDeleteResponseInterruptWindowDaysOfWeek = "Monday"
	ConnectorDeleteResponseInterruptWindowDaysOfWeekTuesday   ConnectorDeleteResponseInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorDeleteResponseInterruptWindowDaysOfWeekWednesday ConnectorDeleteResponseInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorDeleteResponseInterruptWindowDaysOfWeekThursday  ConnectorDeleteResponseInterruptWindowDaysOfWeek = "Thursday"
	ConnectorDeleteResponseInterruptWindowDaysOfWeekFriday    ConnectorDeleteResponseInterruptWindowDaysOfWeek = "Friday"
	ConnectorDeleteResponseInterruptWindowDaysOfWeekSaturday  ConnectorDeleteResponseInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorDeleteResponseInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorDeleteResponseInterruptWindowDaysOfWeekSunday, ConnectorDeleteResponseInterruptWindowDaysOfWeekMonday, ConnectorDeleteResponseInterruptWindowDaysOfWeekTuesday, ConnectorDeleteResponseInterruptWindowDaysOfWeekWednesday, ConnectorDeleteResponseInterruptWindowDaysOfWeekThursday, ConnectorDeleteResponseInterruptWindowDaysOfWeekFriday, ConnectorDeleteResponseInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorDeleteResponseDevice struct {
	ID           string                            `json:"id" api:"required"`
	SerialNumber string                            `json:"serial_number"`
	Type         ConnectorDeleteResponseDeviceType `json:"type"`
	JSON         connectorDeleteResponseDeviceJSON `json:"-"`
}

// connectorDeleteResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorDeleteResponseDevice]
type connectorDeleteResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorDeleteResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponseDeviceType string

const (
	ConnectorDeleteResponseDeviceTypeManaged  ConnectorDeleteResponseDeviceType = "MANAGED"
	ConnectorDeleteResponseDeviceTypeLicensed ConnectorDeleteResponseDeviceType = "LICENSED"
)

func (r ConnectorDeleteResponseDeviceType) IsKnown() bool {
	switch r {
	case ConnectorDeleteResponseDeviceTypeManaged, ConnectorDeleteResponseDeviceTypeLicensed:
		return true
	}
	return false
}

type ConnectorEditResponse struct {
	ID        string `json:"id" api:"required"`
	Activated bool   `json:"activated" api:"required"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    []ConnectorEditResponseInterruptWindowDaysOfWeek `json:"interrupt_window_days_of_week" api:"required"`
	InterruptWindowDurationHours float64                                          `json:"interrupt_window_duration_hours" api:"required"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates []string                    `json:"interrupt_window_embargo_dates" api:"required"`
	InterruptWindowHourOfDay    float64                     `json:"interrupt_window_hour_of_day" api:"required"`
	LastUpdated                 string                      `json:"last_updated" api:"required"`
	Notes                       string                      `json:"notes" api:"required"`
	Timezone                    string                      `json:"timezone" api:"required"`
	Device                      ConnectorEditResponseDevice `json:"device"`
	LastHeartbeat               string                      `json:"last_heartbeat"`
	LastSeenVersion             string                      `json:"last_seen_version"`
	LicenseKey                  string                      `json:"license_key"`
	JSON                        connectorEditResponseJSON   `json:"-"`
}

// connectorEditResponseJSON contains the JSON metadata for the struct
// [ConnectorEditResponse]
type connectorEditResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDaysOfWeek    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowEmbargoDates  apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	LicenseKey                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponseInterruptWindowDaysOfWeek string

const (
	ConnectorEditResponseInterruptWindowDaysOfWeekSunday    ConnectorEditResponseInterruptWindowDaysOfWeek = "Sunday"
	ConnectorEditResponseInterruptWindowDaysOfWeekMonday    ConnectorEditResponseInterruptWindowDaysOfWeek = "Monday"
	ConnectorEditResponseInterruptWindowDaysOfWeekTuesday   ConnectorEditResponseInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorEditResponseInterruptWindowDaysOfWeekWednesday ConnectorEditResponseInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorEditResponseInterruptWindowDaysOfWeekThursday  ConnectorEditResponseInterruptWindowDaysOfWeek = "Thursday"
	ConnectorEditResponseInterruptWindowDaysOfWeekFriday    ConnectorEditResponseInterruptWindowDaysOfWeek = "Friday"
	ConnectorEditResponseInterruptWindowDaysOfWeekSaturday  ConnectorEditResponseInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorEditResponseInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorEditResponseInterruptWindowDaysOfWeekSunday, ConnectorEditResponseInterruptWindowDaysOfWeekMonday, ConnectorEditResponseInterruptWindowDaysOfWeekTuesday, ConnectorEditResponseInterruptWindowDaysOfWeekWednesday, ConnectorEditResponseInterruptWindowDaysOfWeekThursday, ConnectorEditResponseInterruptWindowDaysOfWeekFriday, ConnectorEditResponseInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorEditResponseDevice struct {
	ID           string                          `json:"id" api:"required"`
	SerialNumber string                          `json:"serial_number"`
	Type         ConnectorEditResponseDeviceType `json:"type"`
	JSON         connectorEditResponseDeviceJSON `json:"-"`
}

// connectorEditResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorEditResponseDevice]
type connectorEditResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorEditResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponseDeviceType string

const (
	ConnectorEditResponseDeviceTypeManaged  ConnectorEditResponseDeviceType = "MANAGED"
	ConnectorEditResponseDeviceTypeLicensed ConnectorEditResponseDeviceType = "LICENSED"
)

func (r ConnectorEditResponseDeviceType) IsKnown() bool {
	switch r {
	case ConnectorEditResponseDeviceTypeManaged, ConnectorEditResponseDeviceTypeLicensed:
		return true
	}
	return false
}

type ConnectorGetResponse struct {
	ID        string `json:"id" api:"required"`
	Activated bool   `json:"activated" api:"required"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    []ConnectorGetResponseInterruptWindowDaysOfWeek `json:"interrupt_window_days_of_week" api:"required"`
	InterruptWindowDurationHours float64                                         `json:"interrupt_window_duration_hours" api:"required"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates []string                   `json:"interrupt_window_embargo_dates" api:"required"`
	InterruptWindowHourOfDay    float64                    `json:"interrupt_window_hour_of_day" api:"required"`
	LastUpdated                 string                     `json:"last_updated" api:"required"`
	Notes                       string                     `json:"notes" api:"required"`
	Timezone                    string                     `json:"timezone" api:"required"`
	Device                      ConnectorGetResponseDevice `json:"device"`
	LastHeartbeat               string                     `json:"last_heartbeat"`
	LastSeenVersion             string                     `json:"last_seen_version"`
	LicenseKey                  string                     `json:"license_key"`
	JSON                        connectorGetResponseJSON   `json:"-"`
}

// connectorGetResponseJSON contains the JSON metadata for the struct
// [ConnectorGetResponse]
type connectorGetResponseJSON struct {
	ID                           apijson.Field
	Activated                    apijson.Field
	InterruptWindowDaysOfWeek    apijson.Field
	InterruptWindowDurationHours apijson.Field
	InterruptWindowEmbargoDates  apijson.Field
	InterruptWindowHourOfDay     apijson.Field
	LastUpdated                  apijson.Field
	Notes                        apijson.Field
	Timezone                     apijson.Field
	Device                       apijson.Field
	LastHeartbeat                apijson.Field
	LastSeenVersion              apijson.Field
	LicenseKey                   apijson.Field
	raw                          string
	ExtraFields                  map[string]apijson.Field
}

func (r *ConnectorGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponseInterruptWindowDaysOfWeek string

const (
	ConnectorGetResponseInterruptWindowDaysOfWeekSunday    ConnectorGetResponseInterruptWindowDaysOfWeek = "Sunday"
	ConnectorGetResponseInterruptWindowDaysOfWeekMonday    ConnectorGetResponseInterruptWindowDaysOfWeek = "Monday"
	ConnectorGetResponseInterruptWindowDaysOfWeekTuesday   ConnectorGetResponseInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorGetResponseInterruptWindowDaysOfWeekWednesday ConnectorGetResponseInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorGetResponseInterruptWindowDaysOfWeekThursday  ConnectorGetResponseInterruptWindowDaysOfWeek = "Thursday"
	ConnectorGetResponseInterruptWindowDaysOfWeekFriday    ConnectorGetResponseInterruptWindowDaysOfWeek = "Friday"
	ConnectorGetResponseInterruptWindowDaysOfWeekSaturday  ConnectorGetResponseInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorGetResponseInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorGetResponseInterruptWindowDaysOfWeekSunday, ConnectorGetResponseInterruptWindowDaysOfWeekMonday, ConnectorGetResponseInterruptWindowDaysOfWeekTuesday, ConnectorGetResponseInterruptWindowDaysOfWeekWednesday, ConnectorGetResponseInterruptWindowDaysOfWeekThursday, ConnectorGetResponseInterruptWindowDaysOfWeekFriday, ConnectorGetResponseInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorGetResponseDevice struct {
	ID           string                         `json:"id" api:"required"`
	SerialNumber string                         `json:"serial_number"`
	Type         ConnectorGetResponseDeviceType `json:"type"`
	JSON         connectorGetResponseDeviceJSON `json:"-"`
}

// connectorGetResponseDeviceJSON contains the JSON metadata for the struct
// [ConnectorGetResponseDevice]
type connectorGetResponseDeviceJSON struct {
	ID           apijson.Field
	SerialNumber apijson.Field
	Type         apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ConnectorGetResponseDevice) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseDeviceJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponseDeviceType string

const (
	ConnectorGetResponseDeviceTypeManaged  ConnectorGetResponseDeviceType = "MANAGED"
	ConnectorGetResponseDeviceTypeLicensed ConnectorGetResponseDeviceType = "LICENSED"
)

func (r ConnectorGetResponseDeviceType) IsKnown() bool {
	switch r {
	case ConnectorGetResponseDeviceTypeManaged, ConnectorGetResponseDeviceTypeLicensed:
		return true
	}
	return false
}

type ConnectorNewParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Exactly one of id, serial_number, or provision_license must be provided.
	Device    param.Field[ConnectorNewParamsDevice] `json:"device" api:"required"`
	Activated param.Field[bool]                     `json:"activated"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    param.Field[[]ConnectorNewParamsInterruptWindowDaysOfWeek] `json:"interrupt_window_days_of_week"`
	InterruptWindowDurationHours param.Field[float64]                                       `json:"interrupt_window_duration_hours"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates param.Field[[]string] `json:"interrupt_window_embargo_dates"`
	InterruptWindowHourOfDay    param.Field[float64]  `json:"interrupt_window_hour_of_day"`
	Notes                       param.Field[string]   `json:"notes"`
	Timezone                    param.Field[string]   `json:"timezone"`
}

func (r ConnectorNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Exactly one of id, serial_number, or provision_license must be provided.
type ConnectorNewParamsDevice struct {
	ID param.Field[string] `json:"id"`
	// When true, create and provision a new licence key for the connector.
	ProvisionLicense param.Field[bool]   `json:"provision_license"`
	SerialNumber     param.Field[string] `json:"serial_number"`
}

func (r ConnectorNewParamsDevice) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorNewParamsInterruptWindowDaysOfWeek string

const (
	ConnectorNewParamsInterruptWindowDaysOfWeekSunday    ConnectorNewParamsInterruptWindowDaysOfWeek = "Sunday"
	ConnectorNewParamsInterruptWindowDaysOfWeekMonday    ConnectorNewParamsInterruptWindowDaysOfWeek = "Monday"
	ConnectorNewParamsInterruptWindowDaysOfWeekTuesday   ConnectorNewParamsInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorNewParamsInterruptWindowDaysOfWeekWednesday ConnectorNewParamsInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorNewParamsInterruptWindowDaysOfWeekThursday  ConnectorNewParamsInterruptWindowDaysOfWeek = "Thursday"
	ConnectorNewParamsInterruptWindowDaysOfWeekFriday    ConnectorNewParamsInterruptWindowDaysOfWeek = "Friday"
	ConnectorNewParamsInterruptWindowDaysOfWeekSaturday  ConnectorNewParamsInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorNewParamsInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorNewParamsInterruptWindowDaysOfWeekSunday, ConnectorNewParamsInterruptWindowDaysOfWeekMonday, ConnectorNewParamsInterruptWindowDaysOfWeekTuesday, ConnectorNewParamsInterruptWindowDaysOfWeekWednesday, ConnectorNewParamsInterruptWindowDaysOfWeekThursday, ConnectorNewParamsInterruptWindowDaysOfWeekFriday, ConnectorNewParamsInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorNewResponseEnvelope struct {
	Errors   []ConnectorNewResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ConnectorNewResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ConnectorNewResponse                   `json:"result" api:"required"`
	Success  bool                                   `json:"success" api:"required"`
	JSON     connectorNewResponseEnvelopeJSON       `json:"-"`
}

// connectorNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorNewResponseEnvelope]
type connectorNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewResponseEnvelopeErrors struct {
	Code    float64                                `json:"code" api:"required"`
	Message string                                 `json:"message" api:"required"`
	JSON    connectorNewResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConnectorNewResponseEnvelopeErrors]
type connectorNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorNewResponseEnvelopeMessages struct {
	Code    float64                                  `json:"code" api:"required"`
	Message string                                   `json:"message" api:"required"`
	JSON    connectorNewResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorNewResponseEnvelopeMessages]
type connectorNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Activated param.Field[bool]   `json:"activated"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    param.Field[[]ConnectorUpdateParamsInterruptWindowDaysOfWeek] `json:"interrupt_window_days_of_week"`
	InterruptWindowDurationHours param.Field[float64]                                          `json:"interrupt_window_duration_hours"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates param.Field[[]string] `json:"interrupt_window_embargo_dates"`
	InterruptWindowHourOfDay    param.Field[float64]  `json:"interrupt_window_hour_of_day"`
	Notes                       param.Field[string]   `json:"notes"`
	// When true, regenerate license key for the connector.
	ProvisionLicense param.Field[bool]   `json:"provision_license"`
	Timezone         param.Field[string] `json:"timezone"`
}

func (r ConnectorUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorUpdateParamsInterruptWindowDaysOfWeek string

const (
	ConnectorUpdateParamsInterruptWindowDaysOfWeekSunday    ConnectorUpdateParamsInterruptWindowDaysOfWeek = "Sunday"
	ConnectorUpdateParamsInterruptWindowDaysOfWeekMonday    ConnectorUpdateParamsInterruptWindowDaysOfWeek = "Monday"
	ConnectorUpdateParamsInterruptWindowDaysOfWeekTuesday   ConnectorUpdateParamsInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorUpdateParamsInterruptWindowDaysOfWeekWednesday ConnectorUpdateParamsInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorUpdateParamsInterruptWindowDaysOfWeekThursday  ConnectorUpdateParamsInterruptWindowDaysOfWeek = "Thursday"
	ConnectorUpdateParamsInterruptWindowDaysOfWeekFriday    ConnectorUpdateParamsInterruptWindowDaysOfWeek = "Friday"
	ConnectorUpdateParamsInterruptWindowDaysOfWeekSaturday  ConnectorUpdateParamsInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorUpdateParamsInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorUpdateParamsInterruptWindowDaysOfWeekSunday, ConnectorUpdateParamsInterruptWindowDaysOfWeekMonday, ConnectorUpdateParamsInterruptWindowDaysOfWeekTuesday, ConnectorUpdateParamsInterruptWindowDaysOfWeekWednesday, ConnectorUpdateParamsInterruptWindowDaysOfWeekThursday, ConnectorUpdateParamsInterruptWindowDaysOfWeekFriday, ConnectorUpdateParamsInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorUpdateResponseEnvelope struct {
	Errors   []ConnectorUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ConnectorUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ConnectorUpdateResponse                   `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     connectorUpdateResponseEnvelopeJSON       `json:"-"`
}

// connectorUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorUpdateResponseEnvelope]
type connectorUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponseEnvelopeErrors struct {
	Code    float64                                   `json:"code" api:"required"`
	Message string                                    `json:"message" api:"required"`
	JSON    connectorUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorUpdateResponseEnvelopeErrors]
type connectorUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorUpdateResponseEnvelopeMessages struct {
	Code    float64                                     `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	JSON    connectorUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorUpdateResponseEnvelopeMessages]
type connectorUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Filter connectors by device type.
	DeviceType param.Field[ConnectorListParamsDeviceType] `query:"device_type"`
}

// URLQuery serializes [ConnectorListParams]'s query parameters as `url.Values`.
func (r ConnectorListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter connectors by device type.
type ConnectorListParamsDeviceType string

const (
	ConnectorListParamsDeviceTypeManaged  ConnectorListParamsDeviceType = "MANAGED"
	ConnectorListParamsDeviceTypeLicensed ConnectorListParamsDeviceType = "LICENSED"
)

func (r ConnectorListParamsDeviceType) IsKnown() bool {
	switch r {
	case ConnectorListParamsDeviceTypeManaged, ConnectorListParamsDeviceTypeLicensed:
		return true
	}
	return false
}

type ConnectorDeleteParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConnectorDeleteResponseEnvelope struct {
	Errors   []ConnectorDeleteResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ConnectorDeleteResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ConnectorDeleteResponse                   `json:"result" api:"required"`
	Success  bool                                      `json:"success" api:"required"`
	JSON     connectorDeleteResponseEnvelopeJSON       `json:"-"`
}

// connectorDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorDeleteResponseEnvelope]
type connectorDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponseEnvelopeErrors struct {
	Code    float64                                   `json:"code" api:"required"`
	Message string                                    `json:"message" api:"required"`
	JSON    connectorDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorDeleteResponseEnvelopeErrors]
type connectorDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorDeleteResponseEnvelopeMessages struct {
	Code    float64                                     `json:"code" api:"required"`
	Message string                                      `json:"message" api:"required"`
	JSON    connectorDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorDeleteResponseEnvelopeMessages]
type connectorDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Activated param.Field[bool]   `json:"activated"`
	// Allowed days of the week for upgrades. Default is all days.
	InterruptWindowDaysOfWeek    param.Field[[]ConnectorEditParamsInterruptWindowDaysOfWeek] `json:"interrupt_window_days_of_week"`
	InterruptWindowDurationHours param.Field[float64]                                        `json:"interrupt_window_duration_hours"`
	// List of dates (YYYY-MM-DD) when upgrades are blocked.
	InterruptWindowEmbargoDates param.Field[[]string] `json:"interrupt_window_embargo_dates"`
	InterruptWindowHourOfDay    param.Field[float64]  `json:"interrupt_window_hour_of_day"`
	Notes                       param.Field[string]   `json:"notes"`
	// When true, regenerate license key for the connector.
	ProvisionLicense param.Field[bool]   `json:"provision_license"`
	Timezone         param.Field[string] `json:"timezone"`
}

func (r ConnectorEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectorEditParamsInterruptWindowDaysOfWeek string

const (
	ConnectorEditParamsInterruptWindowDaysOfWeekSunday    ConnectorEditParamsInterruptWindowDaysOfWeek = "Sunday"
	ConnectorEditParamsInterruptWindowDaysOfWeekMonday    ConnectorEditParamsInterruptWindowDaysOfWeek = "Monday"
	ConnectorEditParamsInterruptWindowDaysOfWeekTuesday   ConnectorEditParamsInterruptWindowDaysOfWeek = "Tuesday"
	ConnectorEditParamsInterruptWindowDaysOfWeekWednesday ConnectorEditParamsInterruptWindowDaysOfWeek = "Wednesday"
	ConnectorEditParamsInterruptWindowDaysOfWeekThursday  ConnectorEditParamsInterruptWindowDaysOfWeek = "Thursday"
	ConnectorEditParamsInterruptWindowDaysOfWeekFriday    ConnectorEditParamsInterruptWindowDaysOfWeek = "Friday"
	ConnectorEditParamsInterruptWindowDaysOfWeekSaturday  ConnectorEditParamsInterruptWindowDaysOfWeek = "Saturday"
)

func (r ConnectorEditParamsInterruptWindowDaysOfWeek) IsKnown() bool {
	switch r {
	case ConnectorEditParamsInterruptWindowDaysOfWeekSunday, ConnectorEditParamsInterruptWindowDaysOfWeekMonday, ConnectorEditParamsInterruptWindowDaysOfWeekTuesday, ConnectorEditParamsInterruptWindowDaysOfWeekWednesday, ConnectorEditParamsInterruptWindowDaysOfWeekThursday, ConnectorEditParamsInterruptWindowDaysOfWeekFriday, ConnectorEditParamsInterruptWindowDaysOfWeekSaturday:
		return true
	}
	return false
}

type ConnectorEditResponseEnvelope struct {
	Errors   []ConnectorEditResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ConnectorEditResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ConnectorEditResponse                   `json:"result" api:"required"`
	Success  bool                                    `json:"success" api:"required"`
	JSON     connectorEditResponseEnvelopeJSON       `json:"-"`
}

// connectorEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorEditResponseEnvelope]
type connectorEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponseEnvelopeErrors struct {
	Code    float64                                 `json:"code" api:"required"`
	Message string                                  `json:"message" api:"required"`
	JSON    connectorEditResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ConnectorEditResponseEnvelopeErrors]
type connectorEditResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorEditResponseEnvelopeMessages struct {
	Code    float64                                   `json:"code" api:"required"`
	Message string                                    `json:"message" api:"required"`
	JSON    connectorEditResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorEditResponseEnvelopeMessages]
type connectorEditResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ConnectorGetResponseEnvelope struct {
	Errors   []ConnectorGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []ConnectorGetResponseEnvelopeMessages `json:"messages" api:"required"`
	Result   ConnectorGetResponse                   `json:"result" api:"required"`
	Success  bool                                   `json:"success" api:"required"`
	JSON     connectorGetResponseEnvelopeJSON       `json:"-"`
}

// connectorGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ConnectorGetResponseEnvelope]
type connectorGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponseEnvelopeErrors struct {
	Code    float64                                `json:"code" api:"required"`
	Message string                                 `json:"message" api:"required"`
	JSON    connectorGetResponseEnvelopeErrorsJSON `json:"-"`
}

// connectorGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [ConnectorGetResponseEnvelopeErrors]
type connectorGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ConnectorGetResponseEnvelopeMessages struct {
	Code    float64                                  `json:"code" api:"required"`
	Message string                                   `json:"message" api:"required"`
	JSON    connectorGetResponseEnvelopeMessagesJSON `json:"-"`
}

// connectorGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ConnectorGetResponseEnvelopeMessages]
type connectorGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ConnectorGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r connectorGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}
