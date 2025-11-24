// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package pages

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

// ProjectDomainService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProjectDomainService] method instead.
type ProjectDomainService struct {
	Options []option.RequestOption
}

// NewProjectDomainService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewProjectDomainService(opts ...option.RequestOption) (r *ProjectDomainService) {
	r = &ProjectDomainService{}
	r.Options = opts
	return
}

// Add a new domain for the Pages project.
func (r *ProjectDomainService) New(ctx context.Context, projectName string, params ProjectDomainNewParams, opts ...option.RequestOption) (res *ProjectDomainNewResponse, err error) {
	var env ProjectDomainNewResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", params.AccountID, projectName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) List(ctx context.Context, projectName string, query ProjectDomainListParams, opts ...option.RequestOption) (res *pagination.SinglePage[ProjectDomainListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains", query.AccountID, projectName)
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

// Fetch a list of all domains associated with a Pages project.
func (r *ProjectDomainService) ListAutoPaging(ctx context.Context, projectName string, query ProjectDomainListParams, opts ...option.RequestOption) *pagination.SinglePageAutoPager[ProjectDomainListResponse] {
	return pagination.NewSinglePageAutoPager(r.List(ctx, projectName, query, opts...))
}

// Delete a Pages project's domain.
func (r *ProjectDomainService) Delete(ctx context.Context, projectName string, domainName string, body ProjectDomainDeleteParams, opts ...option.RequestOption) (res *ProjectDomainDeleteResponse, err error) {
	var env ProjectDomainDeleteResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", body.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Retry the validation status of a single domain.
func (r *ProjectDomainService) Edit(ctx context.Context, projectName string, domainName string, body ProjectDomainEditParams, opts ...option.RequestOption) (res *ProjectDomainEditResponse, err error) {
	var env ProjectDomainEditResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", body.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

// Fetch a single domain.
func (r *ProjectDomainService) Get(ctx context.Context, projectName string, domainName string, query ProjectDomainGetParams, opts ...option.RequestOption) (res *ProjectDomainGetResponse, err error) {
	var env ProjectDomainGetResponseEnvelope
	opts = slices.Concat(r.Options, opts)
	if query.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return
	}
	if projectName == "" {
		err = errors.New("missing required project_name parameter")
		return
	}
	if domainName == "" {
		err = errors.New("missing required domain_name parameter")
		return
	}
	path := fmt.Sprintf("accounts/%s/pages/projects/%s/domains/%s", query.AccountID, projectName, domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &env, opts...)
	if err != nil {
		return
	}
	res = &env.Result
	return
}

type ProjectDomainNewResponse struct {
	ID                   string                                       `json:"id,required"`
	CertificateAuthority ProjectDomainNewResponseCertificateAuthority `json:"certificate_authority,required"`
	CreatedOn            string                                       `json:"created_on,required"`
	DomainID             string                                       `json:"domain_id,required"`
	// The domain name.
	Name             string                                   `json:"name,required"`
	Status           ProjectDomainNewResponseStatus           `json:"status,required"`
	ValidationData   ProjectDomainNewResponseValidationData   `json:"validation_data,required"`
	VerificationData ProjectDomainNewResponseVerificationData `json:"verification_data,required"`
	ZoneTag          string                                   `json:"zone_tag,required"`
	JSON             projectDomainNewResponseJSON             `json:"-"`
}

// projectDomainNewResponseJSON contains the JSON metadata for the struct
// [ProjectDomainNewResponse]
type projectDomainNewResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseCertificateAuthority string

const (
	ProjectDomainNewResponseCertificateAuthorityGoogle      ProjectDomainNewResponseCertificateAuthority = "google"
	ProjectDomainNewResponseCertificateAuthorityLetsEncrypt ProjectDomainNewResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainNewResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseCertificateAuthorityGoogle, ProjectDomainNewResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainNewResponseStatus string

const (
	ProjectDomainNewResponseStatusInitializing ProjectDomainNewResponseStatus = "initializing"
	ProjectDomainNewResponseStatusPending      ProjectDomainNewResponseStatus = "pending"
	ProjectDomainNewResponseStatusActive       ProjectDomainNewResponseStatus = "active"
	ProjectDomainNewResponseStatusDeactivated  ProjectDomainNewResponseStatus = "deactivated"
	ProjectDomainNewResponseStatusBlocked      ProjectDomainNewResponseStatus = "blocked"
	ProjectDomainNewResponseStatusError        ProjectDomainNewResponseStatus = "error"
)

func (r ProjectDomainNewResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseStatusInitializing, ProjectDomainNewResponseStatusPending, ProjectDomainNewResponseStatusActive, ProjectDomainNewResponseStatusDeactivated, ProjectDomainNewResponseStatusBlocked, ProjectDomainNewResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainNewResponseValidationData struct {
	Method       ProjectDomainNewResponseValidationDataMethod `json:"method,required"`
	Status       ProjectDomainNewResponseValidationDataStatus `json:"status,required"`
	ErrorMessage string                                       `json:"error_message"`
	TXTName      string                                       `json:"txt_name"`
	TXTValue     string                                       `json:"txt_value"`
	JSON         projectDomainNewResponseValidationDataJSON   `json:"-"`
}

// projectDomainNewResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseValidationData]
type projectDomainNewResponseValidationDataJSON struct {
	Method       apijson.Field
	Status       apijson.Field
	ErrorMessage apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainNewResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseValidationDataMethod string

const (
	ProjectDomainNewResponseValidationDataMethodHTTP ProjectDomainNewResponseValidationDataMethod = "http"
	ProjectDomainNewResponseValidationDataMethodTXT  ProjectDomainNewResponseValidationDataMethod = "txt"
)

func (r ProjectDomainNewResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseValidationDataMethodHTTP, ProjectDomainNewResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainNewResponseValidationDataStatus string

const (
	ProjectDomainNewResponseValidationDataStatusInitializing ProjectDomainNewResponseValidationDataStatus = "initializing"
	ProjectDomainNewResponseValidationDataStatusPending      ProjectDomainNewResponseValidationDataStatus = "pending"
	ProjectDomainNewResponseValidationDataStatusActive       ProjectDomainNewResponseValidationDataStatus = "active"
	ProjectDomainNewResponseValidationDataStatusDeactivated  ProjectDomainNewResponseValidationDataStatus = "deactivated"
	ProjectDomainNewResponseValidationDataStatusError        ProjectDomainNewResponseValidationDataStatus = "error"
)

func (r ProjectDomainNewResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseValidationDataStatusInitializing, ProjectDomainNewResponseValidationDataStatusPending, ProjectDomainNewResponseValidationDataStatusActive, ProjectDomainNewResponseValidationDataStatusDeactivated, ProjectDomainNewResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainNewResponseVerificationData struct {
	Status       ProjectDomainNewResponseVerificationDataStatus `json:"status,required"`
	ErrorMessage string                                         `json:"error_message"`
	JSON         projectDomainNewResponseVerificationDataJSON   `json:"-"`
}

// projectDomainNewResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseVerificationData]
type projectDomainNewResponseVerificationDataJSON struct {
	Status       apijson.Field
	ErrorMessage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainNewResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseVerificationDataStatus string

const (
	ProjectDomainNewResponseVerificationDataStatusPending     ProjectDomainNewResponseVerificationDataStatus = "pending"
	ProjectDomainNewResponseVerificationDataStatusActive      ProjectDomainNewResponseVerificationDataStatus = "active"
	ProjectDomainNewResponseVerificationDataStatusDeactivated ProjectDomainNewResponseVerificationDataStatus = "deactivated"
	ProjectDomainNewResponseVerificationDataStatusBlocked     ProjectDomainNewResponseVerificationDataStatus = "blocked"
	ProjectDomainNewResponseVerificationDataStatusError       ProjectDomainNewResponseVerificationDataStatus = "error"
)

func (r ProjectDomainNewResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseVerificationDataStatusPending, ProjectDomainNewResponseVerificationDataStatusActive, ProjectDomainNewResponseVerificationDataStatusDeactivated, ProjectDomainNewResponseVerificationDataStatusBlocked, ProjectDomainNewResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainListResponse struct {
	ID                   string                                        `json:"id,required"`
	CertificateAuthority ProjectDomainListResponseCertificateAuthority `json:"certificate_authority,required"`
	CreatedOn            string                                        `json:"created_on,required"`
	DomainID             string                                        `json:"domain_id,required"`
	// The domain name.
	Name             string                                    `json:"name,required"`
	Status           ProjectDomainListResponseStatus           `json:"status,required"`
	ValidationData   ProjectDomainListResponseValidationData   `json:"validation_data,required"`
	VerificationData ProjectDomainListResponseVerificationData `json:"verification_data,required"`
	ZoneTag          string                                    `json:"zone_tag,required"`
	JSON             projectDomainListResponseJSON             `json:"-"`
}

// projectDomainListResponseJSON contains the JSON metadata for the struct
// [ProjectDomainListResponse]
type projectDomainListResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseCertificateAuthority string

const (
	ProjectDomainListResponseCertificateAuthorityGoogle      ProjectDomainListResponseCertificateAuthority = "google"
	ProjectDomainListResponseCertificateAuthorityLetsEncrypt ProjectDomainListResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainListResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseCertificateAuthorityGoogle, ProjectDomainListResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainListResponseStatus string

const (
	ProjectDomainListResponseStatusInitializing ProjectDomainListResponseStatus = "initializing"
	ProjectDomainListResponseStatusPending      ProjectDomainListResponseStatus = "pending"
	ProjectDomainListResponseStatusActive       ProjectDomainListResponseStatus = "active"
	ProjectDomainListResponseStatusDeactivated  ProjectDomainListResponseStatus = "deactivated"
	ProjectDomainListResponseStatusBlocked      ProjectDomainListResponseStatus = "blocked"
	ProjectDomainListResponseStatusError        ProjectDomainListResponseStatus = "error"
)

func (r ProjectDomainListResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseStatusInitializing, ProjectDomainListResponseStatusPending, ProjectDomainListResponseStatusActive, ProjectDomainListResponseStatusDeactivated, ProjectDomainListResponseStatusBlocked, ProjectDomainListResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainListResponseValidationData struct {
	Method       ProjectDomainListResponseValidationDataMethod `json:"method,required"`
	Status       ProjectDomainListResponseValidationDataStatus `json:"status,required"`
	ErrorMessage string                                        `json:"error_message"`
	TXTName      string                                        `json:"txt_name"`
	TXTValue     string                                        `json:"txt_value"`
	JSON         projectDomainListResponseValidationDataJSON   `json:"-"`
}

// projectDomainListResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainListResponseValidationData]
type projectDomainListResponseValidationDataJSON struct {
	Method       apijson.Field
	Status       apijson.Field
	ErrorMessage apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainListResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseValidationDataMethod string

const (
	ProjectDomainListResponseValidationDataMethodHTTP ProjectDomainListResponseValidationDataMethod = "http"
	ProjectDomainListResponseValidationDataMethodTXT  ProjectDomainListResponseValidationDataMethod = "txt"
)

func (r ProjectDomainListResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseValidationDataMethodHTTP, ProjectDomainListResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainListResponseValidationDataStatus string

const (
	ProjectDomainListResponseValidationDataStatusInitializing ProjectDomainListResponseValidationDataStatus = "initializing"
	ProjectDomainListResponseValidationDataStatusPending      ProjectDomainListResponseValidationDataStatus = "pending"
	ProjectDomainListResponseValidationDataStatusActive       ProjectDomainListResponseValidationDataStatus = "active"
	ProjectDomainListResponseValidationDataStatusDeactivated  ProjectDomainListResponseValidationDataStatus = "deactivated"
	ProjectDomainListResponseValidationDataStatusError        ProjectDomainListResponseValidationDataStatus = "error"
)

func (r ProjectDomainListResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseValidationDataStatusInitializing, ProjectDomainListResponseValidationDataStatusPending, ProjectDomainListResponseValidationDataStatusActive, ProjectDomainListResponseValidationDataStatusDeactivated, ProjectDomainListResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainListResponseVerificationData struct {
	Status       ProjectDomainListResponseVerificationDataStatus `json:"status,required"`
	ErrorMessage string                                          `json:"error_message"`
	JSON         projectDomainListResponseVerificationDataJSON   `json:"-"`
}

// projectDomainListResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainListResponseVerificationData]
type projectDomainListResponseVerificationDataJSON struct {
	Status       apijson.Field
	ErrorMessage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainListResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainListResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainListResponseVerificationDataStatus string

const (
	ProjectDomainListResponseVerificationDataStatusPending     ProjectDomainListResponseVerificationDataStatus = "pending"
	ProjectDomainListResponseVerificationDataStatusActive      ProjectDomainListResponseVerificationDataStatus = "active"
	ProjectDomainListResponseVerificationDataStatusDeactivated ProjectDomainListResponseVerificationDataStatus = "deactivated"
	ProjectDomainListResponseVerificationDataStatusBlocked     ProjectDomainListResponseVerificationDataStatus = "blocked"
	ProjectDomainListResponseVerificationDataStatusError       ProjectDomainListResponseVerificationDataStatus = "error"
)

func (r ProjectDomainListResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainListResponseVerificationDataStatusPending, ProjectDomainListResponseVerificationDataStatusActive, ProjectDomainListResponseVerificationDataStatusDeactivated, ProjectDomainListResponseVerificationDataStatusBlocked, ProjectDomainListResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainDeleteResponse = interface{}

type ProjectDomainEditResponse struct {
	ID                   string                                        `json:"id,required"`
	CertificateAuthority ProjectDomainEditResponseCertificateAuthority `json:"certificate_authority,required"`
	CreatedOn            string                                        `json:"created_on,required"`
	DomainID             string                                        `json:"domain_id,required"`
	// The domain name.
	Name             string                                    `json:"name,required"`
	Status           ProjectDomainEditResponseStatus           `json:"status,required"`
	ValidationData   ProjectDomainEditResponseValidationData   `json:"validation_data,required"`
	VerificationData ProjectDomainEditResponseVerificationData `json:"verification_data,required"`
	ZoneTag          string                                    `json:"zone_tag,required"`
	JSON             projectDomainEditResponseJSON             `json:"-"`
}

// projectDomainEditResponseJSON contains the JSON metadata for the struct
// [ProjectDomainEditResponse]
type projectDomainEditResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseCertificateAuthority string

const (
	ProjectDomainEditResponseCertificateAuthorityGoogle      ProjectDomainEditResponseCertificateAuthority = "google"
	ProjectDomainEditResponseCertificateAuthorityLetsEncrypt ProjectDomainEditResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainEditResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseCertificateAuthorityGoogle, ProjectDomainEditResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainEditResponseStatus string

const (
	ProjectDomainEditResponseStatusInitializing ProjectDomainEditResponseStatus = "initializing"
	ProjectDomainEditResponseStatusPending      ProjectDomainEditResponseStatus = "pending"
	ProjectDomainEditResponseStatusActive       ProjectDomainEditResponseStatus = "active"
	ProjectDomainEditResponseStatusDeactivated  ProjectDomainEditResponseStatus = "deactivated"
	ProjectDomainEditResponseStatusBlocked      ProjectDomainEditResponseStatus = "blocked"
	ProjectDomainEditResponseStatusError        ProjectDomainEditResponseStatus = "error"
)

func (r ProjectDomainEditResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseStatusInitializing, ProjectDomainEditResponseStatusPending, ProjectDomainEditResponseStatusActive, ProjectDomainEditResponseStatusDeactivated, ProjectDomainEditResponseStatusBlocked, ProjectDomainEditResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainEditResponseValidationData struct {
	Method       ProjectDomainEditResponseValidationDataMethod `json:"method,required"`
	Status       ProjectDomainEditResponseValidationDataStatus `json:"status,required"`
	ErrorMessage string                                        `json:"error_message"`
	TXTName      string                                        `json:"txt_name"`
	TXTValue     string                                        `json:"txt_value"`
	JSON         projectDomainEditResponseValidationDataJSON   `json:"-"`
}

// projectDomainEditResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseValidationData]
type projectDomainEditResponseValidationDataJSON struct {
	Method       apijson.Field
	Status       apijson.Field
	ErrorMessage apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainEditResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseValidationDataMethod string

const (
	ProjectDomainEditResponseValidationDataMethodHTTP ProjectDomainEditResponseValidationDataMethod = "http"
	ProjectDomainEditResponseValidationDataMethodTXT  ProjectDomainEditResponseValidationDataMethod = "txt"
)

func (r ProjectDomainEditResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseValidationDataMethodHTTP, ProjectDomainEditResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainEditResponseValidationDataStatus string

const (
	ProjectDomainEditResponseValidationDataStatusInitializing ProjectDomainEditResponseValidationDataStatus = "initializing"
	ProjectDomainEditResponseValidationDataStatusPending      ProjectDomainEditResponseValidationDataStatus = "pending"
	ProjectDomainEditResponseValidationDataStatusActive       ProjectDomainEditResponseValidationDataStatus = "active"
	ProjectDomainEditResponseValidationDataStatusDeactivated  ProjectDomainEditResponseValidationDataStatus = "deactivated"
	ProjectDomainEditResponseValidationDataStatusError        ProjectDomainEditResponseValidationDataStatus = "error"
)

func (r ProjectDomainEditResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseValidationDataStatusInitializing, ProjectDomainEditResponseValidationDataStatusPending, ProjectDomainEditResponseValidationDataStatusActive, ProjectDomainEditResponseValidationDataStatusDeactivated, ProjectDomainEditResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainEditResponseVerificationData struct {
	Status       ProjectDomainEditResponseVerificationDataStatus `json:"status,required"`
	ErrorMessage string                                          `json:"error_message"`
	JSON         projectDomainEditResponseVerificationDataJSON   `json:"-"`
}

// projectDomainEditResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseVerificationData]
type projectDomainEditResponseVerificationDataJSON struct {
	Status       apijson.Field
	ErrorMessage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainEditResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseVerificationDataStatus string

const (
	ProjectDomainEditResponseVerificationDataStatusPending     ProjectDomainEditResponseVerificationDataStatus = "pending"
	ProjectDomainEditResponseVerificationDataStatusActive      ProjectDomainEditResponseVerificationDataStatus = "active"
	ProjectDomainEditResponseVerificationDataStatusDeactivated ProjectDomainEditResponseVerificationDataStatus = "deactivated"
	ProjectDomainEditResponseVerificationDataStatusBlocked     ProjectDomainEditResponseVerificationDataStatus = "blocked"
	ProjectDomainEditResponseVerificationDataStatusError       ProjectDomainEditResponseVerificationDataStatus = "error"
)

func (r ProjectDomainEditResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseVerificationDataStatusPending, ProjectDomainEditResponseVerificationDataStatusActive, ProjectDomainEditResponseVerificationDataStatusDeactivated, ProjectDomainEditResponseVerificationDataStatusBlocked, ProjectDomainEditResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainGetResponse struct {
	ID                   string                                       `json:"id,required"`
	CertificateAuthority ProjectDomainGetResponseCertificateAuthority `json:"certificate_authority,required"`
	CreatedOn            string                                       `json:"created_on,required"`
	DomainID             string                                       `json:"domain_id,required"`
	// The domain name.
	Name             string                                   `json:"name,required"`
	Status           ProjectDomainGetResponseStatus           `json:"status,required"`
	ValidationData   ProjectDomainGetResponseValidationData   `json:"validation_data,required"`
	VerificationData ProjectDomainGetResponseVerificationData `json:"verification_data,required"`
	ZoneTag          string                                   `json:"zone_tag,required"`
	JSON             projectDomainGetResponseJSON             `json:"-"`
}

// projectDomainGetResponseJSON contains the JSON metadata for the struct
// [ProjectDomainGetResponse]
type projectDomainGetResponseJSON struct {
	ID                   apijson.Field
	CertificateAuthority apijson.Field
	CreatedOn            apijson.Field
	DomainID             apijson.Field
	Name                 apijson.Field
	Status               apijson.Field
	ValidationData       apijson.Field
	VerificationData     apijson.Field
	ZoneTag              apijson.Field
	raw                  string
	ExtraFields          map[string]apijson.Field
}

func (r *ProjectDomainGetResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseCertificateAuthority string

const (
	ProjectDomainGetResponseCertificateAuthorityGoogle      ProjectDomainGetResponseCertificateAuthority = "google"
	ProjectDomainGetResponseCertificateAuthorityLetsEncrypt ProjectDomainGetResponseCertificateAuthority = "lets_encrypt"
)

func (r ProjectDomainGetResponseCertificateAuthority) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseCertificateAuthorityGoogle, ProjectDomainGetResponseCertificateAuthorityLetsEncrypt:
		return true
	}
	return false
}

type ProjectDomainGetResponseStatus string

const (
	ProjectDomainGetResponseStatusInitializing ProjectDomainGetResponseStatus = "initializing"
	ProjectDomainGetResponseStatusPending      ProjectDomainGetResponseStatus = "pending"
	ProjectDomainGetResponseStatusActive       ProjectDomainGetResponseStatus = "active"
	ProjectDomainGetResponseStatusDeactivated  ProjectDomainGetResponseStatus = "deactivated"
	ProjectDomainGetResponseStatusBlocked      ProjectDomainGetResponseStatus = "blocked"
	ProjectDomainGetResponseStatusError        ProjectDomainGetResponseStatus = "error"
)

func (r ProjectDomainGetResponseStatus) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseStatusInitializing, ProjectDomainGetResponseStatusPending, ProjectDomainGetResponseStatusActive, ProjectDomainGetResponseStatusDeactivated, ProjectDomainGetResponseStatusBlocked, ProjectDomainGetResponseStatusError:
		return true
	}
	return false
}

type ProjectDomainGetResponseValidationData struct {
	Method       ProjectDomainGetResponseValidationDataMethod `json:"method,required"`
	Status       ProjectDomainGetResponseValidationDataStatus `json:"status,required"`
	ErrorMessage string                                       `json:"error_message"`
	TXTName      string                                       `json:"txt_name"`
	TXTValue     string                                       `json:"txt_value"`
	JSON         projectDomainGetResponseValidationDataJSON   `json:"-"`
}

// projectDomainGetResponseValidationDataJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseValidationData]
type projectDomainGetResponseValidationDataJSON struct {
	Method       apijson.Field
	Status       apijson.Field
	ErrorMessage apijson.Field
	TXTName      apijson.Field
	TXTValue     apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainGetResponseValidationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseValidationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseValidationDataMethod string

const (
	ProjectDomainGetResponseValidationDataMethodHTTP ProjectDomainGetResponseValidationDataMethod = "http"
	ProjectDomainGetResponseValidationDataMethodTXT  ProjectDomainGetResponseValidationDataMethod = "txt"
)

func (r ProjectDomainGetResponseValidationDataMethod) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseValidationDataMethodHTTP, ProjectDomainGetResponseValidationDataMethodTXT:
		return true
	}
	return false
}

type ProjectDomainGetResponseValidationDataStatus string

const (
	ProjectDomainGetResponseValidationDataStatusInitializing ProjectDomainGetResponseValidationDataStatus = "initializing"
	ProjectDomainGetResponseValidationDataStatusPending      ProjectDomainGetResponseValidationDataStatus = "pending"
	ProjectDomainGetResponseValidationDataStatusActive       ProjectDomainGetResponseValidationDataStatus = "active"
	ProjectDomainGetResponseValidationDataStatusDeactivated  ProjectDomainGetResponseValidationDataStatus = "deactivated"
	ProjectDomainGetResponseValidationDataStatusError        ProjectDomainGetResponseValidationDataStatus = "error"
)

func (r ProjectDomainGetResponseValidationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseValidationDataStatusInitializing, ProjectDomainGetResponseValidationDataStatusPending, ProjectDomainGetResponseValidationDataStatusActive, ProjectDomainGetResponseValidationDataStatusDeactivated, ProjectDomainGetResponseValidationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainGetResponseVerificationData struct {
	Status       ProjectDomainGetResponseVerificationDataStatus `json:"status,required"`
	ErrorMessage string                                         `json:"error_message"`
	JSON         projectDomainGetResponseVerificationDataJSON   `json:"-"`
}

// projectDomainGetResponseVerificationDataJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseVerificationData]
type projectDomainGetResponseVerificationDataJSON struct {
	Status       apijson.Field
	ErrorMessage apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *ProjectDomainGetResponseVerificationData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseVerificationDataJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseVerificationDataStatus string

const (
	ProjectDomainGetResponseVerificationDataStatusPending     ProjectDomainGetResponseVerificationDataStatus = "pending"
	ProjectDomainGetResponseVerificationDataStatusActive      ProjectDomainGetResponseVerificationDataStatus = "active"
	ProjectDomainGetResponseVerificationDataStatusDeactivated ProjectDomainGetResponseVerificationDataStatus = "deactivated"
	ProjectDomainGetResponseVerificationDataStatusBlocked     ProjectDomainGetResponseVerificationDataStatus = "blocked"
	ProjectDomainGetResponseVerificationDataStatusError       ProjectDomainGetResponseVerificationDataStatus = "error"
)

func (r ProjectDomainGetResponseVerificationDataStatus) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseVerificationDataStatusPending, ProjectDomainGetResponseVerificationDataStatusActive, ProjectDomainGetResponseVerificationDataStatusDeactivated, ProjectDomainGetResponseVerificationDataStatusBlocked, ProjectDomainGetResponseVerificationDataStatusError:
		return true
	}
	return false
}

type ProjectDomainNewParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
	// The domain name.
	Name param.Field[string] `json:"name,required"`
}

func (r ProjectDomainNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ProjectDomainNewResponseEnvelope struct {
	Errors   []ProjectDomainNewResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainNewResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDomainNewResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDomainNewResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainNewResponseEnvelopeJSON    `json:"-"`
}

// projectDomainNewResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainNewResponseEnvelope]
type projectDomainNewResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ProjectDomainNewResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDomainNewResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDomainNewResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseEnvelopeErrors]
type projectDomainNewResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    projectDomainNewResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDomainNewResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ProjectDomainNewResponseEnvelopeErrorsSource]
type projectDomainNewResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ProjectDomainNewResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDomainNewResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDomainNewResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDomainNewResponseEnvelopeMessages]
type projectDomainNewResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainNewResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    projectDomainNewResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDomainNewResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ProjectDomainNewResponseEnvelopeMessagesSource]
type projectDomainNewResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainNewResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainNewResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDomainNewResponseEnvelopeSuccess bool

const (
	ProjectDomainNewResponseEnvelopeSuccessTrue ProjectDomainNewResponseEnvelopeSuccess = true
)

func (r ProjectDomainNewResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainNewResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainListParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainDeleteParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainDeleteResponseEnvelope struct {
	Errors   []ProjectDomainDeleteResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainDeleteResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDomainDeleteResponse                   `json:"result,required,nullable"`
	// Whether the API call was successful.
	Success ProjectDomainDeleteResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainDeleteResponseEnvelopeJSON    `json:"-"`
}

// projectDomainDeleteResponseEnvelopeJSON contains the JSON metadata for the
// struct [ProjectDomainDeleteResponseEnvelope]
type projectDomainDeleteResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainDeleteResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainDeleteResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainDeleteResponseEnvelopeErrors struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ProjectDomainDeleteResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDomainDeleteResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDomainDeleteResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainDeleteResponseEnvelopeErrors]
type projectDomainDeleteResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainDeleteResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainDeleteResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainDeleteResponseEnvelopeErrorsSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    projectDomainDeleteResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDomainDeleteResponseEnvelopeErrorsSourceJSON contains the JSON metadata
// for the struct [ProjectDomainDeleteResponseEnvelopeErrorsSource]
type projectDomainDeleteResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainDeleteResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainDeleteResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainDeleteResponseEnvelopeMessages struct {
	Code             int64                                             `json:"code,required"`
	Message          string                                            `json:"message,required"`
	DocumentationURL string                                            `json:"documentation_url"`
	Source           ProjectDomainDeleteResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDomainDeleteResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDomainDeleteResponseEnvelopeMessagesJSON contains the JSON metadata for
// the struct [ProjectDomainDeleteResponseEnvelopeMessages]
type projectDomainDeleteResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainDeleteResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainDeleteResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainDeleteResponseEnvelopeMessagesSource struct {
	Pointer string                                                `json:"pointer"`
	JSON    projectDomainDeleteResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDomainDeleteResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ProjectDomainDeleteResponseEnvelopeMessagesSource]
type projectDomainDeleteResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainDeleteResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainDeleteResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDomainDeleteResponseEnvelopeSuccess bool

const (
	ProjectDomainDeleteResponseEnvelopeSuccessTrue ProjectDomainDeleteResponseEnvelopeSuccess = true
)

func (r ProjectDomainDeleteResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainDeleteResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainEditParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainEditResponseEnvelope struct {
	Errors   []ProjectDomainEditResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainEditResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDomainEditResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDomainEditResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainEditResponseEnvelopeJSON    `json:"-"`
}

// projectDomainEditResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainEditResponseEnvelope]
type projectDomainEditResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseEnvelopeErrors struct {
	Code             int64                                         `json:"code,required"`
	Message          string                                        `json:"message,required"`
	DocumentationURL string                                        `json:"documentation_url"`
	Source           ProjectDomainEditResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDomainEditResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDomainEditResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseEnvelopeErrors]
type projectDomainEditResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseEnvelopeErrorsSource struct {
	Pointer string                                            `json:"pointer"`
	JSON    projectDomainEditResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDomainEditResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ProjectDomainEditResponseEnvelopeErrorsSource]
type projectDomainEditResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseEnvelopeMessages struct {
	Code             int64                                           `json:"code,required"`
	Message          string                                          `json:"message,required"`
	DocumentationURL string                                          `json:"documentation_url"`
	Source           ProjectDomainEditResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDomainEditResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDomainEditResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDomainEditResponseEnvelopeMessages]
type projectDomainEditResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainEditResponseEnvelopeMessagesSource struct {
	Pointer string                                              `json:"pointer"`
	JSON    projectDomainEditResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDomainEditResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ProjectDomainEditResponseEnvelopeMessagesSource]
type projectDomainEditResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainEditResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainEditResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDomainEditResponseEnvelopeSuccess bool

const (
	ProjectDomainEditResponseEnvelopeSuccessTrue ProjectDomainEditResponseEnvelopeSuccess = true
)

func (r ProjectDomainEditResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainEditResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}

type ProjectDomainGetParams struct {
	// Identifier.
	AccountID param.Field[string] `path:"account_id,required"`
}

type ProjectDomainGetResponseEnvelope struct {
	Errors   []ProjectDomainGetResponseEnvelopeErrors   `json:"errors,required"`
	Messages []ProjectDomainGetResponseEnvelopeMessages `json:"messages,required"`
	Result   ProjectDomainGetResponse                   `json:"result,required"`
	// Whether the API call was successful.
	Success ProjectDomainGetResponseEnvelopeSuccess `json:"success,required"`
	JSON    projectDomainGetResponseEnvelopeJSON    `json:"-"`
}

// projectDomainGetResponseEnvelopeJSON contains the JSON metadata for the struct
// [ProjectDomainGetResponseEnvelope]
type projectDomainGetResponseEnvelopeJSON struct {
	Errors      apijson.Field
	Messages    apijson.Field
	Result      apijson.Field
	Success     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelope) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseEnvelopeErrors struct {
	Code             int64                                        `json:"code,required"`
	Message          string                                       `json:"message,required"`
	DocumentationURL string                                       `json:"documentation_url"`
	Source           ProjectDomainGetResponseEnvelopeErrorsSource `json:"source"`
	JSON             projectDomainGetResponseEnvelopeErrorsJSON   `json:"-"`
}

// projectDomainGetResponseEnvelopeErrorsJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseEnvelopeErrors]
type projectDomainGetResponseEnvelopeErrorsJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelopeErrors) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeErrorsJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseEnvelopeErrorsSource struct {
	Pointer string                                           `json:"pointer"`
	JSON    projectDomainGetResponseEnvelopeErrorsSourceJSON `json:"-"`
}

// projectDomainGetResponseEnvelopeErrorsSourceJSON contains the JSON metadata for
// the struct [ProjectDomainGetResponseEnvelopeErrorsSource]
type projectDomainGetResponseEnvelopeErrorsSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelopeErrorsSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeErrorsSourceJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseEnvelopeMessages struct {
	Code             int64                                          `json:"code,required"`
	Message          string                                         `json:"message,required"`
	DocumentationURL string                                         `json:"documentation_url"`
	Source           ProjectDomainGetResponseEnvelopeMessagesSource `json:"source"`
	JSON             projectDomainGetResponseEnvelopeMessagesJSON   `json:"-"`
}

// projectDomainGetResponseEnvelopeMessagesJSON contains the JSON metadata for the
// struct [ProjectDomainGetResponseEnvelopeMessages]
type projectDomainGetResponseEnvelopeMessagesJSON struct {
	Code             apijson.Field
	Message          apijson.Field
	DocumentationURL apijson.Field
	Source           apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelopeMessages) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeMessagesJSON) RawJSON() string {
	return r.raw
}

type ProjectDomainGetResponseEnvelopeMessagesSource struct {
	Pointer string                                             `json:"pointer"`
	JSON    projectDomainGetResponseEnvelopeMessagesSourceJSON `json:"-"`
}

// projectDomainGetResponseEnvelopeMessagesSourceJSON contains the JSON metadata
// for the struct [ProjectDomainGetResponseEnvelopeMessagesSource]
type projectDomainGetResponseEnvelopeMessagesSourceJSON struct {
	Pointer     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ProjectDomainGetResponseEnvelopeMessagesSource) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r projectDomainGetResponseEnvelopeMessagesSourceJSON) RawJSON() string {
	return r.raw
}

// Whether the API call was successful.
type ProjectDomainGetResponseEnvelopeSuccess bool

const (
	ProjectDomainGetResponseEnvelopeSuccessTrue ProjectDomainGetResponseEnvelopeSuccess = true
)

func (r ProjectDomainGetResponseEnvelopeSuccess) IsKnown() bool {
	switch r {
	case ProjectDomainGetResponseEnvelopeSuccessTrue:
		return true
	}
	return false
}
