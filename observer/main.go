package main

import (
	"goDesginPattern/observer/scooterStatus"
)

func main() {

	lng := 127.1111
	lat := 37.1111
	battery := 80
	status := scooterStatus.NewScooterStatus(lng, lat, battery)

	geofenceProcessor := &scooterStatus.Processor{
		Name: "geofence",
	}

	batteryLowProcessor := &scooterStatus.Processor{
		Name: "battery-low",
	}

	tripJurisdictionProcessor := &scooterStatus.Processor{
		Name: "trip-jurisdiction",
	}

	status.Register(geofenceProcessor)
	status.Register(batteryLowProcessor)
	status.Register(tripJurisdictionProcessor)

	status.UpdateScooterBatteryStatus(19, 127.1234, 37.1234)

}
