# BPHS – Classical Vedic Astrology Backend

Go backend that implements classical Vedic astrology principles as described by **Sage Parashara (BPHS)**, **Sage Jaimini**, **Smt. Gayatri Devi Vasudev**, and **Jaya Shekhar**.

It computes planetary positions, house effects, Shadbala, Upagrahas, Karakamsha, conditional dashas, and couple compatibility (Ashtakoota + doshas) using the Swiss Ephemeris via [github.com/mshafiee/swephgo](https://github.com/mshafiee/swephgo).

---

## Features

| Area                       | Capabilities                                                                                      |
| -------------------------- | ------------------------------------------------------------------------------------------------- |
| **Birth chart**            | Graha longitudes, rashis, speeds, declinations, ascendant, True Pushya Paksha ayanamsha           |
| **House effects**          | Bhava-lord based results for all 12 houses                                                        |
| **Shadbala**               | Sthana, Dig, Kala, Chesta, Drik balas + ranking                                                   |
| **Upagrahas & Karakamsha** | Classical upagraha positions and karakamsha effects                                               |
| **Conditional dasha**      | Conditional / special dasha calculations                                                          |
| **Couple matching**        | Varna, Vashya, Tara, Yoni, Graha Maitri, Gana, Bhakuta, Nadi + Rajju / other doshas + Dosha Samya |
| **Auth**                   | PASETO token-based authentication                                                                 |
| **Geocoding**              | Geoapify integration for location lookup                                                          |

---

## Tech Stack

- **Language:** Go 1.25+
- **HTTP:** Gin
- **DB:** PostgreSQL + GORM
- **Auth:** PASETO (github.com/o1egl/paseto)
- **Ephemeris:** Swiss Ephemeris via [mshafiee/swephgo](https://github.com/mshafiee/swephgo) (CGO)
- **Config:** godotenv + environment variables

---

## Prerequisites

1. **Go** >= 1.25  
   https://go.dev/dl/

2. **C compiler & build tools** (required for CGO)
   - Debian/Ubuntu: `sudo apt install build-essential`
   - Fedora: `sudo dnf groupinstall "Development Tools"`
   - macOS: Xcode Command Line Tools (`xcode-select --install`)

3. **PostgreSQL** (running instance)

4. **Geoapify API key**  
   https://www.geoapify.com/

5. **Swiss Ephemeris C library** (`libswe.so` / `libswe.dylib`) – see detailed steps below.

---

## Critical Setup: Swiss Ephemeris + swephgo

This project depends on **github.com/mshafiee/swephgo**, a CGO binding to the Swiss Ephemeris C library.  
You **must** build and install the native library before `go build` / `go run` will succeed.

### 1. Download Swiss Ephemeris source

```bash
# Official source (recommended)
wget https://www.astro.com/ftp/swisseph/swe_unix_src_2.10.03.tar.gz
# or the latest available from https://www.astro.com/ftp/swisseph/

tar -xzf swe_unix_src_*.tar.gz
cd swe_unix_src_*/src
```

### 2. Build the shared library

Edit the Makefile if needed (common on Linux):

```makefile
libswe.so: $(SWEOBJ)
	$(CC) -shared -o libswe.so $(SWEOBJ) -lm -ldl
```

Then:

```bash
make libswe.so
# or on some systems:
# make
```

### 3. Install the library

```bash
sudo cp libswe.so /usr/local/lib/
# On macOS you may produce libswe.dylib instead

# Make the dynamic linker find it
sudo ldconfig          # Linux
# or on macOS:
# export DYLD_LIBRARY_PATH=/usr/local/lib:$DYLD_LIBRARY_PATH
```

### 4. Verify the Go binding can link

```bash
go get github.com/mshafiee/swephgo@v1.1.0
```

A minimal test:

```go
package main

import (
	"fmt"
	"github.com/mshafiee/swephgo"
)

func main() {
	ver := make([]byte, 12)
	swephgo.Version(ver)
	fmt.Printf("Swiss Ephemeris: %s\n", ver)
}
```

```bash
CGO_ENABLED=1 go run main.go
```

### 5. Ephemeris data files (already included)

This repository ships the required data files under `./ephe/`:

- `sepl_18.se1` – planetary ephemeris
- `semo_18.se1` – lunar ephemeris

---

## Environment Variables

Create a `.env` file in the project root (ignored by git):

```env
# Application
APP_ENV=development          # or "production" (skips loading .env)
GIN_MODE=debug               # or "release"
ALLOWED_ORIGIN=http://localhost:4200

# Database (PostgreSQL)
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=bphs
DB_PORT=5432
DB_SSLMODE=disable

# Auth
SYMMETRIC_KEY=32-byte-or-longer-secret-key-here!!

# External services
GEOAPIFY_API_KEY=your_geoapify_api_key
```

In production set the same variables in the environment; the loader skips `.env` when `APP_ENV=production`.

---

## Running Locally

```bash
# 1. Clone
git clone https://github.com/scientist-v08/astrology.git
cd astrology

# 2. Ensure Swiss Ephemeris library is installed (see above)

# 3. Install Go dependencies
go mod tidy

# 4. Create .env (see template above)

# 5. Make sure PostgreSQL is running and the database exists

# 6. Run (CGO must be enabled)
CGO_ENABLED=1 go run main.go
```

The server listens on `:8080` by default (Gin default).

### Hot reload (optional)

```bash
go install github.com/githubnemo/CompileDaemon@latest
CompileDaemon -command="./bphs"
```

(or build first with `CGO_ENABLED=1 go build -o bphs .`)

---

## API Overview (authenticated routes)

Most endpoints require a valid PASETO token in the `Authorization` header.

| Method | Path                             | Description                           |
| ------ | -------------------------------- | ------------------------------------- |
| POST   | `/v1/...` (user routes)          | Registration / login / session        |
| POST   | `/v1/dob...`                     | Birth-chart + full analysis from DOB  |
| POST   | `/v1/...` (house effects)        | Bhava-lord effects                    |
| POST   | `/v1/sthanaBalaOrKaalaBala` etc. | Shadbala calculations                 |
| POST   | `/v1/pairing...`                 | Couple matching (Ashtakoota + doshas) |
| POST   | `/v1/...` (upagraha)             | Upagrahas & Karakamsha                |
| POST   | `/v1/...` (conditional dasha)    | Conditional dasha                     |

See the `routes/` package for the exact paths.

---

## Project Structure

```
.
├── main.go
├── go.mod / go.sum
├── Dockerfile
├── ephe/                    # Swiss Ephemeris data files (required at runtime)
├── constants/               # Rashi, nakshatra, aspect, exaltation maps, etc.
├── controller/              # HTTP handlers
├── dto/                     # Request / response structs
├── initializers/            # DB connection + env loading
├── middleware/              # Auth middleware
├── models/                  # GORM models (User, Session)
├── routes/                  # Route registration
├── services/
│   ├── swiss-ephemeris/     # Planetary positions (swephgo)
│   ├── shadbala/            # Shadbala engine
│   ├── houses/              # House-effect logic
│   ├── bhavaLords/          # Bhava-lord effects
│   ├── couple/              # Matching & doshas
│   ├── geoapify/            # Geocoding
│   └── ...
├── token/                   # PASETO maker
├── utils/
└── validator/               # Custom Gin validators (rashi, graha, nakshatra)
```

---

## License

AGPL-3.0 – see [LICENSE](LICENSE).

Swiss Ephemeris itself is subject to the licensing terms of Astrodienst AG.  
See https://www.astro.com/swisseph/swephinfo_e.htm for details.

---

## Acknowledgements

- Swiss Ephemeris – Astrodienst AG
- [mshafiee/swephgo](https://github.com/mshafiee/swephgo) – Go bindings
- Classical texts: Brihat Parashara Hora Shastra, Jaimini Sutras, and the works of Gayatri Devi Vasudev & Jaya Shekhar
