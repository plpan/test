all:std def;

std:
	@echo "---std test---"
	go run test_std_hello.go

def:
	@echo "---def test---"
	gcc -Wall -shared -o libsayhello.so sayhello.c
	# add `// #cgo LDFLAGS: -L. -lsayhello` in test_def_hello.go
	go run test_def_hello.go

clean:
	-@rm -rf sayhello.o libsayhello.so
