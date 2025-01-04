// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package d1

import (
	"time"

	"github.com/cloudflare/cloudflare-go/v4/internal/apijson"
	"github.com/cloudflare/cloudflare-go/v4/option"
)

// D1Service contains methods and other services that help with interacting with
// the cloudflare API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewD1Service] method instead.
type D1Service struct {
	Options  []option.RequestOption
	Database *DatabaseService
}

// NewD1Service generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewD1Service(opts ...option.RequestOption) (r *D1Service) {
	r = &D1Service{}
	r.Options = opts
	r.Database = NewDatabaseService(opts...)
	return
}

type D1 struct {
	// Specifies the timestamp the resource was created as an ISO8601 string.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// The D1 database's size, in bytes.
	FileSize  float64 `json:"file_size"`
	Name      string  `json:"name"`
	NumTables float64 `json:"num_tables"`
	UUID      string  `json:"uuid"`
	Version   string  `json:"version"`
	JSON      d1JSON  `json:"-"`
}

// d1JSON contains the JSON metadata for the struct [D1]
type d1JSON struct {
	CreatedAt   apijson.Field
	FileSize    apijson.Field
	Name        apijson.Field
	NumTables   apijson.Field
	UUID        apijson.Field
	Version     apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *D1) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r d1JSON) RawJSON() string {
	return r.raw
}
