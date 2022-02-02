//go:build prof

package main

// Will export /debug/pprof
import _ "net/http/pprof"
