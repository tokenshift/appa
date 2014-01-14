.PHONY: all test lib debug clean

all: appa test libappa.a libappa-debug.a
lib: libappa.a
debug: libappa-debug.a
test: appa

appa: main.o libappa-debug.a
	gcc -o appa main.c libappa-debug.a -Wall -Werror -g

libappa.a: src/*.c src/*.h
	gcc -c src/*.c -Wall -Werror
	ar rs libappa.a *.o

libappa-debug.a: src/*.c src/*.h
	gcc -c src/*.c -Wall -Werror -g
	ar rs libappa-debug.a *.o

clean:
	rm -f appa *.a *.o src/*.gch
