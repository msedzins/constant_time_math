# Intro

NOTE: THIS IMPLEMENTATION IS JUST FOR FUN. IT HAS NO PRACTICAL USE.

Module provides constant-time implementation of basic arithmetic operations. 

This is particularly important for cryptographic applications to prevent timing attacks.

# Constant time arithmetic

It is challenging to implement constant time operations. The basic rules are the following:

1. Avoid Conditional Branches:
    1. Do not use conditional statements (if, else, switch) that depend on the input values.
    2. Use bitwise operations to handle conditions instead.
2. Use Bitwise Operations:
    1. Perform arithmetic using bitwise operations (AND, OR, XOR, NOT, SHIFT).
    2. Bitwise operations have a fixed execution time regardless of the input values.
3. Fixed Iteration Loops:
    1. Ensure that loops iterate a fixed number of times, independent of the input values.
    2. Avoid loops that terminate based on input-dependent conditions.
4. Avoid Data-Dependent Memory Access:
    1. Do not access memory locations based on input values.
    2. Use fixed memory access patterns to prevent timing variations due to cache hits/misses.
5. Use Masking Techniques:
    1. Use bitwise masks to conditionally select values without branching.
    2. For example, to conditionally add a value, use a mask that is either all zeros or all ones.
6. Constant-Time Comparison:
    1. Implement comparisons using bitwise operations to avoid branching.
    2. For example, to compare two values, use XOR and check if the result is zero.

The most secure way to implement constant-time operations is to use dedicated processor operations (for example to implement AES encryption one could utilize the AES-NI instruction set that perform several compute intensive parts of the AES algorithm directly by CPU) However, this is not always feasible, so software implementations are often used.

The risk with software implementation is that the compiler may optimize the code in a way that breaks the constant-time property. So, it's usually recommend to verify that the code is constant-time by inspecting the assembly code.

Other risk is that the CPU may optimize the code in a way that breaks the constant-time property. For example, the CPU use cache memory to store frequently accessed data. If the data is accessed in a way that depends on the input values, the cache memory may leak information about the input values.

# Tests

Run tests with:

```bash
go test ./... -v
go test ./... -v -fuzz=<name of Fuzz test>  -fuzztime 10s


# Run benchmarks
go test ./... -v -bench=.
```
