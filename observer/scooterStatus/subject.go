package scooterStatus

import (
	"fmt"
)

type ScooterStatus struct {
	observerList       []Observer
	coordinates        Coordinates
	batteryPercent     int
	isGeofence         bool
	isBatteryLow       bool
	isTripJurisdiction bool
}

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

func NewScooterStatus(lng, lat float64, batteryPercent int) *ScooterStatus {
	return &ScooterStatus{
		coordinates: Coordinates{
			Longitude: lng,
			Latitude:  lat,
		},
		batteryPercent: batteryPercent,
	}
}

func (s *ScooterStatus) UpdateScooterBatteryStatus(batteryPercent int, lng, lat float64) {
	fmt.Printf("scooter battery is changed to %d\n", batteryPercent)
	s.batteryPercent = batteryPercent
	if s.batteryPercent <= 20 {
		fmt.Println("scooter battery is lower than 20%")
		s.isBatteryLow = true
	}
	fmt.Printf("scooter coordinates are changed - lng:%f, lat:%f\n", lng, lat)
	s.coordinates = Coordinates{
		Longitude: lng,
		Latitude:  lat,
	}
	if lng == 127.1234 {
		s.isGeofence = true
	}

	if lat == 37.123 {
		s.isTripJurisdiction = true
	}

	s.NotifyAll()
}

func (s *ScooterStatus) Register(o Observer) {
	fmt.Printf("%s Processor is registered!\n", o.GetProcessorName())
	s.observerList = append(s.observerList, o)
}

func (s *ScooterStatus) Deregister(o Observer) {
	s.observerList = removeFromSlice(s.observerList, o)
}

func (s *ScooterStatus) NotifyAll() {
	for _, observer := range s.observerList {
		observer.Update(s.batteryPercent, s.coordinates)
	}
}

func removeFromSlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.GetProcessorName() == observer.GetProcessorName() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}
