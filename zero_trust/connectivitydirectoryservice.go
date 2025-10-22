// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package zero_trust

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

// ConnectivityDirectoryServiceService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectivityDirectoryServiceService] method instead.
type ConnectivityDirectoryServiceService struct {
	Options []option.RequestOption
}

// NewConnectivityDirectoryServiceService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectivityDirectoryServiceService(opts ...option.RequestOption) (r *ConnectivityDirectoryServiceService) {
	r = &ConnectivityDirectoryServiceService{}
	r.Options = opts
	return
}

// Create connectivity service
func (r *ConnectivityDirectoryServiceService) New(ctx context.Context, params ConnectivityDirectoryServiceNewParams, opts ...option.RequestOption) (err error) {
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
func (r *ConnectivityDirectoryServiceService) Update(ctx context.Context, serviceID string, params ConnectivityDirectoryServiceUpdateParams, opts ...option.RequestOption) (err error) {
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
func (r *ConnectivityDirectoryServiceService) List(ctx context.Context, params ConnectivityDirectoryServiceListParams, opts ...option.RequestOption) (err error) {
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
func (r *ConnectivityDirectoryServiceService) Delete(ctx context.Context, serviceID string, body ConnectivityDirectoryServiceDeleteParams, opts ...option.RequestOption) (err error) {
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
func (r *ConnectivityDirectoryServiceService) Get(ctx context.Context, serviceID string, query ConnectivityDirectoryServiceGetParams, opts ...option.RequestOption) (err error) {
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

type ConnectivityDirectoryServiceNewParams struct {
	// Account identifier
	AccountID param.Field[string]                                    `path:"account_id,required"`
	Host      param.Field[ConnectivityDirectoryServiceNewParamsHost] `json:"host,required"`
	Name      param.Field[string]                                    `json:"name,required"`
	Type      param.Field[ConnectivityDirectoryServiceNewParamsType] `json:"type,required"`
	HTTPPort  param.Field[int64]                                     `json:"http_port"`
	HTTPSPort param.Field[int64]                                     `json:"https_port"`
}

func (r ConnectivityDirectoryServiceNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceNewParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r ConnectivityDirectoryServiceNewParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceNewParamsType string

const (
	ConnectivityDirectoryServiceNewParamsTypeHTTP ConnectivityDirectoryServiceNewParamsType = "http"
)

func (r ConnectivityDirectoryServiceNewParamsType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceNewParamsTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceUpdateParams struct {
	AccountID param.Field[string]                                       `path:"account_id,required"`
	Host      param.Field[ConnectivityDirectoryServiceUpdateParamsHost] `json:"host,required"`
	Name      param.Field[string]                                       `json:"name,required"`
	Type      param.Field[ConnectivityDirectoryServiceUpdateParamsType] `json:"type,required"`
	HTTPPort  param.Field[int64]                                        `json:"http_port"`
	HTTPSPort param.Field[int64]                                        `json:"https_port"`
}

func (r ConnectivityDirectoryServiceUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceUpdateParamsHost struct {
	Hostname        param.Field[string]      `json:"hostname"`
	IPV4            param.Field[string]      `json:"ipv4"`
	IPV6            param.Field[string]      `json:"ipv6"`
	Network         param.Field[interface{}] `json:"network"`
	ResolverNetwork param.Field[interface{}] `json:"resolver_network"`
}

func (r ConnectivityDirectoryServiceUpdateParamsHost) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ConnectivityDirectoryServiceUpdateParamsType string

const (
	ConnectivityDirectoryServiceUpdateParamsTypeHTTP ConnectivityDirectoryServiceUpdateParamsType = "http"
)

func (r ConnectivityDirectoryServiceUpdateParamsType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceUpdateParamsTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceListParams struct {
	// Account identifier
	AccountID param.Field[string] `path:"account_id,required"`
	// Current page in the response
	Page param.Field[int64] `query:"page"`
	// Max amount of entries returned per page
	PerPage param.Field[int64]                                      `query:"per_page"`
	Type    param.Field[ConnectivityDirectoryServiceListParamsType] `query:"type"`
}

// URLQuery serializes [ConnectivityDirectoryServiceListParams]'s query parameters
// as `url.Values`.
func (r ConnectivityDirectoryServiceListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ConnectivityDirectoryServiceListParamsType string

const (
	ConnectivityDirectoryServiceListParamsTypeHTTP ConnectivityDirectoryServiceListParamsType = "http"
)

func (r ConnectivityDirectoryServiceListParamsType) IsKnown() bool {
	switch r {
	case ConnectivityDirectoryServiceListParamsTypeHTTP:
		return true
	}
	return false
}

type ConnectivityDirectoryServiceDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type ConnectivityDirectoryServiceGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}
