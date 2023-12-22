// File generated from our OpenAPI spec by Stainless.

package cloudflare

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudflare/cloudflare-sdk-go/internal/apijson"
	"github.com/cloudflare/cloudflare-sdk-go/internal/apiquery"
	"github.com/cloudflare/cloudflare-sdk-go/internal/param"
	"github.com/cloudflare/cloudflare-sdk-go/internal/requestconfig"
	"github.com/cloudflare/cloudflare-sdk-go/option"
)

// AccountService contains methods and other services that help with interacting
// with the cloudflare API. Note, unlike clients, this service does not read
// variables from the environment automatically. You should not instantiate this
// service directly, and instead use the [NewAccountService] method instead.
type AccountService struct {
	Options          []option.RequestOption
	Addresses        *AccountAddressService
	Billings         *AccountBillingService
	BrandProtections *AccountBrandProtectionService
	CfdTunnels       *AccountCfdTunnelService
	CustomPages      *AccountCustomPageService
	Diagnostics      *AccountDiagnosticService
	Dlps             *AccountDlpService
	DNSFirewalls     *AccountDNSFirewallService
	Emails           *AccountEmailService
	Firewalls        *AccountFirewallService
	Images           *AccountImageService
	Intels           *AccountIntelService
	LoadBalancers    *AccountLoadBalancerService
	Logpushes        *AccountLogpushService
	Logs             *AccountLogService
	Magics           *AccountMagicService
	Members          *AccountMemberService
	Mnms             *AccountMnmService
	MtlsCertificates *AccountMtlsCertificateService
	Pages            *AccountPageService
	Pcaps            *AccountPcapService
	R2s              *AccountR2Service
	Railguns         *AccountRailgunService
	Registrars       *AccountRegistrarService
	RequestTracers   *AccountRequestTracerService
	Roles            *AccountRoleService
	Rules            *AccountRuleService
	SecondaryDNS     *AccountSecondaryDNSService
	Storages         *AccountStorageService
	Streams          *AccountStreamService
	Subscriptions    *AccountSubscriptionService
	Teamnets         *AccountTeamnetService
	Tunnels          *AccountTunnelService
	Workers          *AccountWorkerService
	Gateways         *AccountGatewayService
	Rulesets         *AccountRulesetService
	Accesses         *AccountAccessService
	Alerting         *AccountAlertingService
	CustomNs         *AccountCustomNService
	Devices          *AccountDeviceService
}

// NewAccountService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAccountService(opts ...option.RequestOption) (r *AccountService) {
	r = &AccountService{}
	r.Options = opts
	r.Addresses = NewAccountAddressService(opts...)
	r.Billings = NewAccountBillingService(opts...)
	r.BrandProtections = NewAccountBrandProtectionService(opts...)
	r.CfdTunnels = NewAccountCfdTunnelService(opts...)
	r.CustomPages = NewAccountCustomPageService(opts...)
	r.Diagnostics = NewAccountDiagnosticService(opts...)
	r.Dlps = NewAccountDlpService(opts...)
	r.DNSFirewalls = NewAccountDNSFirewallService(opts...)
	r.Emails = NewAccountEmailService(opts...)
	r.Firewalls = NewAccountFirewallService(opts...)
	r.Images = NewAccountImageService(opts...)
	r.Intels = NewAccountIntelService(opts...)
	r.LoadBalancers = NewAccountLoadBalancerService(opts...)
	r.Logpushes = NewAccountLogpushService(opts...)
	r.Logs = NewAccountLogService(opts...)
	r.Magics = NewAccountMagicService(opts...)
	r.Members = NewAccountMemberService(opts...)
	r.Mnms = NewAccountMnmService(opts...)
	r.MtlsCertificates = NewAccountMtlsCertificateService(opts...)
	r.Pages = NewAccountPageService(opts...)
	r.Pcaps = NewAccountPcapService(opts...)
	r.R2s = NewAccountR2Service(opts...)
	r.Railguns = NewAccountRailgunService(opts...)
	r.Registrars = NewAccountRegistrarService(opts...)
	r.RequestTracers = NewAccountRequestTracerService(opts...)
	r.Roles = NewAccountRoleService(opts...)
	r.Rules = NewAccountRuleService(opts...)
	r.SecondaryDNS = NewAccountSecondaryDNSService(opts...)
	r.Storages = NewAccountStorageService(opts...)
	r.Streams = NewAccountStreamService(opts...)
	r.Subscriptions = NewAccountSubscriptionService(opts...)
	r.Teamnets = NewAccountTeamnetService(opts...)
	r.Tunnels = NewAccountTunnelService(opts...)
	r.Workers = NewAccountWorkerService(opts...)
	r.Gateways = NewAccountGatewayService(opts...)
	r.Rulesets = NewAccountRulesetService(opts...)
	r.Accesses = NewAccountAccessService(opts...)
	r.Alerting = NewAccountAlertingService(opts...)
	r.CustomNs = NewAccountCustomNService(opts...)
	r.Devices = NewAccountDeviceService(opts...)
	return
}

// Get information about a specific account that you are a member of.
func (r *AccountService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing account.
func (r *AccountService) Update(ctx context.Context, identifier interface{}, body AccountUpdateParams, opts ...option.RequestOption) (res *ResponseSingle, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all accounts you have ownership or verified access to.
func (r *AccountService) AccountsListAccounts(ctx context.Context, query AccountAccountsListAccountsParams, opts ...option.RequestOption) (res *ResponseCollectionYgt6DzoC, err error) {
	opts = append(r.Options[:], opts...)
	path := "accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type ResponseCollectionYgt6DzoC struct {
	Errors     []ResponseCollectionYgt6DzoCError    `json:"errors"`
	Messages   []ResponseCollectionYgt6DzoCMessage  `json:"messages"`
	Result     []interface{}                        `json:"result"`
	ResultInfo ResponseCollectionYgt6DzoCResultInfo `json:"result_info"`
	// Whether the API call was successful
	Success ResponseCollectionYgt6DzoCSuccess `json:"success"`
	JSON    responseCollectionYgt6DzoCJSON    `json:"-"`
}

// responseCollectionYgt6DzoCJSON contains the JSON metadata for the struct
// [ResponseCollectionYgt6DzoC]
type responseCollectionYgt6DzoCJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	ResultInfo  apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYgt6DzoC) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionYgt6DzoCError struct {
	Code    int64                               `json:"code,required"`
	Message string                              `json:"message,required"`
	JSON    responseCollectionYgt6DzoCErrorJSON `json:"-"`
}

// responseCollectionYgt6DzoCErrorJSON contains the JSON metadata for the struct
// [ResponseCollectionYgt6DzoCError]
type responseCollectionYgt6DzoCErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYgt6DzoCError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionYgt6DzoCMessage struct {
	Code    int64                                 `json:"code,required"`
	Message string                                `json:"message,required"`
	JSON    responseCollectionYgt6DzoCMessageJSON `json:"-"`
}

// responseCollectionYgt6DzoCMessageJSON contains the JSON metadata for the struct
// [ResponseCollectionYgt6DzoCMessage]
type responseCollectionYgt6DzoCMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYgt6DzoCMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseCollectionYgt6DzoCResultInfo struct {
	// Total number of results for the requested service
	Count float64 `json:"count"`
	// Current page within paginated list of results
	Page float64 `json:"page"`
	// Number of results per page of results
	PerPage float64 `json:"per_page"`
	// Total results available without any search parameters
	TotalCount float64                                  `json:"total_count"`
	JSON       responseCollectionYgt6DzoCResultInfoJSON `json:"-"`
}

// responseCollectionYgt6DzoCResultInfoJSON contains the JSON metadata for the
// struct [ResponseCollectionYgt6DzoCResultInfo]
type responseCollectionYgt6DzoCResultInfoJSON struct {
	Count       apijson.Field
	Page        apijson.Field
	PerPage     apijson.Field
	TotalCount  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseCollectionYgt6DzoCResultInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseCollectionYgt6DzoCSuccess bool

const (
	ResponseCollectionYgt6DzoCSuccessTrue ResponseCollectionYgt6DzoCSuccess = true
)

type ResponseSingle struct {
	Errors   []ResponseSingleError   `json:"errors"`
	Messages []ResponseSingleMessage `json:"messages"`
	Result   interface{}             `json:"result"`
	// Whether the API call was successful
	Success ResponseSingleSuccess `json:"success"`
	JSON    responseSingleJSON    `json:"-"`
}

// responseSingleJSON contains the JSON metadata for the struct [ResponseSingle]
type responseSingleJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseSingleError struct {
	Code    int64                   `json:"code,required"`
	Message string                  `json:"message,required"`
	JSON    responseSingleErrorJSON `json:"-"`
}

// responseSingleErrorJSON contains the JSON metadata for the struct
// [ResponseSingleError]
type responseSingleErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type ResponseSingleMessage struct {
	Code    int64                     `json:"code,required"`
	Message string                    `json:"message,required"`
	JSON    responseSingleMessageJSON `json:"-"`
}

// responseSingleMessageJSON contains the JSON metadata for the struct
// [ResponseSingleMessage]
type responseSingleMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ResponseSingleMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type ResponseSingleSuccess bool

const (
	ResponseSingleSuccessTrue ResponseSingleSuccess = true
)

type AccountUpdateParams struct {
	// Account name
	Name param.Field[string] `json:"name,required"`
	// Account settings
	Settings param.Field[AccountUpdateParamsSettings] `json:"settings"`
}

func (r AccountUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Account settings
type AccountUpdateParamsSettings struct {
	// Indicates whether membership in this account requires that Two-Factor
	// Authentication is enabled
	EnforceTwofactor param.Field[bool] `json:"enforce_twofactor"`
	// Indicates whether new zones should use the account-level custom nameservers by
	// default
	UseAccountCustomNsByDefault param.Field[bool] `json:"use_account_custom_ns_by_default"`
}

func (r AccountUpdateParamsSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AccountAccountsListAccountsParams struct {
	// Direction to order results.
	Direction param.Field[AccountAccountsListAccountsParamsDirection] `query:"direction"`
	// Page number of paginated results.
	Page param.Field[float64] `query:"page"`
	// Maximum number of results per page.
	PerPage param.Field[float64] `query:"per_page"`
}

// URLQuery serializes [AccountAccountsListAccountsParams]'s query parameters as
// `url.Values`.
func (r AccountAccountsListAccountsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction to order results.
type AccountAccountsListAccountsParamsDirection string

const (
	AccountAccountsListAccountsParamsDirectionAsc  AccountAccountsListAccountsParamsDirection = "asc"
	AccountAccountsListAccountsParamsDirectionDesc AccountAccountsListAccountsParamsDirection = "desc"
)
