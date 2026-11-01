#!/bin/bash
set -euo pipefail

# Sender password. Receivers do not get this value.
# Leave it only on this server and in the desktop sender settings.
RELAY_PASS="${RELAY_PASS:-change-me}"
BIN_SRC="${1:-./croc-relay}"
INSTALL_PATH="/usr/local/bin/croc-relay"

if [[ "$RELAY_PASS" == "change-me" ]]; then
  echo "Set RELAY_PASS before running this script."
  exit 1
fi
if [[ ! -f "$BIN_SRC" ]]; then
  echo "Missing relay binary: $BIN_SRC"
  echo "Build it on a machine with this repository:"
  echo "  GOOS=linux GOARCH=amd64 go build -trimpath -ldflags=\"-s -w\" -o croc-relay ./cmd/croc-relay"
  exit 1
fi

echo "[1/5] Installing patched relay to $INSTALL_PATH"
install -m 0755 "$BIN_SRC" "$INSTALL_PATH"

echo "[2/5] Stopping an official croc relay if it holds the ports"
if systemctl list-unit-files croc.service --no-legend >/dev/null 2>&1; then
  systemctl disable --now croc.service || true
fi

echo "[3/5] Writing systemd unit"
cat > /etc/systemd/system/croc-relay.service <<SERVICE
[Unit]
Description=croc relay (sender password required)
After=network.target

[Service]
Type=simple
DynamicUser=yes
CapabilityBoundingSet=CAP_NET_BIND_SERVICE
ExecStart=$INSTALL_PATH -pass "$RELAY_PASS"
Restart=always
RestartSec=5

[Install]
WantedBy=multi-user.target
SERVICE

echo "[4/5] Enabling service and opening 9009-9013/tcp"
systemctl daemon-reload
systemctl enable --now croc-relay
if command -v ufw >/dev/null 2>&1; then
  ufw allow 9009:9013/tcp
else
  echo "ufw was not found. Open 9009-9013/tcp yourself."
  echo "CentOS/RHEL: firewall-cmd --permanent --add-port=9009-9013/tcp && firewall-cmd --reload"
fi

echo "[5/5] Service status"
systemctl status croc-relay --no-pager
