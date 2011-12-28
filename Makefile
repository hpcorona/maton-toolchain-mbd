
include $(GOROOT)/src/Make.inc

TARG = layadd
GOFILES = \
	pos.go \
	layadd.go

include $(GOROOT)/src/Make.cmd

run: all
<<<<<<< HEAD
	./mbd ./output test.xml test2.lang test3.pos
=======
	./layadd test.pos ipod4 QVGA 0.325 out.pos
>>>>>>> c9f8e74ecb11e96786136650a936ba6ac610e488
