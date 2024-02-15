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

// SecondaryDNSIncomingService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewSecondaryDNSIncomingService]
// method instead.
type SecondaryDNSIncomingService struct {
	Options []option.RequestOption
}

// NewSecondaryDNSIncomingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSecondaryDNSIncomingService(opts ...option.RequestOption) (r *SecondaryDNSIncomingService) {
	r = &SecondaryDNSIncomingService{}
	r.Options = opts
	return
}

// Delete secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) Delete(ctx context.Context, zoneID interface{}, opts ...option.RequestOption) (res *SecondaryDNSIncomingDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingDeleteResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Create secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) SecondaryDNSSecondaryZoneNewSecondaryZoneConfiguration(ctx context.Context, zoneID interface{}, body SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationParams, opts ...option.RequestOption) (res *SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) SecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetails(ctx context.Context, zoneID interface{}, opts ...option.RequestOption) (res *SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update secondary zone configuration for incoming zone transfers.
func (r *SecondaryDNSIncomingService) SecondaryDNSSecondaryZoneUpdateSecondaryZoneConfiguration(ctx context.Context, zoneID interface{}, body SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationParams, opts ...option.RequestOption) (res *SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelope
	path := fmt.Sprintf("zones/%v/secondary_dns/incoming", zoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SecondaryDNSIncomingDeleteResponse struct {
	ID   interface{}                            `json:"id"`
	JSON secondaryDNSIncomingDeleteResponseJSON `json:"-"`
}

// secondaryDNSIncomingDeleteResponseJSON contains the JSON metadata for the struct
// [SecondaryDNSIncomingDeleteResponse]
type secondaryDNSIncomingDeleteResponseJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse struct {
	ID interface{} `json:"id"`
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
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                                                                `json:"soa_serial"`
	JSON      secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SoaSerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse struct {
	ID interface{} `json:"id"`
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
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                                                                    `json:"soa_serial"`
	JSON      secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SoaSerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse struct {
	ID interface{} `json:"id"`
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
	Peers []interface{} `json:"peers"`
	// The serial number of the SOA for the given zone.
	SoaSerial float64                                                                                   `json:"soa_serial"`
	JSON      secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseJSON struct {
	ID                 apijson.Field
	AutoRefreshSeconds apijson.Field
	CheckedTime        apijson.Field
	CreatedTime        apijson.Field
	ModifiedTime       apijson.Field
	Name               apijson.Field
	Peers              apijson.Field
	SoaSerial          apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingDeleteResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingDeleteResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingDeleteResponseEnvelopeJSON contains the JSON metadata for
// the struct [SecondaryDNSIncomingDeleteResponseEnvelope]
type secondaryDNSIncomingDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingDeleteResponseEnvelopeErrors struct {
	Code    int64                                                `json:"code,required"`
	Message string                                               `json:"message,required"`
	JSON    secondaryDNSIncomingDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingDeleteResponseEnvelopeErrorsJSON contains the JSON metadata
// for the struct [SecondaryDNSIncomingDeleteResponseEnvelopeErrors]
type secondaryDNSIncomingDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingDeleteResponseEnvelopeMessages struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    secondaryDNSIncomingDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [SecondaryDNSIncomingDeleteResponseEnvelopeMessages]
type secondaryDNSIncomingDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingDeleteResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingDeleteResponseEnvelopeSuccessTrue SecondaryDNSIncomingDeleteResponseEnvelopeSuccess = true
)

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationParams struct {
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelope]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                                `json:"code,required"`
	Message string                                                                                               `json:"message,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeErrors]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                                  `json:"code,required"`
	Message string                                                                                                 `json:"message,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeMessages]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeSuccessTrue SecondaryDNSIncomingSecondaryDNSSecondaryZoneNewSecondaryZoneConfigurationResponseEnvelopeSuccess = true
)

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelope]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeErrors struct {
	Code    int64                                                                                                    `json:"code,required"`
	Message string                                                                                                   `json:"message,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeErrors]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeMessages struct {
	Code    int64                                                                                                      `json:"code,required"`
	Message string                                                                                                     `json:"message,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeMessages]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeSuccessTrue SecondaryDNSIncomingSecondaryDNSSecondaryZoneSecondaryZoneConfigurationDetailsResponseEnvelopeSuccess = true
)

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationParams struct {
	// How often should a secondary zone auto refresh regardless of DNS NOTIFY. Not
	// applicable for primary zones.
	AutoRefreshSeconds param.Field[float64] `json:"auto_refresh_seconds,required"`
	// Zone name.
	Name param.Field[string] `json:"name,required"`
	// A list of peer tags.
	Peers param.Field[[]interface{}] `json:"peers,required"`
}

func (r SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelope struct {
	Errors   []SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeErrors   `json:"errors,required"`
	Messages []SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeMessages `json:"messages,required"`
	Result   SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeSuccess `json:"success,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeJSON    `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelope]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeErrors struct {
	Code    int64                                                                                                   `json:"code,required"`
	Message string                                                                                                  `json:"message,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeErrorsJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeErrors]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeMessages struct {
	Code    int64                                                                                                     `json:"code,required"`
	Message string                                                                                                    `json:"message,required"`
	JSON    secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeMessagesJSON `json:"-"`
}

// secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeMessages]
type secondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeSuccess bool

const (
	SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeSuccessTrue SecondaryDNSIncomingSecondaryDNSSecondaryZoneUpdateSecondaryZoneConfigurationResponseEnvelopeSuccess = true
)
