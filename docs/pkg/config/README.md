# config
--
    import "."


## Usage

```go
const (
	// US represents New Relic's US-based production deployment.
	US = iota

	// EU represents New Relic's EU-based production deployment.
	EU

	// Staging represents New Relic's US-based staging deployment.
	// This is for internal New Relic use only.
	Staging
)
```

```go
var Region = struct {
	US      RegionType
	EU      RegionType
	Staging RegionType
}{
	US:      US,
	EU:      EU,
	Staging: Staging,
}
```
Region specifies the New Relic environment to target.

#### type Config

```go
type Config struct {
	BaseURL       string
	APIKey        string
	Timeout       *time.Duration
	HTTPTransport *http.RoundTripper
	UserAgent     string
	Region        RegionType
}
```

Config contains all the configuration data for the API Client.

#### type RegionType

```go
type RegionType int
```

RegionType represents the members of the Region enumeration.