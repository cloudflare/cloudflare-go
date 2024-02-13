// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// SecondaryDNSOutgoingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSOutgoingService]
// method instead.
type SecondaryDNSOutgoingService struct {
	Options       []option.RequestOption
	Disables      *SecondaryDNSOutgoingDisableService
	Enables       *SecondaryDNSOutgoingEnableService
	ForceNotifies *SecondaryDNSOutgoingForceNotifyService
	Statuses      *SecondaryDNSOutgoingStatusService
}

// NewSecondaryDNSOutgoingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSOutgoingService(opts ...option.RequestOption) (r *SecondaryDNSOutgoingService) {
	r = &SecondaryDNSOutgoingService{}
	r.Options = opts
	r.Disables = NewSecondaryDNSOutgoingDisableService(opts...)
	r.Enables = NewSecondaryDNSOutgoingEnableService(opts...)
	r.ForceNotifies = NewSecondaryDNSOutgoingForceNotifyService(opts...)
	r.Statuses = NewSecondaryDNSOutgoingStatusService(opts...)
	return
}

// Delete primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) Delete(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *SecondaryDNSOutgoingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) SecondaryDNSPrimaryZoneNewPrimaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationParams, opts ...option.RequestOption) (res *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) SecondaryDNSPrimaryZonePrimaryZoneConfigurationDetails(ctx context.Context, zoneIdentifier interface{}, opts ...option.RequestOption) (res *SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update primary zone configuration for outgoing zone transfers.
func (r *SecondaryDNSOutgoingService) SecondaryDNSPrimaryZoneUpdatePrimaryZoneConfiguration(ctx context.Context, zoneIdentifier interface{}, body SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationParams, opts ...option.RequestOption) (res *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/outgoing", zoneIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSOutgoingDeleteResponse struct {
	ID   interface{}                            `json:"id"`
	JSON secondaryDNSOutgoingDeleteResponseJSON `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSOutgoingDeleteResponse]
type secondaryDNSOutgoingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse struct {
	ID interface{} `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                                                            `json:"soa_serial"`
	JSON      secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SoaSerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse struct {
	ID interface{} `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                                                                `json:"soa_serial"`
	JSON      secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse]
type secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SoaSerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse struct {
	ID interface{} `json:"id"`
	// The time for a specific event.
	CheckedTime string `json:"checked_time"`
	// The time for a specific event.
	CreatedTime string `json:"created_time"`
	// The time for a specific event.
	LastTransferredTime string `json:"last_transferred_time"`
	// Zone name.
	Name string `json:"name"`
	// A list of peer tags.
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                                                               `json:"soa_serial"`
	JSON      secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseJSON struct {
	ID                  apijson.Field
	CheckedTime         apijson.Field
	CreatedTime         apijson.Field
	LastTransferredTime apijson.Field
	Name                apijson.Field
	Peers               apijson.Field
	SoaSerial           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDeleteResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSOutgoingDeleteResponseEnvelope]
type secondaryDNSOutgoingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSOutgoingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSOutgoingDeleteResponseEnvelopeErrors]
type secondaryDNSOutgoingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    secondaryDNSOutgoingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSOutgoingDeleteResponseEnvelopeMessages]
type secondaryDNSOutgoingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingDeleteResponseEnvelopeSuccessTrue SecondaryDNSOutgoingDeleteResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationParams struct {
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelope]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                            `json:"code,required"`
	Message string                                                                                           `json:"message,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeErrors]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                              `json:"code,required"`
	Message string                                                                                             `json:"message,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeMessages]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeSuccessTrue SecondaryDNSOutgoingSecondaryDNSPrimaryZoneNewPrimaryZoneConfigurationResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelope]
type secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeErrors]
type secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeMessages]
type secondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeSuccessTrue SecondaryDNSOutgoingSecondaryDNSPrimaryZonePrimaryZoneConfigurationDetailsResponseEnvelopeSuccess = true
)

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationParams struct {
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelope struct {
	Errors   []SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelope]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeErrors]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeMessages]
type secondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeSuccess bool

const (
	SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeSuccessTrue SecondaryDNSOutgoingSecondaryDNSPrimaryZoneUpdatePrimaryZoneConfigurationResponseEnvelopeSuccess = true
)
