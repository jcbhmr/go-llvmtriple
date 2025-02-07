# Mini LLVM TargetParser library for Go

🎯 `LLVMTargetParser` library but it's only `llvm::Triple`-related things

## Installation

```sh
go get github.com/jcbhmr/go-minillvmtargetparser/v19
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/jcbhmr/go-minillvmtargetparser/v19"
    "github.com/jcbhmr/go-minillvmtargetparser/v19/minillvmsupport"
)

func main() {

}
```

## Development

This project focuses on replicating the `llvm::Triple` class and all of its associated members, dependencies, and behaviors in Go. Luckily the `LLVMTargetParser` library has only one dependency, which is `libLLVMSupport`. We only implement a subset of the `livLLVMSupport` library; just enough to get `llvm::Triple` working.

Library filename: `libLLVMTargetParser.a`/`LLVMTargetParser.lib` \
Namespaces: `llvm`, `llvm::sys`, `llvm::ARM`, etc. \
Go package name: `minillvmtargetparser` \
Go package path: `github.com/jcbhmr/go-minillvmtargetparser/v19`
Example Go names: `minillvmtargetparser.LLVM*`, `minillvmtargetparser.LLVMSys*`, `minillvmtargetparser.LLVMARM*`, etc.

Library filename: `libLLVMSupport.a`/`LLVMSupport.lib` \
Namespaces: `llvm`, `llvm::ARMBuildAttrs`, etc. \
Go package name: `minillvmsupport` \
Go package path: `github.com/jcbhmr/go-minillvmtargetparser/v19/minillvmsupport`
Example Go names: `minillvmsupport.LLVM*`, `minillvmsupport.LLVMARMBuildAttrs*`, etc.

Some C++-to-Go translation conventions:

- `enum {}` is mapped without a name prefix.
- C++ names are converted to Go names **with acronyms capitalized according to Go conventions**.
- Names should document their original C++ name.
- Take pointers where C++ would use references. Document that the pointer is not nullable. Document whether it's mutable.
- `std::optional<T>` is mapped to `*T` in parameters. `std::optional<T>` is mapped to `(T, bool)` in return values.

Some examples:

- `llvm::Triple::UnknownArch` ➡ `minillvmtargetparser.LLVMTripleUnknownArch`
- `llvm::Triple::aarch64_be` ➡ `minillvmtargetparser.LLVMTripleAarch64BE`
- `llvm::Triple::x86_64` ➡ `minillvmtargetparser.LLVMTripleX8664`
- `llvm::ARMBuildAttrs::ABI_PCS_wchar_t` ➡ `minillvmsupport.LLVMARMBuildAttrsABIPCSWCharT`
