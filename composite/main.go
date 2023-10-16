package main

import "goDesginPattern/composite/obj"

func main() {
	scooter1 := &obj.Scooter{ModelName: "max pro"}
	scooter2 := &obj.Scooter{ModelName: "max plus"}
	scooter3 := &obj.Scooter{ModelName: "K2"}

	segway := &obj.Company{
		IoTName: "Segway",
	}

	segway.AddPM(scooter1)
	segway.AddPM(scooter2)

	msIoT := &obj.Company{
		IoTName: "Microworks",
	}

	gbike := &obj.Company{
		IoTName: "Gbike",
	}

	gbike.AddPM(segway)
	gbike.AddPM(msIoT)

	msIoT.AddPM(scooter3)

	gbike.Produce()

}
