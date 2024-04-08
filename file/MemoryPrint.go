package file

import (
	"fmt"
	"runtime"
)

// This code belongs to https://gist.github.com/j33ty
// This code was taken from this file https://gist.github.com/j33ty/79e8b736141be19687f565ea4c6f4226

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garbage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", byteToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", byteToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", byteToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

// This function returns the Mb of memory used from the number of bytes
func byteToMb(b uint64) uint64 {
	// return b / 1024 / 1024
	return b
}
