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

// AccountAlertingV3PolicyService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewAccountAlertingV3PolicyService] method instead.
type AccountAlertingV3PolicyService struct {
	Options []option.RequestOption
}

// NewAccountAlertingV3PolicyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccountAlertingV3PolicyService(opts ...option.RequestOption) (r *AccountAlertingV3PolicyService) {
	r = &AccountAlertingV3PolicyService{}
	r.Options = opts
	return
}

// Get details for a single policy.
func (r *AccountAlertingV3PolicyService) Get(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *PoliciesSingleResponse2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a Notification policy.
func (r *AccountAlertingV3PolicyService) Update(ctx context.Context, identifier string, uuid string, body AccountAlertingV3PolicyUpdateParams, opts ...option.RequestOption) (res *PoliciesIDResponse2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Delete a Notification policy.
func (r *AccountAlertingV3PolicyService) Delete(ctx context.Context, identifier string, uuid string, opts ...option.RequestOption) (res *APIResponseCollection, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies/%s", identifier, uuid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates a new Notification policy.
func (r *AccountAlertingV3PolicyService) NotificationPoliciesNewANotificationPolicy(ctx context.Context, identifier string, body AccountAlertingV3PolicyNotificationPoliciesNewANotificationPolicyParams, opts ...option.RequestOption) (res *PoliciesIDResponse2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a list of all Notification policies.
func (r *AccountAlertingV3PolicyService) NotificationPoliciesListNotificationPolicies(ctx context.Context, identifier string, opts ...option.RequestOption) (res *PoliciesResponseCollection2, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%s/alerting/v3/policies", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type PoliciesIDResponse2 struct {
	Errors   []PoliciesIDResponse2Error   `json:"errors"`
	Messages []PoliciesIDResponse2Message `json:"messages"`
	Result   PoliciesIDResponse2Result    `json:"result"`
	// Whether the API call was successful
	Success PoliciesIDResponse2Success `json:"success"`
	JSON    policiesIDResponse2JSON    `json:"-"`
}

// policiesIDResponse2JSON contains the JSON metadata for the struct
// [PoliciesIDResponse2]
type policiesIDResponse2JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesIDResponse2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesIDResponse2Error struct {
	Code    int64                        `json:"code,required"`
	Message string                       `json:"message,required"`
	JSON    policiesIDResponse2ErrorJSON `json:"-"`
}

// policiesIDResponse2ErrorJSON contains the JSON metadata for the struct
// [PoliciesIDResponse2Error]
type policiesIDResponse2ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesIDResponse2Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesIDResponse2Message struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    policiesIDResponse2MessageJSON `json:"-"`
}

// policiesIDResponse2MessageJSON contains the JSON metadata for the struct
// [PoliciesIDResponse2Message]
type policiesIDResponse2MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesIDResponse2Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesIDResponse2Result struct {
	// UUID
	ID   string                        `json:"id"`
	JSON policiesIDResponse2ResultJSON `json:"-"`
}

// policiesIDResponse2ResultJSON contains the JSON metadata for the struct
// [PoliciesIDResponse2Result]
type policiesIDResponse2ResultJSON struct {
	ID          apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesIDResponse2Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PoliciesIDResponse2Success bool

const (
	PoliciesIDResponse2SuccessTrue PoliciesIDResponse2Success = true
)

type PoliciesResponseCollection2 struct {
	Errors     []PoliciesResponseCollection2Error    `json:"errors"`
	Messages   []PoliciesResponseCollection2Message  `json:"messages"`
	Result     []PoliciesResponseCollection2Result   `json:"result"`
	ResultInfo PoliciesResponseCollection2ResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success PoliciesResponseCollection2Success `json:"success"`
	JSON    policiesResponseCollection2JSON    `json:"-"`
}

// policiesResponseCollection2JSON contains the JSON metadata for the struct
// [PoliciesResponseCollection2]
type policiesResponseCollection2JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesResponseCollection2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesResponseCollection2Error struct {
	Code    int64                                `json:"code,required"`
	Message string                               `json:"message,required"`
	JSON    policiesResponseCollection2ErrorJSON `json:"-"`
}

// policiesResponseCollection2ErrorJSON contains the JSON metadata for the struct
// [PoliciesResponseCollection2Error]
type policiesResponseCollection2ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesResponseCollection2Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesResponseCollection2Message struct {
	Code    int64                                  `json:"code,required"`
	Message string                                 `json:"message,required"`
	JSON    policiesResponseCollection2MessageJSON `json:"-"`
}

// policiesResponseCollection2MessageJSON contains the JSON metadata for the struct
// [PoliciesResponseCollection2Message]
type policiesResponseCollection2MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesResponseCollection2Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesResponseCollection2Result struct {
	// UUID
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType string    `json:"alert_type"`
	Created   time.Time `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool        `json:"enabled"`
	Filters interface{} `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms interface{} `json:"mechanisms"`
	Modified   time.Time   `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                                `json:"name"`
	JSON policiesResponseCollection2ResultJSON `json:"-"`
}

// policiesResponseCollection2ResultJSON contains the JSON metadata for the struct
// [PoliciesResponseCollection2Result]
type policiesResponseCollection2ResultJSON struct {
	ID          apijson.Field
	AlertType   apijson.Field
	Created     apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Filters     apijson.Field
	Mechanisms  apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesResponseCollection2Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesResponseCollection2ResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                   `json:"total_count"`
	JSON       policiesResponseCollection2ResultInfoJSON `json:"-"`
}

// policiesResponseCollection2ResultInfoJSON contains the JSON metadata for the
// struct [PoliciesResponseCollection2ResultInfo]
type policiesResponseCollection2ResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesResponseCollection2ResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PoliciesResponseCollection2Success bool

const (
	PoliciesResponseCollection2SuccessTrue PoliciesResponseCollection2Success = true
)

type PoliciesSingleResponse2 struct {
	Errors   []PoliciesSingleResponse2Error   `json:"errors"`
	Messages []PoliciesSingleResponse2Message `json:"messages"`
	Result   PoliciesSingleResponse2Result    `json:"result"`
	// Whether the API call was successful
	Success PoliciesSingleResponse2Success `json:"success"`
	JSON    policiesSingleResponse2JSON    `json:"-"`
}

// policiesSingleResponse2JSON contains the JSON metadata for the struct
// [PoliciesSingleResponse2]
type policiesSingleResponse2JSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesSingleResponse2) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesSingleResponse2Error struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    policiesSingleResponse2ErrorJSON `json:"-"`
}

// policiesSingleResponse2ErrorJSON contains the JSON metadata for the struct
// [PoliciesSingleResponse2Error]
type policiesSingleResponse2ErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesSingleResponse2Error) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesSingleResponse2Message struct {
	Code    int64                              `json:"code,required"`
	Message string                             `json:"message,required"`
	JSON    policiesSingleResponse2MessageJSON `json:"-"`
}

// policiesSingleResponse2MessageJSON contains the JSON metadata for the struct
// [PoliciesSingleResponse2Message]
type policiesSingleResponse2MessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesSingleResponse2Message) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type PoliciesSingleResponse2Result struct {
	// UUID
	ID string `json:"id"`
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType string    `json:"alert_type"`
	Created   time.Time `json:"created" format:"date-time"`
	// Optional description for the Notification policy.
	Description string `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled bool        `json:"enabled"`
	Filters interface{} `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms interface{} `json:"mechanisms"`
	Modified   time.Time   `json:"modified" format:"date-time"`
	// Name of the policy.
	Name string                            `json:"name"`
	JSON policiesSingleResponse2ResultJSON `json:"-"`
}

// policiesSingleResponse2ResultJSON contains the JSON metadata for the struct
// [PoliciesSingleResponse2Result]
type policiesSingleResponse2ResultJSON struct {
	ID          apijson.Field
	AlertType   apijson.Field
	Created     apijson.Field
	Description apijson.Field
	Enabled     apijson.Field
	Filters     apijson.Field
	Mechanisms  apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *PoliciesSingleResponse2Result) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type PoliciesSingleResponse2Success bool

const (
	PoliciesSingleResponse2SuccessTrue PoliciesSingleResponse2Success = true
)

type AccountAlertingV3PolicyUpdateParams struct {
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[string] `json:"alert_type"`
	// Optional description for the Notification policy.
	Description param.Field[string] `json:"description"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool]        `json:"enabled"`
	Filters param.Field[interface{}] `json:"filters"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[interface{}] `json:"mechanisms"`
	// Name of the policy.
	Name param.Field[string] `json:"name"`
}

func (r AccountAlertingV3PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAlertingV3PolicyNotificationPoliciesNewANotificationPolicyParams struct {
	// Refers to which event will trigger a Notification dispatch. You can use the
	// endpoint to get available alert types which then will give you a list of
	// possible values.
	AlertType param.Field[string] `json:"alert_type,required"`
	// Whether or not the Notification policy is enabled.
	Enabled param.Field[bool] `json:"enabled,required"`
	// List of IDs that will be used when dispatching a notification. IDs for email
	// type will be the email address.
	Mechanisms param.Field[interface{}] `json:"mechanisms,required"`
	// Name of the policy.
	Name param.Field[string] `json:"name,required"`
	// Optional description for the Notification policy.
	Description param.Field[string]      `json:"description"`
	Filters     param.Field[interface{}] `json:"filters"`
}

func (r AccountAlertingV3PolicyNotificationPoliciesNewANotificationPolicyParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
