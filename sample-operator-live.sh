#!/bin/bash

for i in {1..100}
do
    ./multi-operator-manager sample-operator apply-configuration --input-dir ./must-gather  --output-dir=./sample-operator-output --controllers-to-run=demo-controller

    go run cmd/must-gather-syncer/must-gather-syncer.go

    sleep 1
done

