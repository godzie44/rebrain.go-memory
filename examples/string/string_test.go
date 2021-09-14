package string

import (
	"bytes"
	"runtime"
	"testing"
	"unsafe"
)

type PackageType string

const (
	TypeControl = "control"
	TypeData    = "data"
	TypeUnknown = "unknown"
)

type RawPackage struct {
	typ []byte
}

//go:noinline
func ParseType(p *RawPackage) PackageType {
	strType := string(p.typ)
	switch strType {
	case TypeControl, TypeData:
		return PackageType(strType)
	default:
		return TypeUnknown
	}
}

var (
	TypeControlBytes = []byte("control")
	TypeDataBytes    = []byte("data")
)

//go:noinline
func ParseTypeNoAlloc(p *RawPackage) PackageType {
	if bytes.Compare(p.typ, TypeControlBytes) == 0 {
		return TypeControl
	}
	if bytes.Compare(p.typ, TypeDataBytes) == 0 {
		return TypeData
	}
	return TypeUnknown
}

var testPackages = []*RawPackage{{typ: []byte("control")}, {typ: []byte("data")}, {typ: []byte("foo")}}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tp := ParseType(testPackages[i % 3])
		runtime.KeepAlive(tp)
	}
}

func BenchmarkParseNoAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tp := ParseTypeNoAlloc(testPackages[i % 3])
		runtime.KeepAlive(tp)
	}
}

func ExternalLibFunc(typ string)  {
}

func (p *RawPackage) TypeAsStringUnsafe() string  {
	return *(*string)(unsafe.Pointer(&p.typ))
}

func BenchmarkExternal(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str = string(testPackages[i % 3].typ)
		ExternalLibFunc(str)
	}
}

func BenchmarkExternalNoAlloc(b *testing.B) {
	var str string
	for i := 0; i < b.N; i++ {
		str = testPackages[i % 3].TypeAsStringUnsafe()
		ExternalLibFunc(str)
	}
}