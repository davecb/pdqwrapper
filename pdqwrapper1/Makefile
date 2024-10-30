#all: trivial
all: pdq simple_test


pdq: pdq.c # this the newest pdq wrapper
	cc -g -o pdq pdq.c ./pdq5/lib/*.o -lm
	# sudo ln -s `pwd`/pdq /usr/local/bin/pdq

simple_test: pdq
	./pdq -z 0.9 -s 0.1 -v 1 10 1
	./pdq -z 0.9 -s 0.1 -d 0.09 -v 1 10 1
	./pdq -c 2 -z 0.9 -s 0.1 -d 0.09 -v 1 10 1

flattened:
	cc -g -o flattened flattened.c /usr/local/pdq/lib/*.o -lm
	./flattened -v 1 20 1
net:
	cc -g -o net net.c /usr/local/pdq/lib/*.o -lm

clean: # do not remove pdq!
	rm -f net general *.o core
trivial: trivial.c
	cc -g -o trivial trivial.c /usr/local/pdq/lib/*.o -lm
	./trivial 1 20 1 
	rm ./trivial

general: general.c # this the immediate predecessor to pdq
	cc -g -o general general.c /usr/local/pdq/lib/*.o -lm
	# in FUTURE: ln -s `pwd` ${HOME}/davecb/bin/pdq

fail: # writes wrong message
	pdq -v -s 0.1 -x 0.9 1 10 1
