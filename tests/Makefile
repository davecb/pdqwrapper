# test of go and c versions
#
default: c # golang

golang:
	go test -v -run Test_goPdq

c:
	go test -v -run  Test_cPdq
	# took 2.8 seconds to printf the file to stdout
	# took 0.6 seconds to write it to cPdq_test.csv
	ls -l cPdq_test.csv
	# test belongs here
