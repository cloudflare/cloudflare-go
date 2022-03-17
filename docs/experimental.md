# README (experimental)

An experimental and incremental update of the library focusing on a more modern
and consistent experience.

## Improvements

### Consistent CRUD method signatures

Majority of entities follow a standard method signature.

- `Get(ctx, id)`: fetches a single entity by an identifer
- `List(ctx, ...params)`: fetches all entities and automatically paginates
- `Create(ctx, ...params)`: creates a new entity with the provided parameters
- `Update(ctx, id, ...params)`: updates an existing entity
- `Delete(ctx, id)`: deletes a single entity

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

**initialising a new client with options like your own `http.Client`**

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
  HTTPClient: myCustomHTTPClient,
  // ...
}
c, err := cloudflare.NewExperimental(params)
```

**create a new zone**

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

**fetching a known zone ID**

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

z, _ := c.Zones.Get(ctx, "3e7705498e8be60520841409ebc69bc1")
```

**fetching all zones matching a single account ID**

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

zParams := &cloudflare.ZoneParams{
  AccountID: "d8e8fca2dc0f896fd7cb4cb0031ba249"
}
z, _ := c.Zones.List(ctx, zParams)
```

**update a zone**

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

zParams := &cloudflare.ZoneParams{
  Nameservers: cloudflare.StringSlice([]string{
    "ns1.example.com",
    "ns2.example.com"
  })
}
z, _ := c.Zones.Update(ctx, "b5163cf270a3fbac34827c4a2713eef4", zParams)
```

**delete a zone**

```go
params := cloudflare.ClientParams{
  Key: "3bc3be114fb6323adc5b0ad7422d193a",
  Email: "someone@example.com",
}
c, err := cloudflare.NewExperimental(params)

z, _ := c.Zones.Delete(ctx, "b5163cf270a3fbac34827c4a2713eef4")
```
