# Gommon v0.0.14 dcli

## TODO

- [ ] merge w/ existing design doc in [dcli/doc/design](../../../dcli/doc/design)
- [ ] split up features
- [ ] list implementation order
 
## Overview

A commandline application builder `dcli` that replaces [spf13/cobra](https://github.com/spf13/cobra).
Minor fix to update small util packages e.g. `mathutil`, `stringutil`, `envutil`.

## Motivation

`dcli`

- less dependencies
- more customization
- type safe
- learn from other command line builders, e.g. clap (w/ structopt)

## Specs

- support git style flag and subcommand
- use `dcli` for `gommon` command