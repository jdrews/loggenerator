#!/bin/bash
# Usage: ./run-many-loggenerators.sh [count]
# Example: ./run-many-loggenerators.sh 20

# Default to 5 logs if no argument provided
COUNT=${1:-5}

# Array to store PIDs
PIDS=()

# Cleanup function to kill all background processes
cleanup() {
    echo ""
    echo "Stopping log generators..."
    for pid in "${PIDS[@]}"; do
        if kill -0 "$pid" 2>/dev/null; then
            kill "$pid"
        fi
    done
    echo "All log generators stopped."
    exit 0
}

# Trap SIGINT (Ctrl+C) and SIGTERM
trap cleanup SIGINT SIGTERM

echo "Starting $COUNT log generators..."

for ((i=1; i<=COUNT; i++)); do
    # Format the log file name with zero-padding (e.g., logfile-01.log)
    LOGFILE=$(printf "test/logfile-%02d.log" "$i")
    
    echo "Starting loggenerator for $LOGFILE"
    # Run in background and redirect output to null
    ./loggenerator -logfile "$LOGFILE" > /dev/null 2>&1 &
    # Store the PID
    PIDS+=($!)
done

echo "Log generators running. Press Ctrl+C to stop."

# Wait for all background processes to finish (or until signal received)
wait