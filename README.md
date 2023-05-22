# log generator

Generates logs with randomized lorem ipsum content.

Binaries available in [releases](https://github.com/jdrews/loggenerator/releases). See [usage](https://github.com/jdrews/loggenerator#usage) below.


## Usage ##
You can use this as a library/package or as a standalone executable.

### Library Usage ###
See [generator/generator.go](generator/generator.go) for all the public functions available.

### Standalone Executable Usage ###
Call `loggenerator` with the following flags
* `--logfile` Name of logfile that receives the generated log lines
* `--interval` Log line generation interval in milliseconds
* `--prependlogname` Prepend the name of the logfile to the loglines

## Building ##
You can simply run `go build` if building for your local machine. 

loggenerator also uses  [goreleaser](https://github.com/goreleaser/goreleaser) for all releases

To build all targets locally you can run   
`goreleaser build --snapshot --clean`

## Releasing ##

* Push a tag following [semver](https://semver.org/)
    * `git tag -a 2.0.0 -m "2.0.0"`
    * `git push origin 2.0.0`
* Ensure you have an environment variable with `GITHUB_TOKEN="YOUR_GH_TOKEN"` and minimum of `write:packages` permissions
* Release!
    * `goreleaser release`
* Read the [goreleaser quickstart](https://goreleaser.com/quick-start/) for more details

## History ##
This tool was created to help test [jdrews/logstation](https://github.com/jdrews/logstation). I needed a way to generate random logs to test out log watching functionality. 

