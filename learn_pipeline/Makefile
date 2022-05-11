all: bounded md5sum_serial md5sum_parallel_naive md5sum_parallel_fixed_readers

bounded: bounded.go
	go build -o $@ bounded.go

md5sum_serial: serial.go
	go build -o $@ serial.go

md5sum_parallel_naive: parallel_naive.go
	go build -o $@ parallel_naive.go

md5sum_parallel_fixed_readers: parallel_fixed_readers.go
	go build -o $@ parallel_fixed_readers.go

clean:
	rm md5sum_serial; rm bounded; rm md5sum_parallel_naive; rm md5sum_parallel_fixed_readers
