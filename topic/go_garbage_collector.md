# The Go Blog: Getting to Go: The Journey of Go's Garbage Collector

- [Original Article Summary or Guidance](#original-article-summary-or-guidance)

## Original Article Summary or Guidance

### GC's view of Go

- Go is value-oriented

Value-orientation also helps with the foreign function interfaces, (FFI). We have a fast FFI with C and C++.
This one design decision has led to some of the more amazing things that have to go on with the runtime.
It is probably the most important thing that differentiates Go from other GCed languages.

- Go allows interior pointers

Of course Go can have pointers and in fact they can have interior pointers. Such pointers keep the entire value live, 
and they are fairly common.

- Static ahead of time compilation

Binary contains entire runtime. No JIT recompilation.

- The two GC knobs `SetGCPercent`, `SetMaxHeap`

Go comes with two knobs to control the GC. The first one is GCPresent. Basically this is a knob that adjusts how much CPU
you want to use and how much memory you want to use. The default is 100 which means that half the heap is dedicated to 
live memory and  half is dedicated to allocation. Tou can modify this in either direction.

MaxHeap, which is not yet released but being used and evaluated internally, lets the programmer set what the maximum heap 
size should be. Out of memory, OOMs, are tough on Go; temporary spikes in memory useage should be handled by increasing
CPU costs, not by aborting.

- Go runtime How we got here

GC latency is an existential threat to Go.

...

The decision was to do a tri-color concurrent algorithm.

more content, omit...
