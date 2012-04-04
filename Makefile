

TARG = mbd

all:
	go build -o $(TARG)

install:
	cp $(TARG) $(GOROOT)/bin

run: all
	./mbd ./output test.xml test2.lang test3.pos
