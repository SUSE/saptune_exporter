#!/usr/bin/env bash

cat <<EOF
## Path:           SAP/System Tuning/General
## Description:    Global settings for saptune - the comprehensive optimisation management utility for SAP solutions
## ServiceRestart: tuned

## Type:    string
## Default: ""
#
# When saptune is activated, apply optimisations for these SAP solutions.
# The value is a list of solution names, separated by spaces.
# Run "saptune solution list" to get a comprehensive list of solution names.
TUNE_FOR_SOLUTIONS="HANA"

## Type:    string
## Default: ""
#
# When saptune is activated, apply tuning for these SAP notes in addition to those
# already recommended by the above list of SAP solutions.
# The value is a list of note numbers, separated by spaces.
# Run "saptune note list" to get a comprehensive list of note numbers.
TUNE_FOR_NOTES=""

## Type:    string
## Default: ""
#
# When saptune is activated, apply tuning for the notes in exactly the below
# order
# The value is a list of note numbers, separated by spaces.
NOTE_APPLY_ORDER="941735 1771258 1980196 2578899 2684254 2382421 2534844"

## Type:    string
## Default: "2"
#
# Version of saptune
SAPTUNE_VERSION="2"
EOF
