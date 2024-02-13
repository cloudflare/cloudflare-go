// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"net/http"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// UserTokenPermissionGroupService contains methods and other services that help
// with interacting with the cloudflare API. Note, unlike clients, this service
// does not read variables from the environment automatically. You should not
// instantiate this service directly, and instead use the
// [NewUserTokenPermissionGroupService] method instead.
type UserTokenPermissionGroupService struct {
	Options []option.RequestOption
}

// NewUserTokenPermissionGroupService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewUserTokenPermissionGroupService(opts ...option.RequestOption) (r *UserTokenPermissionGroupService) {
	r = &UserTokenPermissionGroupService{}
	r.Options = opts
	return
}

// Find all available permission groups.
func (r *UserTokenPermissionGroupService) PermissionGroupsListPermissionGroups(ctx context.Context, opts ...option.RequestOption) (res *[]UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelope
	path := "user/tokens/permission_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponse = interface{}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelope struct {
	Errors   []UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeJSON       `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelope]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeErrors struct {
	Code    int64                                                                                  `json:"code,required"`
	Message string                                                                                 `json:"message,required"`
	JSON    userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeErrorsJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeErrors]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeMessages struct {
	Code    int64                                                                                    `json:"code,required"`
	Message string                                                                                   `json:"message,required"`
	JSON    userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeMessagesJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeMessages]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeSuccess bool

const (
	UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeSuccessTrue UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeSuccess = true
)

type UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                                                    `json:"total_count"`
	JSON       userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeResultInfoJSON `json:"-"`
}

// userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeResultInfoJSON
// contains the JSON metadata for the struct
// [UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeResultInfo]
type userTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupPermissionGroupsListPermissionGroupsResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
