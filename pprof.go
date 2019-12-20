package pprof

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
)

func StartProfile(cpuprofile, memprofile string) func() {
	//cpu
	cpufile, err := os.Create(cpuprofile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := pprof.StartCPUProfile(cpufile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//mem
	memfile, err := os.Create(memprofile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	runtime.GC()
	if err := pprof.WriteHeapProfile(memfile); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return func() {
		pprof.StopCPUProfile()
		cpufile.Close()
		memfile.Close()
	}
}
