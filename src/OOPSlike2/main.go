package main

import ("fmt")

type vehicle struct {
	vType string
}
type landVehicle struct {
	vehicle 
	numWheels int
}
type airVehicle struct {
	vehicle
	numLeaf int
}
func (av airVehicle) running() {
	fmt.Printf("This is: %v flying on %v leaves\n",av.vehicle.vType, av.numLeaf)
}
func (lv landVehicle) running() {
	fmt.Printf("This is: %v riding on %v wheels\n",lv.vehicle.vType, lv.numWheels)
}

type run interface{
	running()
}

func start(v run){
	v.running()
}

func main() {
	v := vehicle {
		vType : "Air Vehicle",
	}
	fmt.Println(v.vType)
	
	lv := landVehicle {
		vehicle : vehicle {
			vType : "Land Vehicle",
		},
		numWheels : 4,
	}
	
	av := airVehicle {
		vehicle : vehicle {
			vType : "Air Vehicle",
		},
		numLeaf : 2,
	}
	start(lv)
	start(av)
	 	
	
	 	
	
}
