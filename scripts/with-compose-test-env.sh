#!/usr/bin/env bash
set -euo pipefail

if [ "$#" -eq 0 ]; then
  echo "Usage: $0 <command> [args...]" >&2
  exit 2
fi

export USE_COMPOSE="${USE_COMPOSE:-1}"
export TARANTOOL_USER="${TARANTOOL_USER:-userT}"
export TARANTOOL_PASS="${TARANTOOL_PASS:-passT}"
export TARANTOOL_HOST="${TARANTOOL_HOST:-127.0.0.1}"
export TARANTOOL_PORT="${TARANTOOL_PORT:-3303}"
export CLICKHOUSE_USER="${CLICKHOUSE_USER:-userC}"
export CLICKHOUSE_PASS="${CLICKHOUSE_PASS:-passC}"
export CLICKHOUSE_HOST="${CLICKHOUSE_HOST:-127.0.0.1}"
export CLICKHOUSE_PORT="${CLICKHOUSE_PORT:-9002}"
export CLICKHOUSE_HTTP_PORT="${CLICKHOUSE_HTTP_PORT:-8125}"
export CLICKHOUSE_DB="${CLICKHOUSE_DB:-default}"

if [ "$USE_COMPOSE" = "1" ]; then
  wait_for_tcp() {
    local name="$1"
    local host="$2"
    local port="$3"

    for _ in {1..60}; do
      if timeout 1 bash -c "</dev/tcp/${host}/${port}" 2>/dev/null; then
        return 0
      fi
      sleep 1
    done

    echo "Timed out waiting for ${name} at ${host}:${port}" >&2
    return 1
  }

  wait_for_http() {
    local name="$1"
    local url="$2"

    for _ in {1..60}; do
      if curl -fsS --max-time 1 "$url" >/dev/null 2>&1; then
        return 0
      fi
      sleep 1
    done

    echo "Timed out waiting for ${name} at ${url}" >&2
    return 1
  }

  wait_for_tcp "Tarantool" "$TARANTOOL_HOST" "$TARANTOOL_PORT"
  wait_for_http "ClickHouse" "http://${CLICKHOUSE_HOST}:${CLICKHOUSE_HTTP_PORT}/ping"
fi

exec "$@"
