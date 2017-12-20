package cskv

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		kv string
	}
	tests := []struct {
		name string
		args args
		want KV
	}{
		{"empty string", args{kv: ""}, KV{kv: map[string]string{}}},
		{"bad single value string", args{kv: "key"}, KV{kv: map[string]string{}}},
		{"key but no value", args{kv: "key:"}, KV{kv: map[string]string{"key": ""},Str:"key:"}},
		{"value but no key", args{kv: ":value"}, KV{kv: map[string]string{"": "value"},Str:":value"}},
		{"key:value", args{kv: "key:value"}, KV{kv: map[string]string{"key": "value"},Str:"key:value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.kv); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestKV_String(t *testing.T) {
	knv := make(map[string]string)
	knv["key"] = ""
	vnk := make(map[string]string)
	vnk[""] = "value"
	kv := make(map[string]string)
	kv["key"] = "value"

	tests := []struct {
		name string
		kv   *KV
		want string
	}{
		{"empty KV", &KV{kv: make(map[string]string)}, ""},
		{"key but no value", &KV{kv: knv}, "key:"},
		{"no key but value", &KV{kv: vnk}, ":value"},
		{"key:value", &KV{kv: kv}, "key:value"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.kv.String(); got != tt.want {
				t.Errorf("KV.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUse(t *testing.T) {
	knv := make(map[string]string)
	knv["key"] = ""
	vnk := make(map[string]string)
	vnk[""] = "value"
	kv := make(map[string]string)
	kv["key"] = "value"

	type args struct {
		kvmap map[string]string
	}
	tests := []struct {
		name string
		args args
		want KV
	}{
		{"empty map", args{kvmap: make(map[string]string)}, KV{kv: make(map[string]string)}},
		{"key no value", args{kvmap: knv}, KV{kv: knv,Str:"key:"}},
		{"value no key", args{kvmap: vnk}, KV{kv: vnk,Str:":value"}},
		{"key:value", args{kvmap: kv}, KV{kv: kv,Str:"key:value"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Use(tt.args.kvmap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Use() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
