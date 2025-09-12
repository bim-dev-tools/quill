#!/usr/bin/env bash
set -e

REPO="bim-dev-tools/quill"
LATEST=$(curl -s https://api.github.com/repos/$REPO/releases/latest | grep tag_name | cut -d '"' -f4)

if [ -z "$LATEST" ]; then
  echo "Could not fetch latest release."
  exit 1
fi


OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

case "$OS" in
  linux) OS="linux" ;;
  darwin) OS="darwin" ;;
  msys*|mingw*|cygwin*) OS="windows" ;;
  *) echo "Unsupported OS: $OS"; exit 1 ;;
esac

case "$ARCH" in
  x86_64|amd64) ARCH="amd64" ;;
  arm64|aarch64) ARCH="arm64" ;;
  *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

BINARY="quill_${LATEST}_${OS}_${ARCH}"
TMPDIR=$(mktemp -d)
cd "$TMPDIR"

if [ "$OS" = "windows" ]; then
  ZIP="${BINARY}.zip"
  URL="https://github.com/$REPO/releases/download/$LATEST/$ZIP"
  echo "Downloading $URL..."
  curl -L -o "$ZIP" "$URL"
  unzip "$ZIP"
  INSTALL_PATH="$HOME/quill.exe"
  mv "${BINARY}.exe" "$INSTALL_PATH"
else
  URL="https://github.com/$REPO/releases/download/$LATEST/$BINARY"
  echo "Downloading $URL..."
  curl -L -o "$BINARY" "$URL"
  chmod +x "$BINARY"
  INSTALL_PATH="/usr/local/bin/quill"
  sudo mv "$BINARY" "$INSTALL_PATH"
fi

cd -
echo "Cleaning up..."
rm -rf "$TMPDIR"

echo "Quill installed!"
$INSTALL_PATH --version
