#!/usr/bin/env bash

source ./scripts/env.sh

if [ ! -f $BINARY ]; then
    echo "Binaries not found, building first..."
    ./scripts/build.sh
fi

echo "Press CTRL+C to exit..."
$BINARY
