# go-scan
a simple port scanner in go

# compile
simply run:
`go build go-scan.go`

# usage
```
USAGE: go-scan [IP] [ARGS]
	IP         : ip of the target
	-p         : range to scan (default: 1-1000)
	-v         : verbrose, shows ports that are closed as well as other info
	-t         : scan type (default: tcp)
	example: go-scan xxx.xxx.xxx.xxx -p 1-1000 -t tcp -v
  ```
