[![GoDoc](https://godoc.org/github.com/julianfrank/cskv?status.svg)](https://godoc.org/github.com/julianfrank/cskv)
 [![Coverage Status](https://coveralls.io/repos/github/julianfrank/cskv/badge.svg?branch=master)](https://coveralls.io/github/julianfrank/cskv?branch=master) [![Build Status](https://travis-ci.org/julianfrank/cskv.svg?branch=master)](https://travis-ci.org/julianfrank/cskv)

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

### Error
In your Application use the familiar cskv.Error ...Just remember that the E is capital
```go
comment:="Awsome!"
dude:="Julian Frank"
myError:=cskv.Error("This comment (%s) was made by %s")
//Output myError=> "This comment (Awsome!) was made by Julian Frank" 
//of type error
//Use err.Error() to retreive error string 
```

### cskv Display Flags 
If You want to selectively enable or disable cskv display of Log or Error you can simply do that by setting cskv.LogMode and cskv.ErrorMode to true or false. True is default
```go
//To make screen completely quiet
cskv.LogMode=false
cskv.ErrorMode=false

//To Display only errors and not the Logs
cskv.LogMode=false
cskv.ErrorMode=true
```

## License
MIT