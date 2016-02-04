HASH=`git rev-parse HEAD`
TIME=`date -u '+%Y-%m-%d %H:%M:%S'`

prepare:
	rm -f templates.rice-box.go
	rm -f static.rice-box.go
	rice embed-go

check:
	go test -cover -v github.com/denisbr/filebin/app/api github.com/denisbr/filebin/app/model github.com/denisbr/filebin/app/config

get-deps:
	go get github.com/dustin/go-humanize
	go get github.com/gorilla/mux
	go get github.com/rwcarlsen/goexif/exif
	go get github.com/disintegration/imaging
	go get github.com/GeertJohan/go.rice
	go get github.com/GeertJohan/go.rice/rice

build: prepare
	go build -ldflags "-X main.buildstamp \"${TIME}\" -X main.githash \"${HASH}\""

install: prepare
	go install -ldflags "-X main.buildstamp \"${TIME}\" -X main.githash \"${HASH}\""
