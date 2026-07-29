#!/usr/bin/env bash
# Démarre tiny-service.
# Le token attendu est lu depuis API_TOKEN (comme une ANTHROPIC_API_KEY).
set -euo pipefail

cd "$(dirname "$0")"

export API_TOKEN="${API_TOKEN:-sk-tiny-1234567890}"
export PORT="${PORT:-6565}"

echo "🔑 API_TOKEN = ${API_TOKEN}"
echo "🚀 Build..."
go build -o tiny-service .

echo "🐳 Démarrage sur le port ${PORT} (bind 0.0.0.0)"
exec ./tiny-service
