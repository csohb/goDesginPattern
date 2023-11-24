package scooterStatus

import (
	"fmt"
)

type ScooterStatus struct {
	ObserverList       []Observer
	Coordinates        Coordinates
	BatteryPercent     int
	IsGeofence         bool
	IsBatteryLow       bool
	IsTripJurisdiction bool
}

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

func NewScooterStatus(lng, lat float64, batteryPercent int) *ScooterStatus {
	return &ScooterStatus{
		Coordinates: Coordinates{
			Longitude: lng,
			Latitude:  lat,
		},
		BatteryPercent: batteryPercent,
	}
}

func (s *ScooterStatus) UpdateScooterBatteryStatus(batteryPercent int, lng, lat float64) {
	fmt.Printf("scooter battery is changed to %d\n", batteryPercent)
	s.BatteryPercent = batteryPercent
	if s.BatteryPercent <= 20 {
		fmt.Println("scooter battery is lower than 20%")
		s.IsBatteryLow = true
	}
	fmt.Printf("scooter Coordinates are changed - lng:%f, lat:%f\n", lng, lat)
	s.Coordinates = Coordinates{
		Longitude: lng,
		Latitude:  lat,
	}
	if lng == 127.1234 {
		s.IsGeofence = true
	}

	if lat == 37.123 {
		s.IsTripJurisdiction = true
	}

	s.NotifyAll()
}

func (s *ScooterStatus) Register(o Observer) {
	fmt.Printf("%s Processor is registered!\n", o.GetProcessorName())
	s.ObserverList = append(s.ObserverList, o)
}

func (s *ScooterStatus) Deregister(o Observer) {
	s.ObserverList = removeFromSlice(s.ObserverList, o)
}

func (s *ScooterStatus) NotifyAll() {
	for _, observer := range s.ObserverList {
		observer.Update(s.BatteryPercent, s.Coordinates)
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
