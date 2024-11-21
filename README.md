# Globals vs. Pointers in Go: Performance Comparison

This project demonstrates the performance differences between using **global variables** and passing **pointers to arrays** in Go. It highlights the trade-offs in speed and design decisions when managing large arrays in a program.

---

## Test Description

The test compares two methods for modifying an array of size 1,000,000:

1. **Global Array**:
   - The array is declared as a global variable and modified directly within a function.
   - Example:
     ```go
     var globalArray [1000000]int
     func modifyGlobalArray() {
         for i := 0; i < 1000000; i++ {
             globalArray[i]++
         }
     }
     ```

2. **Pointer to Array**:
   - The array is declared locally, and a pointer to it is passed to a function for modification.
   - Example:
     ```go
     func modifyPointerArray(arr *[1000000]int) {
         for i := 0; i < 1000000; i++ {
             arr[i]++
         }
     }
     ```

---

## Results

| Method                | Execution Time  |
|-----------------------|-----------------|
| Global Array          | **4.65 ms**     |
| Pointer to Array      | **1.53 ms**     |

---

## Analysis

### Global Array
The global array takes longer to modify because global variables reside in a separate memory region (data segment), which may introduce additional overhead when accessed by functions.

### Pointer to Array
Passing a pointer to a local array is faster in this test. Local arrays are allocated on the stack, making access faster due to better cache locality. The pointer itself adds negligible overhead.

---

## Key Takeaways

1. **Performance**:  
   In this test, **pointer-based array modification is faster**. This result suggests that stack-allocated arrays accessed via pointers can outperform global arrays due to cache efficiency.

2. **Design Considerations**:  
   While global variables may seem convenient, they come with significant downsides:
   - Reduced modularity.
   - Harder-to-read and maintain code.
   - Potential issues in concurrent environments.

   Using pointers makes dependencies explicit, improving code clarity and modularity.

3. **When to Use Globals**:
   - For data that is constant and accessed by multiple parts of the program.
   - When global state simplifies implementation (though this should be avoided if possible).

4. **When to Use Pointers**:
   - For arrays or data structures that are local to a function or module.
   - When aiming for better maintainability and testability.

---

## Conclusion

Although global arrays may introduce minimal overhead, the **pointer-based approach** is generally preferred due to its faster access time and better code organization. Always prioritize maintainability over micro-optimizations unless performance is critical.

---

## Usage

To run the program and observe the results:

```bash
go run main.go
```
