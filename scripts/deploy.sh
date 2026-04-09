#!/bin/bash
set -e

ENV=$1
APP_DIR="/opt/apps/invaders-$ENV"
SERVICE="invaders-$ENV"

echo "🚀 Deploy invaders ($ENV)"

cd "$APP_DIR"

echo "▶ Actualizando código desde el repo..."
git fetch origin
git reset --hard origin/$(
  if [ "$ENV" = "dev" ]; then echo develop;
  elif [ "$ENV" = "qa" ]; then echo release;
  else echo master; fi
)

echo "▶ Compilando binario..."
go build -o "$SERVICE"

echo "▶ Reiniciando servicio systemd..."
sudo systemctl restart "$SERVICE"

echo "✅ Deploy $ENV terminado"

