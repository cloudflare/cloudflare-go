// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package intel

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudflare/cloudflare-go/v2/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v2/internal/param"
	"github.com/cloudflare/cloudflare-go/v2/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v2/option"
	"github.com/cloudflare/cloudflare-go/v2/shared"
)

// AttackSurfaceReportIssueTypeService contains methods and other services that
// help with interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAttackSurfaceReportIssueTypeService] method instead.
type AttackSurfaceReportIssueTypeService struct {
	Options []option.RequestOption
}

// NewAttackSurfaceReportIssueTypeService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAttackSurfaceReportIssueTypeService(opts ...option.RequestOption) (r *AttackSurfaceReportIssueTypeService) {
	r = &AttackSurfaceReportIssueTypeService{}
	r.Options = opts
	return
}

// Get Security Center Issues Types
func (r *AttackSurfaceReportIssueTypeService) Get(ctx context.Context, query AttackSurfaceReportIssueTypeGetParams, opts ...option.RequestOption) (res *[]string, err error) {
	opts = append(r.Options[:], opts...)
	var env AttackSurfaceReportIssueTypeGetResponseEnvelope
	path := fmt.Sprintf("accounts/%s/intel/attack-surface-report/issue-types", query.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type AttackSurfaceReportIssueTypeGetParams struct {
	// Identifier
	AccountID param.Field[string] `path:"account_id,required"`
}

type AttackSurfaceReportIssueTypeGetResponseEnvelope struct {
	Errors   []shared.ResponseInfo `json:"errors,required"`
	Messages []shared.ResponseInfo `json:"messages,required"`
	Result   []string              `json:"result,required"`
	// Whether the API call was successful
	Success AttackSurfaceReportIssueTypeGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    attackSurfaceReportIssueTypeGetResponseEnvelopeJSON    `json:"-"`
}

// attackSurfaceReportIssueTypeGetResponseEnvelopeJSON contains the JSON metadata
// for the struct [AttackSurfaceReportIssueTypeGetResponseEnvelope]
type attackSurfaceReportIssueTypeGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AttackSurfaceReportIssueTypeGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r attackSurfaceReportIssueTypeGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful
type AttackSurfaceReportIssueTypeGetResponseEnvelopeSuccess bool

const (
	AttackSurfaceReportIssueTypeGetResponseEnvelopeSuccessTrue AttackSurfaceReportIssueTypeGetResponseEnvelopeSuccess = true
)

func (r AttackSurfaceReportIssueTypeGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case AttackSurfaceReportIssueTypeGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
