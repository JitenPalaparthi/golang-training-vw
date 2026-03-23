package main

const TICK int = 1 // with type

const PI = 3.14 // type inference

var Global int = 100 // Global or Package Level variables. In some programming they are called as static variables

var Float = 32.76

const (
	MINUTE = 60 * TICK // expression, evaualted by the compiler at compiler time
	HOUR   = 60 * MINUTE
	DAY    = 24 * HOUR
)

const (
	A, B = 10, 20
	C    = 100 + (A+B)/2 + (A-B)*2 + 25 - 10 + MINUTE // expression
)

func main() {

	//var num = 10 // when will it assign the value to num?

	// const MAX = 9999 * num // type inferences , when will it be evaluated

	const NUM = 10
	const MAX = 9999 * NUM

	a, b, ok, str := 10, 20.5, true, "Hello World"
	println(a, b, ok, str)
	Global++
	println(Global)
	println(DAY)
}

/* Memory model of the process

Text/Code Segment
-----------------

- Your binary is loaded into this segment, all your machine code is stored in thsi memory, when you run the application
-- Some programing / Os types --> constants are also loaded here, but now majorly changed

Data Segment
------------

RO Data Segment --> Constants , String literals etc

Static/Global variables are stored here

They can be initialized or uninitiazed variables

Stack memory
------------

Local variables, functions  are stored in stack memory (FILO/LIFO) --> SP (Stack Pointer)
The size of the stack is 2MB( in majority of OS, deffer based on OS and Settings but once assigned , application starts running , cannot be changed)

Heap Memory
-----------

Heap is unlimited amount of memory, obiously if memory is available and allocatable
// malloc, jelloac, calloc --> advanced allocators like page allocator, arena allocator etc..

*/

// In Go unlike any other , a variable can be on stack or on heap
// Escape Analysis --> The compiler decides where to allocate the varialbe on stack or on heap

// What compiler does
// Compiler --> Parsing, tokens, AST
// -- Place semicolan by the end of statement or expression
// -- Constants evaluation
// -- Escape Analysis
// -- Many other things
