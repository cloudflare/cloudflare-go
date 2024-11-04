// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package calls

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/cloudflare/cloudflare-go/v3/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v3/internal/param"
	"github.com/cloudflare/cloudflare-go/v3/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v3/option"
	"github.com/cloudflare/cloudflare-go/v3/packages/pagination"
	"github.com/cloudflare/cloudflare-go/v3/shared"
)

// SfuService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSfuService] method instead.
type SfuService struct {
	Options []option.RequestOption
}

// NewSfuService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSfuService(opts ...option.RequestOption) (r *SfuService) {
	r = &SfuService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare calls app. An app is an unique enviroment where each
// Session can access all Tracks within the app.
func (r *SfuService) New(ctx context.Context, params SfuNewParams, opts ...option.RequestOption) (res *SfuNewResponse, err error) {
	var env SfuNewResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Edit details for a single app.
func (r *SfuService) Update(ctx context.Context, appID string, params SfuUpdateParams, opts ...option.RequestOption) (res *SfuUpdateResponse, err error) {
	var env SfuUpdateResponseEnvelope
	opts = append(r.Options[:], opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", params.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Lists all apps in the Cloudflare account
func (r *SfuService) List(ctx context.Context, query SfuListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SfuListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps", query.AccountID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, nil, &res, opts...)
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

// Lists all apps in the Cloudflare account
func (r *SfuService) ListAutoPaging(ctx context.Context, query SfuListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SfuListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an app from Cloudflare Calls
func (r *SfuService) Delete(ctx context.Context, appID string, body SfuDeleteParams, opts ...option.RequestOption) (res *SfuDeleteResponse, err error) {
	var env SfuDeleteResponseEnvelope
	opts = append(r.Options[:], opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", body.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetches details for a single Calls app.
func (r *SfuService) Get(ctx context.Context, appID string, query SfuGetParams, opts ...option.RequestOption) (res *SfuGetResponse, err error) {
	var env SfuGetResponseEnvelope
	opts = append(r.Options[:], opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if appID == "" {
		err = errors.New("missing required app_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/calls/apps/%s", query.AccountID, appID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type SfuNewResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// Bearer token
	Secret string `json:"secret"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string             `json:"uid"`
	JSON sfuNewResponseJSON `json:"-"`
}

// sfuNewResponseJSON contains the JSON metadata for the struct [SfuNewResponse]
type sfuNewResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuNewResponseJSON) RawJSON() string {
	return r.raw
}

type SfuUpdateResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string                `json:"uid"`
	JSON sfuUpdateResponseJSON `json:"-"`
}

// sfuUpdateResponseJSON contains the JSON metadata for the struct
// [SfuUpdateResponse]
type sfuUpdateResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SfuListResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string              `json:"uid"`
	JSON sfuListResponseJSON `json:"-"`
}

// sfuListResponseJSON contains the JSON metadata for the struct [SfuListResponse]
type sfuListResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuListResponseJSON) RawJSON() string {
	return r.raw
}

type SfuDeleteResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string                `json:"uid"`
	JSON sfuDeleteResponseJSON `json:"-"`
}

// sfuDeleteResponseJSON contains the JSON metadata for the struct
// [SfuDeleteResponse]
type sfuDeleteResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SfuGetResponse struct {
	// The date and time the item was created.
	Created time.Time `json:"created" format:"date-time"`
	// The date and time the item was last modified.
	Modified time.Time `json:"modified" format:"date-time"`
	// A short description of Calls app, not shown to end users.
	Name string `json:"name"`
	// A Cloudflare-generated unique identifier for a item.
	UID  string             `json:"uid"`
	JSON sfuGetResponseJSON `json:"-"`
}

// sfuGetResponseJSON contains the JSON metadata for the struct [SfuGetResponse]
type sfuGetResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuGetResponseJSON) RawJSON() string {
	return r.raw
}

type SfuNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r SfuNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SfuNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SfuNewResponseEnvelopeSuccess `json:"success,required"`
	Result  SfuNewResponse                `json:"result"`
	JSON    sfuNewResponseEnvelopeJSON    `json:"-"`
}

// sfuNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SfuNewResponseEnvelope]
type sfuNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SfuNewResponseEnvelopeSuccess bool

const (
	SfuNewResponseEnvelopeSuccessTrue SfuNewResponseEnvelopeSuccess = true
)

func (r SfuNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SfuNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SfuUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r SfuUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SfuUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SfuUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  SfuUpdateResponse                `json:"result"`
	JSON    sfuUpdateResponseEnvelopeJSON    `json:"-"`
}

// sfuUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SfuUpdateResponseEnvelope]
type sfuUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SfuUpdateResponseEnvelopeSuccess bool

const (
	SfuUpdateResponseEnvelopeSuccessTrue SfuUpdateResponseEnvelopeSuccess = true
)

func (r SfuUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SfuUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SfuListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SfuDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SfuDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SfuDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  SfuDeleteResponse                `json:"result"`
	JSON    sfuDeleteResponseEnvelopeJSON    `json:"-"`
}

// sfuDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SfuDeleteResponseEnvelope]
type sfuDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SfuDeleteResponseEnvelopeSuccess bool

const (
	SfuDeleteResponseEnvelopeSuccessTrue SfuDeleteResponseEnvelopeSuccess = true
)

func (r SfuDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SfuDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SfuGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SfuGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SfuGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SfuGetResponse                `json:"result"`
	JSON    sfuGetResponseEnvelopeJSON    `json:"-"`
}

// sfuGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SfuGetResponseEnvelope]
type sfuGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SfuGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SfuGetResponseEnvelopeSuccess bool

const (
	SfuGetResponseEnvelopeSuccessTrue SfuGetResponseEnvelopeSuccess = true
)

func (r SfuGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SfuGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
