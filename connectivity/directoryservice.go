// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package connectivity

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
)

// DirectoryServiceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDirectoryServiceService] method instead.
type DirectoryServiceService struct {
	Options []option.RequestOption
}

// NewDirectoryServiceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDirectoryServiceService(opts ...option.RequestOption) (r *DirectoryServiceService) {
	r = &DirectoryServiceService{}
	r.Options = opts
	return
}

// Create connectivity service
func (r *DirectoryServiceService) New(ctx context.Context, params DirectoryServiceNewParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return
}

// Update connectivity service
func (r *DirectoryServiceService) Update(ctx context.Context, serviceID string, params DirectoryServiceUpdateParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", params.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, nil, opts...)
	return
}

// List connectivity services
func (r *DirectoryServiceService) List(ctx context.Context, params DirectoryServiceListParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, nil, opts...)
	return
}

// Delete connectivity service
func (r *DirectoryServiceService) Delete(ctx context.Context, serviceID string, body DirectoryServiceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", body.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get connectivity service
func (r *DirectoryServiceService) Get(ctx context.Context, serviceID string, query DirectoryServiceGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if serviceID == "" {
		err = errors.New("missing required service_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/connectivity/directory/services/%s", query.AccountID, serviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type DirectoryServiceNewParams struct {
	// Account identifier
	AccountID param.Field[string]                        `path:"account_id,required"`
	Host      param.Field[DirectoryServiceNewParamsHost] `json:"host,required"`
	Name      param.Field[string]                        `json:"name,required"`
	Type      param.Field[DirectoryServiceNewParamsType] `json:"type,required"`
	HTTPPort  param.Field[int64]                         `json:"http_port"`
	HTTPSPort param.Field[int64]                         `json:"https_port"`
}

func (r DirectoryServiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceNewParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceNewParamsType string

const (
	DirectoryServiceNewParamsTypeHTTP DirectoryServiceNewParamsType = "http"
)

func (r DirectoryServiceNewParamsType) IsKnown() bool {
	switch r {
	case DirectoryServiceNewParamsTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceUpdateParams struct {
	AccountID param.Field[string]                           `path:"account_id,required"`
	Host      param.Field[DirectoryServiceUpdateParamsHost] `json:"host,required"`
	Name      param.Field[string]                           `json:"name,required"`
	Type      param.Field[DirectoryServiceUpdateParamsType] `json:"type,required"`
	HTTPPort  param.Field[int64]                            `json:"http_port"`
	HTTPSPort param.Field[int64]                            `json:"https_port"`
}

func (r DirectoryServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r DirectoryServiceUpdateParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DirectoryServiceUpdateParamsType string

const (
	DirectoryServiceUpdateParamsTypeHTTP DirectoryServiceUpdateParamsType = "http"
)

func (r DirectoryServiceUpdateParamsType) IsKnown() bool {
	switch r {
	case DirectoryServiceUpdateParamsTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Current page in the response
	Page param.Field[int64] `query:"page"`
	// Max amount of entries returned per page
	PerPage param.Field[int64]                          `query:"per_page"`
	Type    param.Field[DirectoryServiceListParamsType] `query:"type"`
}

// URLQuery serializes [DirectoryServiceListParams]'s query parameters as
// `url.Values`.
func (r DirectoryServiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DirectoryServiceListParamsType string

const (
	DirectoryServiceListParamsTypeHTTP DirectoryServiceListParamsType = "http"
)

func (r DirectoryServiceListParamsType) IsKnown() bool {
	switch r {
	case DirectoryServiceListParamsTypeHTTP:
		return true
	}
	return false
}

type DirectoryServiceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DirectoryServiceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
