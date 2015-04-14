package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

//A test program to compare the performance difference between
//using many goroutines managed by the go scheduler or matching
//the amount of goroutines to threads without using a scheduler.
//Comment the method you don't want to use out.
func main() {
	procNo := runtime.NumCPU()
	runtime.GOMAXPROCS(procNo)
	fmt.Println("Using", procNo, "processors for maximum thread count")

	//goroutinesWithScheduler()
	goroutinesCpuMatch()
}

func goroutinesWithScheduler() {
	//We'll use an unsigned integer as a counter.
	var ops uint64 = 0

	//To simulate concurrent updates, we'll start 10
	//goroutines that each increment the counter as often
	//as they can
	for i := 0; i < 10; i++ {
		go func() {
			for {
				//To atomically increment the counter we
				//use `AddUint64`, giving it the memory
				//address of our `ops` counter with the
				//`&` syntax.
				atomic.AddUint64(&ops, 1)

				//Allow other goroutines to proceed.
				runtime.Gosched()
			}
		}()
	}

	//Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)

	//In order to safely use the counter while it's still
	//being updated by other goroutines, we extract a
	//copy of the current value into `opsFinal` via
	//`LoadUint64`. As above we need to give this
	//function the memory address `&ops` from which to
	//fetch the value.
	opsFinal := atomic.LoadUint64(&ops)

	fmt.Println("Many goroutines using runtime.Gosched:", opsFinal, "ops per second")
}

func goroutinesCpuMatch() {
	//This benchmark is much the same as above so I have only
	//commented the differences
	var ops uint64 = 0

	//Only produce as many goroutines as there are free threads
	for i := 0; i < runtime.NumCPU()-1; i++ {
		go func() {
			for {
				//this time we just increment the atomic counter
				//as Gosched isn't needed
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	//Wait a second to allow some ops to accumulate.
	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)

	fmt.Println("goroutines match CPU count:", opsFinal, "ops per second")
}
