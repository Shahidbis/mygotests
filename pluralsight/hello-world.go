package main

import(
	"fmt"
)

func main(){

	plantCapacities := []float64{30, 30, 30, 60, 60, 100}

	activePlants := []int{0,1}

	gridLoad := 75.

	fmt.Println("1) Generate Power Plant report")
	fmt.Println("2) Generate Power Grid report")
	fmt.Print("Please choose an option: ")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		plantReport(plantCapacities...)
	case "2":
		gridRepor(activePlants, plantCapacities, gridLoad)
	default:
		fmt.Println("Unknown choice. No action taken")

	}
}
func plantReport(plantCapacities ...float64){
	for idx, v := range plantCapacities{
		fmt.Printf("Plant %d capacity: %.0f\n", idx, v)
	}
}

func gridRepor(activePlants []int, plantCapacities []float64, gridLoad float64){
	capacity := 0.
	for _, plantId := range activePlants{
		capacity += plantCapacities[plantId]
	}
	fmt.Println("Capacity: ", capacity)
	fmt.Println("Load: ", gridLoad)
	fmt.Println("Utilization: ", gridLoad/capacity)
}


type PlantType string

const(
	hydro PlantType = "Hydro"
	wind PlantType = "Wind"
	solar PlantType = "Solar"
)

type PlantStatus string

const(
	active PlantStatus = "Active"
	inactive PlantStatus = "Inactive"
	unavailable PlantStatus = "Unavailable"
)

type PowerPlant struct {
	plantType PlantType
	capacity float64
	status PlantStatus
}

type PowerGrid struct {
	load float64
	plants []PowerPlant
}



