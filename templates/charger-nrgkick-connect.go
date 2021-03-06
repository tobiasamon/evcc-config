package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "charger",
		Type:   "nrgkick-connect",
		Name:   "NRGKick Connect",
		Sample: `uri: http://192.168.1.4
mac: 00:99:22 # MAC address
password: # password`,
	}

	registry.Add(template)
}
