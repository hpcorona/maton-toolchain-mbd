
include $(GOROOT)/src/Make.inc

TARG = mbd
GOFILES = \
	xml.go \
	lang.go \
	mbd.go

include $(GOROOT)/src/Make.cmd

run: all
	./mbd test.xml test2.lang
