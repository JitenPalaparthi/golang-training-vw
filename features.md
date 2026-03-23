
- Open Source , backed by Google and other big companies

- General Purpose ,Compiled Programming Language 

- Systems Language 

- Fast Compilation 

- Simple Language (25 keyrwords and 18+ builtin funcs)

- Not OOPs, Not Functional, Not Procedural , it is Hybrid

- Fit for Cloud Native applications 

- Fast Startup times

- No need to install any runtime

- Cross Compilation 

- GC --> Concurrent GC, Green Tea GC, Tri Colour GC 

- Concurrency - CSP 

- Static Binaries 

- Statically Linked

- Very stable, thousands of libraries available 

- Near to C++ Performance 

- Docker, Kubernetes, Openshift, DropBox, Consul,Vault, Nomad, Etcd, Minio, seawoodFS ... 


### Compilation

-- .NET Based

C# Source Code --> C# Compiler --> IR Code --> .Net Runtime --> JIT Compilation + Execute 

-- JVM based

Java Source Code --> Java Compiler --> Byte Code --> JVM --> JIT Compilation + Execute 

Kotlin Source Code --> Kotlin Compiler --> Byte Code --> JVM --> JIT Compilation + Execute 

Scale Source Code --> Scala Compiler --> Byte Code --> JVM --> JIT Compilation + Execute 

Closure Source Code --> Closure Compiler --> Byte Code --> JVM --> JIT Compilation + Execute 


-- Programming languages without JIT which is AOT 

C/C++ Source Code --> GCC/LLVM [-> IR -> FE --> BE(Target Compilation --> Assembler + Linker)] -> Binary/Executable

Rust Source Code --> LLVM [-> IR -> FE --> BE(Target Compilation --> Assembler + Linker)] -> Binary/Executable

Go Source Code --> Go compiler -> Plan9 Assemlber (Assemply Language) + Linker -> Binary/Executable 

Note: There is not emitted IR code in Go, some kind of IR is in memory IR but not emitted one