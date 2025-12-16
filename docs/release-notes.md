# Release Notes

## APRS-IS Passcode Generator – v1.0.0 Release Notes

**Release Date:** 2025-12-16

### Overview

`aprs-passcode` v1.0.0 is the first official release of a lightweight Go-based tool for generating **APRS-IS passcodes** from amateur radio callsigns. Passcodes are required to log in to APRS-IS servers and are computed using a simple checksum algorithm.

This release supports:

* **Standalone CLI** for quick passcode generation
* **Go library** for integration into other Go projects

> ⚠️ **Note:** The passcode algorithm is **not cryptographically secure**. It is suitable only for APRS login purposes.

---

### Features in v1.0.0

* Generate APRS-IS passcodes from any valid callsign
* Ignore SSID suffixes (`-1`, `-9`, etc.) automatically
* Case-insensitive input handling
* 15-bit integer output suitable for APRS-IS
* Fully tested with known APRS passcode values
* GitHub Actions workflow for CI:

  * Automated tests
  * CLI build
  * Code quality checks

---

**v1.0.0** marks the first stable release of `aprs-passcode`, providing reliable, easy-to-use APRS-IS passcode generation for both CLI and Go library users.