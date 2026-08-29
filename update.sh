#!/bin/sh
# Re-download the MDI SVG set if npm has a newer @mdi/svg than VERSION.
set -eu
cd "$(dirname "$0")"

meta=$(curl -sf https://registry.npmjs.org/@mdi/svg/latest)
latest=$(printf '%s' "$meta" | grep -o '"version":"[^"]*"' | head -1 | cut -d'"' -f4)
current=$(cat VERSION 2>/dev/null || echo none)

if [ "$latest" = "$current" ]; then
	echo "up to date ($current)"
	exit 0
fi

tarball=$(printf '%s' "$meta" | grep -o '"tarball":"[^"]*"' | cut -d'"' -f4)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
curl -sfL "$tarball" | tar xz -C "$tmp"

rm -rf svg
cp -r "$tmp/package/svg" svg
cp "$tmp/package/LICENSE" LICENSE
echo "$latest" >VERSION
sed -i "s/(v[0-9.]*,/(v$latest,/" mdi.go

go run gen.go # regenerates icons.go and all/all.go
echo "updated $current -> $latest"
go test ./...
