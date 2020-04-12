cd cli

BUILDDATE=$(date '+%Y/%m/%d %H:%M:%S %Z')

LDFLAGS1="github.com/masaki-linkode/tglo/cli/subcommand.version=$(git describe --tags --abbrev=0)"
LDFLAGS2="github.com/masaki-linkode/tglo/cli/subcommand.revision=$(git rev-parse --short HEAD)"
LDFLAGS3="github.com/masaki-linkode/tglo/cli/subcommand.buildDate=${BUILDDATE}"
echo $LDFLAGS1
echo $LDFLAGS2
echo $LDFLAGS3
go build -ldflags "-X \"${LDFLAGS1}\" -X \"${LDFLAGS2}\" -X \"${LDFLAGS3}\"" -o ../bin/linux/tglo
GOOS=windows GOARCH=amd64 go build -ldflags "-X \"${LDFLAGS1}\" -X \"${LDFLAGS2}\" -X \"${LDFLAGS3}\"" -o ../bin/windows/tglo.exe
GOOS=darwin GOARCH=amd64 go build -ldflags "-X \"${LDFLAGS1}\" -X \"${LDFLAGS2}\" -X \"${LDFLAGS3}\"" -o ../bin/osx/tglo