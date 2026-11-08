// This file contains the main entry point for the program
//
// Author: Josh McIntyre
//

package main

import (
	"fmt"
	"sync"
)

// Output formatting for instrument string tensions
func output_tensions(instrument_config InstrumentConfig, tensions []float64) {

	// Output formatting
	tension_table := ""
	total_tension := 0.0
	for i, tension := range tensions {
		tension_table += "\t" + instrument_config.Tuning[i] + ": " + fmt.Sprintf("%f", tension) + "\n"
		total_tension += tension
	}
	tension_table += fmt.Sprintf("\tTotal: %f", total_tension)

	gauge_join := ""
	for i, gauge := range instrument_config.Gauges {

		if i == (len(instrument_config.Gauges) - 1) {
			gauge_join += fmt.Sprintf("%f", gauge)
		} else {
			gauge_join += fmt.Sprintf("%f", gauge) + ", "
		}

	}

	sep := "----------------------------------------\n"
	output := fmt.Sprintf("%sTensions for: %s\nTuning: %s\nString set: %s\nGauges: %s\nTensions: \n%s\n%s",
		sep, instrument_config.Name, instrument_config.TuningName, instrument_config.StringSet,
		gauge_join, tension_table, sep)

	fmt.Println(output)
}

// The main entry point for the program
func main() {

	instrument_configs := instrument_config_parse(CONFIG_FILENAME)

	// Use goroutines and a WaitGroup to parallelize calculating tension for each instrument in the config
	var wg sync.WaitGroup
	tensions_results := make([][]float64, len(instrument_configs))

	fmt.Println("Calculating tension for instruments: ")
	for i, config := range instrument_configs {
		wg.Add(1)
		go calculate_instrument_tension(config, &wg, tensions_results, i)
	}

	wg.Wait()

	// Output the final calculated tensions once all the workers have finished
	for i, config := range instrument_configs {
		output_tensions(config, tensions_results[i])
	}
}
