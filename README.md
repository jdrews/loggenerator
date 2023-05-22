# log generator

Generates logs with randomized lorem ipsum content.

## Usage ##
You can use this as a library or as a standalone executable.

### Library Usage ###
See [generator/generator.go](generator/generator.go) for all the public functions available.

### Standalone Executable Usage ###
Call `loggenerator` with the following flags
* `--logfile` Name of logfile that receives the generated log lines
* `--interval` Log line generation interval in milliseconds
* `--prependlogname` Prepend the name of the logfile to the loglines

## Building ##
Just run `go build`

## History ##
This tool was created to help test [jdrews/logstation](https://github.com/jdrews/logstation). I needed a way to generate random logs to test out log watching functionality. 

