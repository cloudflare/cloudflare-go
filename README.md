# Cloudflare Go API Library

<a href="https://pkg.go.dev/github.com/cloudflare/cloudflare-sdk-go"><img src="https://pkg.go.dev/badge/github.com/cloudflare/cloudflare-sdk-go.svg" alt="Go Reference"></a>

The Cloudflare Go library provides convenient access to the Cloudflare REST
API from applications written in Go.

## Installation

Within a Go module, you can just import this package and let the Go compiler
resolve this module.

```go
import (
	"github.com/cloudflare/cloudflare-sdk-go" // imported as cloudflare
)
```

Or, explicitly import this package with

```
go get -u 'github.com/cloudflare/cloudflare-sdk-go'
```

## Documentation

The API documentation can be found [here](https://developers.cloudflare.com/api/).

## Usage

```go
package main

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-sdk-go"
	"github.com/cloudflare/cloudflare-sdk-go/fields"
	"github.com/cloudflare/cloudflare-sdk-go/requests"
)

func main() {
	client := cloudflare.NewCloudflare()
	res, err := client.Zones.New(context.TODO(), &requests.ZoneNewParams{
		Account: fields.F(requests.ZoneNewParamsAccount{
			ID: fields.F("023e105f4ecef8ad9ca31a8372d0c353"), Name: "example.com", Type: "full",
		}),
	})
	if err != nil {
		panic(err)
	}
	fmt.Sprintf("%s", res.Result.ID)
}

```

### Request Fields

Types for requests look like the following:

```go
type FooParams struct {
	ID     fields.Field[string] `json:"id,required"`
	Number fields.Field[int64]  `json:"number,required"`
	Name   fields.Field[string] `json:"name"`
	Other  fields.Field[Bar]    `json:"other"`
}

type Bar struct {
	Number fields.Field[int64]  `json:"number"`
	Name   fields.Field[string] `json:"name"`
}
```

For each field, you can either supply a value field with `fields.F(...)`, a
`null` value with `fields.NullField()`, or some raw JSON value with
`fields.RawField(...)` that you specify as a byte slice. If you do not supply a
value, then we do not populate the field. An example request may look like

```go
params := &FooParams{
	// Normally populates this field as `"id": "food_id"`
	ID: fields.F("foo_id"),

	// Integer helper casts integer values and literals to fields.Field[int64]
	Number: fields.Int(12),

	// Explicitly sends this field as null, e.g., `"name": null`
	Name: fields.NullField[string](),

	// Overrides this field as `"other": "ovveride_this_field"`
	Other: fields.RawField[Bar]("override_this_field")
}
```

If you want to add or override some field in the JSON, then you can use the `options.WithJSONSet(key string, value interface{})` RequestOption, which you can read more about [here](#requestoptions). Internally, this uses 'github.com/tidwall/sjson' library, so you can compose complex access as seen [here](https://github.com/tidwall/sjson).

### Response Objects

Response objects in this SDK have value type members. Accessing properties on
response objects is as simple as:

```go
res, err := client.Service.Foo(context.TODO())
res.Name // is just some string value
```

If null, not present, or invalid, all fields will simply be their empty values.

If you want to be able to tell that the value is either `null`, not present, or
invalid, we provide metadata in the `JSON` property of every response object.

```go
// This is true if `name` is _either_ not present or explicitly null
res.JSON.Name.IsNull()

// This is true if `name` is not present
res.JSON.Name.IsMissing()

// This is true if `name` is present, but not coercable
res.JSON.Name.IsMissing()

// If needed, you can access the Raw JSON value of the field by accessing
res.JSON.Name.Raw()
```

You can also access the JSON value of the entire object with `res.JSON.Raw`.

There may be instances where we provide experimental or private API features
for some customers. In those cases, the related features will not be exposed to
the SDK as typed fields, and are instead deserialized to an internal map. We
provide methods to get and set these json fields in API objects.

```go
// Access the JSON value as
body := res.JSON.Extras["extra_field"].Raw()

// You can `Unmarshal` the JSON into a struct as needed
custom := struct{A string, B int64}{}
json.Unmarshal(body, &custom)
```

### RequestOptions

This library uses the functional options pattern. `RequestOptions` are closures
that mutate a `RequestConfig`. These options can be supplied to the client or
at individual requests, and they will cascade appropriately.

At each request, these closures will be run in the order that they are
supplied, after the defaults for that particular request.

For example:

```go
client := Cloudflare.NewCloudflare(
	// Adds header to every request made by client
	options.WithHeader("X-Some-Header", "custom_header_info"),

	// Overrides APIkey read from environment
	options.WithAPIKey("api_key"),
)

client.Zones.New(
	context.TODO(),
	...,
	// These options override the client options
	options.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	options.WithAPIKey("other_api_key"),
)
```

### Pagination

List methods in the Cloudflare API are paginated.

In addition to exposing standard response values, this library provides an
auto-paginating iterator with each list response, so you do not have to request
successive pages manually:

To iterate over all elements in a paginated list, we call the paginated
endpoint which returns a `Page`, then iterate while there is a `Next()`
element. When `Next()` is called, it will load the next element into
`Current()`, which is also aliased to the endpoint-specific name. If we are at
the end of the current page, `Next()` will fetch the next page.

```go
page, err := <Service>(
	context.TODO(),
	...,
)
if err != nil {
	panic(err.Error())
}
for page.Next() {
	item := page.Current()
	// ...
}
if page.Err() != nil {
	panic(page.Err().Error())
}
```

### Errors

For the errors generated by the SDK, we provide extra convenience methods for debugging.

### Middleware

You may apply any middleware you wish by overriding the `http.Client` with
`options.WithClient(client)`. An example of a basic logging middleware is given
below:

```go
TODO
```

## Status

This package is in beta. Its internals and interfaces are not stable and
subject to change without a major semver bump; please reach out if you rely on
any undocumented behavior.

We are keen for your feedback; please email us at
[dev-feedback@todo.com](mailto:dev-feedback@todo.com) or open an issue with questions, bugs, or
suggestions.

## Requirements

This library requires Go 1.18+.