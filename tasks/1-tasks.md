# Golang Tasks: Basic to Advanced

## 1) Basics: Syntax, Variables, Types, and Control Flow

### Task 1 — Hello Go
**Topic:** Program structure, `package main`, `func main()`  
**Task:** Write a Go program that prints `Hello, Go!` and explains the role of `package main` and `main()` in comments.

### Task 2 — Variables and Constants
**Topic:** `var`, short declaration `:=`, constants  
**Task:** Create a program that declares variables using different styles and constants for application metadata like app name, version, and port.

### Task 3 — Zero Values
**Topic:** Default values in Go  
**Task:** Declare variables of type `int`, `float64`, `string`, `bool`, pointer, slice, map, and interface, then print their zero values.

### Task 4 — Basic Data Types
**Topic:** Numeric, string, boolean types  
**Task:** Build a program that takes hardcoded values and performs arithmetic, comparisons, and boolean operations.

### Task 5 — Type Conversion
**Topic:** Explicit conversions  
**Task:** Convert between `int`, `float64`, `byte`, and `string`, and print the results while noting where conversions are safe or lossy.

### Task 6 — Formatted Printing
**Topic:** `fmt.Printf`, format verbs  
**Task:** Print the same values using `%v`, `%+v`, `%#v`, `%T`, `%d`, `%f`, `%s`, and `%q`.

### Task 7 — User Input
**Topic:** Reading from standard input  
**Task:** Read a user's name and age from input and print a formatted greeting.

### Task 8 — If / Else Logic
**Topic:** Conditional statements  
**Task:** Write a program that checks whether a number is positive, negative, or zero, and whether it is even or odd.

### Task 9 — Switch Statement
**Topic:** `switch`, multiple cases, expression-less switch  
**Task:** Build a simple command classifier that maps input strings like `start`, `stop`, and `pause` to actions.

### Task 10 — Loops
**Topic:** `for` loop variations  
**Task:** Use all three loop forms in Go to print numbers, sum ranges, and simulate a while loop.

---

## 2) Functions and Core Language Mechanics

### Task 11 — Functions with Parameters and Return Values
**Topic:** Function basics  
**Task:** Create functions for add, subtract, multiply, and divide, then call them from `main()`.

### Task 12 — Multiple Return Values
**Topic:** Idiomatic Go returns  
**Task:** Write a function that returns quotient and remainder for two integers.

### Task 13 — Named Return Values
**Topic:** Named return parameters  
**Task:** Implement a function that calculates rectangle area and perimeter using named returns.

### Task 14 — Variadic Functions
**Topic:** `...` parameters  
**Task:** Create a variadic function that accepts any number of integers and returns sum, average, min, and max.

### Task 15 — Anonymous Functions
**Topic:** Function literals  
**Task:** Use an anonymous function to calculate square and cube values for numbers in a slice.

### Task 16 — Closures
**Topic:** Capturing outer variables  
**Task:** Build a counter generator function that returns another function which increments and returns a private counter.

### Task 17 — Recursion
**Topic:** Recursive functions  
**Task:** Implement factorial and Fibonacci recursively, then compare with iterative versions.

### Task 18 — Defer
**Topic:** Deferred execution  
**Task:** Write a program showing how `defer` works with multiple calls and explain LIFO execution.

### Task 19 — Panic and Recover
**Topic:** Runtime failure handling  
**Task:** Create a safe division function that uses `panic` for invalid internal states and `recover` in the caller.

### Task 20 — Error Values
**Topic:** `error` handling  
**Task:** Write a function that parses an integer from string and returns a proper error when parsing fails.

---

## 3) Collections: Arrays, Slices, Maps, and Strings

### Task 21 — Arrays
**Topic:** Fixed-size collections  
**Task:** Create an array of 10 integers, initialize it, and compute sum, average, and reverse order.

### Task 22 — Slice Basics
**Topic:** Dynamic views over arrays  
**Task:** Build a program demonstrating slice creation, slicing, `len`, and `cap`.

### Task 23 — Append and Growth
**Topic:** Slice capacity behavior  
**Task:** Track how slice length and capacity change while repeatedly appending elements.

### Task 24 — Copying Slices
**Topic:** `copy` behavior  
**Task:** Show the difference between assigning a slice and copying a slice into a new backing array.

### Task 25 — 2D Slices
**Topic:** Nested slices  
**Task:** Build a matrix using slices and print it in row-column format.

### Task 26 — Map Basics
**Topic:** Key-value storage  
**Task:** Create a student marks map and support insert, update, delete, and lookup operations.

### Task 27 — Map Iteration
**Topic:** `range` over maps  
**Task:** Count the frequency of words in a sentence using a map.

### Task 28 — Strings and Runes
**Topic:** UTF-8, bytes vs runes  
**Task:** Write a program that prints byte length, rune count, and iterates over Unicode characters in a string.

### Task 29 — String Utilities
**Topic:** `strings` package  
**Task:** Implement a simple text analyzer using `Contains`, `Split`, `TrimSpace`, `ToUpper`, and `ReplaceAll`.

### Task 30 — Sorting Collections
**Topic:** `sort` package  
**Task:** Sort integers, strings, and a slice of structs by different fields.

---

## 4) Pointers, Structs, and Methods

### Task 31 — Pointers
**Topic:** Address and dereference  
**Task:** Write functions that modify a value with and without pointers and compare the results.

### Task 32 — Struct Basics
**Topic:** Custom types with fields  
**Task:** Define a `Book` struct and create several instances using different initialization styles.

### Task 33 — Struct Methods
**Topic:** Methods on types  
**Task:** Add methods to `Book` such as `Display()` and `IsExpensive()`.

### Task 34 — Pointer Receivers vs Value Receivers
**Topic:** Method receiver semantics  
**Task:** Create a `Counter` struct with methods using both receiver types and observe the difference.

### Task 35 — Nested Structs
**Topic:** Composition through fields  
**Task:** Model `Address` inside `Employee` and print formatted employee details.

### Task 36 — Anonymous Structs
**Topic:** Temporary structured values  
**Task:** Use anonymous structs to represent configuration loaded in memory for a simple app.

### Task 37 — Struct Tags
**Topic:** Metadata via tags  
**Task:** Define JSON tags on a struct and marshal/unmarshal sample data.

### Task 38 — Comparing Structs
**Topic:** Comparable types  
**Task:** Experiment with comparing structs directly and note when comparison fails due to slice or map fields.

### Task 39 — Constructor Pattern
**Topic:** Factory-like helpers  
**Task:** Create `NewUser()` and validate required fields before returning the struct.

### Task 40 — Methods as Behavior
**Topic:** Domain modeling  
**Task:** Build a `BankAccount` struct with `Deposit`, `Withdraw`, and `Balance` methods including validation.

---

## 5) Interfaces and Polymorphism

### Task 41 — Basic Interface
**Topic:** Interfaces by behavior  
**Task:** Define a `Shape` interface with `Area()` and implement it for circle, rectangle, and triangle.

### Task 42 — Empty Interface and Type Assertions
**Topic:** `any`, type assertion  
**Task:** Write a function that accepts `any` and behaves differently for `int`, `string`, and `[]int`.

### Task 43 — Type Switch
**Topic:** Dynamic dispatch  
**Task:** Build a generic printer using a type switch for multiple supported input types.

### Task 44 — Interface Composition
**Topic:** Combining interfaces  
**Task:** Create `Reader`, `Writer`, and `ReadWriter` style interfaces and implement them in a mock type.

### Task 45 — Custom Stringer
**Topic:** `fmt.Stringer`  
**Task:** Implement the `String()` method for a custom struct and print it using `fmt.Println`.

### Task 46 — Error Interface
**Topic:** Custom errors  
**Task:** Create a custom error type for validation failures and return it from input validation logic.

### Task 47 — Dependency Inversion with Interfaces
**Topic:** Abstraction for testability  
**Task:** Build a service that depends on an interface, then swap implementations for console logging and file logging.

### Task 48 — Interface Nil Pitfall
**Topic:** Typed nil inside interface  
**Task:** Demonstrate the difference between a nil interface and an interface holding a nil pointer.

### Task 49 — io.Reader Practice
**Topic:** Standard interfaces  
**Task:** Write a function that reads all data from any `io.Reader` and prints the content length and data preview.

### Task 50 — Polymorphic Payment Processor
**Topic:** Real use of interfaces  
**Task:** Implement a payment processor supporting card, UPI, and wallet methods via a shared interface.

---

## 6) Packages, Modules, and Project Structure

### Task 51 — Create a Multi-file Package
**Topic:** Package organization  
**Task:** Split a small application into multiple files under one package and call shared functions from `main()`.

### Task 52 — Create a Reusable Package
**Topic:** Exported vs unexported identifiers  
**Task:** Build a `mathutil` package exposing only selected helpers.

### Task 53 — Go Modules
**Topic:** `go mod init`, dependency tracking  
**Task:** Create a new module, add a third-party package, and inspect `go.mod` and `go.sum`.

### Task 54 — Internal Packages
**Topic:** Access boundaries  
**Task:** Create an `internal` package and verify it cannot be imported from outside its allowed scope.

### Task 55 — Build Tags
**Topic:** Conditional compilation  
**Task:** Create OS-specific files or debug-only behavior using build tags.

### Task 56 — Package Documentation
**Topic:** Go doc comments  
**Task:** Add proper package, type, and function documentation comments for a reusable library.

### Task 57 — Init Function
**Topic:** Package initialization  
**Task:** Demonstrate how `init()` runs across multiple files and packages.

### Task 58 — Environment-based Configuration
**Topic:** Config loading  
**Task:** Build a config package that reads values from environment variables with defaults.

### Task 59 — CLI Flags
**Topic:** `flag` package  
**Task:** Create a command-line tool that accepts name, age, and verbose flags.

### Task 60 — Project Layout Exercise
**Topic:** Organizing real projects  
**Task:** Restructure a small app into `cmd/`, `internal/`, `pkg/`, and `configs/` folders.

---

## 7) File Handling, JSON, and Data Processing

### Task 61 — Read a Text File
**Topic:** File I/O  
**Task:** Read a text file line by line and print line numbers.

### Task 62 — Write and Append to a File
**Topic:** Output to files  
**Task:** Create a log file writer that appends timestamped entries.

### Task 63 — Buffered I/O
**Topic:** `bufio` package  
**Task:** Use buffered readers and writers to process a large text file efficiently.

### Task 64 — JSON Marshal / Unmarshal
**Topic:** `encoding/json` basics  
**Task:** Convert a struct to JSON and back again.

### Task 65 — JSON from File
**Topic:** External data parsing  
**Task:** Read a JSON file containing users, unmarshal it, validate fields, and print a summary.

### Task 66 — CSV Processing
**Topic:** `encoding/csv`  
**Task:** Read a CSV of products and compute total inventory value.

### Task 67 — XML Basics
**Topic:** `encoding/xml`  
**Task:** Marshal and unmarshal simple XML data for a catalog.

### Task 68 — Temporary Files and Directories
**Topic:** OS utilities  
**Task:** Create temp files, write content, read it back, and clean up.

### Task 69 — Walk Directory Tree
**Topic:** `filepath.WalkDir`  
**Task:** Build a utility that lists all `.go` files and counts them by folder.

### Task 70 — Mini Log Analyzer
**Topic:** Combining strings, files, maps  
**Task:** Parse a log file and count occurrences of `INFO`, `WARN`, and `ERROR`.

---

## 8) Concurrency: Goroutines, Channels, Context

### Task 71 — First Goroutine
**Topic:** Concurrent execution  
**Task:** Launch multiple goroutines that print messages and coordinate completion correctly.

### Task 72 — WaitGroup
**Topic:** Synchronization  
**Task:** Use `sync.WaitGroup` to run 10 tasks concurrently and wait for all to finish.

### Task 73 — Channels Basics
**Topic:** Communication between goroutines  
**Task:** Send integers from one goroutine to another and compute their squares.

### Task 74 — Buffered Channels
**Topic:** Capacity and blocking  
**Task:** Compare behavior of buffered and unbuffered channels with simple producer-consumer code.

### Task 75 — Range and Close on Channels
**Topic:** Channel lifecycle  
**Task:** Generate numbers in one goroutine, close the channel, and consume them using `range`.

### Task 76 — Select Statement
**Topic:** Multiplexing channels  
**Task:** Build a program listening on two channels and handling whichever is ready first.

### Task 77 — Timeout with `time.After`
**Topic:** Controlling wait time  
**Task:** Add timeout handling to a channel receive operation.

### Task 78 — Context Cancellation
**Topic:** Cooperative cancellation  
**Task:** Start worker goroutines that stop when a context is cancelled.

### Task 79 — Worker Pool
**Topic:** Controlled concurrency  
**Task:** Implement a worker pool that processes jobs from a channel and returns results.

### Task 80 — Fan-out / Fan-in Pipeline
**Topic:** Concurrent pipelines  
**Task:** Create a pipeline that distributes input work to multiple workers and merges the results.

---

## 9) Synchronization, Memory, and Performance

### Task 81 — Mutex
**Topic:** Protecting shared state  
**Task:** Implement a counter incremented by multiple goroutines safely using `sync.Mutex`.

### Task 82 — RWMutex
**Topic:** Read-heavy workloads  
**Task:** Build an in-memory store using `sync.RWMutex` and simulate concurrent readers and writers.

### Task 83 — Once
**Topic:** One-time initialization  
**Task:** Use `sync.Once` to initialize a singleton configuration object.

### Task 84 — Atomic Operations
**Topic:** `sync/atomic`  
**Task:** Compare atomic counters with mutex-based counters.

### Task 85 — Race Condition Detection
**Topic:** `go test -race`  
**Task:** Intentionally create a race condition, detect it with the race detector, then fix it.

### Task 86 — Benchmarking
**Topic:** Performance measurement  
**Task:** Write benchmarks comparing two implementations of string concatenation.

### Task 87 — Profiling
**Topic:** CPU and memory profiling  
**Task:** Create a CPU-heavy function and capture profiling data using `pprof`.

### Task 88 — Memory Allocation Awareness
**Topic:** Escape analysis, allocations  
**Task:** Compare value vs pointer-heavy code and inspect compiler output for escape behavior.

### Task 89 — Efficient Slice and Map Usage
**Topic:** Preallocation and optimization  
**Task:** Rewrite a slow program to use preallocated slices and maps, then benchmark the difference.

### Task 90 — Object Pooling
**Topic:** `sync.Pool`  
**Task:** Use `sync.Pool` in a workload with repeated temporary objects and measure effects.

---

## 10) Testing, Generics, Reflection, and Advanced Topics

### Task 91 — Unit Testing
**Topic:** `testing` package  
**Task:** Write unit tests for arithmetic and validation functions, including table-driven tests.

### Task 92 — Subtests and Test Helpers
**Topic:** Better test structure  
**Task:** Refactor tests to use `t.Run()` and helper functions.

### Task 93 — Mocking via Interfaces
**Topic:** Test doubles  
**Task:** Test a service by injecting a fake repository implementation.

### Task 94 — Fuzz Testing
**Topic:** Robustness testing  
**Task:** Add a fuzz test for a parser or sanitizer function.

### Task 95 — Generics Basics
**Topic:** Type parameters  
**Task:** Write generic functions for min, max, reverse slice, and contains.

### Task 96 — Generic Data Structure
**Topic:** Reusable typed containers  
**Task:** Implement a generic stack or queue.

### Task 97 — Reflection
**Topic:** `reflect` package  
**Task:** Build a utility that prints struct field names, types, values, and tags dynamically.

### Task 98 — HTTP Server
**Topic:** Networking, handlers, JSON APIs  
**Task:** Build a REST API with routes for creating, listing, updating, and deleting tasks in memory.

### Task 99 — Database Integration
**Topic:** SQL with Go  
**Task:** Connect to PostgreSQL or SQLite, create a table, and implement CRUD operations.
