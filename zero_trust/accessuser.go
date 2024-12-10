// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v4/internal/param"
	"github.com/cloudflare/cloudflare-go/v4/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v4/option"
	"github.com/cloudflare/cloudflare-go/v4/packages/pagination"
)

// AccessUserService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAccessUserService] method instead.
type AccessUserService struct {
	Options          []option.RequestOption
	ActiveSessions   *AccessUserActiveSessionService
	LastSeenIdentity *AccessUserLastSeenIdentityService
	FailedLogins     *AccessUserFailedLoginService
}

// NewAccessUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAccessUserService(opts ...option.RequestOption) (r *AccessUserService) {
	r = &AccessUserService{}
	r.Options = opts
	r.ActiveSessions = NewAccessUserActiveSessionService(opts...)
	r.LastSeenIdentity = NewAccessUserLastSeenIdentityService(opts...)
	r.FailedLogins = NewAccessUserFailedLoginService(opts...)
	return
}

// Gets a list of users for an account.
func (r *AccessUserService) List(ctx context.Context, params AccessUserListParams, opts ...option.RequestOption) (res *pagination.SinglePage[AccessUser], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/access/users", params.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Gets a list of users for an account.
func (r *AccessUserService) ListAutoPaging(ctx context.Context, params AccessUserListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[AccessUser] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, params, opts...))
}

type AccessUser struct {
	// UUID
	ID string `json:"id"`
	// True if the user has authenticated with Cloudflare Access.
	AccessSeat bool `json:"access_seat"`
	// The number of active devices registered to the user.
	ActiveDeviceCount float64   `json:"active_device_count"`
	CreatedAt         time.Time `json:"created_at" format:"date-time"`
	// The email of the user.
	Email string `json:"email" format:"email"`
	// True if the user has logged into the WARP client.
	GatewaySeat bool `json:"gateway_seat"`
	// The time at which the user last successfully logged in.
	LastSuccessfulLogin time.Time `json:"last_successful_login" format:"date-time"`
	// The name of the user.
	Name string `json:"name"`
	// The unique API identifier for the Zero Trust seat.
	SeatUID string `json:"seat_uid"`
	// The unique API identifier for the user.
	UID       string         `json:"uid"`
	UpdatedAt time.Time      `json:"updated_at" format:"date-time"`
	JSON      accessUserJSON `json:"-"`
}

// accessUserJSON contains the JSON metadata for the struct [AccessUser]
type accessUserJSON struct {
	ID                  apijson.Field
	AccessSeat          apijson.Field
	ActiveDeviceCount   apijson.Field
	CreatedAt           apijson.Field
	Email               apijson.Field
	GatewaySeat         apijson.Field
	LastSuccessfulLogin apijson.Field
	Name                apijson.Field
	SeatUID             apijson.Field
	UID                 apijson.Field
	UpdatedAt           apijson.Field
	raw                 string
	ExtraFields         map[string]apijson.Field
}

func (r *AccessUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r accessUserJSON) RawJSON() string {
	return r.raw
}

type AccessUserListParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// The email of the user.
	Email param.Field[string] `query:"email"`
	// The name of the user.
	Name param.Field[string] `query:"name"`
	// Search for users by other listed query parameters.
	Search param.Field[string] `query:"search"`
}

// URLQuery serializes [AccessUserListParams]'s query parameters as `url.Values`.
func (r AccessUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}
