.PHONY: all test lib debug clean

all: appa test libappa.a libappa-debug.a
lib: libappa.a
debug: libappa-debug.a
test: appa

appa: main.o libappa-debug.a
	gcc -o appa main.c libappa-debug.a -Wall -g

libappa.a: src/*.c src/*.h
	gcc -c src/*.c -Wall
	ar rs libappa.a *.o

libappa-debug.a: src/*.c src/*.h
	gcc -c src/*.c -Wall -g
	ar rs libappa-debug.a *.o

clean:
	rm -f appa *.a *.o src/*.gch
