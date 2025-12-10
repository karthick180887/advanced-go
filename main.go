package main

import (
	"errors"
	"fmt"
)

var (
	ErrNotImplemented = errors.New("not implemented")
	ErrTruckNotFound  = errors.New("truck not found")
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
		return fmt.Errorf("error loading cargo: %w", err)
	}

	err = truck.UnLoadCargo() // ✅ reuse, DO NOT use :=
	if err != nil {
		return fmt.Errorf("error unloading cargo: %w", err)
	}

	return nil
}
