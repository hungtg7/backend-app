#!/usr/bin/env bash
goose -dir ./migrations up
./main
