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
export CLICKHOUSE_DB="${CLICKHOUSE_DB:-default}"

exec "$@"
