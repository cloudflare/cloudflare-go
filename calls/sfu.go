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

// SFUService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSFUService] method instead.
type SFUService struct {
	Options []option.RequestOption
}

// NewSFUService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSFUService(opts ...option.RequestOption) (r *SFUService) {
	r = &SFUService{}
	r.Options = opts
	return
}

// Creates a new Cloudflare calls app. An app is an unique enviroment where each
// Session can access all Tracks within the app.
func (r *SFUService) New(ctx context.Context, params SFUNewParams, opts ...option.RequestOption) (res *SFUNewResponse, err error) {
	var env SFUNewResponseEnvelope
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
func (r *SFUService) Update(ctx context.Context, appID string, params SFUUpdateParams, opts ...option.RequestOption) (res *SFUUpdateResponse, err error) {
	var env SFUUpdateResponseEnvelope
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
func (r *SFUService) List(ctx context.Context, query SFUListParams, opts ...option.RequestOption) (res *pagination.SinglePage[SFUListResponse], err error) {
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
func (r *SFUService) ListAutoPaging(ctx context.Context, query SFUListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[SFUListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, query, opts...))
}

// Deletes an app from Cloudflare Calls
func (r *SFUService) Delete(ctx context.Context, appID string, body SFUDeleteParams, opts ...option.RequestOption) (res *SFUDeleteResponse, err error) {
	var env SFUDeleteResponseEnvelope
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
func (r *SFUService) Get(ctx context.Context, appID string, query SFUGetParams, opts ...option.RequestOption) (res *SFUGetResponse, err error) {
	var env SFUGetResponseEnvelope
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

type SFUNewResponse struct {
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

// sfuNewResponseJSON contains the JSON metadata for the struct [SFUNewResponse]
type sfuNewResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	Secret      apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuNewResponseJSON) RawJSON() string {
	return r.raw
}

type SFUUpdateResponse struct {
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
// [SFUUpdateResponse]
type sfuUpdateResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type SFUListResponse struct {
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

// sfuListResponseJSON contains the JSON metadata for the struct [SFUListResponse]
type sfuListResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuListResponseJSON) RawJSON() string {
	return r.raw
}

type SFUDeleteResponse struct {
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
// [SFUDeleteResponse]
type sfuDeleteResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type SFUGetResponse struct {
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

// sfuGetResponseJSON contains the JSON metadata for the struct [SFUGetResponse]
type sfuGetResponseJSON struct {
	Created     apijson.Field
	Modified    apijson.Field
	Name        apijson.Field
	UID         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuGetResponseJSON) RawJSON() string {
	return r.raw
}

type SFUNewParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r SFUNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SFUNewResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SFUNewResponseEnvelopeSuccess `json:"success,required"`
	Result  SFUNewResponse                `json:"result"`
	JSON    sfuNewResponseEnvelopeJSON    `json:"-"`
}

// sfuNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [SFUNewResponseEnvelope]
type sfuNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SFUNewResponseEnvelopeSuccess bool

const (
	SFUNewResponseEnvelopeSuccessTrue SFUNewResponseEnvelopeSuccess = true
)

func (r SFUNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SFUNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SFUUpdateParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
	// A short description of Calls app, not shown to end users.
	Name param.Field[string] `json:"name"`
}

func (r SFUUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type SFUUpdateResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SFUUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  SFUUpdateResponse                `json:"result"`
	JSON    sfuUpdateResponseEnvelopeJSON    `json:"-"`
}

// sfuUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [SFUUpdateResponseEnvelope]
type sfuUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SFUUpdateResponseEnvelopeSuccess bool

const (
	SFUUpdateResponseEnvelopeSuccessTrue SFUUpdateResponseEnvelopeSuccess = true
)

func (r SFUUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SFUUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SFUListParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SFUDeleteParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SFUDeleteResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SFUDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  SFUDeleteResponse                `json:"result"`
	JSON    sfuDeleteResponseEnvelopeJSON    `json:"-"`
}

// sfuDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [SFUDeleteResponseEnvelope]
type sfuDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SFUDeleteResponseEnvelopeSuccess bool

const (
	SFUDeleteResponseEnvelopeSuccessTrue SFUDeleteResponseEnvelopeSuccess = true
)

func (r SFUDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SFUDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type SFUGetParams struct {
	// The account identifier tag.
	AccountID param.Field[string] `path:"account_id,required"`
}

type SFUGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	// Whether the API call was successful
	Success SFUGetResponseEnvelopeSuccess `json:"success,required"`
	Result  SFUGetResponse                `json:"result"`
	JSON    sfuGetResponseEnvelopeJSON    `json:"-"`
}

// sfuGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [SFUGetResponseEnvelope]
type sfuGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *SFUGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r sfuGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type SFUGetResponseEnvelopeSuccess bool

const (
	SFUGetResponseEnvelopeSuccessTrue SFUGetResponseEnvelopeSuccess = true
)

func (r SFUGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case SFUGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
