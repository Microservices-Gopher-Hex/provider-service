#!/usr/bin/env bash
set -euo pipefail

echo "[provider] waiting db..."
until pg_isready -d "$DB_DSN" >/dev/null 2>&1; do
sleep 1
done

  # GORM hará AutoMigrate en main.go
echo "[provider] starting on :${HTTP_PORT}"
exec /app/provider