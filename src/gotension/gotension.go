// This file contains code for a string tension calculator
//
// Author: Josh McIntyre
//

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"sync"
)

// Constants and constant-ish data
const CONFIG_FILENAME = "bin/gotension/instrument_config.json"

const A4_BASE_FREQUENCY = 440.0
const INCH_TO_METERS = 0.0254
const NEWTON_TO_LBFT = 4.44822

var MATERIAL_WEIGHTS = map[string]float64{
	"plain_steel":    7800,
	"nickel_wound":   7300,
	"8020_wound":     8200,
	"phosphor_wound": 8500,
}

var SPN_OFFSETS = map[string]int{
	"C": 0, "B#": 0,
	"C#": 1, "Db": 1,
	"D":  2,
	"D#": 3, "Eb": 3,
	"E": 4, "Fb": 4,
	"F": 5, "E#": 5,
	"F#": 6, "Gb": 6,
	"G":  7,
	"G#": 8, "Ab": 8,
	"A":  9,
	"A#": 10, "Bb": 10,
	"B": 11, "Cb": 11,
}

// Types
type InstrumentConfig struct {
	Name        string    `json:"name"`
	StringSet   string    `json:"string_set"`
	TuningName  string    `json:"tuning_name"`
	ScaleLength float64   `json:"scale_length"`
	Tuning      []string  `json:"tuning"`
	Gauges      []float64 `json:"gauges"`
	Material    []string  `json:"material"`
}

// This function parses a given instrument config JSON file and Marshals into structs
func instrument_config_parse(instrument_config_filename string) []InstrumentConfig {

	file, err := os.Open(instrument_config_filename)

	if err != nil {
		fmt.Println(err)
	}

	json_bytes, err := io.ReadAll(file)

	if err != nil {
		fmt.Println(err)
	}

	var instrument_configs []InstrumentConfig
	err = json.Unmarshal(json_bytes, &instrument_configs)

	file.Close()

	return instrument_configs
}

// This function gets the frequency in Hertz (Hz) for a given Scientific Pitch Notation note (SPN)
func spn_to_hz(spn_note string) float64 {

	// First, calculate the MIDI number using the formula
	// midi = 12(octave + 1) + offset
	// Octave number should always be last in the note string, so array slice it out
	octave_num_str := spn_note[len(spn_note)-1:]
	octave_num, err := strconv.Atoi(octave_num_str)

	if err != nil {
		fmt.Println(err)
	}

	note_name := spn_note[:len(spn_note)-1]

	// Calculate the midi number from SPN, then convert that to hz
	offset, ok := SPN_OFFSETS[note_name]
	if !ok {
		fmt.Printf("Cannot determine midi offset for:%s\n", note_name)
	}

	midi_num := 12*(octave_num+1) + offset

	var power float64 = (float64(midi_num) - 69) / 12.0
	frequency := A4_BASE_FREQUENCY * math.Pow(2, power)

	return frequency
}

// Get the linear mass density of the string using an approximated effective density
func calculate_linear_mass_density(diameter float64, material string) float64 {

	diameter_meters := diameter * INCH_TO_METERS
	area := math.Pi * math.Pow((diameter_meters/2), 2)

	effective_density, ok := MATERIAL_WEIGHTS[material]
	if !ok {
		fmt.Printf("Cannot determine material density for: %s", material)
	}

	linear_mass_density := effective_density * area

	return linear_mass_density
}

// Calculate the tension of a given string with the given configuration and pitch
func calculate_string_tension(linear_mass_density float64, scale_length float64, frequency float64) float64 {

	scale_length_meters := scale_length * INCH_TO_METERS
	tension_newtons := math.Pow((2*scale_length_meters*frequency), 2) * linear_mass_density
	tension := tension_newtons / NEWTON_TO_LBFT

	return tension
}

// This wrapper will be the main function for calculating the tension for a given instrument
func calculate_instrument_tension(instrument_config InstrumentConfig, wg *sync.WaitGroup, tensions_results [][]float64, results_index int) {

	defer wg.Done()

	var tensions []float64 = make([]float64, len(instrument_config.Tuning))
	for i, spn := range instrument_config.Tuning {

		hz := spn_to_hz(spn)
		lmd := calculate_linear_mass_density(instrument_config.Gauges[i], instrument_config.Material[i])
		tension := calculate_string_tension(lmd, instrument_config.ScaleLength, hz)

		tensions[i] = tension
	}

	tensions_results[results_index] = tensions
}
