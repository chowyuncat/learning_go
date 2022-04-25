#!/bin/sh
set -eux

SOURCE_FILE="wordcount.go"
INPUT_FILE="romeo_and_juliet.txt"

time go run ${SOURCE_FILE} fields ${INPUT_FILE}
time go run ${SOURCE_FILE} naive ${INPUT_FILE}
