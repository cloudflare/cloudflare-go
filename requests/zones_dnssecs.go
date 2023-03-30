package requests

import (
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-sdk-go/core"
	pjson "github.com/cloudflare/cloudflare-sdk-go/core/json"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
)

type DeleteDnssecResponseSingle struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	Result   fields.Field[string]     `json:"result"`
	// Whether the API call was successful
	Success fields.Field[DeleteDnssecResponseSingleSuccess] `json:"success"`
}

// MarshalJSON serializes DeleteDnssecResponseSingle into an array of bytes using
// the gjson library. Members of the `jsonFields` field are serialized into the
// top-level, and will overwrite known members of the same name.
func (r *DeleteDnssecResponseSingle) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DeleteDnssecResponseSingle) String() (result string) {
	return fmt.Sprintf("&DeleteDnssecResponseSingle{Errors:%s Messages:%s Result:%s Success:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Result, r.Success)
}

type DeleteDnssecResponseSingleSuccess bool

const (
	DeleteDnssecResponseSingleSuccessTrue DeleteDnssecResponseSingleSuccess = true
)

type DnssecSingle struct {
	Errors   fields.Field[[]Messages] `json:"errors"`
	Messages fields.Field[[]Messages] `json:"messages"`
	Result   fields.Field[Dnssec]     `json:"result"`
	// Whether the API call was successful
	Success fields.Field[DnssecSingleSuccess] `json:"success"`
}

// MarshalJSON serializes DnssecSingle into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DnssecSingle) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DnssecSingle) String() (result string) {
	return fmt.Sprintf("&DnssecSingle{Errors:%s Messages:%s Result:%s Success:%s}", core.Fmt(r.Errors), core.Fmt(r.Messages), r.Result, r.Success)
}

type Dnssec struct {
	// Algorithm key code.
	Algorithm fields.Field[string] `json:"algorithm,nullable"`
	// Digest hash.
	Digest fields.Field[string] `json:"digest,nullable"`
	// Type of digest algorithm.
	DigestAlgorithm fields.Field[string] `json:"digest_algorithm,nullable"`
	// Coded type for digest algorithm.
	DigestType fields.Field[string] `json:"digest_type,nullable"`
	// Full DS record.
	Ds fields.Field[string] `json:"ds,nullable"`
	// Flag for DNSSEC record.
	Flags fields.Field[float64] `json:"flags,nullable"`
	// Code for key tag.
	KeyTag fields.Field[float64] `json:"key_tag,nullable"`
	// Algorithm key type.
	KeyType fields.Field[string] `json:"key_type,nullable"`
	// When DNSSEC was last modified.
	ModifiedOn fields.Field[time.Time] `json:"modified_on,nullable" format:"date-time"`
	// Public key for DS record.
	PublicKey fields.Field[string] `json:"public_key,nullable"`
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status fields.Field[Status_0xnKqZfH] `json:"status"`
}

// MarshalJSON serializes Dnssec into an array of bytes using the gjson library.
// Members of the `jsonFields` field are serialized into the top-level, and will
// overwrite known members of the same name.
func (r *Dnssec) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r Dnssec) String() (result string) {
	return fmt.Sprintf("&Dnssec{Algorithm:%s Digest:%s DigestAlgorithm:%s DigestType:%s Ds:%s Flags:%s KeyTag:%s KeyType:%s ModifiedOn:%s PublicKey:%s Status:%s}", r.Algorithm, r.Digest, r.DigestAlgorithm, r.DigestType, r.Ds, r.Flags, r.KeyTag, r.KeyType, r.ModifiedOn, r.PublicKey, r.Status)
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

type DnssecUpdateParams struct {
	// Status of DNSSEC, based on user-desired state and presence of necessary records.
	Status fields.Field[DnssecUpdateParamsStatus] `json:"status,required"`
}

// MarshalJSON serializes DnssecUpdateParams into an array of bytes using the gjson
// library. Members of the `jsonFields` field are serialized into the top-level,
// and will overwrite known members of the same name.
func (r *DnssecUpdateParams) MarshalJSON() (data []byte, err error) {
	return pjson.MarshalRoot(r)
}

func (r DnssecUpdateParams) String() (result string) {
	return fmt.Sprintf("&DnssecUpdateParams{Status:%s}", r.Status)
}

type DnssecUpdateParamsStatus string

const (
	DnssecUpdateParamsStatusActive   DnssecUpdateParamsStatus = "active"
	DnssecUpdateParamsStatusDisabled DnssecUpdateParamsStatus = "disabled"
)
