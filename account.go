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
	"github.com/cloudflare/cloudflare-sdk-go/internal/shared"
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
	Dlp              *AccountDlpService
	DNSFirewalls     *AccountDNSFirewallService
	Emails           *AccountEmailService
	Firewalls        *AccountFirewallService
	Images           *AccountImageService
	Intel            *AccountIntelService
	LoadBalancers    *AccountLoadBalancerService
	Logpushes        *AccountLogpushService
	Logs             *AccountLogService
	Magic            *AccountMagicService
	Members          *AccountMemberService
	Mnms             *AccountMnmService
	MtlsCertificates *AccountMtlsCertificateService
	Pages            *AccountPageService
	Pcaps            *AccountPcapService
	R2               *AccountR2Service
	Railguns         *AccountRailgunService
	Registrar        *AccountRegistrarService
	RequestTracers   *AccountRequestTracerService
	Roles            *AccountRoleService
	Rules            *AccountRuleService
	SecondaryDNS     *AccountSecondaryDNSService
	Storages         *AccountStorageService
	Streams          *AccountStreamService
	Subscriptions    *AccountSubscriptionService
	Teamnet          *AccountTeamnetService
	Tunnels          *AccountTunnelService
	Workers          *AccountWorkerService
	Gateway          *AccountGatewayService
	Rulesets         *AccountRulesetService
	Access           *AccountAccessService
	Alerting         *AccountAlertingService
	CustomNs         *AccountCustomNService
	Devices          *AccountDeviceService
	Addressing       *AccountAddressingService
	AI               *AccountAIService
	Challenges       *AccountChallengeService
	D1               *AccountD1Service
	Dex              *AccountDexService
	Hyperdrive       *AccountHyperdriveService
	PagesProjects    *AccountPagesProjectService
	Rum              *AccountRumService
	VectorizeIndexes *AccountVectorizeIndexService
	WarpConnector    *AccountWarpConnectorService
	Zerotrust        *AccountZerotrustService
	Urlscanner       *AccountUrlscannerService
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
	r.Dlp = NewAccountDlpService(opts...)
	r.DNSFirewalls = NewAccountDNSFirewallService(opts...)
	r.Emails = NewAccountEmailService(opts...)
	r.Firewalls = NewAccountFirewallService(opts...)
	r.Images = NewAccountImageService(opts...)
	r.Intel = NewAccountIntelService(opts...)
	r.LoadBalancers = NewAccountLoadBalancerService(opts...)
	r.Logpushes = NewAccountLogpushService(opts...)
	r.Logs = NewAccountLogService(opts...)
	r.Magic = NewAccountMagicService(opts...)
	r.Members = NewAccountMemberService(opts...)
	r.Mnms = NewAccountMnmService(opts...)
	r.MtlsCertificates = NewAccountMtlsCertificateService(opts...)
	r.Pages = NewAccountPageService(opts...)
	r.Pcaps = NewAccountPcapService(opts...)
	r.R2 = NewAccountR2Service(opts...)
	r.Railguns = NewAccountRailgunService(opts...)
	r.Registrar = NewAccountRegistrarService(opts...)
	r.RequestTracers = NewAccountRequestTracerService(opts...)
	r.Roles = NewAccountRoleService(opts...)
	r.Rules = NewAccountRuleService(opts...)
	r.SecondaryDNS = NewAccountSecondaryDNSService(opts...)
	r.Storages = NewAccountStorageService(opts...)
	r.Streams = NewAccountStreamService(opts...)
	r.Subscriptions = NewAccountSubscriptionService(opts...)
	r.Teamnet = NewAccountTeamnetService(opts...)
	r.Tunnels = NewAccountTunnelService(opts...)
	r.Workers = NewAccountWorkerService(opts...)
	r.Gateway = NewAccountGatewayService(opts...)
	r.Rulesets = NewAccountRulesetService(opts...)
	r.Access = NewAccountAccessService(opts...)
	r.Alerting = NewAccountAlertingService(opts...)
	r.CustomNs = NewAccountCustomNService(opts...)
	r.Devices = NewAccountDeviceService(opts...)
	r.Addressing = NewAccountAddressingService(opts...)
	r.AI = NewAccountAIService(opts...)
	r.Challenges = NewAccountChallengeService(opts...)
	r.D1 = NewAccountD1Service(opts...)
	r.Dex = NewAccountDexService(opts...)
	r.Hyperdrive = NewAccountHyperdriveService(opts...)
	r.PagesProjects = NewAccountPagesProjectService(opts...)
	r.Rum = NewAccountRumService(opts...)
	r.VectorizeIndexes = NewAccountVectorizeIndexService(opts...)
	r.WarpConnector = NewAccountWarpConnectorService(opts...)
	r.Zerotrust = NewAccountZerotrustService(opts...)
	r.Urlscanner = NewAccountUrlscannerService(opts...)
	return
}

// Get information about a specific account that you are a member of.
func (r *AccountService) Get(ctx context.Context, identifier interface{}, opts ...option.RequestOption) (res *AccountGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update an existing account.
func (r *AccountService) Update(ctx context.Context, identifier interface{}, body AccountUpdateParams, opts ...option.RequestOption) (res *AccountUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("accounts/%v", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all accounts you have ownership or verified access to.
func (r *AccountService) AccountsListAccounts(ctx context.Context, query AccountAccountsListAccountsParams, opts ...option.RequestOption) (res *shared.Page[AccountAccountsListAccountsResponse], err error) {
	var raw *http.Response
	opts = append(r.Options, opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "accounts"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
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

type AccountGetResponse struct {
	Errors   []AccountGetResponseError   `json:"errors"`
	Messages []AccountGetResponseMessage `json:"messages"`
	Result   interface{}                 `json:"result"`
	// Whether the API call was successful
	Success AccountGetResponseSuccess `json:"success"`
	JSON    accountGetResponseJSON    `json:"-"`
}

// accountGetResponseJSON contains the JSON metadata for the struct
// [AccountGetResponse]
type accountGetResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseError struct {
	Code    int64                       `json:"code,required"`
	Message string                      `json:"message,required"`
	JSON    accountGetResponseErrorJSON `json:"-"`
}

// accountGetResponseErrorJSON contains the JSON metadata for the struct
// [AccountGetResponseError]
type accountGetResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountGetResponseMessage struct {
	Code    int64                         `json:"code,required"`
	Message string                        `json:"message,required"`
	JSON    accountGetResponseMessageJSON `json:"-"`
}

// accountGetResponseMessageJSON contains the JSON metadata for the struct
// [AccountGetResponseMessage]
type accountGetResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountGetResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountGetResponseSuccess bool

const (
	AccountGetResponseSuccessTrue AccountGetResponseSuccess = true
)

type AccountUpdateResponse struct {
	Errors   []AccountUpdateResponseError   `json:"errors"`
	Messages []AccountUpdateResponseMessage `json:"messages"`
	Result   interface{}                    `json:"result"`
	// Whether the API call was successful
	Success AccountUpdateResponseSuccess `json:"success"`
	JSON    accountUpdateResponseJSON    `json:"-"`
}

// accountUpdateResponseJSON contains the JSON metadata for the struct
// [AccountUpdateResponse]
type accountUpdateResponseJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseError struct {
	Code    int64                          `json:"code,required"`
	Message string                         `json:"message,required"`
	JSON    accountUpdateResponseErrorJSON `json:"-"`
}

// accountUpdateResponseErrorJSON contains the JSON metadata for the struct
// [AccountUpdateResponseError]
type accountUpdateResponseErrorJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseError) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

type AccountUpdateResponseMessage struct {
	Code    int64                            `json:"code,required"`
	Message string                           `json:"message,required"`
	JSON    accountUpdateResponseMessageJSON `json:"-"`
}

// accountUpdateResponseMessageJSON contains the JSON metadata for the struct
// [AccountUpdateResponseMessage]
type accountUpdateResponseMessageJSON struct {
	Code        apijson.Field
	Message     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AccountUpdateResponseMessage) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

// Whether the API call was successful
type AccountUpdateResponseSuccess bool

const (
	AccountUpdateResponseSuccessTrue AccountUpdateResponseSuccess = true
)

type AccountAccountsListAccountsResponse = interface{}

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
	// Specifies the default nameservers to be used for new zones added to this
	// account.
	//
	// - `cloudflare.standard` for Cloudflare-branded nameservers
	// - `custom.account` for account custom nameservers
	// - `custom.tenant` for tenant custom nameservers
	//
	// See
	// [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)
	// for more information.
	DefaultNameservers param.Field[AccountUpdateParamsSettingsDefaultNameservers] `json:"default_nameservers"`
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

// Specifies the default nameservers to be used for new zones added to this
// account.
//
// - `cloudflare.standard` for Cloudflare-branded nameservers
// - `custom.account` for account custom nameservers
// - `custom.tenant` for tenant custom nameservers
//
// See
// [Custom Nameservers](https://developers.cloudflare.com/dns/additional-options/custom-nameservers/)
// for more information.
type AccountUpdateParamsSettingsDefaultNameservers string

const (
	AccountUpdateParamsSettingsDefaultNameserversCloudflareStandard AccountUpdateParamsSettingsDefaultNameservers = "cloudflare.standard"
	AccountUpdateParamsSettingsDefaultNameserversCustomAccount      AccountUpdateParamsSettingsDefaultNameservers = "custom.account"
	AccountUpdateParamsSettingsDefaultNameserversCustomTenant       AccountUpdateParamsSettingsDefaultNameservers = "custom.tenant"
)

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
