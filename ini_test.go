package main

import (
	"bytes"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const correctFig string = `dev:
    image: btburke/golang-dev
    ports:
        - "10001:10001"
    volumes_from:
        - code
    working_dir: /golang/src/github.com/BTBurke/dev
code:
    image: busybox
    volumes:
        - /Users/btb/project/golang/src/github.com/BTBurke/dev:/golang/src/github.com/BTBurke/dev`

func TestFig(t *testing.T) {
	Convey("it should write the fig.yml stub correctly", t, func() {
		config := Config{
			Port:         "10001:10001",
			ContainerDir: "/golang/src/github.com/BTBurke/dev",
			LocalDir:     "/Users/btb/project/golang/src/github.com/BTBurke/dev",
			DevImage:     "btburke/golang-dev",
		}

		var out bytes.Buffer
		err := renderToFig(&config, &out)
		So(err, ShouldBeNil)
		So(out.String(), ShouldEqual, correctFig)
	})
}
