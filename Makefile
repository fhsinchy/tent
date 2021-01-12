V=0.0.2
LDFLAGS=-ldflags="-X 'github.com/fhsinchy/tent/cmd.version=v${V}'"

build:
	go build ${LDFLAGS} -o bin/tent
install:
	go install ${LDFLAGS}