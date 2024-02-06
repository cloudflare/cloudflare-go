// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccessLogAccessRequestService contains methods and other services that help with
// interacting with the cloudflare API. Note, unlike clients, this service does not
// read variables from the environment automatically. You should not instantiate
// this service directly, and instead use the [NewAccessLogAccessRequestService]
// method instead.
type AccessLogAccessRequestService struct {
	Options []option.RequestOption
}

// NewAccessLogAccessRequestService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAccessLogAccessRequestService(opts ...option.RequestOption) (r *AccessLogAccessRequestService) {
	r = &AccessLogAccessRequestService{}
	r.Options = opts
	return
}

// Gets a list of Access authentication audit logs for an account.
func (r *AccessLogAccessRequestService) AccessAuthenticationLogsGetAccessAuthenticationLogs(ctx context.Context, identifier string, opts ...option.RequestOption) (res *[]AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelope
	path := fmt.Sprintf("accounts/%s/access/logs/access_requests", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse struct {
	// The event that occurred, such as a login attempt.
	Action string `json:"action"`
	// The result of the authentication event.
	Allowed bool `json:"allowed"`
	// The URL of the Access application.
	AppDomain string `json:"app_domain"`
	// The unique identifier for the Access application.
	AppUid interface{} `json:"app_uid"`
	// The IdP used to authenticate.
	Connection string    `json:"connection"`
	CreatedAt  time.Time `json:"created_at" format:"date-time"`
	// The IP address of the authenticating user.
	IPAddress string `json:"ip_address"`
	// The unique identifier for the request to Cloudflare.
	RayID string `json:"ray_id"`
	// The email address of the authenticating user.
	UserEmail string                                                                                `json:"user_email" format:"email"`
	JSON      accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseJSON `json:"-"`
}

// accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseJSON
// contains the JSON metadata for the struct
// [AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse]
type accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseJSON struct {
	Action      apijson.Field
	Allowed     apijson.Field
	AppDomain   apijson.Field
	AppUid      apijson.Field
	Connection  apijson.Field
	CreatedAt   apijson.Field
	IPAddress   apijson.Field
	RayID       apijson.Field
	UserEmail   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelope struct {
	Errors     []AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeErrors   `json:"errors"`
	Messages   []AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeMessages `json:"messages"`
	Result     []AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponse                 `json:"result"`
	ResultInfo AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeSuccess `json:"success"`
	JSON    accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeJSON    `json:"-"`
}

// accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelope]
type accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeErrors struct {
	Code    int64                                                                                               `json:"code,required"`
	Message string                                                                                              `json:"message,required"`
	JSON    accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeErrorsJSON `json:"-"`
}

// accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeErrors]
type accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeMessages struct {
	Code    int64                                                                                                 `json:"code,required"`
	Message string                                                                                                `json:"message,required"`
	JSON    accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeMessagesJSON `json:"-"`
}

// accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeMessages]
type accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                                 `json:"total_count"`
	JSON       accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeResultInfoJSON `json:"-"`
}

// accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeResultInfo]
type accessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeSuccess bool

const (
	AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeSuccessTrue AccessLogAccessRequestAccessAuthenticationLogsGetAccessAuthenticationLogsResponseEnvelopeSuccess = true
)
