// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloudforce_one

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/cloudflare/cloudflare-go/v7/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v7/internal/apiquery"
	"github.com/cloudflare/cloudflare-go/v7/internal/param"
	"github.com/cloudflare/cloudflare-go/v7/internal/requestconfig"
	"github.com/cloudflare/cloudflare-go/v7/option"
)

// ThreatEventTagCategoryService contains methods and other services that help with
// interacting with the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewThreatEventTagCategoryService] method instead.
type ThreatEventTagCategoryService struct {
	Options []option.RequestOption
}

// NewThreatEventTagCategoryService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewThreatEventTagCategoryService(opts ...option.RequestOption) (r *ThreatEventTagCategoryService) {
	r = &ThreatEventTagCategoryService{}
	r.Options = opts
	return
}

// Creates a new Source-of-Truth tag category for an account.
func (r *ThreatEventTagCategoryService) New(ctx context.Context, params ThreatEventTagCategoryNewParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories/create", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns all Source-of-Truth tag categories for an account.
func (r *ThreatEventTagCategoryService) List(ctx context.Context, params ThreatEventTagCategoryListParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories", params.AccountID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Deletes a Source-of-Truth tag category by UUID.
func (r *ThreatEventTagCategoryService) Delete(ctx context.Context, categoryUUID string, body ThreatEventTagCategoryDeleteParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryUUID == "" {
		err = errors.New("missing required category_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories/%s", body.AccountID, categoryUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Updates a Source-of-Truth tag category by UUID.
func (r *ThreatEventTagCategoryService) Edit(ctx context.Context, categoryUUID string, params ThreatEventTagCategoryEditParams, opts ...option.RequestOption) (res *ThreatEventTagCategoryEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountID.Value == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	if categoryUUID == "" {
		err = errors.New("missing required category_uuid parameter")
		return nil, err
	}
	path := fmt.Sprintf("accounts/%s/cloudforce-one/events/tags/categories/%s", params.AccountID, categoryUUID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type ThreatEventTagCategoryNewResponse struct {
	Name        string `json:"name" api:"required"`
	UUID        string `json:"uuid" api:"required"`
	CreatedAt   string `json:"createdAt"`
	Description string `json:"description"`
	// Parsed FieldDefinition[] defining custom fields for this category, or null if
	// none.
	Schema    []ThreatEventTagCategoryNewResponseSchema `json:"schema" api:"nullable"`
	UpdatedAt string                                    `json:"updatedAt"`
	JSON      threatEventTagCategoryNewResponseJSON     `json:"-"`
}

// threatEventTagCategoryNewResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagCategoryNewResponse]
type threatEventTagCategoryNewResponseJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Schema      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryNewResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryNewResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryNewResponseSchema struct {
	Key              string                                                  `json:"key" api:"required"`
	Kind             ThreatEventTagCategoryNewResponseSchemaKind             `json:"kind" api:"required"`
	AllowedValues    []string                                                `json:"allowedValues"`
	Annotations      ThreatEventTagCategoryNewResponseSchemaAnnotations      `json:"annotations"`
	Element          interface{}                                             `json:"element"`
	Enforcement      ThreatEventTagCategoryNewResponseSchemaEnforcement      `json:"enforcement"`
	Format           ThreatEventTagCategoryNewResponseSchemaFormat           `json:"format"`
	Label            string                                                  `json:"label"`
	MaxLength        int64                                                   `json:"maxLength"`
	NumberConstraint ThreatEventTagCategoryNewResponseSchemaNumberConstraint `json:"numberConstraint"`
	// Map of property key to FieldDefinition for object fields. Required when kind is
	// 'object'. See FieldDefinition (recursive).
	Properties map[string]interface{}                      `json:"properties"`
	Required   bool                                        `json:"required"`
	JSON       threatEventTagCategoryNewResponseSchemaJSON `json:"-"`
}

// threatEventTagCategoryNewResponseSchemaJSON contains the JSON metadata for the
// struct [ThreatEventTagCategoryNewResponseSchema]
type threatEventTagCategoryNewResponseSchemaJSON struct {
	Key              apijson.Field
	Kind             apijson.Field
	AllowedValues    apijson.Field
	Annotations      apijson.Field
	Element          apijson.Field
	Enforcement      apijson.Field
	Format           apijson.Field
	Label            apijson.Field
	MaxLength        apijson.Field
	NumberConstraint apijson.Field
	Properties       apijson.Field
	Required         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreatEventTagCategoryNewResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryNewResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryNewResponseSchemaKind string

const (
	ThreatEventTagCategoryNewResponseSchemaKindString ThreatEventTagCategoryNewResponseSchemaKind = "string"
	ThreatEventTagCategoryNewResponseSchemaKindNumber ThreatEventTagCategoryNewResponseSchemaKind = "number"
	ThreatEventTagCategoryNewResponseSchemaKindEnum   ThreatEventTagCategoryNewResponseSchemaKind = "enum"
	ThreatEventTagCategoryNewResponseSchemaKindDate   ThreatEventTagCategoryNewResponseSchemaKind = "date"
	ThreatEventTagCategoryNewResponseSchemaKindArray  ThreatEventTagCategoryNewResponseSchemaKind = "array"
	ThreatEventTagCategoryNewResponseSchemaKindObject ThreatEventTagCategoryNewResponseSchemaKind = "object"
)

func (r ThreatEventTagCategoryNewResponseSchemaKind) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryNewResponseSchemaKindString, ThreatEventTagCategoryNewResponseSchemaKindNumber, ThreatEventTagCategoryNewResponseSchemaKindEnum, ThreatEventTagCategoryNewResponseSchemaKindDate, ThreatEventTagCategoryNewResponseSchemaKindArray, ThreatEventTagCategoryNewResponseSchemaKindObject:
		return true
	}
	return false
}

type ThreatEventTagCategoryNewResponseSchemaAnnotations struct {
	Confidence bool                                                   `json:"confidence"`
	TLP        bool                                                   `json:"tlp"`
	JSON       threatEventTagCategoryNewResponseSchemaAnnotationsJSON `json:"-"`
}

// threatEventTagCategoryNewResponseSchemaAnnotationsJSON contains the JSON
// metadata for the struct [ThreatEventTagCategoryNewResponseSchemaAnnotations]
type threatEventTagCategoryNewResponseSchemaAnnotationsJSON struct {
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryNewResponseSchemaAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryNewResponseSchemaAnnotationsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryNewResponseSchemaEnforcement string

const (
	ThreatEventTagCategoryNewResponseSchemaEnforcementError ThreatEventTagCategoryNewResponseSchemaEnforcement = "error"
	ThreatEventTagCategoryNewResponseSchemaEnforcementWarn  ThreatEventTagCategoryNewResponseSchemaEnforcement = "warn"
	ThreatEventTagCategoryNewResponseSchemaEnforcementOff   ThreatEventTagCategoryNewResponseSchemaEnforcement = "off"
)

func (r ThreatEventTagCategoryNewResponseSchemaEnforcement) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryNewResponseSchemaEnforcementError, ThreatEventTagCategoryNewResponseSchemaEnforcementWarn, ThreatEventTagCategoryNewResponseSchemaEnforcementOff:
		return true
	}
	return false
}

type ThreatEventTagCategoryNewResponseSchemaFormat string

const (
	ThreatEventTagCategoryNewResponseSchemaFormatDate     ThreatEventTagCategoryNewResponseSchemaFormat = "date"
	ThreatEventTagCategoryNewResponseSchemaFormatURL      ThreatEventTagCategoryNewResponseSchemaFormat = "url"
	ThreatEventTagCategoryNewResponseSchemaFormatDuration ThreatEventTagCategoryNewResponseSchemaFormat = "duration"
	ThreatEventTagCategoryNewResponseSchemaFormatCountry  ThreatEventTagCategoryNewResponseSchemaFormat = "country"
)

func (r ThreatEventTagCategoryNewResponseSchemaFormat) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryNewResponseSchemaFormatDate, ThreatEventTagCategoryNewResponseSchemaFormatURL, ThreatEventTagCategoryNewResponseSchemaFormatDuration, ThreatEventTagCategoryNewResponseSchemaFormatCountry:
		return true
	}
	return false
}

type ThreatEventTagCategoryNewResponseSchemaNumberConstraint struct {
	Integer bool                                                        `json:"integer"`
	Max     float64                                                     `json:"max"`
	Min     float64                                                     `json:"min"`
	JSON    threatEventTagCategoryNewResponseSchemaNumberConstraintJSON `json:"-"`
}

// threatEventTagCategoryNewResponseSchemaNumberConstraintJSON contains the JSON
// metadata for the struct
// [ThreatEventTagCategoryNewResponseSchemaNumberConstraint]
type threatEventTagCategoryNewResponseSchemaNumberConstraintJSON struct {
	Integer     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryNewResponseSchemaNumberConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryNewResponseSchemaNumberConstraintJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryListResponse struct {
	Categories []ThreatEventTagCategoryListResponseCategory `json:"categories" api:"required"`
	JSON       threatEventTagCategoryListResponseJSON       `json:"-"`
}

// threatEventTagCategoryListResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagCategoryListResponse]
type threatEventTagCategoryListResponseJSON struct {
	Categories  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryListResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryListResponseCategory struct {
	Name        string `json:"name" api:"required"`
	UUID        string `json:"uuid" api:"required"`
	CreatedAt   string `json:"createdAt"`
	Description string `json:"description"`
	// Parsed FieldDefinition[] defining custom fields for this category, or null if
	// none.
	Schema    []ThreatEventTagCategoryListResponseCategoriesSchema `json:"schema" api:"nullable"`
	UpdatedAt string                                               `json:"updatedAt"`
	JSON      threatEventTagCategoryListResponseCategoryJSON       `json:"-"`
}

// threatEventTagCategoryListResponseCategoryJSON contains the JSON metadata for
// the struct [ThreatEventTagCategoryListResponseCategory]
type threatEventTagCategoryListResponseCategoryJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Schema      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryListResponseCategory) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryListResponseCategoryJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryListResponseCategoriesSchema struct {
	Key              string                                                             `json:"key" api:"required"`
	Kind             ThreatEventTagCategoryListResponseCategoriesSchemaKind             `json:"kind" api:"required"`
	AllowedValues    []string                                                           `json:"allowedValues"`
	Annotations      ThreatEventTagCategoryListResponseCategoriesSchemaAnnotations      `json:"annotations"`
	Element          interface{}                                                        `json:"element"`
	Enforcement      ThreatEventTagCategoryListResponseCategoriesSchemaEnforcement      `json:"enforcement"`
	Format           ThreatEventTagCategoryListResponseCategoriesSchemaFormat           `json:"format"`
	Label            string                                                             `json:"label"`
	MaxLength        int64                                                              `json:"maxLength"`
	NumberConstraint ThreatEventTagCategoryListResponseCategoriesSchemaNumberConstraint `json:"numberConstraint"`
	// Map of property key to FieldDefinition for object fields. Required when kind is
	// 'object'. See FieldDefinition (recursive).
	Properties map[string]interface{}                                 `json:"properties"`
	Required   bool                                                   `json:"required"`
	JSON       threatEventTagCategoryListResponseCategoriesSchemaJSON `json:"-"`
}

// threatEventTagCategoryListResponseCategoriesSchemaJSON contains the JSON
// metadata for the struct [ThreatEventTagCategoryListResponseCategoriesSchema]
type threatEventTagCategoryListResponseCategoriesSchemaJSON struct {
	Key              apijson.Field
	Kind             apijson.Field
	AllowedValues    apijson.Field
	Annotations      apijson.Field
	Element          apijson.Field
	Enforcement      apijson.Field
	Format           apijson.Field
	Label            apijson.Field
	MaxLength        apijson.Field
	NumberConstraint apijson.Field
	Properties       apijson.Field
	Required         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreatEventTagCategoryListResponseCategoriesSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryListResponseCategoriesSchemaJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryListResponseCategoriesSchemaKind string

const (
	ThreatEventTagCategoryListResponseCategoriesSchemaKindString ThreatEventTagCategoryListResponseCategoriesSchemaKind = "string"
	ThreatEventTagCategoryListResponseCategoriesSchemaKindNumber ThreatEventTagCategoryListResponseCategoriesSchemaKind = "number"
	ThreatEventTagCategoryListResponseCategoriesSchemaKindEnum   ThreatEventTagCategoryListResponseCategoriesSchemaKind = "enum"
	ThreatEventTagCategoryListResponseCategoriesSchemaKindDate   ThreatEventTagCategoryListResponseCategoriesSchemaKind = "date"
	ThreatEventTagCategoryListResponseCategoriesSchemaKindArray  ThreatEventTagCategoryListResponseCategoriesSchemaKind = "array"
	ThreatEventTagCategoryListResponseCategoriesSchemaKindObject ThreatEventTagCategoryListResponseCategoriesSchemaKind = "object"
)

func (r ThreatEventTagCategoryListResponseCategoriesSchemaKind) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryListResponseCategoriesSchemaKindString, ThreatEventTagCategoryListResponseCategoriesSchemaKindNumber, ThreatEventTagCategoryListResponseCategoriesSchemaKindEnum, ThreatEventTagCategoryListResponseCategoriesSchemaKindDate, ThreatEventTagCategoryListResponseCategoriesSchemaKindArray, ThreatEventTagCategoryListResponseCategoriesSchemaKindObject:
		return true
	}
	return false
}

type ThreatEventTagCategoryListResponseCategoriesSchemaAnnotations struct {
	Confidence bool                                                              `json:"confidence"`
	TLP        bool                                                              `json:"tlp"`
	JSON       threatEventTagCategoryListResponseCategoriesSchemaAnnotationsJSON `json:"-"`
}

// threatEventTagCategoryListResponseCategoriesSchemaAnnotationsJSON contains the
// JSON metadata for the struct
// [ThreatEventTagCategoryListResponseCategoriesSchemaAnnotations]
type threatEventTagCategoryListResponseCategoriesSchemaAnnotationsJSON struct {
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryListResponseCategoriesSchemaAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryListResponseCategoriesSchemaAnnotationsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryListResponseCategoriesSchemaEnforcement string

const (
	ThreatEventTagCategoryListResponseCategoriesSchemaEnforcementError ThreatEventTagCategoryListResponseCategoriesSchemaEnforcement = "error"
	ThreatEventTagCategoryListResponseCategoriesSchemaEnforcementWarn  ThreatEventTagCategoryListResponseCategoriesSchemaEnforcement = "warn"
	ThreatEventTagCategoryListResponseCategoriesSchemaEnforcementOff   ThreatEventTagCategoryListResponseCategoriesSchemaEnforcement = "off"
)

func (r ThreatEventTagCategoryListResponseCategoriesSchemaEnforcement) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryListResponseCategoriesSchemaEnforcementError, ThreatEventTagCategoryListResponseCategoriesSchemaEnforcementWarn, ThreatEventTagCategoryListResponseCategoriesSchemaEnforcementOff:
		return true
	}
	return false
}

type ThreatEventTagCategoryListResponseCategoriesSchemaFormat string

const (
	ThreatEventTagCategoryListResponseCategoriesSchemaFormatDate     ThreatEventTagCategoryListResponseCategoriesSchemaFormat = "date"
	ThreatEventTagCategoryListResponseCategoriesSchemaFormatURL      ThreatEventTagCategoryListResponseCategoriesSchemaFormat = "url"
	ThreatEventTagCategoryListResponseCategoriesSchemaFormatDuration ThreatEventTagCategoryListResponseCategoriesSchemaFormat = "duration"
	ThreatEventTagCategoryListResponseCategoriesSchemaFormatCountry  ThreatEventTagCategoryListResponseCategoriesSchemaFormat = "country"
)

func (r ThreatEventTagCategoryListResponseCategoriesSchemaFormat) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryListResponseCategoriesSchemaFormatDate, ThreatEventTagCategoryListResponseCategoriesSchemaFormatURL, ThreatEventTagCategoryListResponseCategoriesSchemaFormatDuration, ThreatEventTagCategoryListResponseCategoriesSchemaFormatCountry:
		return true
	}
	return false
}

type ThreatEventTagCategoryListResponseCategoriesSchemaNumberConstraint struct {
	Integer bool                                                                   `json:"integer"`
	Max     float64                                                                `json:"max"`
	Min     float64                                                                `json:"min"`
	JSON    threatEventTagCategoryListResponseCategoriesSchemaNumberConstraintJSON `json:"-"`
}

// threatEventTagCategoryListResponseCategoriesSchemaNumberConstraintJSON contains
// the JSON metadata for the struct
// [ThreatEventTagCategoryListResponseCategoriesSchemaNumberConstraint]
type threatEventTagCategoryListResponseCategoriesSchemaNumberConstraintJSON struct {
	Integer     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryListResponseCategoriesSchemaNumberConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryListResponseCategoriesSchemaNumberConstraintJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryDeleteResponse struct {
	UUID string                                   `json:"uuid" api:"required"`
	JSON threatEventTagCategoryDeleteResponseJSON `json:"-"`
}

// threatEventTagCategoryDeleteResponseJSON contains the JSON metadata for the
// struct [ThreatEventTagCategoryDeleteResponse]
type threatEventTagCategoryDeleteResponseJSON struct {
	UUID        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryEditResponse struct {
	Name        string `json:"name" api:"required"`
	UUID        string `json:"uuid" api:"required"`
	CreatedAt   string `json:"createdAt"`
	Description string `json:"description"`
	// Parsed FieldDefinition[] defining custom fields for this category, or null if
	// none.
	Schema    []ThreatEventTagCategoryEditResponseSchema `json:"schema" api:"nullable"`
	UpdatedAt string                                     `json:"updatedAt"`
	JSON      threatEventTagCategoryEditResponseJSON     `json:"-"`
}

// threatEventTagCategoryEditResponseJSON contains the JSON metadata for the struct
// [ThreatEventTagCategoryEditResponse]
type threatEventTagCategoryEditResponseJSON struct {
	Name        apijson.Field
	UUID        apijson.Field
	CreatedAt   apijson.Field
	Description apijson.Field
	Schema      apijson.Field
	UpdatedAt   apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryEditResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryEditResponseJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryEditResponseSchema struct {
	Key              string                                                   `json:"key" api:"required"`
	Kind             ThreatEventTagCategoryEditResponseSchemaKind             `json:"kind" api:"required"`
	AllowedValues    []string                                                 `json:"allowedValues"`
	Annotations      ThreatEventTagCategoryEditResponseSchemaAnnotations      `json:"annotations"`
	Element          interface{}                                              `json:"element"`
	Enforcement      ThreatEventTagCategoryEditResponseSchemaEnforcement      `json:"enforcement"`
	Format           ThreatEventTagCategoryEditResponseSchemaFormat           `json:"format"`
	Label            string                                                   `json:"label"`
	MaxLength        int64                                                    `json:"maxLength"`
	NumberConstraint ThreatEventTagCategoryEditResponseSchemaNumberConstraint `json:"numberConstraint"`
	// Map of property key to FieldDefinition for object fields. Required when kind is
	// 'object'. See FieldDefinition (recursive).
	Properties map[string]interface{}                       `json:"properties"`
	Required   bool                                         `json:"required"`
	JSON       threatEventTagCategoryEditResponseSchemaJSON `json:"-"`
}

// threatEventTagCategoryEditResponseSchemaJSON contains the JSON metadata for the
// struct [ThreatEventTagCategoryEditResponseSchema]
type threatEventTagCategoryEditResponseSchemaJSON struct {
	Key              apijson.Field
	Kind             apijson.Field
	AllowedValues    apijson.Field
	Annotations      apijson.Field
	Element          apijson.Field
	Enforcement      apijson.Field
	Format           apijson.Field
	Label            apijson.Field
	MaxLength        apijson.Field
	NumberConstraint apijson.Field
	Properties       apijson.Field
	Required         apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *ThreatEventTagCategoryEditResponseSchema) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryEditResponseSchemaJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryEditResponseSchemaKind string

const (
	ThreatEventTagCategoryEditResponseSchemaKindString ThreatEventTagCategoryEditResponseSchemaKind = "string"
	ThreatEventTagCategoryEditResponseSchemaKindNumber ThreatEventTagCategoryEditResponseSchemaKind = "number"
	ThreatEventTagCategoryEditResponseSchemaKindEnum   ThreatEventTagCategoryEditResponseSchemaKind = "enum"
	ThreatEventTagCategoryEditResponseSchemaKindDate   ThreatEventTagCategoryEditResponseSchemaKind = "date"
	ThreatEventTagCategoryEditResponseSchemaKindArray  ThreatEventTagCategoryEditResponseSchemaKind = "array"
	ThreatEventTagCategoryEditResponseSchemaKindObject ThreatEventTagCategoryEditResponseSchemaKind = "object"
)

func (r ThreatEventTagCategoryEditResponseSchemaKind) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryEditResponseSchemaKindString, ThreatEventTagCategoryEditResponseSchemaKindNumber, ThreatEventTagCategoryEditResponseSchemaKindEnum, ThreatEventTagCategoryEditResponseSchemaKindDate, ThreatEventTagCategoryEditResponseSchemaKindArray, ThreatEventTagCategoryEditResponseSchemaKindObject:
		return true
	}
	return false
}

type ThreatEventTagCategoryEditResponseSchemaAnnotations struct {
	Confidence bool                                                    `json:"confidence"`
	TLP        bool                                                    `json:"tlp"`
	JSON       threatEventTagCategoryEditResponseSchemaAnnotationsJSON `json:"-"`
}

// threatEventTagCategoryEditResponseSchemaAnnotationsJSON contains the JSON
// metadata for the struct [ThreatEventTagCategoryEditResponseSchemaAnnotations]
type threatEventTagCategoryEditResponseSchemaAnnotationsJSON struct {
	Confidence  apijson.Field
	TLP         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryEditResponseSchemaAnnotations) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryEditResponseSchemaAnnotationsJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryEditResponseSchemaEnforcement string

const (
	ThreatEventTagCategoryEditResponseSchemaEnforcementError ThreatEventTagCategoryEditResponseSchemaEnforcement = "error"
	ThreatEventTagCategoryEditResponseSchemaEnforcementWarn  ThreatEventTagCategoryEditResponseSchemaEnforcement = "warn"
	ThreatEventTagCategoryEditResponseSchemaEnforcementOff   ThreatEventTagCategoryEditResponseSchemaEnforcement = "off"
)

func (r ThreatEventTagCategoryEditResponseSchemaEnforcement) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryEditResponseSchemaEnforcementError, ThreatEventTagCategoryEditResponseSchemaEnforcementWarn, ThreatEventTagCategoryEditResponseSchemaEnforcementOff:
		return true
	}
	return false
}

type ThreatEventTagCategoryEditResponseSchemaFormat string

const (
	ThreatEventTagCategoryEditResponseSchemaFormatDate     ThreatEventTagCategoryEditResponseSchemaFormat = "date"
	ThreatEventTagCategoryEditResponseSchemaFormatURL      ThreatEventTagCategoryEditResponseSchemaFormat = "url"
	ThreatEventTagCategoryEditResponseSchemaFormatDuration ThreatEventTagCategoryEditResponseSchemaFormat = "duration"
	ThreatEventTagCategoryEditResponseSchemaFormatCountry  ThreatEventTagCategoryEditResponseSchemaFormat = "country"
)

func (r ThreatEventTagCategoryEditResponseSchemaFormat) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryEditResponseSchemaFormatDate, ThreatEventTagCategoryEditResponseSchemaFormatURL, ThreatEventTagCategoryEditResponseSchemaFormatDuration, ThreatEventTagCategoryEditResponseSchemaFormatCountry:
		return true
	}
	return false
}

type ThreatEventTagCategoryEditResponseSchemaNumberConstraint struct {
	Integer bool                                                         `json:"integer"`
	Max     float64                                                      `json:"max"`
	Min     float64                                                      `json:"min"`
	JSON    threatEventTagCategoryEditResponseSchemaNumberConstraintJSON `json:"-"`
}

// threatEventTagCategoryEditResponseSchemaNumberConstraintJSON contains the JSON
// metadata for the struct
// [ThreatEventTagCategoryEditResponseSchemaNumberConstraint]
type threatEventTagCategoryEditResponseSchemaNumberConstraintJSON struct {
	Integer     apijson.Field
	Max         apijson.Field
	Min         apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *ThreatEventTagCategoryEditResponseSchemaNumberConstraint) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r threatEventTagCategoryEditResponseSchemaNumberConstraintJSON) RawJSON() string {
	return r.raw
}

type ThreatEventTagCategoryNewParams struct {
	// Account ID.
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Name        param.Field[string] `json:"name" api:"required"`
	Description param.Field[string] `json:"description"`
	// Optional array of FieldDefinition objects defining custom fields for tags in
	// this category. Persisted as JSON; returned as a parsed array.
	Schema param.Field[[]ThreatEventTagCategoryNewParamsSchema] `json:"schema"`
}

func (r ThreatEventTagCategoryNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryNewParamsSchema struct {
	Key              param.Field[string]                                                `json:"key" api:"required"`
	Kind             param.Field[ThreatEventTagCategoryNewParamsSchemaKind]             `json:"kind" api:"required"`
	AllowedValues    param.Field[[]string]                                              `json:"allowedValues"`
	Annotations      param.Field[ThreatEventTagCategoryNewParamsSchemaAnnotations]      `json:"annotations"`
	Element          param.Field[interface{}]                                           `json:"element"`
	Enforcement      param.Field[ThreatEventTagCategoryNewParamsSchemaEnforcement]      `json:"enforcement"`
	Format           param.Field[ThreatEventTagCategoryNewParamsSchemaFormat]           `json:"format"`
	Label            param.Field[string]                                                `json:"label"`
	MaxLength        param.Field[int64]                                                 `json:"maxLength"`
	NumberConstraint param.Field[ThreatEventTagCategoryNewParamsSchemaNumberConstraint] `json:"numberConstraint"`
	// Map of property key to FieldDefinition for object fields. Required when kind is
	// 'object'. See FieldDefinition (recursive).
	Properties param.Field[map[string]interface{}] `json:"properties"`
	Required   param.Field[bool]                   `json:"required"`
}

func (r ThreatEventTagCategoryNewParamsSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryNewParamsSchemaKind string

const (
	ThreatEventTagCategoryNewParamsSchemaKindString ThreatEventTagCategoryNewParamsSchemaKind = "string"
	ThreatEventTagCategoryNewParamsSchemaKindNumber ThreatEventTagCategoryNewParamsSchemaKind = "number"
	ThreatEventTagCategoryNewParamsSchemaKindEnum   ThreatEventTagCategoryNewParamsSchemaKind = "enum"
	ThreatEventTagCategoryNewParamsSchemaKindDate   ThreatEventTagCategoryNewParamsSchemaKind = "date"
	ThreatEventTagCategoryNewParamsSchemaKindArray  ThreatEventTagCategoryNewParamsSchemaKind = "array"
	ThreatEventTagCategoryNewParamsSchemaKindObject ThreatEventTagCategoryNewParamsSchemaKind = "object"
)

func (r ThreatEventTagCategoryNewParamsSchemaKind) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryNewParamsSchemaKindString, ThreatEventTagCategoryNewParamsSchemaKindNumber, ThreatEventTagCategoryNewParamsSchemaKindEnum, ThreatEventTagCategoryNewParamsSchemaKindDate, ThreatEventTagCategoryNewParamsSchemaKindArray, ThreatEventTagCategoryNewParamsSchemaKindObject:
		return true
	}
	return false
}

type ThreatEventTagCategoryNewParamsSchemaAnnotations struct {
	Confidence param.Field[bool] `json:"confidence"`
	TLP        param.Field[bool] `json:"tlp"`
}

func (r ThreatEventTagCategoryNewParamsSchemaAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryNewParamsSchemaEnforcement string

const (
	ThreatEventTagCategoryNewParamsSchemaEnforcementError ThreatEventTagCategoryNewParamsSchemaEnforcement = "error"
	ThreatEventTagCategoryNewParamsSchemaEnforcementWarn  ThreatEventTagCategoryNewParamsSchemaEnforcement = "warn"
	ThreatEventTagCategoryNewParamsSchemaEnforcementOff   ThreatEventTagCategoryNewParamsSchemaEnforcement = "off"
)

func (r ThreatEventTagCategoryNewParamsSchemaEnforcement) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryNewParamsSchemaEnforcementError, ThreatEventTagCategoryNewParamsSchemaEnforcementWarn, ThreatEventTagCategoryNewParamsSchemaEnforcementOff:
		return true
	}
	return false
}

type ThreatEventTagCategoryNewParamsSchemaFormat string

const (
	ThreatEventTagCategoryNewParamsSchemaFormatDate     ThreatEventTagCategoryNewParamsSchemaFormat = "date"
	ThreatEventTagCategoryNewParamsSchemaFormatURL      ThreatEventTagCategoryNewParamsSchemaFormat = "url"
	ThreatEventTagCategoryNewParamsSchemaFormatDuration ThreatEventTagCategoryNewParamsSchemaFormat = "duration"
	ThreatEventTagCategoryNewParamsSchemaFormatCountry  ThreatEventTagCategoryNewParamsSchemaFormat = "country"
)

func (r ThreatEventTagCategoryNewParamsSchemaFormat) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryNewParamsSchemaFormatDate, ThreatEventTagCategoryNewParamsSchemaFormatURL, ThreatEventTagCategoryNewParamsSchemaFormatDuration, ThreatEventTagCategoryNewParamsSchemaFormatCountry:
		return true
	}
	return false
}

type ThreatEventTagCategoryNewParamsSchemaNumberConstraint struct {
	Integer param.Field[bool]    `json:"integer"`
	Max     param.Field[float64] `json:"max"`
	Min     param.Field[float64] `json:"min"`
}

func (r ThreatEventTagCategoryNewParamsSchemaNumberConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryListParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
	Search    param.Field[string] `query:"search"`
}

// URLQuery serializes [ThreatEventTagCategoryListParams]'s query parameters as
// `url.Values`.
func (r ThreatEventTagCategoryListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

type ThreatEventTagCategoryDeleteParams struct {
	// Account ID.
	AccountID param.Field[string] `path:"account_id" api:"required"`
}

type ThreatEventTagCategoryEditParams struct {
	// Account ID.
	AccountID   param.Field[string] `path:"account_id" api:"required"`
	Description param.Field[string] `json:"description"`
	Name        param.Field[string] `json:"name"`
	// Optional array of FieldDefinition objects. When provided, replaces the existing
	// field schema. When omitted, the existing schema is preserved.
	Schema param.Field[[]ThreatEventTagCategoryEditParamsSchema] `json:"schema"`
}

func (r ThreatEventTagCategoryEditParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryEditParamsSchema struct {
	Key              param.Field[string]                                                 `json:"key" api:"required"`
	Kind             param.Field[ThreatEventTagCategoryEditParamsSchemaKind]             `json:"kind" api:"required"`
	AllowedValues    param.Field[[]string]                                               `json:"allowedValues"`
	Annotations      param.Field[ThreatEventTagCategoryEditParamsSchemaAnnotations]      `json:"annotations"`
	Element          param.Field[interface{}]                                            `json:"element"`
	Enforcement      param.Field[ThreatEventTagCategoryEditParamsSchemaEnforcement]      `json:"enforcement"`
	Format           param.Field[ThreatEventTagCategoryEditParamsSchemaFormat]           `json:"format"`
	Label            param.Field[string]                                                 `json:"label"`
	MaxLength        param.Field[int64]                                                  `json:"maxLength"`
	NumberConstraint param.Field[ThreatEventTagCategoryEditParamsSchemaNumberConstraint] `json:"numberConstraint"`
	// Map of property key to FieldDefinition for object fields. Required when kind is
	// 'object'. See FieldDefinition (recursive).
	Properties param.Field[map[string]interface{}] `json:"properties"`
	Required   param.Field[bool]                   `json:"required"`
}

func (r ThreatEventTagCategoryEditParamsSchema) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryEditParamsSchemaKind string

const (
	ThreatEventTagCategoryEditParamsSchemaKindString ThreatEventTagCategoryEditParamsSchemaKind = "string"
	ThreatEventTagCategoryEditParamsSchemaKindNumber ThreatEventTagCategoryEditParamsSchemaKind = "number"
	ThreatEventTagCategoryEditParamsSchemaKindEnum   ThreatEventTagCategoryEditParamsSchemaKind = "enum"
	ThreatEventTagCategoryEditParamsSchemaKindDate   ThreatEventTagCategoryEditParamsSchemaKind = "date"
	ThreatEventTagCategoryEditParamsSchemaKindArray  ThreatEventTagCategoryEditParamsSchemaKind = "array"
	ThreatEventTagCategoryEditParamsSchemaKindObject ThreatEventTagCategoryEditParamsSchemaKind = "object"
)

func (r ThreatEventTagCategoryEditParamsSchemaKind) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryEditParamsSchemaKindString, ThreatEventTagCategoryEditParamsSchemaKindNumber, ThreatEventTagCategoryEditParamsSchemaKindEnum, ThreatEventTagCategoryEditParamsSchemaKindDate, ThreatEventTagCategoryEditParamsSchemaKindArray, ThreatEventTagCategoryEditParamsSchemaKindObject:
		return true
	}
	return false
}

type ThreatEventTagCategoryEditParamsSchemaAnnotations struct {
	Confidence param.Field[bool] `json:"confidence"`
	TLP        param.Field[bool] `json:"tlp"`
}

func (r ThreatEventTagCategoryEditParamsSchemaAnnotations) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type ThreatEventTagCategoryEditParamsSchemaEnforcement string

const (
	ThreatEventTagCategoryEditParamsSchemaEnforcementError ThreatEventTagCategoryEditParamsSchemaEnforcement = "error"
	ThreatEventTagCategoryEditParamsSchemaEnforcementWarn  ThreatEventTagCategoryEditParamsSchemaEnforcement = "warn"
	ThreatEventTagCategoryEditParamsSchemaEnforcementOff   ThreatEventTagCategoryEditParamsSchemaEnforcement = "off"
)

func (r ThreatEventTagCategoryEditParamsSchemaEnforcement) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryEditParamsSchemaEnforcementError, ThreatEventTagCategoryEditParamsSchemaEnforcementWarn, ThreatEventTagCategoryEditParamsSchemaEnforcementOff:
		return true
	}
	return false
}

type ThreatEventTagCategoryEditParamsSchemaFormat string

const (
	ThreatEventTagCategoryEditParamsSchemaFormatDate     ThreatEventTagCategoryEditParamsSchemaFormat = "date"
	ThreatEventTagCategoryEditParamsSchemaFormatURL      ThreatEventTagCategoryEditParamsSchemaFormat = "url"
	ThreatEventTagCategoryEditParamsSchemaFormatDuration ThreatEventTagCategoryEditParamsSchemaFormat = "duration"
	ThreatEventTagCategoryEditParamsSchemaFormatCountry  ThreatEventTagCategoryEditParamsSchemaFormat = "country"
)

func (r ThreatEventTagCategoryEditParamsSchemaFormat) IsKnown() bool {
	switch r {
	case ThreatEventTagCategoryEditParamsSchemaFormatDate, ThreatEventTagCategoryEditParamsSchemaFormatURL, ThreatEventTagCategoryEditParamsSchemaFormatDuration, ThreatEventTagCategoryEditParamsSchemaFormatCountry:
		return true
	}
	return false
}

type ThreatEventTagCategoryEditParamsSchemaNumberConstraint struct {
	Integer param.Field[bool]    `json:"integer"`
	Max     param.Field[float64] `json:"max"`
	Min     param.Field[float64] `json:"min"`
}

func (r ThreatEventTagCategoryEditParamsSchemaNumberConstraint) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
