// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustAccessServiceTokenService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewZeroTrustAccessServiceTokenService] method instead.
type ZeroTrustAccessServiceTokenService struct {
	Options []option.RequestOption
}

// NewZeroTrustAccessServiceTokenService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewZeroTrustAccessServiceTokenService(opts ...option.RequestOption) (r *ZeroTrustAccessServiceTokenService) {
	r = &ZeroTrustAccessServiceTokenService{}
	r.Options = opts
	return
}

// Generates a new service token. **Note:** This is the only time you can get the
// Client Secret. If you lose the Client Secret, you will have to rotate the Client
// Secret or create a new service token.
func (r *ZeroTrustAccessServiceTokenService) New(ctx context.Context, params ZeroTrustAccessServiceTokenNewParams, opts ...option.RequestOption) (res *ZeroTrustAccessServiceTokenNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessServiceTokenNewResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Updates a configured service token.
func (r *ZeroTrustAccessServiceTokenService) Update(ctx context.Context, uuid string, params ZeroTrustAccessServiceTokenUpdateParams, opts ...option.RequestOption) (res *ZeroTrustAccessServiceTokenUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessServiceTokenUpdateResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if params.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = params.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = params.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all service tokens.
func (r *ZeroTrustAccessServiceTokenService) List(ctx context.Context, query ZeroTrustAccessServiceTokenListParams, opts ...option.RequestOption) (res *[]ZeroTrustAccessServiceTokenListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessServiceTokenListResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if query.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = query.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = query.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens", accountOrZone, accountOrZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Deletes a service token.
func (r *ZeroTrustAccessServiceTokenService) Delete(ctx context.Context, uuid string, body ZeroTrustAccessServiceTokenDeleteParams, opts ...option.RequestOption) (res *ZeroTrustAccessServiceTokenDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessServiceTokenDeleteResponseEnvelope
	var accountOrZone string
	var accountOrZoneID param.Field[string]
	if body.AccountID.Present {
		accountOrZone = "accounts"
		accountOrZoneID = body.AccountID
	} else {
		accountOrZone = "zones"
		accountOrZoneID = body.ZoneID
	}
	path := fmt.Sprintf("%s/%s/access/service_tokens/%s", accountOrZone, accountOrZoneID, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Refreshes the expiration of a service token.
func (r *ZeroTrustAccessServiceTokenService) Refresh(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZeroTrustAccessServiceTokenRefreshResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessServiceTokenRefreshResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/refresh", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Generates a new Client Secret for a service token and revokes the old one.
func (r *ZeroTrustAccessServiceTokenService) Rotate(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *ZeroTrustAccessServiceTokenRotateResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustAccessServiceTokenRotateResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/service_tokens/%s/rotate", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ZeroTrustAccessServiceTokenNewResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                     `json:"name"`
	UpdatedAt time.Time                                  `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessServiceTokenNewResponseJSON `json:"-"`
}

// zeroTrustAccessServiceTokenNewResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessServiceTokenNewResponse]
type zeroTrustAccessServiceTokenNewResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenNewResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenUpdateResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                        `json:"name"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessServiceTokenUpdateResponseJSON `json:"-"`
}

// zeroTrustAccessServiceTokenUpdateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessServiceTokenUpdateResponse]
type zeroTrustAccessServiceTokenUpdateResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenListResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                      `json:"name"`
	UpdatedAt time.Time                                   `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessServiceTokenListResponseJSON `json:"-"`
}

// zeroTrustAccessServiceTokenListResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessServiceTokenListResponse]
type zeroTrustAccessServiceTokenListResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenListResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenDeleteResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                        `json:"name"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessServiceTokenDeleteResponseJSON `json:"-"`
}

// zeroTrustAccessServiceTokenDeleteResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessServiceTokenDeleteResponse]
type zeroTrustAccessServiceTokenDeleteResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenRefreshResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID  string    `json:"client_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                         `json:"name"`
	UpdatedAt time.Time                                      `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessServiceTokenRefreshResponseJSON `json:"-"`
}

// zeroTrustAccessServiceTokenRefreshResponseJSON contains the JSON metadata for
// the struct [ZeroTrustAccessServiceTokenRefreshResponse]
type zeroTrustAccessServiceTokenRefreshResponseJSON struct {
	ID          apijson.Field
	ClientID    apijson.Field
	CreatedAt   apijson.Field
	Duration    apijson.Field
	Name        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRefreshResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRefreshResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenRotateResponse struct {
	// The ID of the service token.
	ID interface{} `json:"id"`
	// The Client ID for the service token. Access will check for this value in the
	// `CF-Access-Client-ID` request header.
	ClientID string `json:"client_id"`
	// The Client Secret for the service token. Access will check for this value in the
	// `CF-Access-Client-Secret` request header.
	ClientSecret string    `json:"client_secret"`
	CreatedAt    time.Time `json:"created_at" format:"date-time"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration string `json:"duration"`
	// The name of the service token.
	Name      string                                        `json:"name"`
	UpdatedAt time.Time                                     `json:"updated_at" format:"date-time"`
	JSON      zeroTrustAccessServiceTokenRotateResponseJSON `json:"-"`
}

// zeroTrustAccessServiceTokenRotateResponseJSON contains the JSON metadata for the
// struct [ZeroTrustAccessServiceTokenRotateResponse]
type zeroTrustAccessServiceTokenRotateResponseJSON struct {
	ID           apijson.Field
	ClientID     apijson.Field
	ClientSecret apijson.Field
	CreatedAt    apijson.Field
	Duration     apijson.Field
	Name         apijson.Field
	UpdatedAt    apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRotateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRotateResponseJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenNewParams struct {
	// The name of the service token.
	Name param.Field[string] `json:"name,required"`
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
}

func (r ZeroTrustAccessServiceTokenNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessServiceTokenNewResponseEnvelope struct {
	Errors   []ZeroTrustAccessServiceTokenNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessServiceTokenNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessServiceTokenNewResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessServiceTokenNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessServiceTokenNewResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessServiceTokenNewResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessServiceTokenNewResponseEnvelope]
type zeroTrustAccessServiceTokenNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenNewResponseEnvelopeErrors struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenNewResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessServiceTokenNewResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessServiceTokenNewResponseEnvelopeErrors]
type zeroTrustAccessServiceTokenNewResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenNewResponseEnvelopeMessages struct {
	Code    int64                                                      `json:"code,required"`
	Message string                                                     `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenNewResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessServiceTokenNewResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [ZeroTrustAccessServiceTokenNewResponseEnvelopeMessages]
type zeroTrustAccessServiceTokenNewResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessServiceTokenNewResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessServiceTokenNewResponseEnvelopeSuccessTrue ZeroTrustAccessServiceTokenNewResponseEnvelopeSuccess = true
)

type ZeroTrustAccessServiceTokenUpdateParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
	// The duration for how long the service token will be valid. Must be in the format
	// `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h. The
	// default is 1 year in hours (8760h).
	Duration param.Field[string] `json:"duration"`
	// The name of the service token.
	Name param.Field[string] `json:"name"`
}

func (r ZeroTrustAccessServiceTokenUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ZeroTrustAccessServiceTokenUpdateResponseEnvelope struct {
	Errors   []ZeroTrustAccessServiceTokenUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessServiceTokenUpdateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessServiceTokenUpdateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessServiceTokenUpdateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessServiceTokenUpdateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessServiceTokenUpdateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessServiceTokenUpdateResponseEnvelope]
type zeroTrustAccessServiceTokenUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenUpdateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenUpdateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessServiceTokenUpdateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenUpdateResponseEnvelopeErrors]
type zeroTrustAccessServiceTokenUpdateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenUpdateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenUpdateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessServiceTokenUpdateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenUpdateResponseEnvelopeMessages]
type zeroTrustAccessServiceTokenUpdateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessServiceTokenUpdateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessServiceTokenUpdateResponseEnvelopeSuccessTrue ZeroTrustAccessServiceTokenUpdateResponseEnvelopeSuccess = true
)

type ZeroTrustAccessServiceTokenListParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessServiceTokenListResponseEnvelope struct {
	Errors   []ZeroTrustAccessServiceTokenListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessServiceTokenListResponseEnvelopeMessages `json:"messages,required"`
	Result   []ZeroTrustAccessServiceTokenListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustAccessServiceTokenListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustAccessServiceTokenListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustAccessServiceTokenListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustAccessServiceTokenListResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessServiceTokenListResponseEnvelope]
type zeroTrustAccessServiceTokenListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenListResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenListResponseEnvelopeErrors struct {
	Code    int64                                                     `json:"code,required"`
	Message string                                                    `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessServiceTokenListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [ZeroTrustAccessServiceTokenListResponseEnvelopeErrors]
type zeroTrustAccessServiceTokenListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenListResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenListResponseEnvelopeMessages struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessServiceTokenListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenListResponseEnvelopeMessages]
type zeroTrustAccessServiceTokenListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenListResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessServiceTokenListResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessServiceTokenListResponseEnvelopeSuccessTrue ZeroTrustAccessServiceTokenListResponseEnvelopeSuccess = true
)

type ZeroTrustAccessServiceTokenListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                       `json:"total_count"`
	JSON       zeroTrustAccessServiceTokenListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustAccessServiceTokenListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenListResponseEnvelopeResultInfo]
type zeroTrustAccessServiceTokenListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenListResponseEnvelopeResultInfoJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenDeleteParams struct {
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID param.Field[string] `path:"account_id"`
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID param.Field[string] `path:"zone_id"`
}

type ZeroTrustAccessServiceTokenDeleteResponseEnvelope struct {
	Errors   []ZeroTrustAccessServiceTokenDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessServiceTokenDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessServiceTokenDeleteResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessServiceTokenDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessServiceTokenDeleteResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessServiceTokenDeleteResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessServiceTokenDeleteResponseEnvelope]
type zeroTrustAccessServiceTokenDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenDeleteResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenDeleteResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessServiceTokenDeleteResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenDeleteResponseEnvelopeErrors]
type zeroTrustAccessServiceTokenDeleteResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenDeleteResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenDeleteResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessServiceTokenDeleteResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenDeleteResponseEnvelopeMessages]
type zeroTrustAccessServiceTokenDeleteResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessServiceTokenDeleteResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessServiceTokenDeleteResponseEnvelopeSuccessTrue ZeroTrustAccessServiceTokenDeleteResponseEnvelopeSuccess = true
)

type ZeroTrustAccessServiceTokenRefreshResponseEnvelope struct {
	Errors   []ZeroTrustAccessServiceTokenRefreshResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessServiceTokenRefreshResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessServiceTokenRefreshResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessServiceTokenRefreshResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessServiceTokenRefreshResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessServiceTokenRefreshResponseEnvelopeJSON contains the JSON
// metadata for the struct [ZeroTrustAccessServiceTokenRefreshResponseEnvelope]
type zeroTrustAccessServiceTokenRefreshResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRefreshResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRefreshResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenRefreshResponseEnvelopeErrors struct {
	Code    int64                                                        `json:"code,required"`
	Message string                                                       `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenRefreshResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessServiceTokenRefreshResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenRefreshResponseEnvelopeErrors]
type zeroTrustAccessServiceTokenRefreshResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRefreshResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRefreshResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenRefreshResponseEnvelopeMessages struct {
	Code    int64                                                          `json:"code,required"`
	Message string                                                         `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenRefreshResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessServiceTokenRefreshResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenRefreshResponseEnvelopeMessages]
type zeroTrustAccessServiceTokenRefreshResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRefreshResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRefreshResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessServiceTokenRefreshResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessServiceTokenRefreshResponseEnvelopeSuccessTrue ZeroTrustAccessServiceTokenRefreshResponseEnvelopeSuccess = true
)

type ZeroTrustAccessServiceTokenRotateResponseEnvelope struct {
	Errors   []ZeroTrustAccessServiceTokenRotateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustAccessServiceTokenRotateResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustAccessServiceTokenRotateResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustAccessServiceTokenRotateResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustAccessServiceTokenRotateResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustAccessServiceTokenRotateResponseEnvelopeJSON contains the JSON metadata
// for the struct [ZeroTrustAccessServiceTokenRotateResponseEnvelope]
type zeroTrustAccessServiceTokenRotateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRotateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRotateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenRotateResponseEnvelopeErrors struct {
	Code    int64                                                       `json:"code,required"`
	Message string                                                      `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenRotateResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustAccessServiceTokenRotateResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenRotateResponseEnvelopeErrors]
type zeroTrustAccessServiceTokenRotateResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRotateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRotateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ZeroTrustAccessServiceTokenRotateResponseEnvelopeMessages struct {
	Code    int64                                                         `json:"code,required"`
	Message string                                                        `json:"message,required"`
	JSON    zeroTrustAccessServiceTokenRotateResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustAccessServiceTokenRotateResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct
// [ZeroTrustAccessServiceTokenRotateResponseEnvelopeMessages]
type zeroTrustAccessServiceTokenRotateResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustAccessServiceTokenRotateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r zeroTrustAccessServiceTokenRotateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type ZeroTrustAccessServiceTokenRotateResponseEnvelopeSuccess bool

const (
	ZeroTrustAccessServiceTokenRotateResponseEnvelopeSuccessTrue ZeroTrustAccessServiceTokenRotateResponseEnvelopeSuccess = true
)
