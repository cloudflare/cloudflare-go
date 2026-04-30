// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package fraud

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// FraudService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFraudService] method instead.
type FraudService struct {
	Options []option.RequestOption
}

// NewFraudService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFraudService(opts ...option.RequestOption) (r *FraudService) {
	r = &FraudService{}
	r.Options = opts
	return
}

// Update Fraud Detection settings for a zone.
//
// Notes on `username_expressions` behavior:
//
// - If omitted or set to null, expressions are not modified.
// - If provided as an empty array `[]`, all expressions will be cleared.
func (r *FraudService) Update(ctx context.Context, params FraudUpdateParams, opts ...option.RequestOption) (res *FraudSettings, err error) {
	var env FraudUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/fraud_detection/settings", params.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

// Retrieve Fraud Detection settings for a zone.
func (r *FraudService) Get(ctx context.Context, query FraudGetParams, opts ...option.RequestOption) (res *FraudSettings, err error) {
	var env FraudGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("zones/%s/fraud_detection/settings", query.ZoneID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return nil, err
	}
	res = &env.Result
	return res, nil
}

type FraudSettings struct {
	// Configuration for classifying login authentication outcomes based on the origin
	// response. Requires `user_profiles` to be enabled.
	//
	//   - Success and failure criteria are independently updatable — sending only
	//     `success_criteria` leaves failure codes untouched, and vice versa.
	//   - Omit `authentication_settings` entirely to leave both unchanged.
	//   - Status codes must not overlap between success and failure criteria.
	AuthenticationSettings FraudSettingsAuthenticationSettings `json:"authentication_settings"`
	// Whether Fraud User Profiles is enabled for the zone.
	UserProfiles FraudSettingsUserProfiles `json:"user_profiles"`
	// List of expressions to detect usernames in write HTTP requests.
	//
	//   - Maximum of 10 expressions.
	//   - Omit or set to null to leave unchanged on update.
	//   - Provide an empty array `[]` to clear all expressions on update.
	//   - Invalid expressions will result in a 10400 Bad Request with details in the
	//     `messages` array.
	UsernameExpressions []string          `json:"username_expressions"`
	JSON                fraudSettingsJSON `json:"-"`
}

// fraudSettingsJSON contains the JSON metadata for the struct [FraudSettings]
type fraudSettingsJSON struct {
	AuthenticationSettings apijson.Field
	UserProfiles           apijson.Field
	UsernameExpressions    apijson.Field
	raw                    string
	ExtraFields            map[string]apijson.Field
}

func (r *FraudSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudSettingsJSON) RawJSON() string {
	return r.raw
}

// Configuration for classifying login authentication outcomes based on the origin
// response. Requires `user_profiles` to be enabled.
//
//   - Success and failure criteria are independently updatable — sending only
//     `success_criteria` leaves failure codes untouched, and vice versa.
//   - Omit `authentication_settings` entirely to leave both unchanged.
//   - Status codes must not overlap between success and failure criteria.
type FraudSettingsAuthenticationSettings struct {
	// Criterion for identifying failed login responses.
	FailureCriteria FraudSettingsAuthenticationSettingsFailureCriteria `json:"failure_criteria"`
	// Criterion for identifying successful login responses.
	SuccessCriteria FraudSettingsAuthenticationSettingsSuccessCriteria `json:"success_criteria"`
	JSON            fraudSettingsAuthenticationSettingsJSON            `json:"-"`
}

// fraudSettingsAuthenticationSettingsJSON contains the JSON metadata for the
// struct [FraudSettingsAuthenticationSettings]
type fraudSettingsAuthenticationSettingsJSON struct {
	FailureCriteria apijson.Field
	SuccessCriteria apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *FraudSettingsAuthenticationSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudSettingsAuthenticationSettingsJSON) RawJSON() string {
	return r.raw
}

// Criterion for identifying failed login responses.
type FraudSettingsAuthenticationSettingsFailureCriteria struct {
	// The type of criterion. Currently only `status_code` is supported.
	Kind FraudSettingsAuthenticationSettingsFailureCriteriaKind `json:"kind" api:"required"`
	// HTTP status codes to match against the origin response.
	//
	// - Maximum of 10 codes per criterion.
	// - Each code must be a valid HTTP status code (100-599).
	// - Codes are deduplicated and sorted on save.
	// - Omit to leave unchanged on update.
	// - Provide an empty array `[]` to clear codes on update.
	StatusCodes []int64                                                `json:"status_codes"`
	JSON        fraudSettingsAuthenticationSettingsFailureCriteriaJSON `json:"-"`
}

// fraudSettingsAuthenticationSettingsFailureCriteriaJSON contains the JSON
// metadata for the struct [FraudSettingsAuthenticationSettingsFailureCriteria]
type fraudSettingsAuthenticationSettingsFailureCriteriaJSON struct {
	Kind        apijson.Field
	StatusCodes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudSettingsAuthenticationSettingsFailureCriteria) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudSettingsAuthenticationSettingsFailureCriteriaJSON) RawJSON() string {
	return r.raw
}

// The type of criterion. Currently only `status_code` is supported.
type FraudSettingsAuthenticationSettingsFailureCriteriaKind string

const (
	FraudSettingsAuthenticationSettingsFailureCriteriaKindStatusCode FraudSettingsAuthenticationSettingsFailureCriteriaKind = "status_code"
)

func (r FraudSettingsAuthenticationSettingsFailureCriteriaKind) IsKnown() bool {
	switch r {
	case FraudSettingsAuthenticationSettingsFailureCriteriaKindStatusCode:
		return true
	}
	return false
}

// Criterion for identifying successful login responses.
type FraudSettingsAuthenticationSettingsSuccessCriteria struct {
	// The type of criterion. Currently only `status_code` is supported.
	Kind FraudSettingsAuthenticationSettingsSuccessCriteriaKind `json:"kind" api:"required"`
	// HTTP status codes to match against the origin response.
	//
	// - Maximum of 10 codes per criterion.
	// - Each code must be a valid HTTP status code (100-599).
	// - Codes are deduplicated and sorted on save.
	// - Omit to leave unchanged on update.
	// - Provide an empty array `[]` to clear codes on update.
	StatusCodes []int64                                                `json:"status_codes"`
	JSON        fraudSettingsAuthenticationSettingsSuccessCriteriaJSON `json:"-"`
}

// fraudSettingsAuthenticationSettingsSuccessCriteriaJSON contains the JSON
// metadata for the struct [FraudSettingsAuthenticationSettingsSuccessCriteria]
type fraudSettingsAuthenticationSettingsSuccessCriteriaJSON struct {
	Kind        apijson.Field
	StatusCodes apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudSettingsAuthenticationSettingsSuccessCriteria) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudSettingsAuthenticationSettingsSuccessCriteriaJSON) RawJSON() string {
	return r.raw
}

// The type of criterion. Currently only `status_code` is supported.
type FraudSettingsAuthenticationSettingsSuccessCriteriaKind string

const (
	FraudSettingsAuthenticationSettingsSuccessCriteriaKindStatusCode FraudSettingsAuthenticationSettingsSuccessCriteriaKind = "status_code"
)

func (r FraudSettingsAuthenticationSettingsSuccessCriteriaKind) IsKnown() bool {
	switch r {
	case FraudSettingsAuthenticationSettingsSuccessCriteriaKindStatusCode:
		return true
	}
	return false
}

// Whether Fraud User Profiles is enabled for the zone.
type FraudSettingsUserProfiles string

const (
	FraudSettingsUserProfilesEnabled  FraudSettingsUserProfiles = "enabled"
	FraudSettingsUserProfilesDisabled FraudSettingsUserProfiles = "disabled"
)

func (r FraudSettingsUserProfiles) IsKnown() bool {
	switch r {
	case FraudSettingsUserProfilesEnabled, FraudSettingsUserProfilesDisabled:
		return true
	}
	return false
}

type FraudSettingsParam struct {
	// Configuration for classifying login authentication outcomes based on the origin
	// response. Requires `user_profiles` to be enabled.
	//
	//   - Success and failure criteria are independently updatable — sending only
	//     `success_criteria` leaves failure codes untouched, and vice versa.
	//   - Omit `authentication_settings` entirely to leave both unchanged.
	//   - Status codes must not overlap between success and failure criteria.
	AuthenticationSettings param.Field[FraudSettingsAuthenticationSettingsParam] `json:"authentication_settings"`
	// Whether Fraud User Profiles is enabled for the zone.
	UserProfiles param.Field[FraudSettingsUserProfiles] `json:"user_profiles"`
	// List of expressions to detect usernames in write HTTP requests.
	//
	//   - Maximum of 10 expressions.
	//   - Omit or set to null to leave unchanged on update.
	//   - Provide an empty array `[]` to clear all expressions on update.
	//   - Invalid expressions will result in a 10400 Bad Request with details in the
	//     `messages` array.
	UsernameExpressions param.Field[[]string] `json:"username_expressions"`
}

func (r FraudSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Configuration for classifying login authentication outcomes based on the origin
// response. Requires `user_profiles` to be enabled.
//
//   - Success and failure criteria are independently updatable — sending only
//     `success_criteria` leaves failure codes untouched, and vice versa.
//   - Omit `authentication_settings` entirely to leave both unchanged.
//   - Status codes must not overlap between success and failure criteria.
type FraudSettingsAuthenticationSettingsParam struct {
	// Criterion for identifying failed login responses.
	FailureCriteria param.Field[FraudSettingsAuthenticationSettingsFailureCriteriaParam] `json:"failure_criteria"`
	// Criterion for identifying successful login responses.
	SuccessCriteria param.Field[FraudSettingsAuthenticationSettingsSuccessCriteriaParam] `json:"success_criteria"`
}

func (r FraudSettingsAuthenticationSettingsParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Criterion for identifying failed login responses.
type FraudSettingsAuthenticationSettingsFailureCriteriaParam struct {
	// The type of criterion. Currently only `status_code` is supported.
	Kind param.Field[FraudSettingsAuthenticationSettingsFailureCriteriaKind] `json:"kind" api:"required"`
	// HTTP status codes to match against the origin response.
	//
	// - Maximum of 10 codes per criterion.
	// - Each code must be a valid HTTP status code (100-599).
	// - Codes are deduplicated and sorted on save.
	// - Omit to leave unchanged on update.
	// - Provide an empty array `[]` to clear codes on update.
	StatusCodes param.Field[[]int64] `json:"status_codes"`
}

func (r FraudSettingsAuthenticationSettingsFailureCriteriaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Criterion for identifying successful login responses.
type FraudSettingsAuthenticationSettingsSuccessCriteriaParam struct {
	// The type of criterion. Currently only `status_code` is supported.
	Kind param.Field[FraudSettingsAuthenticationSettingsSuccessCriteriaKind] `json:"kind" api:"required"`
	// HTTP status codes to match against the origin response.
	//
	// - Maximum of 10 codes per criterion.
	// - Each code must be a valid HTTP status code (100-599).
	// - Codes are deduplicated and sorted on save.
	// - Omit to leave unchanged on update.
	// - Provide an empty array `[]` to clear codes on update.
	StatusCodes param.Field[[]int64] `json:"status_codes"`
}

func (r FraudSettingsAuthenticationSettingsSuccessCriteriaParam) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type FraudUpdateParams struct {
	// Identifier.
	ZoneID        param.Field[string] `path:"zone_id" api:"required"`
	FraudSettings FraudSettingsParam  `json:"fraud_settings" api:"required"`
}

func (r FraudUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r.FraudSettings)
}

type FraudUpdateResponseEnvelope struct {
	Errors   []FraudUpdateResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []FraudUpdateResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success FraudUpdateResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  FraudSettings                      `json:"result"`
	JSON    fraudUpdateResponseEnvelopeJSON    `json:"-"`
}

// fraudUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [FraudUpdateResponseEnvelope]
type fraudUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FraudUpdateResponseEnvelopeErrors struct {
	Code             int64                                   `json:"code" api:"required"`
	Message          string                                  `json:"message" api:"required"`
	DocumentationURL string                                  `json:"documentation_url"`
	Source           FraudUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             fraudUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// fraudUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FraudUpdateResponseEnvelopeErrors]
type fraudUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FraudUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FraudUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                      `json:"pointer"`
	JSON    fraudUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fraudUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [FraudUpdateResponseEnvelopeErrorsSource]
type fraudUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FraudUpdateResponseEnvelopeMessages struct {
	Code             int64                                     `json:"code" api:"required"`
	Message          string                                    `json:"message" api:"required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           FraudUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             fraudUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// fraudUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [FraudUpdateResponseEnvelopeMessages]
type fraudUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FraudUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FraudUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    fraudUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// fraudUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [FraudUpdateResponseEnvelopeMessagesSource]
type fraudUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type FraudUpdateResponseEnvelopeSuccess bool

const (
	FraudUpdateResponseEnvelopeSuccessTrue FraudUpdateResponseEnvelopeSuccess = true
)

func (r FraudUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FraudUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type FraudGetParams struct {
	// Identifier.
	ZoneID param.Field[string] `path:"zone_id" api:"required"`
}

type FraudGetResponseEnvelope struct {
	Errors   []FraudGetResponseEnvelopeErrors   `json:"errors" api:"required"`
	Messages []FraudGetResponseEnvelopeMessages `json:"messages" api:"required"`
	// Whether the API call was successful.
	Success FraudGetResponseEnvelopeSuccess `json:"success" api:"required"`
	Result  FraudSettings                   `json:"result"`
	JSON    fraudGetResponseEnvelopeJSON    `json:"-"`
}

// fraudGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [FraudGetResponseEnvelope]
type fraudGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type FraudGetResponseEnvelopeErrors struct {
	Code             int64                                `json:"code" api:"required"`
	Message          string                               `json:"message" api:"required"`
	DocumentationURL string                               `json:"documentation_url"`
	Source           FraudGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             fraudGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// fraudGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [FraudGetResponseEnvelopeErrors]
type fraudGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FraudGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type FraudGetResponseEnvelopeErrorsSource struct {
	Pointer string                                   `json:"pointer"`
	JSON    fraudGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// fraudGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [FraudGetResponseEnvelopeErrorsSource]
type fraudGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type FraudGetResponseEnvelopeMessages struct {
	Code             int64                                  `json:"code" api:"required"`
	Message          string                                 `json:"message" api:"required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           FraudGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             fraudGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// fraudGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [FraudGetResponseEnvelopeMessages]
type fraudGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *FraudGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type FraudGetResponseEnvelopeMessagesSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    fraudGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// fraudGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [FraudGetResponseEnvelopeMessagesSource]
type fraudGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *FraudGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r fraudGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type FraudGetResponseEnvelopeSuccess bool

const (
	FraudGetResponseEnvelopeSuccessTrue FraudGetResponseEnvelopeSuccess = true
)

func (r FraudGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case FraudGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
