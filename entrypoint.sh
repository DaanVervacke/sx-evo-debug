#!/bin/bash

echo "Starting CAN interface setup..."

if [ ! -c "$CAN_DEVICE" ]; then
    echo "ERROR: Device $CAN_DEVICE does not exist or is not accessible"
    exit 1
fi

echo "Setting up CAN interface..."

if ! sudo slcand -o -c -s4 $CAN_DEVICE $CAN_INTERFACE; then
    echo "ERROR: Failed to setup interface on $CAN_DEVICE"
    exit 1
fi

echo "Bringing up $CAN_INTERFACE interface..."

if ! sudo ifconfig $CAN_INTERFACE up; then
    echo "ERROR: Failed to bring up $CAN_INTERFACE interface"
    exit 1
fi

echo "Setting txqueuelen to 1000..."

if ! sudo ifconfig $CAN_INTERFACE txqueuelen 1000; then
    echo "ERROR: Failed to set txqueuelen for can0"
    exit 1
fi

echo "CAN interface setup completed successfully"

echo "Starting sx-evo-debug..."

CMD=( "./sx-evo-debug" "--interface" "$CAN_INTERFACE" )
[ -n "$DOBISS_MODULE" ] && CMD+=( "--module" "$DOBISS_MODULE" )
[ -n "$DOBISS_OUTPUT" ] && CMD+=( "--output" "$DOBISS_OUTPUT" )

exec "${CMD[@]}"
