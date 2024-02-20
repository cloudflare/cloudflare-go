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
func (r *UserTokenPermissionGroupService) List(ctx context.Context, opts ...option.RequestOption) (res *[]UserTokenPermissionGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	var env UserTokenPermissionGroupListResponseEnvelope
	path := "user/tokens/permission_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type UserTokenPermissionGroupListResponse = interface{}

type UserTokenPermissionGroupListResponseEnvelope struct {
	Errors   []UserTokenPermissionGroupListResponseEnvelopeErrors   `json:"errors,required"`
	Messages []UserTokenPermissionGroupListResponseEnvelopeMessages `json:"messages,required"`
	Result   []UserTokenPermissionGroupListResponse                 `json:"result,required,nullable"`
	// Whether the API call was successful
	Success    UserTokenPermissionGroupListResponseEnvelopeSuccess    `json:"success,required"`
	ResultInfo UserTokenPermissionGroupListResponseEnvelopeResultInfo `json:"result_info"`
	JSON       userTokenPermissionGroupListResponseEnvelopeJSON       `json:"-"`
}

// userTokenPermissionGroupListResponseEnvelopeJSON contains the JSON metadata for
// the struct [UserTokenPermissionGroupListResponseEnvelope]
type userTokenPermissionGroupListResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	ResultInfo  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupListResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenPermissionGroupListResponseEnvelopeErrors struct {
	Code    int64                                                  `json:"code,required"`
	Message string                                                 `json:"message,required"`
	JSON    userTokenPermissionGroupListResponseEnvelopeErrorsJSON `json:"-"`
}

// userTokenPermissionGroupListResponseEnvelopeErrorsJSON contains the JSON
// metadata for the struct [UserTokenPermissionGroupListResponseEnvelopeErrors]
type userTokenPermissionGroupListResponseEnvelopeErrorsJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupListResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type UserTokenPermissionGroupListResponseEnvelopeMessages struct {
	Code    int64                                                    `json:"code,required"`
	Message string                                                   `json:"message,required"`
	JSON    userTokenPermissionGroupListResponseEnvelopeMessagesJSON `json:"-"`
}

// userTokenPermissionGroupListResponseEnvelopeMessagesJSON contains the JSON
// metadata for the struct [UserTokenPermissionGroupListResponseEnvelopeMessages]
type userTokenPermissionGroupListResponseEnvelopeMessagesJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupListResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type UserTokenPermissionGroupListResponseEnvelopeSuccess bool

const (
	UserTokenPermissionGroupListResponseEnvelopeSuccessTrue UserTokenPermissionGroupListResponseEnvelopeSuccess = true
)

type UserTokenPermissionGroupListResponseEnvelopeResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                                    `json:"total_count"`
	JSON       userTokenPermissionGroupListResponseEnvelopeResultInfoJSON `json:"-"`
}

// userTokenPermissionGroupListResponseEnvelopeResultInfoJSON contains the JSON
// metadata for the struct [UserTokenPermissionGroupListResponseEnvelopeResultInfo]
type userTokenPermissionGroupListResponseEnvelopeResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserTokenPermissionGroupListResponseEnvelopeResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}
