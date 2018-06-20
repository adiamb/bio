package bam

// This import is needed to use go:linkname.
import _ "unsafe"

// The following functions are defined in go runtime.  To use them, we need to
// import "unsafe", and elsewhere in this package, import "C" to force compiler
// to recognize the "go:linktime" directive. Some of the details are explained
// in the below blog post.
//
// procPin() pins the caller to the current processor, and returns the processor
// id in range [0,GOMAXPROCS). procUnpin() undos the effect of procPin().
//
// http://www.alangpierce.com/blog/2016/03/17/adventures-in-go-accessing-unexported-functions/

//go:linkname runtime_procPin sync.runtime_procPin
//go:nosplit
func runtime_procPin() int

//go:linkname runtime_procUnpin sync.runtime_procUnpin
//go:nosplit
func runtime_procUnpin()

//go:linkname fastrand sync.fastrand
func fastrand() uint32
