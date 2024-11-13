## Memory Management in Golang.

Memory allocation and deallocation happens automatically.

### Memory allocation methods.
- new()
    - Allocate memory, but no INIT.
    - You will get a memory address.
    - ZEROED STORAGE.

- make()
    - Allocate memory, and INIT.
    - You will get a memory address.
    - No ZEROED STORAGE.

### Garbage Collection in golang.
- GC happens automatically.
- Out of Scope or nil