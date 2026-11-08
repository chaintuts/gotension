// This file contains unit tests for the GoTension application
//
// Author: Josh McIntyre
package main

import (
	"math"
	"testing"
)

// Constants
const A4_EXPECTED = 440.0
const E1_EXPECTED = 41.203445
const C2_EXPECTED = 65.406391
const F2_EXPECTED = 87.307058
const E2_EXPECTED = 82.406889
const E4_EXPECTED = 329.627557

var EXPECTED_MAP = map[string]float64{
	"E1": E1_EXPECTED,
	"C2": C2_EXPECTED,
	"F2": F2_EXPECTED,
	"E2": E2_EXPECTED,
	"E4": E4_EXPECTED,
}

const TEST_MATERIAL = "plain_steel"
const TEST_DIAMETER = 0.0105
const TEST_SCALE_LENGTH = 24.75

const LMD_EXPECTED = 0.000436
const TENSION_EXPECTED = 16.825544

const EPSILON = 0.00001

// Unit test helpers
func float_comp(num1 float64, num2 float64, epsilon float64) bool {

	res := math.Abs(num2-num1) < epsilon

	return res
}

// Test functions
func TestSpnToHz(t *testing.T) {

	a4_res := spn_to_hz("A4")
	if a4_res != A4_EXPECTED {
		t.Errorf("Expected for A4: %f, got: %f", A4_EXPECTED, a4_res)
	}

	for spn, freq := range EXPECTED_MAP {
		res := spn_to_hz(spn)
		comp_res := float_comp(res, freq, EPSILON)
		if !comp_res {
			t.Errorf("Expected for %s: %f, got: %f", spn, freq, res)
		}
	}
}

func TestCalculateLinearMassDensity(t *testing.T) {

	lmd_res := calculate_linear_mass_density(TEST_DIAMETER, TEST_MATERIAL)
	comp_res := float_comp(lmd_res, LMD_EXPECTED, EPSILON)
	if !comp_res {
		t.Errorf("Expected for %f %s: %f, got: %f", TEST_DIAMETER, TEST_MATERIAL, LMD_EXPECTED, lmd_res)
	}

}

func TestCalculateStringTension(t *testing.T) {

	lmd := calculate_linear_mass_density(TEST_DIAMETER, TEST_MATERIAL)
	tension_res := calculate_string_tension(lmd, TEST_SCALE_LENGTH, E4_EXPECTED)
	comp_res := float_comp(tension_res, TENSION_EXPECTED, EPSILON)

	if !comp_res {
		t.Errorf("Expected for %f %f %f: %f, got: %f", lmd, TEST_SCALE_LENGTH, E4_EXPECTED, TENSION_EXPECTED, tension_res)
	}
}
