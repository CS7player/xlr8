#!/bin/bash

set -e

APP=xlr8
VERSION=1.0.0

rm -rf pkg
rm -f ${APP}.deb

echo "Building binary..."
go build -o ${APP} ./cmd/xlr8

echo "Creating package structure..."

mkdir -p pkg/DEBIAN
mkdir -p pkg/usr/bin
mkdir -p pkg/etc/profile.d

cp ${APP} pkg/usr/bin/
chmod 755 pkg/usr/bin/${APP}

cat > pkg/etc/profile.d/xlr8.sh <<'EOF'
xlr8_go() {
    xlr8

    if [ -f /tmp/xlr8-cwd ]; then
        cd "$(cat /tmp/xlr8-cwd)"
    fi
}
EOF

chmod 644 pkg/etc/profile.d/xlr8.sh

cat > pkg/DEBIAN/control <<EOF
Package: ${APP}
Version: ${VERSION}
Section: utils
Priority: optional
Architecture: amd64
Maintainer: CS7player
Description: Terminal file manager written in Go
EOF

echo "Building .deb package..."
dpkg-deb --build pkg

mv pkg.deb ${APP}.deb

echo ""
echo "Created ${APP}.deb"
echo ""
echo "Install with:"
echo "  sudo apt install ./${APP}.deb"
echo ""
echo "After installation open a new terminal or run:"
echo "  source /etc/profile.d/xlr8.sh"
echo ""
echo "Then use:"
echo "  xlr8_go"