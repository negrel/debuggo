ssa:
	go clean -modcache -cache
	go build main.go
	go tool objdump main > main.s
	go clean -modcache -cache
	go build -tags debugo main.go
	go tool objdump main > main.debugo.s

clean: 
	rm -f main main.debugo.s main.s