# APRS-IS Passcode Generator – User Guide

## Overview

`aprs-passcode` is a small Go-based tool that generates the **APRS-IS passcode** for a given amateur radio callsign.
The passcode is required when logging into APRS-IS servers.

This app can be used as:

* **Standalone CLI** tool
* **Go library** imported into other projects

> **Note:** The passcode algorithm is a lightweight checksum for APRS login; it is **not cryptographically secure**.

---

## 1. Installation

### Option A: Download prebuilt CLI

If you publish releases with GitHub Releases, download the binary for your OS and place it in your `PATH`.

### Option B: Build from source

```bash
git clone https://github.com/cmgann/aprs-passcode.git
cd aprsis-passcode
# Build the CLI
go build ./cmd/aprs-passcode
# Optional: move it to a directory in your PATH
mv aprs-passcode /usr/local/bin/
```

---

## 2. Using the CLI

### Syntax

```bash
aprs-passcode <CALLSIGN>
```

### Examples

```bash
# Generate passcode for N0CALL
aprs-passcode N0CALL
# Output:
13023

# Generate passcode for a callsign with SSID
aprs-passcode N0CALL-9
# Output:
13023
```

### Notes

* Callsigns are **case-insensitive**
* Any SSID suffix (`-1`, `-9`, etc.) is ignored during calculation
* Output is a **15-bit integer**, suitable for APRS-IS login

---

## 3. Using the Go Library

You can also import the library in your Go projects.

### Import

```go
import "github.com/cmgann/aprs-passcode/aprs"
```

### Generate passcode

```go
package main

import (
	"fmt"
	"github.com/cmgann/aprs-passcode/aprs"
)

func main() {
	call := "N0CALL"
	passcode := aprs.Passcode(call)
	fmt.Printf("APRS passcode for %s is %d\n", call, passcode)
}
```

### Output

```
APRS passcode for N0CALL is 30671
```

---

## 4. Testing

The library includes tests for correctness.

```bash
# Run all tests
go test ./...
```

Sample test cases cover:

* Uppercase and lowercase callsigns
* Callsigns with SSIDs
* Known APRS passcode results

---

## 5. Workflow / CI

* **GitHub Actions workflow** is included to:

  * Run tests
  * Build the CLI
  * Ensure code quality

* Status checks can be required on `main` or `development` branches to prevent merging broken code.

---

## 6. Tips and Best Practices

* Use the **library** if integrating into another Go project.
* Use the **CLI** for quick passcode lookup.
* Avoid storing passcodes in source control; generate them dynamically when needed.
* Verify that your APRS client supports 15-bit passcodes.

---

## 7. Contribution Guide (Optional)

* Fork the repository
* Create a feature branch
* Submit a pull request to `development`
* Ensure **all tests pass**
* Maintain Go doc comments for exported functions