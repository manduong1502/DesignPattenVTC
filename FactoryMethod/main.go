package main

import "fmt"

type iTransport interface{
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

type transport struct {
	name string
	power int
}

func (g* transport) setName(name string) {
	g.name = name
}

func (g*transport) getName() string {
	return g.name
}

func (g* transport) setPower(power int) {
	g.power = power
}

func (g*transport) getPower() int {
	return g.power
}

type truck struct{
	transport
}

func newTruck() iTransport {
	return &truck{
		transport: transport{
			name: "Truck",
			power: 4,
		},
	}
}

type ship struct{
	transport
}

func newShip() iTransport {
	return &ship{
		transport: transport{
			name: "Ship",
			power: 5,
		},
	}
}

type grab struct{
	transport
}

func newGrab() iTransport {
	return &grab{
		transport: transport{
			name: "grab",
			power: 6,
		},
	}
}

func getTransport(transportType string) (iTransport, error){
	if transportType == "truck" {
		return newTruck(), nil
	}
	if transportType == "ship" {
		return newShip(), nil
	}
	if transportType == "grab" {
		return newGrab(), nil
	}
	return nil, fmt.Errorf("Wrong transport type")
}

func main()  {
	truck, _ := getTransport("truck")
	ship, _ := getTransport("ship")
	grab, _ := getTransport("grab")
	printDetail(truck)
	printDetail(ship)
	printDetail(grab)
}

func printDetail(g iTransport)  {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}