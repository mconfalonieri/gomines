#!/usr/bin/env bash
set -e

# Run only during Linux builds
if [ "$4" != "linux" ]; then
  exit 0
fi


# Ensure each parallel run has an isolated workspace to avoid "Text file busy"
export TMPDIR=$(mktemp -d)
trap 'rm -rf "$TMPDIR"' EXIT

BINARY_PATH=$1  # Target binary built by GoReleaser
ARCH=$2         # amd64 or arm64
VERSION=$3

APP_DIR="dist/gomines-${ARCH}.AppDir"
mkdir -p "$APP_DIR/usr/bin"
mkdir -p "$APP_DIR/usr/share/icons/hicolor/256x256/apps"

# 1. Copy binary
cp "$BINARY_PATH" "$APP_DIR/usr/bin/gomines"

# 2. Copy application icon from img/ directory
cp "img/gomines-icon.png" "$APP_DIR/gomines.png"
cp "img/gomines-icon.png" "$APP_DIR/.DirIcon"
cp "img/gomines-icon.png" "$APP_DIR/usr/share/icons/hicolor/256x256/apps/gomines.png"

# 3. Create standard AppRun entrypoint
cat << 'EOF' > "$APP_DIR/AppRun"
#!/bin/sh
exec "$APPDIR/usr/bin/gomines" "$@"
EOF
chmod +x "$APP_DIR/AppRun"

# 4. Create desktop entry pointing to the icon
cat << EOF > "$APP_DIR/gomines.desktop"
[Desktop Entry]
Name=gomines
Exec=gomines
Icon=gomines
Type=Application
Categories=Game;
Terminal=true
EOF

# 5. Package AppImage using appimagetool
ARCH=$ARCH appimagetool "$APP_DIR" "dist/gomines_${VERSION}_linux_${ARCH}.AppImage"
