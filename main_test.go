package main

import (
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("Process Truck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			// TODO: Implement test assertions for truck processing
			nt := &NormalTruck{id: "1", cargo: 42}
			et := &ElectricTruck{id: "2"}

			person := make(map[string]any, 0)
			person["name"] = "Tiago"
			person["age"] = 42

			err := processTruck(nt)
			if err != nil {
				t.Fatalf("Error Processing Normal Truck: %s", err)
			}

			err = processTruck(et)
			if err != nil {
				t.Fatalf("Error Processing Electric Truck: %s", err)
			}

			if nt.cargo != 0 {
				t.Fatalf("Expected NormalTruck cargo to be 0 after unloading, got %d", nt.cargo)
			}

			if et.battery != -2 {
				t.Fatalf("Expected ElectricTruck cargo to be 0 after unloading, got %d", et.cargo)
			}
		})
	})
}
