# README (experimental)

An experimental and incremental update of the library focusing on a more modern
and consistent experience.

## Improvements

### Consistent CRUD method signatures

Majority of entities follow a standard method signature (where `$entity` is the
resource such as `Zone`, `DNSRecord`, ...).

| Signature                                             | Purpose                                                                                                                              | Return value                                                                                                                   |
| ----------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `Get(ctx, *Resource) ($entity, error)`                | Fetches a single entity by a `Resource`. Should use `ZoneIdentifier` and `AccountIdentifier` respectively (not \*Resource directly). | Returns the entity and `error` based on the listing parameters.                                                                |
| `List(ctx, ...params) ([]$entity, ResultInfo, error)` | Fetches all entities.                                                                                                                | Returns the list of matching entities, the result information (pagination, fail/success and additional metadata), and `error`. |
| `New(ctx, ...params) ($entity, error)`                | Creates a new entity with the provided parameters.                                                                                   | Returns the newly created entity and `error`.                                                                                  |
| `Update(ctx, ...params) ($entity, error)`             | Updates an existing entity.                                                                                                          | Returns the updated entity and `error`.                                                                                        |
| `Delete(ctx, *Resource) (error)`                      | Deletes a single entity by a `Resource`. Should use `ZoneIdentifier` and `AccountIdentifier` respectively (not \*Resource directly). | Returns `error`.                                                                                                               |

### Why no iterator for `List` operations?

Initially, there was proposal to either: 1) automatically paginate all results
or 2) return an `Iterator` object with `Next()` methods. I've opted to instead
return the pagination information to the operator as using an `Iterator` object
has some hidden complexities when attempting to cover all use cases. Some
examples:

- forward and backwards iteration
- further filtering the iterations
- only returning subsets

## Nested methods and services

Not all methods are defined at the top level. Instead, they are nested under
service objects.

```golang
// old
client.ListZones(...)
client.ZoneLevelAccessServiceTokens(...)

// new
client.Zones.List()
client.Access.ServiceTokens(...)
```

This avoids polluting the global namespace and having more specific methods
for services.

## Examples

A zone is used below for the examples however, all entites will implement the
same methods and interfaces.

### Initialising a new client with options like your own `http.Client`

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
  HTTPClient: myCustomHTTPClient,
  // ...
}
c, err := cloudflare.NewExperimental(params)
```

### Create a new zone

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

zParams := &cloudflare.ZoneParams{
  Name: "example.com",
  AccountID: "d8e8fca2dc0f896fd7cb4cb0031ba249"
}
z, _ := c.Zones.New(ctx, zParams)
```

### Fetching a known zone ID

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

z, _ := c.Zones.Get(ctx, cloudflare.ZoneIdentifier("3e7705498e8be60520841409ebc69bc1"))
```

### Fetching all zones matching a single account ID

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

zParams := &cloudflare.ZoneParams{
  AccountID: "d8e8fca2dc0f896fd7cb4cb0031ba249"
}
z, _, _ := c.Zones.List(ctx, zParams)
```

### Update a zone

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

zParams := &cloudflare.ZoneParams{
  ID: "3e7705498e8be60520841409ebc69bc1",
  Nameservers: cloudflare.StringSlice([]string{
    "ns1.example.com",
    "ns2.example.com"
  })
}
z, _ := c.Zones.Update(ctx, zParams)
```

### Delete a zone

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

z, _ := c.Zones.Delete(ctx, cloudflare.ZoneIdentifier("b5163cf270a3fbac34827c4a2713eef4"))
```
