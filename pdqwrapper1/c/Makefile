all: pdq simple_test


pdq: pdq.c # this the newest pdq wrapper
	cc -g -o ./pdq pdq.c ../pdq5/lib/*.o -lm

install:
	sudo cp `pwd`/pdq /usr/local/bin/pdq
	-/usr/local/bin/pdq -h

simple_test: pdq
	./pdq -t 0.9 -s 0.1 -v 1 10 1

gdb:
	gdb --args ./pdq -x 0 1 1

clean:
	rm -f pdq *.o core
