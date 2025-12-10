package main

import (
	"errors"
	"fmt"
	"log"
)

var (
	ErrNotImplemented = errors.New("Not Implemented")
	ErrTruckNotFound  = errors.New("Truck Not Found")
)

type Truck interface {
	LoadCargo() error
	UnLoadCargo() error
}

// ========================================================================
type NormalTruck struct {
	id    string
	cargo int
}

// Implementation of LoadCargo Method
func (t *NormalTruck) LoadCargo() error {
	t.cargo += 1
	return nil
}

// Implementation of LoadCargo Method
func (t *NormalTruck) UnLoadCargo() error {
	t.cargo = 0
	return nil
}

// ========================================================================

// ========================================================================
type ElectricTruck struct {
	id      string
	cargo   int
	battery float64
}

func (e *ElectricTruck) LoadCargo() error {
	e.cargo += 1
	e.battery = -1
	return nil
}
func (e *ElectricTruck) UnLoadCargo() error {
	e.cargo = 0
	e.battery += -1
	return nil
}

// ========================================================================

// processTruck handles the loading and unloading of a truck.
func processTruck(truck Truck) error {
	fmt.Printf("Processing Truck %+v \n", truck)
	err := truck.LoadCargo() // ✅ declare once
	if err != nil {
		return fmt.Errorf("Error Loading Cargo: %w", err)
	}

	err = truck.UnLoadCargo() // ✅ reuse, DO NOT use :=
	if err != nil {
		return fmt.Errorf("Error unloading Cargo: %w", err)
	}

	return nil
}

func main() {
	nt := &NormalTruck{id: "1"}
	et := &ElectricTruck{id: "2"}

	person := make(map[string]any, 0)
	person["name"] = "Tiago"
	person["age"] = 42

	age, exists := person["width"].(int)
	if !exists {
		log.Fatal("age not exists")
		return
	}
	log.Println(age)

	err := processTruck(nt)
	if err != nil {
		log.Fatalf("Error Processing Normal Truck: %s", err)
	}

	err = processTruck(et)
	if err != nil {
		log.Fatalf("Error Processing Electric Truck: %s", err)
	}

	log.Println(nt.cargo)
	log.Println(et.battery)
}
