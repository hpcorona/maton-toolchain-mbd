
include $(GOROOT)/src/Make.inc

TARG = mbd
GOFILES = \
	xml.go \
	lang.go \
	pos.go \
	mbd.go

include $(GOROOT)/src/Make.cmd

run: all
	./mbd ./output test.xml test2.lang test3.pos
