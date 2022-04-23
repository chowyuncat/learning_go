#!/bin/sh
set -eux

time go run wordcount.go fields romeo_and_juliet.txt
time go run wordcount.go naive romeo_and_juliet.txt
