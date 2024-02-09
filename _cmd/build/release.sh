version=`git describe --tags --abbrev=0`
go build -ldflags "-X main.version=$version" prokishi/_cmd/prokishi
go build -ldflags "-X main.version=$version" prokishi/_cmd/prokishi-server

echo "$version build"
