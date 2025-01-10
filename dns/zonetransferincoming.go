// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package dns

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/shared"
)

// ZoneTransferIncomingService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewZoneTransferIncomingService] method instead.
type ZoneTransferIncomingService struct {
	Options []option.RequestOption
}

// NewZoneTransferIncomingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZoneTransferIncomingService(opts ...option.RequestOption) (r *ZoneTransferIncomingService) {
	r = &ZoneTransferIncomingService{}
	r.Options = opts
	return
}

// Create secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) New(ctx context.Context, params ZoneTransferIncomingNewParams, opts ...option.RequestOption) (res *ZoneTransferIncomingNewResponse, err error) {
	var env ZoneTransferIncomingNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) Update(ctx context.Context, params ZoneTransferIncomingUpdateParams, opts ...option.RequestOption) (res *ZoneTransferIncomingUpdateResponse, err error) {
	var env ZoneTransferIncomingUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Delete secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) Delete(ctx context.Context, body ZoneTransferIncomingDeleteParams, opts ...option.RequestOption) (res *ZoneTransferIncomingDeleteResponse, err error) {
	var env ZoneTransferIncomingDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", body.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get secondary zone configuration for incoming zone transfers.
func (r *ZoneTransferIncomingService) Get(ctx context.Context, query ZoneTransferIncomingGetParams, opts ...option.RequestOption) (res *ZoneTransferIncomingGetResponse, err error) {
	var env ZoneTransferIncomingGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/secondary_dns/incoming", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZoneTransferIncomingNewResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                             `json:"soa_serial"`
	JSON      zoneTransferIncomingNewResponseJSON `json:"-"`
}

// zoneTransferIncomingNewResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingNewResponse]
type zoneTransferIncomingNewResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingUpdateResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                                `json:"soa_serial"`
	JSON      zoneTransferIncomingUpdateResponseJSON `json:"-"`
}

// zoneTransferIncomingUpdateResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingUpdateResponse]
type zoneTransferIncomingUpdateResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingDeleteResponse struct {
	ID   string                                 `json:"id"`
	JSON zoneTransferIncomingDeleteResponseJSON `json:"-"`
}

// zoneTransferIncomingDeleteResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingDeleteResponse]
type zoneTransferIncomingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingGetResponse struct {
	ID string `json:"id"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds float64 `json:"auto_refresh_seconds"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	ModifiedTime string `json:"modified_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []string `json:"peers"`
	// The serial number of the SOA for the given zone.
	SOASerial float64                             `json:"soa_serial"`
	JSON      zoneTransferIncomingGetResponseJSON `json:"-"`
}

// zoneTransferIncomingGetResponseJSON contains the JSON metadata for the struct
// [ZoneTransferIncomingGetResponse]
type zoneTransferIncomingGetResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SOASerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseJSON) RawJSON() string {
	return r.raw
}

type ZoneTransferIncomingNewParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r ZoneTransferIncomingNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferIncomingNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferIncomingNewResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingNewResponse                `json:"result"`
	JSON    zoneTransferIncomingNewResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingNewResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferIncomingNewResponseEnvelope]
type zoneTransferIncomingNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferIncomingNewResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingNewResponseEnvelopeSuccessTrue ZoneTransferIncomingNewResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferIncomingUpdateParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]string] `json:"peers,required"`
}

func (r ZoneTransferIncomingUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZoneTransferIncomingUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferIncomingUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingUpdateResponse                `json:"result"`
	JSON    zoneTransferIncomingUpdateResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingUpdateResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferIncomingUpdateResponseEnvelope]
type zoneTransferIncomingUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferIncomingUpdateResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingUpdateResponseEnvelopeSuccessTrue ZoneTransferIncomingUpdateResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferIncomingDeleteParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferIncomingDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferIncomingDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingDeleteResponse                `json:"result"`
	JSON    zoneTransferIncomingDeleteResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [ZoneTransferIncomingDeleteResponseEnvelope]
type zoneTransferIncomingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferIncomingDeleteResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingDeleteResponseEnvelopeSuccessTrue ZoneTransferIncomingDeleteResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ZoneTransferIncomingGetParams struct {
	ZoneID param.Field[string] `path:"zone_id,required"`
}

type ZoneTransferIncomingGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success ZoneTransferIncomingGetResponseEnvelopeSuccess `json:"success,required"`
	Result  ZoneTransferIncomingGetResponse                `json:"result"`
	JSON    zoneTransferIncomingGetResponseEnvelopeJSON    `json:"-"`
}

// zoneTransferIncomingGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZoneTransferIncomingGetResponseEnvelope]
type zoneTransferIncomingGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZoneTransferIncomingGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zoneTransferIncomingGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZoneTransferIncomingGetResponseEnvelopeSuccess bool

const (
	ZoneTransferIncomingGetResponseEnvelopeSuccessTrue ZoneTransferIncomingGetResponseEnvelopeSuccess = true
)

func (r ZoneTransferIncomingGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ZoneTransferIncomingGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
