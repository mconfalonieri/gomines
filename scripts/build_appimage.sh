#!/usr/bin/env bash
set -e

# Run only during Linux builds
if [ "$4" != "linux" ]; then
  exit 0
fi

BINARY_PATH=$1  # Target binary built by GoReleaser (absolute or relative)
ARCH=$2         # amd64 or arm64
VERSION=$3

# Ensure output directory exists
mkdir -p dist

# Temporary build workspace
APP_DIR="dist/gomines-${ARCH}.AppDir"
RDNS_NAME="com.github.mconfalonieri.gomines"

rm -rf "$APP_DIR"

# Ensure each parallel run has an isolated workspace for appimagetool mounts
export TMPDIR=$(mktemp -d)
trap 'fusermount -u "$TMPDIR"/.mount_* 2>/dev/null || umount -l "$TMPDIR"/.mount_* 2>/dev/null || true; rm -rf "$TMPDIR"' EXIT

mkdir -p "$APP_DIR/usr/bin"
mkdir -p "$APP_DIR/usr/share/applications"
mkdir -p "$APP_DIR/usr/share/icons/hicolor/256x256/apps"
mkdir -p "$APP_DIR/usr/share/metainfo"

# Copy binary and icons
cp "$BINARY_PATH" "$APP_DIR/usr/bin/gomines"
cp "img/gomines-icon.png" "$APP_DIR/${RDNS_NAME}.png"
cp "img/gomines-icon.png" "$APP_DIR/.DirIcon"
cp "img/gomines-icon.png" "$APP_DIR/usr/share/icons/hicolor/256x256/apps/${RDNS_NAME}.png"

# AppRun entrypoint
cat << 'EOF' > "$APP_DIR/AppRun"
#!/bin/sh
exec "$APPDIR/usr/bin/gomines" "$@"
EOF
chmod +x "$APP_DIR/AppRun"

# Desktop entry
cat << EOF > "$APP_DIR/usr/share/applications/${RDNS_NAME}.desktop"
[Desktop Entry]
Name=gomines
Exec=gomines
Icon=${RDNS_NAME}
Type=Application
Categories=Game;
Terminal=true
X-AppImage-Version=${VERSION}
EOF

# Symlink desktop file at the AppDir root
ln -sr "$APP_DIR/usr/share/applications/${RDNS_NAME}.desktop" "$APP_DIR/${RDNS_NAME}.desktop"

# AppStream Metadata
cat << EOF > "$APP_DIR/usr/share/metainfo/${RDNS_NAME}.appdata.xml"
<?xml version="1.0" encoding="UTF-8"?>
<component type="desktop-application">
  <id>${RDNS_NAME}</id>
  <metadata_license>CC0-1.0</metadata_license>
  <project_license>GPL-3.0-or-later</project_license>
  <name>Gomines</name>
  <summary>A minesweeper game written in Go</summary>
  <developer id="@f00l@mastodon.bida.im">
    <name>Marco Confalonieri</name>
  </developer>
  <description>
    <p>
      A little minesweeper game written in Go to play around the BubbleTea
      framework. It runs in the terminal and it is controlled only using the
      keyboard.
    </p>
  </description>
  <launchable type="desktop-id">${RDNS_NAME}.desktop</launchable>
  <provides>
    <binary>gomines</binary>
  </provides>
  <content_rating type="oars-1.1" />
  <url type="homepage">https://github.com/mconfalonieri/gomines</url>
  <screenshots>
    <screenshot type="default">
      <caption>Gomines gameplay</caption>
      <image type="source">https://raw.githubusercontent.com/mconfalonieri/gomines/main/img/gomines.png</image>
    </screenshot>
  </screenshots>
</component>
EOF

# Update information
UPDATE_INFO="gh-releases-zsync|mconfalonieri|gomines|latest|gomines_*_linux_${ARCH}.AppImage.zsync"

# Package AppImage inside dist directory
PROJECT_ROOT=$(pwd)
cd dist

ARCH=$ARCH appimagetool -u "$UPDATE_INFO" "gomines-${ARCH}.AppDir" "gomines_${VERSION}_linux_${ARCH}.AppImage"

cd "$PROJECT_ROOT"
