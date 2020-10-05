#!/usr/bin/env bash

if [ ! -f $BINARY ]; then
    echo "Binaries not found, building first..."
    ./scripts/build.sh
fi

echo "Press CTRL+C to exit..."
$BINARY
