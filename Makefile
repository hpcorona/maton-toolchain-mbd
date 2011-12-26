
include $(GOROOT)/src/Make.inc

TARG = layadd
GOFILES = \
	pos.go \
	layadd.go

include $(GOROOT)/src/Make.cmd

run: all
	./layadd test.pos ipod4 QVGA 0.325 out.pos
