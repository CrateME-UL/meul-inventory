#!/bin/sh

wait_for_db() {
  local host="$1"
  local port="$2"
  local timeout="$3"
  local start_time=$(date +%s)

  while true; do
    if nc -z "$host" "$port"; then
      echo "Database connection established."
      break
    else
      local current_time=$(date +%s)
      if [ $((current_time - start_time)) -ge "$timeout" ]; then
        echo "Timed out waiting for DB: $host:$port"
        exit 1
      fi
      echo "Waiting for database connection..."
      sleep 1
    fi
  done
}

DATABASE_TIMEOUT=30
echo $DATABASE_HOST
echo $DATABASE_URL
echo $DATABASE_PORT
echo $DATABASE_TIMEOUT
wait_for_db "$DATABASE_HOST" "$DATABASE_PORT" "$DATABASE_TIMEOUT"

exec "$@"