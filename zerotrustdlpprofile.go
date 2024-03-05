// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// ZeroTrustDLPProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewZeroTrustDLPProfileService]
// method instead.
type ZeroTrustDLPProfileService struct {
	Options     []option.RequestOption
	Customs     *ZeroTrustDLPProfileCustomService
	Predefineds *ZeroTrustDLPProfilePredefinedService
}

// NewZeroTrustDLPProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewZeroTrustDLPProfileService(opts ...option.RequestOption) (r *ZeroTrustDLPProfileService) {
	r = &ZeroTrustDLPProfileService{}
	r.Options = opts
	r.Customs = NewZeroTrustDLPProfileCustomService(opts...)
	r.Predefineds = NewZeroTrustDLPProfilePredefinedService(opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *ZeroTrustDLPProfileService) List(ctx context.Context, query ZeroTrustDLPProfileListParams, opts ...option.RequestOption) (res *[]DLPProfiles, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfileListResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches a DLP profile by ID. Supports both predefined and custom profiles
func (r *ZeroTrustDLPProfileService) Get(ctx context.Context, profileID string, query ZeroTrustDLPProfileGetParams, opts ...option.RequestOption) (res *ZeroTrustDLPProfileGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env ZeroTrustDLPProfileGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", query.AccountID, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Union satisfied by [DLPPredefinedProfile], [DLPCustomProfile] or
// [DLPProfilesDLPIntegrationProfile].
type DLPProfiles interface {
	implementsDLPProfiles()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*DLPProfiles)(nil)).Elem(), "")
}

type DLPProfilesDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []DLPProfilesDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      DLPProfilesDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                            `json:"updated_at" format:"date-time"`
	JSON      dlpProfilesDLPIntegrationProfileJSON `json:"-"`
}

// dlpProfilesDLPIntegrationProfileJSON contains the JSON metadata for the struct
// [DLPProfilesDLPIntegrationProfile]
type dlpProfilesDLPIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilesDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r DLPProfilesDLPIntegrationProfile) implementsDLPProfiles() {}

// An entry derived from an integration
type DLPProfilesDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                               `json:"profile_id"`
	UpdatedAt time.Time                                 `json:"updated_at" format:"date-time"`
	JSON      dlpProfilesDLPIntegrationProfileEntryJSON `json:"-"`
}

// dlpProfilesDLPIntegrationProfileEntryJSON contains the JSON metadata for the
// struct [DLPProfilesDLPIntegrationProfileEntry]
type dlpProfilesDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DLPProfilesDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type DLPProfilesDLPIntegrationProfileType string

const (
	DLPProfilesDLPIntegrationProfileTypeIntegration DLPProfilesDLPIntegrationProfileType = "integration"
)

// Union satisfied by [DLPPredefinedProfile], [DLPCustomProfile] or
// [ZeroTrustDLPProfileGetResponseDLPIntegrationProfile].
type ZeroTrustDLPProfileGetResponse interface {
	implementsZeroTrustDLPProfileGetResponse()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ZeroTrustDLPProfileGetResponse)(nil)).Elem(), "")
}

type ZeroTrustDLPProfileGetResponseDLPIntegrationProfile struct {
	// The ID for this profile
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The description of the profile.
	Description string `json:"description"`
	// The entries for this profile.
	Entries []ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry `json:"entries"`
	// The name of the profile.
	Name string `json:"name"`
	// The type of the profile.
	Type      ZeroTrustDLPProfileGetResponseDLPIntegrationProfileType `json:"type"`
	UpdatedAt time.Time                                               `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileGetResponseDLPIntegrationProfileJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPIntegrationProfileJSON contains the JSON
// metadata for the struct [ZeroTrustDLPProfileGetResponseDLPIntegrationProfile]
type zeroTrustDLPProfileGetResponseDLPIntegrationProfileJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Entries     apijson.Field
	Name        apijson.Field
	Type        apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPIntegrationProfile) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r ZeroTrustDLPProfileGetResponseDLPIntegrationProfile) implementsZeroTrustDLPProfileGetResponse() {
}

// An entry derived from an integration
type ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry struct {
	// The ID for this entry
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Whether the entry is enabled or not.
	Enabled bool `json:"enabled"`
	// The name of the entry.
	Name string `json:"name"`
	// ID of the parent profile
	ProfileID interface{}                                                  `json:"profile_id"`
	UpdatedAt time.Time                                                    `json:"updated_at" format:"date-time"`
	JSON      zeroTrustDLPProfileGetResponseDLPIntegrationProfileEntryJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseDLPIntegrationProfileEntryJSON contains the JSON
// metadata for the struct
// [ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry]
type zeroTrustDLPProfileGetResponseDLPIntegrationProfileEntryJSON struct {
	ID          apijson.Field
	CreatedAt   apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	ProfileID   apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseDLPIntegrationProfileEntry) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// The type of the profile.
type ZeroTrustDLPProfileGetResponseDLPIntegrationProfileType string

const (
	ZeroTrustDLPProfileGetResponseDLPIntegrationProfileTypeIntegration ZeroTrustDLPProfileGetResponseDLPIntegrationProfileType = "integration"
)

type ZeroTrustDLPProfileListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPProfileListResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfileListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfileListResponseEnvelopeMessages `json:"messages,required"`
	Result   []DLPProfiles                                     `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    ZeroTrustDLPProfileListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo ZeroTrustDLPProfileListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       zeroTrustDLPProfileListResponseEnvelopeJSON       `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileListResponseEnvelope]
type zeroTrustDLPProfileListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPProfileListResponseEnvelopeErrors struct {
	Code    int64                                             `json:"code,required"`
	Message string                                            `json:"message,required"`
	JSON    zeroTrustDLPProfileListResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDLPProfileListResponseEnvelopeErrors]
type zeroTrustDLPProfileListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPProfileListResponseEnvelopeMessages struct {
	Code    int64                                               `json:"code,required"`
	Message string                                              `json:"message,required"`
	JSON    zeroTrustDLPProfileListResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileListResponseEnvelopeMessages]
type zeroTrustDLPProfileListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDLPProfileListResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfileListResponseEnvelopeSuccessTrue ZeroTrustDLPProfileListResponseEnvelopeSuccess = true
)

type ZeroTrustDLPProfileListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                               `json:"total_count"`
	JSON       zeroTrustDLPProfileListResponseEnvelopeResultInfoJSON `json:"-"`
}

// zeroTrustDLPProfileListResponseEnvelopeResultInfoJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileListResponseEnvelopeResultInfo]
type zeroTrustDLPProfileListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPProfileGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type ZeroTrustDLPProfileGetResponseEnvelope struct {
	Errors   []ZeroTrustDLPProfileGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ZeroTrustDLPProfileGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ZeroTrustDLPProfileGetResponse                   `json:"result,required"`
	// Whether the API call was successful
	Success ZeroTrustDLPProfileGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    zeroTrustDLPProfileGetResponseEnvelopeJSON    `json:"-"`
}

// zeroTrustDLPProfileGetResponseEnvelopeJSON contains the JSON metadata for the
// struct [ZeroTrustDLPProfileGetResponseEnvelope]
type zeroTrustDLPProfileGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPProfileGetResponseEnvelopeErrors struct {
	Code    int64                                            `json:"code,required"`
	Message string                                           `json:"message,required"`
	JSON    zeroTrustDLPProfileGetResponseEnvelopeErrorsJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseEnvelopeErrorsJSON contains the JSON metadata for
// the struct [ZeroTrustDLPProfileGetResponseEnvelopeErrors]
type zeroTrustDLPProfileGetResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ZeroTrustDLPProfileGetResponseEnvelopeMessages struct {
	Code    int64                                              `json:"code,required"`
	Message string                                             `json:"message,required"`
	JSON    zeroTrustDLPProfileGetResponseEnvelopeMessagesJSON `json:"-"`
}

// zeroTrustDLPProfileGetResponseEnvelopeMessagesJSON contains the JSON metadata
// for the struct [ZeroTrustDLPProfileGetResponseEnvelopeMessages]
type zeroTrustDLPProfileGetResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ZeroTrustDLPProfileGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ZeroTrustDLPProfileGetResponseEnvelopeSuccess bool

const (
	ZeroTrustDLPProfileGetResponseEnvelopeSuccessTrue ZeroTrustDLPProfileGetResponseEnvelopeSuccess = true
)
