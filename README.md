## General
____________

### Author
* Josh McIntyre

### Website
* jmcintyre.net

### Overview
* GoTension is a string tension calculator for multi-instrument use

## Development
________________

### Git Workflow
* development for bugfixes and new features

### Building
* make build
Build the application - wraps `go build`
* make test
Build the unit tests and run - wraps `go test`
* make clean
Clean the build directory

### Features
* Take an instrument configuration JSON file with data on instruments, tunings, and strings
* Use estimated material densities for various common string types (electric & acoustic, wound & plain)
* Calculate individual string tension for each string along with a total, based on scale length, string density, and note frequency from the tuning
* Calculations are done in parallel using goroutines
* Output a final table of tension results, per-instrument

### Requirements
* Requires Go language build tools

### Platforms
* Windows
* MacOSX
* Linux

## Usage
____________

### Command Line Usage
* Configure your own `instrument_config.json` or use the included sample
* Run `gotension.exe` to calculate tensions and output to the command line

### Unit tests
* Run `make test` or directly run `go test ./src/gotension -o bin/gotension -v`