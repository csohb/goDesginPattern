package main

import "goDesginPattern/composite/obj"

func main() {
	scooter1 := &obj.Scooter{ModelName: "max pro"}
	scooter2 := &obj.Scooter{ModelName: "max plus"}
	scooter3 := &obj.Scooter{ModelName: "K2"}

	segwayIoT := &obj.IoT{
		IoTName: "Segway IoT",
	}

	segwayIoT.AddIot(scooter1)
	segwayIoT.AddIot(scooter2)

	msIoT := &obj.IoT{
		IoTName: "Microworks IoT",
	}

	gbikeIoT := &obj.IoT{
		IoTName: "Gbike IoT",
	}

	gbikeIoT.AddIot(segwayIoT)
	gbikeIoT.AddIot(msIoT)

	msIoT.AddIot(scooter3)

	gbikeIoT.Produce()

}
