// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package load_balancers

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
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// MonitorGroupReferenceService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMonitorGroupReferenceService] method instead.
type MonitorGroupReferenceService struct {
	Options []option.RequestOption
}

// NewMonitorGroupReferenceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMonitorGroupReferenceService(opts ...option.RequestOption) (r *MonitorGroupReferenceService) {
	r = &MonitorGroupReferenceService{}
	r.Options = opts
	return
}

// Get the list of resources that reference the provided monitor group.
func (r *MonitorGroupReferenceService) Get(ctx context.Context, monitorGroupID string, query MonitorGroupReferenceGetParams, opts ...option.RequestOption) (res *pagination.SinglePage[MonitorGroupReferenceGetResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if monitorGroupID == "" {
		err = errors.New("missing required monitor_group_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/load_balancers/monitor_groups/%s/references", query.AccountID, monitorGroupID)
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

// Get the list of resources that reference the provided monitor group.
func (r *MonitorGroupReferenceService) GetAutoPaging(ctx context.Context, monitorGroupID string, query MonitorGroupReferenceGetParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[MonitorGroupReferenceGetResponse] {
	return pagination.NewSinglePageAutoPager(r.Get(ctx, monitorGroupID, query, opts...))
}

type MonitorGroupReferenceGetResponse struct {
	ReferenceType MonitorGroupReferenceGetResponseReferenceType `json:"reference_type"`
	ResourceID    string                                        `json:"resource_id"`
	ResourceName  string                                        `json:"resource_name"`
	ResourceType  string                                        `json:"resource_type"`
	JSON          monitorGroupReferenceGetResponseJSON          `json:"-"`
}

// monitorGroupReferenceGetResponseJSON contains the JSON metadata for the struct
// [MonitorGroupReferenceGetResponse]
type monitorGroupReferenceGetResponseJSON struct {
	ReferenceType apijson.Field
	ResourceID    apijson.Field
	ResourceName  apijson.Field
	ResourceType  apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MonitorGroupReferenceGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r monitorGroupReferenceGetResponseJSON) RawJSON() string {
	return r.raw
}

type MonitorGroupReferenceGetResponseReferenceType string

const (
	MonitorGroupReferenceGetResponseReferenceTypeStar     MonitorGroupReferenceGetResponseReferenceType = "*"
	MonitorGroupReferenceGetResponseReferenceTypeReferral MonitorGroupReferenceGetResponseReferenceType = "referral"
	MonitorGroupReferenceGetResponseReferenceTypeReferrer MonitorGroupReferenceGetResponseReferenceType = "referrer"
)

func (r MonitorGroupReferenceGetResponseReferenceType) IsKnown() bool {
	switch r {
	case MonitorGroupReferenceGetResponseReferenceTypeStar, MonitorGroupReferenceGetResponseReferenceTypeReferral, MonitorGroupReferenceGetResponseReferenceTypeReferrer:
		return true
	}
	return false
}

type MonitorGroupReferenceGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}
