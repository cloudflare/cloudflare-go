// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package resource_tagging

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v6/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v6/internal/param"
	"github.com/cloudflare/cloudflare-go/v6/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v6/option"
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// ValueService contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewValueService] method instead.
type ValueService struct {
	Options []option.RequestOption
}

// NewValueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewValueService(opts ...option.RequestOption) (r *ValueService) {
	r = &ValueService{}
	r.Options = opts
	return
}

// Lists all distinct values for a given tag key, optionally filtered by resource
// type.
func (r *ValueService) List(ctx context.Context, tagKey string, params ValueListParams, opts ...option.RequestOption) (res *pagination.CursorPaginationAfter[string], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if tagKey == "" {
		err = errors.New("missing required tag_key parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/tags/values/%s", params.AccountID, tagKey)
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

// Lists all distinct values for a given tag key, optionally filtered by resource
// type.
func (r *ValueService) ListAutoPaging(ctx context.Context, tagKey string, params ValueListParams, opts ...option.RequestOption) *pagination.CursorPaginationAfterAutoPager[string] {
	return pagination.NewCursorPaginationAfterAutoPager(r.List(ctx, tagKey, params, opts...))
}

type ValueListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	// Cursor for pagination.
	Cursor param.Field[string] `query:"cursor"`
	// Filter by resource type.
	Type param.Field[ValueListParamsType] `query:"type"`
}

// URLQuery serializes [ValueListParams]'s query parameters as `url.Values`.
func (r ValueListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Filter by resource type.
type ValueListParamsType string

const (
	ValueListParamsTypeAccessApplication        ValueListParamsType = "access_application"
	ValueListParamsTypeAccessApplicationPolicy  ValueListParamsType = "access_application_policy"
	ValueListParamsTypeAccessGroup              ValueListParamsType = "access_group"
	ValueListParamsTypeAccount                  ValueListParamsType = "account"
	ValueListParamsTypeAIGateway                ValueListParamsType = "ai_gateway"
	ValueListParamsTypeAlertingPolicy           ValueListParamsType = "alerting_policy"
	ValueListParamsTypeAlertingWebhook          ValueListParamsType = "alerting_webhook"
	ValueListParamsTypeAPIGatewayOperation      ValueListParamsType = "api_gateway_operation"
	ValueListParamsTypeCloudflaredTunnel        ValueListParamsType = "cloudflared_tunnel"
	ValueListParamsTypeCustomCertificate        ValueListParamsType = "custom_certificate"
	ValueListParamsTypeCustomHostname           ValueListParamsType = "custom_hostname"
	ValueListParamsTypeD1Database               ValueListParamsType = "d1_database"
	ValueListParamsTypeDNSRecord                ValueListParamsType = "dns_record"
	ValueListParamsTypeDurableObjectNamespace   ValueListParamsType = "durable_object_namespace"
	ValueListParamsTypeGatewayList              ValueListParamsType = "gateway_list"
	ValueListParamsTypeGatewayRule              ValueListParamsType = "gateway_rule"
	ValueListParamsTypeImage                    ValueListParamsType = "image"
	ValueListParamsTypeKVNamespace              ValueListParamsType = "kv_namespace"
	ValueListParamsTypeManagedClientCertificate ValueListParamsType = "managed_client_certificate"
	ValueListParamsTypeQueue                    ValueListParamsType = "queue"
	ValueListParamsTypeR2Bucket                 ValueListParamsType = "r2_bucket"
	ValueListParamsTypeResourceShare            ValueListParamsType = "resource_share"
	ValueListParamsTypeStreamLiveInput          ValueListParamsType = "stream_live_input"
	ValueListParamsTypeStreamVideo              ValueListParamsType = "stream_video"
	ValueListParamsTypeWorker                   ValueListParamsType = "worker"
	ValueListParamsTypeWorkerVersion            ValueListParamsType = "worker_version"
	ValueListParamsTypeZone                     ValueListParamsType = "zone"
)

func (r ValueListParamsType) IsKnown() bool {
	switch r {
	case ValueListParamsTypeAccessApplication, ValueListParamsTypeAccessApplicationPolicy, ValueListParamsTypeAccessGroup, ValueListParamsTypeAccount, ValueListParamsTypeAIGateway, ValueListParamsTypeAlertingPolicy, ValueListParamsTypeAlertingWebhook, ValueListParamsTypeAPIGatewayOperation, ValueListParamsTypeCloudflaredTunnel, ValueListParamsTypeCustomCertificate, ValueListParamsTypeCustomHostname, ValueListParamsTypeD1Database, ValueListParamsTypeDNSRecord, ValueListParamsTypeDurableObjectNamespace, ValueListParamsTypeGatewayList, ValueListParamsTypeGatewayRule, ValueListParamsTypeImage, ValueListParamsTypeKVNamespace, ValueListParamsTypeManagedClientCertificate, ValueListParamsTypeQueue, ValueListParamsTypeR2Bucket, ValueListParamsTypeResourceShare, ValueListParamsTypeStreamLiveInput, ValueListParamsTypeStreamVideo, ValueListParamsTypeWorker, ValueListParamsTypeWorkerVersion, ValueListParamsTypeZone:
		return true
	}
	return false
}
