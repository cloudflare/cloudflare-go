// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudflare

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

// DNSRecordService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDNSRecordService] method instead.
type DNSRecordService struct {
	Options []option.RequestOption
}

// NewDNSRecordService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDNSRecordService(opts ...option.RequestOption) (r *DNSRecordService) {
	r = &DNSRecordService{}
	r.Options = opts
	return
}

// List, search, sort, and filter a zones' DNS records.
func (r *DNSRecordService) List(ctx context.Context, params DNSRecordListParams, opts ...option.RequestOption) (res *pagination.V4PagePaginationArray[DNSRecordListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ZoneID.Value == "" {
		err = errors.New("missing required zone_id parameter")
		return
	}
	path := fmt.Sprintf("zones/%s/dns_records", params.ZoneID)
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

// List, search, sort, and filter a zones' DNS records.
func (r *DNSRecordService) ListAutoPaging(ctx context.Context, params DNSRecordListParams, opts ...option.RequestOption) *pagination.V4PagePaginationArrayAutoPager[DNSRecordListResponse] {
	return pagination.NewV4PagePaginationArrayAutoPager(r.List(ctx, params, opts...))
}

type DNSRecordListResponse = interface{}

type DNSRecordListParams struct {
	// Identifier.
	ZoneID  param.Field[string]                     `path:"zone_id,required"`
	Comment param.Field[DNSRecordListParamsComment] `query:"comment"`
	Content param.Field[DNSRecordListParamsContent] `query:"content"`
	// Direction to order DNS records in.
	Direction param.Field[DNSRecordListParamsDirection] `query:"direction"`
	// Whether to match all search requirements or at least one (any). If set to `all`,
	// acts like a logical AND between filters. If set to `any`, acts like a logical OR
	// instead. Note that the interaction between tag filters is controlled by the
	// `tag-match` parameter instead.
	Match param.Field[DNSRecordListParamsMatch] `query:"match"`
	Name  param.Field[DNSRecordListParamsName]  `query:"name"`
	// Field to order DNS records by.
	Order param.Field[DNSRecordListParamsOrder] `query:"order"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Number of DNS records per page.
	PerPage param.Field[float64] `query:"per_page"`
	// Whether the record is receiving the performance and security benefits of
	// Cloudflare.
	Proxied param.Field[bool] `query:"proxied"`
	// Allows searching in multiple properties of a DNS record simultaneously. This
	// parameter is intended for human users, not automation. Its exact behavior is
	// intentionally left unspecified and is subject to change in the future. This
	// parameter works independently of the `match` setting. For automated searches,
	// please use the other available parameters.
	Search param.Field[string]                 `query:"search"`
	Tag    param.Field[DNSRecordListParamsTag] `query:"tag"`
	// Whether to match all tag search requirements or at least one (any). If set to
	// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
	// logical OR instead. Note that the regular `match` parameter is still used to
	// combine the resulting condition with other filters that aren't related to tags.
	TagMatch param.Field[DNSRecordListParamsTagMatch] `query:"tag_match"`
	// Record type.
	Type param.Field[DNSRecordListParamsType] `query:"type"`
}

// URLQuery serializes [DNSRecordListParams]'s query parameters as `url.Values`.
func (r DNSRecordListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DNSRecordListParamsComment struct {
	// If this parameter is present, only records _without_ a comment are returned.
	Absent param.Field[string] `query:"absent"`
	// Substring of the DNS record comment. Comment filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record comment. Comment filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record comment. Comment filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// If this parameter is present, only records _with_ a comment are returned.
	Present param.Field[string] `query:"present"`
	// Prefix of the DNS record comment. Comment filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [DNSRecordListParamsComment]'s query parameters as
// `url.Values`.
func (r DNSRecordListParamsComment) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type DNSRecordListParamsContent struct {
	// Substring of the DNS record content. Content filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record content. Content filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record content. Content filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS record content. Content filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [DNSRecordListParamsContent]'s query parameters as
// `url.Values`.
func (r DNSRecordListParamsContent) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Direction to order DNS records in.
type DNSRecordListParamsDirection string

const (
	DNSRecordListParamsDirectionAsc  DNSRecordListParamsDirection = "asc"
	DNSRecordListParamsDirectionDesc DNSRecordListParamsDirection = "desc"
)

func (r DNSRecordListParamsDirection) IsKnown() bool {
	switch r {
	case DNSRecordListParamsDirectionAsc, DNSRecordListParamsDirectionDesc:
		return true
	}
	return false
}

// Whether to match all search requirements or at least one (any). If set to `all`,
// acts like a logical AND between filters. If set to `any`, acts like a logical OR
// instead. Note that the interaction between tag filters is controlled by the
// `tag-match` parameter instead.
type DNSRecordListParamsMatch string

const (
	DNSRecordListParamsMatchAny DNSRecordListParamsMatch = "any"
	DNSRecordListParamsMatchAll DNSRecordListParamsMatch = "all"
)

func (r DNSRecordListParamsMatch) IsKnown() bool {
	switch r {
	case DNSRecordListParamsMatchAny, DNSRecordListParamsMatchAll:
		return true
	}
	return false
}

type DNSRecordListParamsName struct {
	// Substring of the DNS record name. Name filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// Suffix of the DNS record name. Name filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// Exact value of the DNS record name. Name filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Prefix of the DNS record name. Name filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [DNSRecordListParamsName]'s query parameters as
// `url.Values`.
func (r DNSRecordListParamsName) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Field to order DNS records by.
type DNSRecordListParamsOrder string

const (
	DNSRecordListParamsOrderType    DNSRecordListParamsOrder = "type"
	DNSRecordListParamsOrderName    DNSRecordListParamsOrder = "name"
	DNSRecordListParamsOrderContent DNSRecordListParamsOrder = "content"
	DNSRecordListParamsOrderTtl     DNSRecordListParamsOrder = "ttl"
	DNSRecordListParamsOrderProxied DNSRecordListParamsOrder = "proxied"
)

func (r DNSRecordListParamsOrder) IsKnown() bool {
	switch r {
	case DNSRecordListParamsOrderType, DNSRecordListParamsOrderName, DNSRecordListParamsOrderContent, DNSRecordListParamsOrderTtl, DNSRecordListParamsOrderProxied:
		return true
	}
	return false
}

type DNSRecordListParamsTag struct {
	// Name of a tag which must _not_ be present on the DNS record. Tag filters are
	// case-insensitive.
	Absent param.Field[string] `query:"absent"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value contains
	// `<tag-value>`. Tag filters are case-insensitive.
	Contains param.Field[string] `query:"contains"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value ends with
	// `<tag-value>`. Tag filters are case-insensitive.
	Endswith param.Field[string] `query:"endswith"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value is `<tag-value>`. Tag
	// filters are case-insensitive.
	Exact param.Field[string] `query:"exact"`
	// Name of a tag which must be present on the DNS record. Tag filters are
	// case-insensitive.
	Present param.Field[string] `query:"present"`
	// A tag and value, of the form `<tag-name>:<tag-value>`. The API will only return
	// DNS records that have a tag named `<tag-name>` whose value starts with
	// `<tag-value>`. Tag filters are case-insensitive.
	Startswith param.Field[string] `query:"startswith"`
}

// URLQuery serializes [DNSRecordListParamsTag]'s query parameters as `url.Values`.
func (r DNSRecordListParamsTag) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Whether to match all tag search requirements or at least one (any). If set to
// `all`, acts like a logical AND between tag filters. If set to `any`, acts like a
// logical OR instead. Note that the regular `match` parameter is still used to
// combine the resulting condition with other filters that aren't related to tags.
type DNSRecordListParamsTagMatch string

const (
	DNSRecordListParamsTagMatchAny DNSRecordListParamsTagMatch = "any"
	DNSRecordListParamsTagMatchAll DNSRecordListParamsTagMatch = "all"
)

func (r DNSRecordListParamsTagMatch) IsKnown() bool {
	switch r {
	case DNSRecordListParamsTagMatchAny, DNSRecordListParamsTagMatchAll:
		return true
	}
	return false
}

// Record type.
type DNSRecordListParamsType string

const (
	DNSRecordListParamsTypeA          DNSRecordListParamsType = "A"
	DNSRecordListParamsTypeAaaa       DNSRecordListParamsType = "AAAA"
	DNSRecordListParamsTypeCaa        DNSRecordListParamsType = "CAA"
	DNSRecordListParamsTypeCert       DNSRecordListParamsType = "CERT"
	DNSRecordListParamsTypeCname      DNSRecordListParamsType = "CNAME"
	DNSRecordListParamsTypeDnskey     DNSRecordListParamsType = "DNSKEY"
	DNSRecordListParamsTypeDs         DNSRecordListParamsType = "DS"
	DNSRecordListParamsTypeHTTPS      DNSRecordListParamsType = "HTTPS"
	DNSRecordListParamsTypeLoc        DNSRecordListParamsType = "LOC"
	DNSRecordListParamsTypeMx         DNSRecordListParamsType = "MX"
	DNSRecordListParamsTypeNaptr      DNSRecordListParamsType = "NAPTR"
	DNSRecordListParamsTypeNs         DNSRecordListParamsType = "NS"
	DNSRecordListParamsTypeOpenpgpkey DNSRecordListParamsType = "OPENPGPKEY"
	DNSRecordListParamsTypePtr        DNSRecordListParamsType = "PTR"
	DNSRecordListParamsTypeSmimea     DNSRecordListParamsType = "SMIMEA"
	DNSRecordListParamsTypeSrv        DNSRecordListParamsType = "SRV"
	DNSRecordListParamsTypeSshfp      DNSRecordListParamsType = "SSHFP"
	DNSRecordListParamsTypeSvcb       DNSRecordListParamsType = "SVCB"
	DNSRecordListParamsTypeTlsa       DNSRecordListParamsType = "TLSA"
	DNSRecordListParamsTypeTxt        DNSRecordListParamsType = "TXT"
	DNSRecordListParamsTypeUri        DNSRecordListParamsType = "URI"
)

func (r DNSRecordListParamsType) IsKnown() bool {
	switch r {
	case DNSRecordListParamsTypeA, DNSRecordListParamsTypeAaaa, DNSRecordListParamsTypeCaa, DNSRecordListParamsTypeCert, DNSRecordListParamsTypeCname, DNSRecordListParamsTypeDnskey, DNSRecordListParamsTypeDs, DNSRecordListParamsTypeHTTPS, DNSRecordListParamsTypeLoc, DNSRecordListParamsTypeMx, DNSRecordListParamsTypeNaptr, DNSRecordListParamsTypeNs, DNSRecordListParamsTypeOpenpgpkey, DNSRecordListParamsTypePtr, DNSRecordListParamsTypeSmimea, DNSRecordListParamsTypeSrv, DNSRecordListParamsTypeSshfp, DNSRecordListParamsTypeSvcb, DNSRecordListParamsTypeTlsa, DNSRecordListParamsTypeTxt, DNSRecordListParamsTypeUri:
		return true
	}
	return false
}
