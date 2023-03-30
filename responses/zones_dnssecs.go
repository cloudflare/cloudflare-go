package responses

import (
	"time"

	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
)

type DeleteDnssecResponseSingle struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	Result   string     `json:"result"`
	// Whether the API call was successful
	Success DeleteDnssecResponseSingleSuccess `json:"success"`
	JSON    DeleteDnssecResponseSingleJSON
}

type DeleteDnssecResponseSingleJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Result   pjson.Metadata
	Success  pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DeleteDnssecResponseSingle
// using the internal pjson library. Unrecognized fields are stored in the
// `jsonFields` property.
func (r *DeleteDnssecResponseSingle) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type DeleteDnssecResponseSingleSuccess bool

const (
	DeleteDnssecResponseSingleSuccessTrue DeleteDnssecResponseSingleSuccess = true
)

type DnssecSingle struct {
	Errors   []Messages `json:"errors"`
	Messages []Messages `json:"messages"`
	Result   Dnssec     `json:"result"`
	// Whether the API call was successful
	Success DnssecSingleSuccess `json:"success"`
	JSON    DnssecSingleJSON
}

type DnssecSingleJSON struct {
	Errors   pjson.Metadata
	Messages pjson.Metadata
	Result   pjson.Metadata
	Success  pjson.Metadata
	Raw      []byte
	Extras   map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into DnssecSingle using the
// internal pjson library. Unrecognized fields are stored in the `jsonFields`
// property.
func (r *DnssecSingle) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Dnssec struct {
	// Algorithm key code.
	Algorithm string `json:"algorithm,nullable"`
	// Digest hash.
	Digest string `json:"digest,nullable"`
	// Type of digest algorithm.
	DigestAlgorithm string `json:"digest_algorithm,nullable"`
	// Coded type for digest algorithm.
	DigestType string `json:"digest_type,nullable"`
	// Full DS record.
	Ds string `json:"ds,nullable"`
	// Flag for DNSSEC record.
	Flags float64 `json:"flags,nullable"`
	// Code for key tag.
	KeyTag float64 `json:"key_tag,nullable"`
	// Algorithm key type.
	KeyType string `json:"key_type,nullable"`
	// When DNSSEC was last modified.
	ModifiedOn time.Time `json:"modified_on,nullable" format:"date-time"`
	// Public key for DS record.
	PublicKey string `json:"public_key,nullable"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status Status_0xnKqZfH `json:"status"`
	JSON   DnssecJSON
}

type DnssecJSON struct {
	Algorithm       pjson.Metadata
	Digest          pjson.Metadata
	DigestAlgorithm pjson.Metadata
	DigestType      pjson.Metadata
	Ds              pjson.Metadata
	Flags           pjson.Metadata
	KeyTag          pjson.Metadata
	KeyType         pjson.Metadata
	ModifiedOn      pjson.Metadata
	PublicKey       pjson.Metadata
	Status          pjson.Metadata
	Raw             []byte
	Extras          map[string]pjson.Metadata
}

// UnmarshalJSON deserializes the provided bytes into Dnssec using the internal
// pjson library. Unrecognized fields are stored in the `jsonFields` property.
func (r *Dnssec) UnmarshalJSON(data []byte) (err error) {
	return pjson.UnmarshalRoot(data, r)
}

type Status_0xnKqZfH string

const (
	Status_0xnKqZfHActive          Status_0xnKqZfH = "active"
	Status_0xnKqZfHPending         Status_0xnKqZfH = "pending"
	Status_0xnKqZfHDisabled        Status_0xnKqZfH = "disabled"
	Status_0xnKqZfHPendingDisabled Status_0xnKqZfH = "pending-disabled"
	Status_0xnKqZfHError           Status_0xnKqZfH = "error"
)

type DnssecSingleSuccess bool

const (
	DnssecSingleSuccessTrue DnssecSingleSuccess = true
)
