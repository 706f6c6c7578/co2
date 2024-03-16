package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Constants for CO2 emission rates
const (
	globalGridIntensity = 442 // Global average grid intensity in grams CO2 per kWh
	energyPerByte = 1.5e-8 // Energy consumption in kWh per byte transferred
)

func main() {
	// Check if input is provided
	if len(os.Args) > 1 {
		fmt.Println("Usage: co2 < input file")
		os.Exit(1)
	}

	// Read all data from stdin
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	// Calculate the CO2 produced based on global grid intensity
	bytesTransferred := float64(len(data))
	energyConsumed := bytesTransferred * energyPerByte // Energy consumption in kWh
	co2ProducedGlobal := energyConsumed * globalGridIntensity

	// Output the result
	fmt.Printf("Transferring %.f bytes produces approximately %.6f grams of CO2 based on global average grid intensity.\n", bytesTransferred, co2ProducedGlobal)
}
