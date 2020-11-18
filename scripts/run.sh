#!/usr/bin/env bash

source ./scripts/env.sh

export $(cat ./.env)

if [ ! -f $BINARY ]; then
    echo "Binaries not found, building first..."
    ./scripts/build.sh
fi

echo "Attempting to stop if already running..."
pkill -f $BINARY

echo "Press CTRL+C to exit..."
$BINARY
