package templates 

import (
	"github.com/andig/evcc-config/registry"
)

func init() {
	template := registry.Template{
		Class:  "meter",
		Type:   "modbus",
		Name:   "Modbus (RTU)",
		Sample: `model: sdm
uri: rs485.fritz.box:23
rtu: true # rs485 device connected using ethernet adapter
id: 2
power: Power # default value, optionally override
energy: Sum # energy value (Zählerstand)`,
	}

	registry.Add(template)
}
