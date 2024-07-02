// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceBuilder(t *testing.T) {
	for _, test := range []string{"default", "all_set", "none_set"} {
		t.Run(test, func(t *testing.T) {
			cfg := loadResourceAttributesConfig(t, test)
			rb := NewResourceBuilder(cfg)
			rb.SetDatabase("database-val")
			rb.SetServerAddress("server.address-val")
			rb.SetServerPort(11)

			res := rb.Emit()
			assert.Equal(t, 0, rb.Emit().Attributes().Len()) // Second call should return empty Resource

			switch test {
			case "default":
				assert.Equal(t, 2, res.Attributes().Len())
			case "all_set":
				assert.Equal(t, 3, res.Attributes().Len())
			case "none_set":
				assert.Equal(t, 0, res.Attributes().Len())
				return
			default:
				assert.Failf(t, "unexpected test case: %s", test)
			}

			val, ok := res.Attributes().Get("database")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "database-val", val.Str())
			}
			val, ok = res.Attributes().Get("server.address")
			assert.True(t, ok)
			if ok {
				assert.EqualValues(t, "server.address-val", val.Str())
			}
			val, ok = res.Attributes().Get("server.port")
			assert.Equal(t, test == "all_set", ok)
			if ok {
				assert.EqualValues(t, 11, val.Int())
			}
		})
	}
}
