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
	"github.com/cloudflare/cloudflare-go/v6/packages/pagination"
)

// DEXRuleService contains methods and other services that help with interacting
// with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDEXRuleService] method instead.
type DEXRuleService struct {
	Options []option.RequestOption
}

// NewDEXRuleService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDEXRuleService(opts ...option.RequestOption) (r *DEXRuleService) {
	r = &DEXRuleService{}
	r.Options = opts
	return
}

// Create a DEX Rule
func (r *DEXRuleService) New(ctx context.Context, params DEXRuleNewParams, opts ...option.RequestOption) (res *DEXRuleNewResponse, err error) {
	var env DEXRuleNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/rules", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Update a DEX Rule
func (r *DEXRuleService) Update(ctx context.Context, ruleID string, params DEXRuleUpdateParams, opts ...option.RequestOption) (res *DEXRuleUpdateResponse, err error) {
	var env DEXRuleUpdateResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/rules/%s", params.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// List DEX Rules
func (r *DEXRuleService) List(ctx context.Context, params DEXRuleListParams, opts ...option.RequestOption) (res *pagination.V4PagePagination[DEXRuleListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/rules", params.AccountID)
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

// List DEX Rules
func (r *DEXRuleService) ListAutoPaging(ctx context.Context, params DEXRuleListParams, opts ...option.RequestOption) *pagination.V4PagePaginationAutoPager[DEXRuleListResponse] {
	return pagination.NewV4PagePaginationAutoPager(r.List(ctx, params, opts...))
}

// Delete a DEX Rule
func (r *DEXRuleService) Delete(ctx context.Context, ruleID string, body DEXRuleDeleteParams, opts ...option.RequestOption) (res *bool, err error) {
	var env DEXRuleDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/rules/%s", body.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Get details for a DEX Rule
func (r *DEXRuleService) Get(ctx context.Context, ruleID string, query DEXRuleGetParams, opts ...option.RequestOption) (res *DEXRuleGetResponse, err error) {
	var env DEXRuleGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if ruleID == "" {
		err = errors.New("missing required rule_id parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/dex/rules/%s", query.AccountID, ruleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type DEXRuleNewResponse struct {
	// API Resource UUID tag.
	ID            string                           `json:"id,required"`
	CreatedAt     string                           `json:"created_at,required"`
	Match         string                           `json:"match,required"`
	Name          string                           `json:"name,required"`
	Description   string                           `json:"description"`
	TargetedTests []DEXRuleNewResponseTargetedTest `json:"targeted_tests"`
	UpdatedAt     string                           `json:"updated_at"`
	JSON          dexRuleNewResponseJSON           `json:"-"`
}

// dexRuleNewResponseJSON contains the JSON metadata for the struct
// [DEXRuleNewResponse]
type dexRuleNewResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Match         apijson.Field
	Name          apijson.Field
	Description   apijson.Field
	TargetedTests apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXRuleNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseJSON) RawJSON() string {
	return r.raw
}

type DEXRuleNewResponseTargetedTest struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data    DEXRuleNewResponseTargetedTestsData `json:"data,required"`
	Enabled bool                                `json:"enabled,required"`
	Name    string                              `json:"name,required"`
	TestID  string                              `json:"test_id,required"`
	JSON    dexRuleNewResponseTargetedTestJSON  `json:"-"`
}

// dexRuleNewResponseTargetedTestJSON contains the JSON metadata for the struct
// [DEXRuleNewResponseTargetedTest]
type dexRuleNewResponseTargetedTestJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	TestID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleNewResponseTargetedTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseTargetedTestJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DEXRuleNewResponseTargetedTestsData struct {
	// The desired endpoint to test.
	Host string `json:"host,required"`
	// The type of test.
	Kind DEXRuleNewResponseTargetedTestsDataKind `json:"kind,required"`
	// The HTTP request method type.
	Method DEXRuleNewResponseTargetedTestsDataMethod `json:"method"`
	JSON   dexRuleNewResponseTargetedTestsDataJSON   `json:"-"`
}

// dexRuleNewResponseTargetedTestsDataJSON contains the JSON metadata for the
// struct [DEXRuleNewResponseTargetedTestsData]
type dexRuleNewResponseTargetedTestsDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleNewResponseTargetedTestsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseTargetedTestsDataJSON) RawJSON() string {
	return r.raw
}

// The type of test.
type DEXRuleNewResponseTargetedTestsDataKind string

const (
	DEXRuleNewResponseTargetedTestsDataKindHTTP       DEXRuleNewResponseTargetedTestsDataKind = "http"
	DEXRuleNewResponseTargetedTestsDataKindTraceroute DEXRuleNewResponseTargetedTestsDataKind = "traceroute"
)

func (r DEXRuleNewResponseTargetedTestsDataKind) IsKnown() bool {
	switch r {
	case DEXRuleNewResponseTargetedTestsDataKindHTTP, DEXRuleNewResponseTargetedTestsDataKindTraceroute:
		return true
	}
	return false
}

// The HTTP request method type.
type DEXRuleNewResponseTargetedTestsDataMethod string

const (
	DEXRuleNewResponseTargetedTestsDataMethodGet DEXRuleNewResponseTargetedTestsDataMethod = "GET"
)

func (r DEXRuleNewResponseTargetedTestsDataMethod) IsKnown() bool {
	switch r {
	case DEXRuleNewResponseTargetedTestsDataMethodGet:
		return true
	}
	return false
}

type DEXRuleUpdateResponse struct {
	// API Resource UUID tag.
	ID            string                              `json:"id,required"`
	CreatedAt     string                              `json:"created_at,required"`
	Match         string                              `json:"match,required"`
	Name          string                              `json:"name,required"`
	Description   string                              `json:"description"`
	TargetedTests []DEXRuleUpdateResponseTargetedTest `json:"targeted_tests"`
	UpdatedAt     string                              `json:"updated_at"`
	JSON          dexRuleUpdateResponseJSON           `json:"-"`
}

// dexRuleUpdateResponseJSON contains the JSON metadata for the struct
// [DEXRuleUpdateResponse]
type dexRuleUpdateResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Match         apijson.Field
	Name          apijson.Field
	Description   apijson.Field
	TargetedTests apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXRuleUpdateResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseJSON) RawJSON() string {
	return r.raw
}

type DEXRuleUpdateResponseTargetedTest struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data    DEXRuleUpdateResponseTargetedTestsData `json:"data,required"`
	Enabled bool                                   `json:"enabled,required"`
	Name    string                                 `json:"name,required"`
	TestID  string                                 `json:"test_id,required"`
	JSON    dexRuleUpdateResponseTargetedTestJSON  `json:"-"`
}

// dexRuleUpdateResponseTargetedTestJSON contains the JSON metadata for the struct
// [DEXRuleUpdateResponseTargetedTest]
type dexRuleUpdateResponseTargetedTestJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	TestID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleUpdateResponseTargetedTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseTargetedTestJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DEXRuleUpdateResponseTargetedTestsData struct {
	// The desired endpoint to test.
	Host string `json:"host,required"`
	// The type of test.
	Kind DEXRuleUpdateResponseTargetedTestsDataKind `json:"kind,required"`
	// The HTTP request method type.
	Method DEXRuleUpdateResponseTargetedTestsDataMethod `json:"method"`
	JSON   dexRuleUpdateResponseTargetedTestsDataJSON   `json:"-"`
}

// dexRuleUpdateResponseTargetedTestsDataJSON contains the JSON metadata for the
// struct [DEXRuleUpdateResponseTargetedTestsData]
type dexRuleUpdateResponseTargetedTestsDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleUpdateResponseTargetedTestsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseTargetedTestsDataJSON) RawJSON() string {
	return r.raw
}

// The type of test.
type DEXRuleUpdateResponseTargetedTestsDataKind string

const (
	DEXRuleUpdateResponseTargetedTestsDataKindHTTP       DEXRuleUpdateResponseTargetedTestsDataKind = "http"
	DEXRuleUpdateResponseTargetedTestsDataKindTraceroute DEXRuleUpdateResponseTargetedTestsDataKind = "traceroute"
)

func (r DEXRuleUpdateResponseTargetedTestsDataKind) IsKnown() bool {
	switch r {
	case DEXRuleUpdateResponseTargetedTestsDataKindHTTP, DEXRuleUpdateResponseTargetedTestsDataKindTraceroute:
		return true
	}
	return false
}

// The HTTP request method type.
type DEXRuleUpdateResponseTargetedTestsDataMethod string

const (
	DEXRuleUpdateResponseTargetedTestsDataMethodGet DEXRuleUpdateResponseTargetedTestsDataMethod = "GET"
)

func (r DEXRuleUpdateResponseTargetedTestsDataMethod) IsKnown() bool {
	switch r {
	case DEXRuleUpdateResponseTargetedTestsDataMethodGet:
		return true
	}
	return false
}

type DEXRuleListResponse struct {
	Rules []DEXRuleListResponseRule `json:"rules"`
	JSON  dexRuleListResponseJSON   `json:"-"`
}

// dexRuleListResponseJSON contains the JSON metadata for the struct
// [DEXRuleListResponse]
type dexRuleListResponseJSON struct {
	Rules       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleListResponseJSON) RawJSON() string {
	return r.raw
}

type DEXRuleListResponseRule struct {
	// API Resource UUID tag.
	ID            string                                 `json:"id,required"`
	CreatedAt     string                                 `json:"created_at,required"`
	Match         string                                 `json:"match,required"`
	Name          string                                 `json:"name,required"`
	Description   string                                 `json:"description"`
	TargetedTests []DEXRuleListResponseRulesTargetedTest `json:"targeted_tests"`
	UpdatedAt     string                                 `json:"updated_at"`
	JSON          dexRuleListResponseRuleJSON            `json:"-"`
}

// dexRuleListResponseRuleJSON contains the JSON metadata for the struct
// [DEXRuleListResponseRule]
type dexRuleListResponseRuleJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Match         apijson.Field
	Name          apijson.Field
	Description   apijson.Field
	TargetedTests apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXRuleListResponseRule) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleListResponseRuleJSON) RawJSON() string {
	return r.raw
}

type DEXRuleListResponseRulesTargetedTest struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data    DEXRuleListResponseRulesTargetedTestsData `json:"data,required"`
	Enabled bool                                      `json:"enabled,required"`
	Name    string                                    `json:"name,required"`
	TestID  string                                    `json:"test_id,required"`
	JSON    dexRuleListResponseRulesTargetedTestJSON  `json:"-"`
}

// dexRuleListResponseRulesTargetedTestJSON contains the JSON metadata for the
// struct [DEXRuleListResponseRulesTargetedTest]
type dexRuleListResponseRulesTargetedTestJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	TestID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleListResponseRulesTargetedTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleListResponseRulesTargetedTestJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DEXRuleListResponseRulesTargetedTestsData struct {
	// The desired endpoint to test.
	Host string `json:"host,required"`
	// The type of test.
	Kind DEXRuleListResponseRulesTargetedTestsDataKind `json:"kind,required"`
	// The HTTP request method type.
	Method DEXRuleListResponseRulesTargetedTestsDataMethod `json:"method"`
	JSON   dexRuleListResponseRulesTargetedTestsDataJSON   `json:"-"`
}

// dexRuleListResponseRulesTargetedTestsDataJSON contains the JSON metadata for the
// struct [DEXRuleListResponseRulesTargetedTestsData]
type dexRuleListResponseRulesTargetedTestsDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleListResponseRulesTargetedTestsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleListResponseRulesTargetedTestsDataJSON) RawJSON() string {
	return r.raw
}

// The type of test.
type DEXRuleListResponseRulesTargetedTestsDataKind string

const (
	DEXRuleListResponseRulesTargetedTestsDataKindHTTP       DEXRuleListResponseRulesTargetedTestsDataKind = "http"
	DEXRuleListResponseRulesTargetedTestsDataKindTraceroute DEXRuleListResponseRulesTargetedTestsDataKind = "traceroute"
)

func (r DEXRuleListResponseRulesTargetedTestsDataKind) IsKnown() bool {
	switch r {
	case DEXRuleListResponseRulesTargetedTestsDataKindHTTP, DEXRuleListResponseRulesTargetedTestsDataKindTraceroute:
		return true
	}
	return false
}

// The HTTP request method type.
type DEXRuleListResponseRulesTargetedTestsDataMethod string

const (
	DEXRuleListResponseRulesTargetedTestsDataMethodGet DEXRuleListResponseRulesTargetedTestsDataMethod = "GET"
)

func (r DEXRuleListResponseRulesTargetedTestsDataMethod) IsKnown() bool {
	switch r {
	case DEXRuleListResponseRulesTargetedTestsDataMethodGet:
		return true
	}
	return false
}

type DEXRuleGetResponse struct {
	// API Resource UUID tag.
	ID            string                           `json:"id,required"`
	CreatedAt     string                           `json:"created_at,required"`
	Match         string                           `json:"match,required"`
	Name          string                           `json:"name,required"`
	Description   string                           `json:"description"`
	TargetedTests []DEXRuleGetResponseTargetedTest `json:"targeted_tests"`
	UpdatedAt     string                           `json:"updated_at"`
	JSON          dexRuleGetResponseJSON           `json:"-"`
}

// dexRuleGetResponseJSON contains the JSON metadata for the struct
// [DEXRuleGetResponse]
type dexRuleGetResponseJSON struct {
	ID            apijson.Field
	CreatedAt     apijson.Field
	Match         apijson.Field
	Name          apijson.Field
	Description   apijson.Field
	TargetedTests apijson.Field
	UpdatedAt     apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *DEXRuleGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseJSON) RawJSON() string {
	return r.raw
}

type DEXRuleGetResponseTargetedTest struct {
	// The configuration object which contains the details for the WARP client to
	// conduct the test.
	Data    DEXRuleGetResponseTargetedTestsData `json:"data,required"`
	Enabled bool                                `json:"enabled,required"`
	Name    string                              `json:"name,required"`
	TestID  string                              `json:"test_id,required"`
	JSON    dexRuleGetResponseTargetedTestJSON  `json:"-"`
}

// dexRuleGetResponseTargetedTestJSON contains the JSON metadata for the struct
// [DEXRuleGetResponseTargetedTest]
type dexRuleGetResponseTargetedTestJSON struct {
	Data        apijson.Field
	Enabled     apijson.Field
	Name        apijson.Field
	TestID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleGetResponseTargetedTest) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseTargetedTestJSON) RawJSON() string {
	return r.raw
}

// The configuration object which contains the details for the WARP client to
// conduct the test.
type DEXRuleGetResponseTargetedTestsData struct {
	// The desired endpoint to test.
	Host string `json:"host,required"`
	// The type of test.
	Kind DEXRuleGetResponseTargetedTestsDataKind `json:"kind,required"`
	// The HTTP request method type.
	Method DEXRuleGetResponseTargetedTestsDataMethod `json:"method"`
	JSON   dexRuleGetResponseTargetedTestsDataJSON   `json:"-"`
}

// dexRuleGetResponseTargetedTestsDataJSON contains the JSON metadata for the
// struct [DEXRuleGetResponseTargetedTestsData]
type dexRuleGetResponseTargetedTestsDataJSON struct {
	Host        apijson.Field
	Kind        apijson.Field
	Method      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleGetResponseTargetedTestsData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseTargetedTestsDataJSON) RawJSON() string {
	return r.raw
}

// The type of test.
type DEXRuleGetResponseTargetedTestsDataKind string

const (
	DEXRuleGetResponseTargetedTestsDataKindHTTP       DEXRuleGetResponseTargetedTestsDataKind = "http"
	DEXRuleGetResponseTargetedTestsDataKindTraceroute DEXRuleGetResponseTargetedTestsDataKind = "traceroute"
)

func (r DEXRuleGetResponseTargetedTestsDataKind) IsKnown() bool {
	switch r {
	case DEXRuleGetResponseTargetedTestsDataKindHTTP, DEXRuleGetResponseTargetedTestsDataKindTraceroute:
		return true
	}
	return false
}

// The HTTP request method type.
type DEXRuleGetResponseTargetedTestsDataMethod string

const (
	DEXRuleGetResponseTargetedTestsDataMethodGet DEXRuleGetResponseTargetedTestsDataMethod = "GET"
)

func (r DEXRuleGetResponseTargetedTestsDataMethod) IsKnown() bool {
	switch r {
	case DEXRuleGetResponseTargetedTestsDataMethodGet:
		return true
	}
	return false
}

type DEXRuleNewParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// The wirefilter expression to match.
	Match param.Field[string] `json:"match,required"`
	// The name of the Rule.
	Name        param.Field[string] `json:"name,required"`
	Description param.Field[string] `json:"description"`
}

func (r DEXRuleNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DEXRuleNewResponseEnvelope struct {
	Errors   []DEXRuleNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXRuleNewResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DEXRuleNewResponseEnvelopeSuccess `json:"success,required"`
	Result  DEXRuleNewResponse                `json:"result"`
	JSON    dexRuleNewResponseEnvelopeJSON    `json:"-"`
}

// dexRuleNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXRuleNewResponseEnvelope]
type dexRuleNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXRuleNewResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           DEXRuleNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexRuleNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexRuleNewResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DEXRuleNewResponseEnvelopeErrors]
type dexRuleNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXRuleNewResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    dexRuleNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexRuleNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DEXRuleNewResponseEnvelopeErrorsSource]
type dexRuleNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXRuleNewResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           DEXRuleNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexRuleNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexRuleNewResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DEXRuleNewResponseEnvelopeMessages]
type dexRuleNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXRuleNewResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    dexRuleNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexRuleNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DEXRuleNewResponseEnvelopeMessagesSource]
type dexRuleNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXRuleNewResponseEnvelopeSuccess bool

const (
	DEXRuleNewResponseEnvelopeSuccessTrue DEXRuleNewResponseEnvelopeSuccess = true
)

func (r DEXRuleNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXRuleNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXRuleUpdateParams struct {
	AccountID   param.Field[string] `path:"account_id,required"`
	Description param.Field[string] `json:"description"`
	// The wirefilter expression to match.
	Match param.Field[string] `json:"match"`
	// The name of the Rule.
	Name param.Field[string] `json:"name"`
}

func (r DEXRuleUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type DEXRuleUpdateResponseEnvelope struct {
	Errors   []DEXRuleUpdateResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXRuleUpdateResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DEXRuleUpdateResponseEnvelopeSuccess `json:"success,required"`
	Result  DEXRuleUpdateResponse                `json:"result"`
	JSON    dexRuleUpdateResponseEnvelopeJSON    `json:"-"`
}

// dexRuleUpdateResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXRuleUpdateResponseEnvelope]
type dexRuleUpdateResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleUpdateResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXRuleUpdateResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DEXRuleUpdateResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexRuleUpdateResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexRuleUpdateResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DEXRuleUpdateResponseEnvelopeErrors]
type dexRuleUpdateResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleUpdateResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXRuleUpdateResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dexRuleUpdateResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexRuleUpdateResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DEXRuleUpdateResponseEnvelopeErrorsSource]
type dexRuleUpdateResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleUpdateResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXRuleUpdateResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DEXRuleUpdateResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexRuleUpdateResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexRuleUpdateResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DEXRuleUpdateResponseEnvelopeMessages]
type dexRuleUpdateResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleUpdateResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXRuleUpdateResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dexRuleUpdateResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexRuleUpdateResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DEXRuleUpdateResponseEnvelopeMessagesSource]
type dexRuleUpdateResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleUpdateResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleUpdateResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXRuleUpdateResponseEnvelopeSuccess bool

const (
	DEXRuleUpdateResponseEnvelopeSuccessTrue DEXRuleUpdateResponseEnvelopeSuccess = true
)

func (r DEXRuleUpdateResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXRuleUpdateResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXRuleListParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
	// Page number of paginated results
	Page param.Field[float64] `query:"page,required"`
	// Number of items per page
	PerPage param.Field[float64] `query:"per_page,required"`
	// Filter results by rule name
	Name param.Field[string] `query:"name"`
	// Which property to sort results by
	SortBy param.Field[DEXRuleListParamsSortBy] `query:"sort_by"`
	// Sort direction for sort_by property
	SortOrder param.Field[DEXRuleListParamsSortOrder] `query:"sort_order"`
}

// URLQuery serializes [DEXRuleListParams]'s query parameters as `url.Values`.
func (r DEXRuleListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Which property to sort results by
type DEXRuleListParamsSortBy string

const (
	DEXRuleListParamsSortByName      DEXRuleListParamsSortBy = "name"
	DEXRuleListParamsSortByCreatedAt DEXRuleListParamsSortBy = "created_at"
	DEXRuleListParamsSortByUpdatedAt DEXRuleListParamsSortBy = "updated_at"
)

func (r DEXRuleListParamsSortBy) IsKnown() bool {
	switch r {
	case DEXRuleListParamsSortByName, DEXRuleListParamsSortByCreatedAt, DEXRuleListParamsSortByUpdatedAt:
		return true
	}
	return false
}

// Sort direction for sort_by property
type DEXRuleListParamsSortOrder string

const (
	DEXRuleListParamsSortOrderAsc  DEXRuleListParamsSortOrder = "ASC"
	DEXRuleListParamsSortOrderDesc DEXRuleListParamsSortOrder = "DESC"
)

func (r DEXRuleListParamsSortOrder) IsKnown() bool {
	switch r {
	case DEXRuleListParamsSortOrderAsc, DEXRuleListParamsSortOrderDesc:
		return true
	}
	return false
}

type DEXRuleDeleteParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DEXRuleDeleteResponseEnvelope struct {
	Errors   []DEXRuleDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXRuleDeleteResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DEXRuleDeleteResponseEnvelopeSuccess `json:"success,required"`
	Result  bool                                 `json:"result,nullable"`
	JSON    dexRuleDeleteResponseEnvelopeJSON    `json:"-"`
}

// dexRuleDeleteResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXRuleDeleteResponseEnvelope]
type dexRuleDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXRuleDeleteResponseEnvelopeErrors struct {
	Code             int64                                     `json:"code,required"`
	Message          string                                    `json:"message,required"`
	DocumentationURL string                                    `json:"documentation_url"`
	Source           DEXRuleDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexRuleDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexRuleDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [DEXRuleDeleteResponseEnvelopeErrors]
type dexRuleDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXRuleDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                        `json:"pointer"`
	JSON    dexRuleDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexRuleDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DEXRuleDeleteResponseEnvelopeErrorsSource]
type dexRuleDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXRuleDeleteResponseEnvelopeMessages struct {
	Code             int64                                       `json:"code,required"`
	Message          string                                      `json:"message,required"`
	DocumentationURL string                                      `json:"documentation_url"`
	Source           DEXRuleDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexRuleDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexRuleDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [DEXRuleDeleteResponseEnvelopeMessages]
type dexRuleDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXRuleDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                          `json:"pointer"`
	JSON    dexRuleDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexRuleDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata for
// the struct [DEXRuleDeleteResponseEnvelopeMessagesSource]
type dexRuleDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXRuleDeleteResponseEnvelopeSuccess bool

const (
	DEXRuleDeleteResponseEnvelopeSuccessTrue DEXRuleDeleteResponseEnvelopeSuccess = true
)

func (r DEXRuleDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXRuleDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type DEXRuleGetParams struct {
	AccountID param.Field[string] `path:"account_id,required"`
}

type DEXRuleGetResponseEnvelope struct {
	Errors   []DEXRuleGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []DEXRuleGetResponseEnvelopeMessages `json:"messages,required"`
	// Whether the API call was successful.
	Success DEXRuleGetResponseEnvelopeSuccess `json:"success,required"`
	Result  DEXRuleGetResponse                `json:"result"`
	JSON    dexRuleGetResponseEnvelopeJSON    `json:"-"`
}

// dexRuleGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [DEXRuleGetResponseEnvelope]
type dexRuleGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Success     apijson.Field
	Result      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type DEXRuleGetResponseEnvelopeErrors struct {
	Code             int64                                  `json:"code,required"`
	Message          string                                 `json:"message,required"`
	DocumentationURL string                                 `json:"documentation_url"`
	Source           DEXRuleGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             dexRuleGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// dexRuleGetResponseEnvelopeErrorsJSON contains the JSON metadata for the struct
// [DEXRuleGetResponseEnvelopeErrors]
type dexRuleGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type DEXRuleGetResponseEnvelopeErrorsSource struct {
	Pointer string                                     `json:"pointer"`
	JSON    dexRuleGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// dexRuleGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for the
// struct [DEXRuleGetResponseEnvelopeErrorsSource]
type dexRuleGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type DEXRuleGetResponseEnvelopeMessages struct {
	Code             int64                                    `json:"code,required"`
	Message          string                                   `json:"message,required"`
	DocumentationURL string                                   `json:"documentation_url"`
	Source           DEXRuleGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             dexRuleGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// dexRuleGetResponseEnvelopeMessagesJSON contains the JSON metadata for the struct
// [DEXRuleGetResponseEnvelopeMessages]
type dexRuleGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *DEXRuleGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type DEXRuleGetResponseEnvelopeMessagesSource struct {
	Pointer string                                       `json:"pointer"`
	JSON    dexRuleGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// dexRuleGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata for the
// struct [DEXRuleGetResponseEnvelopeMessagesSource]
type dexRuleGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *DEXRuleGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r dexRuleGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type DEXRuleGetResponseEnvelopeSuccess bool

const (
	DEXRuleGetResponseEnvelopeSuccessTrue DEXRuleGetResponseEnvelopeSuccess = true
)

func (r DEXRuleGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case DEXRuleGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
