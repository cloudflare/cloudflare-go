// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountDlpProfileService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccountDlpProfileService] method
// instead.
type AccountDlpProfileService struct {
	Options     []option.RequestOption
	Customs     *AccountDlpProfileCustomService
	Predefineds *AccountDlpProfilePredefinedService
}

// NewAccountDlpProfileService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountDlpProfileService(opts ...option.RequestOption) (r *AccountDlpProfileService) {
	r = &AccountDlpProfileService{}
	r.Options = opts
	r.Customs = NewAccountDlpProfileCustomService(opts...)
	r.Predefineds = NewAccountDlpProfilePredefinedService(opts...)
	return
}

// Fetches a DLP profile by ID. Supports both predefined and custom profiles
func (r *AccountDlpProfileService) Get(ctx context.Context, accountIdentifier string, profileID string, opts ...option.RequestOption) (res *EitherProfileResponseSNgAtLbh, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles/%s", accountIdentifier, profileID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists all DLP profiles in an account.
func (r *AccountDlpProfileService) List(ctx context.Context, accountIdentifier string, opts ...option.RequestOption) (res *ProfilesResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/dlp/profiles", accountIdentifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EitherProfileResponseSNgAtLbh struct {
	Errors   []EitherProfileResponseSNgAtLbhError   `json:"errors"`
	Messages []EitherProfileResponseSNgAtLbhMessage `json:"messages"`
	Result   EitherProfileResponseSNgAtLbhResult    `json:"result"`
	// Whether the API call was successful
	Success EitherProfileResponseSNgAtLbhSuccess `json:"success"`
	JSON    eitherProfileResponseSNgAtLbhJSON    `json:"-"`
}

// eitherProfileResponseSNgAtLbhJSON contains the JSON metadata for the struct
// [EitherProfileResponseSNgAtLbh]
type eitherProfileResponseSNgAtLbhJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EitherProfileResponseSNgAtLbh) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EitherProfileResponseSNgAtLbhError struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    eitherProfileResponseSNgAtLbhErrorJSON `json:"-"`
}

// eitherProfileResponseSNgAtLbhErrorJSON contains the JSON metadata for the struct
// [EitherProfileResponseSNgAtLbhError]
type eitherProfileResponseSNgAtLbhErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EitherProfileResponseSNgAtLbhError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type EitherProfileResponseSNgAtLbhMessage struct {
	Code    int64                                    `json:"code,required"`
	Message string                                   `json:"message,required"`
	JSON    eitherProfileResponseSNgAtLbhMessageJSON `json:"-"`
}

// eitherProfileResponseSNgAtLbhMessageJSON contains the JSON metadata for the
// struct [EitherProfileResponseSNgAtLbhMessage]
type eitherProfileResponseSNgAtLbhMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *EitherProfileResponseSNgAtLbhMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [PredefinedProfile] or [CustomProfile].
type EitherProfileResponseSNgAtLbhResult interface {
	implementsEitherProfileResponseSNgAtLbhResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*EitherProfileResponseSNgAtLbhResult)(nil)).Elem(), "")
}

// Whether the API call was successful
type EitherProfileResponseSNgAtLbhSuccess bool

const (
	EitherProfileResponseSNgAtLbhSuccessTrue EitherProfileResponseSNgAtLbhSuccess = true
)

type ProfilesResponseCollection struct {
	Errors     []ProfilesResponseCollectionError    `json:"errors"`
	Messages   []ProfilesResponseCollectionMessage  `json:"messages"`
	Result     []ProfilesResponseCollectionResult   `json:"result"`
	ResultInfo ProfilesResponseCollectionResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ProfilesResponseCollectionSuccess `json:"success"`
	JSON    profilesResponseCollectionJSON    `json:"-"`
}

// profilesResponseCollectionJSON contains the JSON metadata for the struct
// [ProfilesResponseCollection]
type profilesResponseCollectionJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilesResponseCollection) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilesResponseCollectionError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    profilesResponseCollectionErrorJSON `json:"-"`
}

// profilesResponseCollectionErrorJSON contains the JSON metadata for the struct
// [ProfilesResponseCollectionError]
type profilesResponseCollectionErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilesResponseCollectionError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ProfilesResponseCollectionMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    profilesResponseCollectionMessageJSON `json:"-"`
}

// profilesResponseCollectionMessageJSON contains the JSON metadata for the struct
// [ProfilesResponseCollectionMessage]
type profilesResponseCollectionMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilesResponseCollectionMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Union satisfied by [PredefinedProfile] or [CustomProfile].
type ProfilesResponseCollectionResult interface {
	implementsProfilesResponseCollectionResult()
}

func init() {
	apijson.RegisterUnion(reflect.TypeOf((*ProfilesResponseCollectionResult)(nil)).Elem(), "")
}

type ProfilesResponseCollectionResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       profilesResponseCollectionResultInfoJSON `json:"-"`
}

// profilesResponseCollectionResultInfoJSON contains the JSON metadata for the
// struct [ProfilesResponseCollectionResultInfo]
type profilesResponseCollectionResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProfilesResponseCollectionResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ProfilesResponseCollectionSuccess bool

const (
	ProfilesResponseCollectionSuccessTrue ProfilesResponseCollectionSuccess = true
)
