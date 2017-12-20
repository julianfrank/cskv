package cskv

import "strings"

//KV Key Value Store Structure
type KV struct {
	kv  map[string]string
	Str string
}

//New Decode the provided String and return a with a 1 layer key-value map
func New(kv string) KV {
	var o KV
	o.kv = make(map[string]string)
	if len(kv) == 0 {
		return o
	}
	h := strings.Split(kv, ",")
	for _, r := range h {
		v := strings.Split(r, ":")
		if len(v) == 2 {
			o.kv[v[0]] = v[1]
		}
	}
	o.Str = o.String()
	return o
}

// Use Parse the Provided map and return a CSKV Object
func Use(kvmap map[string]string) KV {
	var o KV
	o.kv = make(map[string]string)
	for k, v := range kvmap {
		o.kv[k] = v
	}
	o.Str = o.String()
	return o
}

//String Return the String Representation of the provided map
func (kv *KV) String() string {
	o := ""
	c := 0
	for k, v := range kv.kv {
		o += (k + ":" + v)
		if c != 0 {
			o += ","
		}
		c++
	}

	return o
}
