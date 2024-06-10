package maxmind

import (
	"context"
	"errors"
	"fmt"
	"net"

	conventions "github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/convention"
	"github.com/open-telemetry/opentelemetry-collector-contrib/processor/geoipprocessor/internal/provider"
	"github.com/oschwald/geoip2-golang"
	"go.opentelemetry.io/otel/attribute"
)

var (
	// defaultLocale specifies English as the default Geolocation name, see https://dev.maxmind.com/geoip/docs/web-services/responses#languages
	defaultLocale      = "en"
	geoIP2CityDBType   = "GeoIP2-City"
	geoLite2CityDBType = "GeoLite2-City"

	errUnsupportedDB = errors.New("unsupported geo IP database type")
)

type maxMindProvider struct {
	geoReader *geoip2.Reader
	// language/locale to be used in name retrieval
	name string
}

var _ provider.GeoIPProvider = (*maxMindProvider)(nil)

func newMaxMindProvider(cfg *Config) (*maxMindProvider, error) {
	geoReader, err := geoip2.Open(cfg.GeoIPDatabasePath)
	if err != nil {
		return nil, err
	}

	return &maxMindProvider{geoReader: geoReader, name: defaultLocale}, nil
}

// Location implements provider.GeoIPProvider.
func (g *maxMindProvider) Location(_ context.Context, ipAddress net.IP) (attribute.Set, error) {
	switch g.geoReader.Metadata().DatabaseType {
	case geoIP2CityDBType, geoLite2CityDBType:
		attrs, err := g.cityAttributes(ipAddress)
		if err != nil {
			return attribute.Set{}, err
		}
		return attribute.NewSet(*attrs...), nil
	default:
		return attribute.Set{}, fmt.Errorf("%w type: %s", errUnsupportedDB, g.geoReader.Metadata().DatabaseType)
	}
}

func (g *maxMindProvider) cityAttributes(ipAddress net.IP) (*[]attribute.KeyValue, error) {
	attributes := make([]attribute.KeyValue, 0, 11)

	city, err := g.geoReader.City(ipAddress)
	if err != nil {
		return nil, err
	}

	// The exact set of top-level keys varies based on the particular GeoIP2 web service you are using. If a key maps to an undefined or empty value, it is not included in the JSON object. The following anonymous function appends the given key-value only if the value is not empty.
	appendIfNotEmpty := func(keyName, value string) {
		if value != "" {
			attributes = append(attributes, attribute.String(keyName, value))
		}
	}

	// city
	appendIfNotEmpty(conventions.AttributeGeoCityName, city.City.Names[g.name])
	// country
	appendIfNotEmpty(conventions.AttributeGeoCountryName, city.Country.Names[g.name])
	appendIfNotEmpty(conventions.AttributeGeoCountryIsoCode, city.Country.IsoCode)
	// continent
	appendIfNotEmpty(conventions.AttributeGeoContinentName, city.Continent.Names[g.name])
	appendIfNotEmpty(conventions.AttributeGeoContinentCode, city.Continent.Code)
	// postal code
	appendIfNotEmpty(conventions.AttributeGeoPostalCode, city.Postal.Code)
	// region
	if len(city.Subdivisions) > 0 {
		// The most specific subdivision is located at the last array position, see https://github.com/maxmind/GeoIP2-java/blob/2fe4c65424fed2c3c2449e5530381b6452b0560f/src/main/java/com/maxmind/geoip2/model/AbstractCityResponse.java#L112
		mostSpecificSubdivision := city.Subdivisions[len(city.Subdivisions)-1]
		appendIfNotEmpty(conventions.AttributeGeoRegionName, mostSpecificSubdivision.Names[g.name])
		appendIfNotEmpty(conventions.AttributeGeoRegionIsoCode, mostSpecificSubdivision.IsoCode)
	}

	// location
	appendIfNotEmpty(conventions.AttributeGeoTimezone, city.Location.TimeZone)
	if city.Location.Latitude != 0 && city.Location.Longitude != 0 {
		attributes = append(attributes, attribute.Float64(conventions.AttributeGeoLocationLat, city.Location.Latitude), attribute.Float64(conventions.AttributeGeoLocationLon, city.Location.Longitude))
	}

	return &attributes, err
}
