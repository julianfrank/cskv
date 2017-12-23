[![GoDoc](https://godoc.org/github.com/julianfrank/cskv?status.svg)](https://godoc.org/github.com/julianfrank/cskv)
 [![Coverage Status](https://coveralls.io/repos/github/julianfrank/cskv/badge.svg?branch=master)](https://coveralls.io/github/julianfrank/cskv?branch=master) [![Build Status](https://travis-ci.org/julianfrank/cskv.svg?branch=master)](https://travis-ci.org/julianfrank/cskv)
 [![Go Report Card](https://goreportcard.com/badge/github.com/julianfrank/cskv)](https://goreportcard.com/report/github.com/julianfrank/cskv)

# Comma Separate Key Value 
package for Quick Key-Value Array Sharing and usage in Text based messages

## Usage
Just import this package
```go
import "github.com/julianfrank/cskv"
```


### New(string)
New(string) => KV Object

```go
mystr := "one:onevalue,two:twovalue"

mykv:=cskv.New(mystr)

// Output : KV Object
```

### KV.String()
```go
log.Print(mykv.String())
// Output : one:onevalue,two:twovalue

Same Output for log.Print(mykv.Str)
```

### Use
Use(map[string]string) => string in <key>:<value>,<key>:<value> format

```go
mymap:=make(map[string]string)
mymap["one"]="onevalue"
mymap["two"]="twovalue"

mystr := cskv.Use(mymap).String()
[or]
mystr := cskv.Use(mymap).Str

// Output : "one:onevalue,two:twovalue"
```

### To Be Done
- [ ] Get/Set per key on KV Object
- [ ] JSON Export

## License
MIT
