# GEMINI.md

## Project Overview

This project, `llama-swap`, is a lightweight, transparent proxy server written in Go. Its primary purpose is to provide automatic model swapping for `llama.cpp`'s server. It is designed to be easy to deploy as a single binary with no dependencies, and it is configured using a single YAML file.

The proxy intercepts OpenAI-compatible API requests, inspects the `model` parameter, and ensures the correct `llama.cpp` server configuration is running to handle the request. If a different model is required, `llama-swap` will automatically stop the current server and start the correct one.

The project also includes a web-based UI for real-time monitoring of logs and models.

## Building and Running

### Building the Project

The project uses a `Makefile` for building the application. The following commands can be used:

*   `make all`: Builds the application for macOS and Linux.
*   `make mac`: Builds the application for macOS (ARM64).
*   `make linux`: Builds the application for Linux (AMD64 and ARM64).
*   `make windows`: Builds the application for Windows (AMD64).
*   `make clean`: Removes the build directory.

The compiled binaries are placed in the `build/` directory.

### Running the Project

To run `llama-swap`, you need a configuration file. An example is provided in `config.example.yaml`.

The main command to run the application is:

```bash
./build/llama-swap-darwin-arm64 --config path/to/config.yaml --listen localhost:8080
```

Replace `./build/llama-swap-darwin-arm64` with the appropriate binary for your system.

The following command-line flags are available:

*   `--config`: Path to the configuration file (default: `config.yaml`).
*   `--listen`: Address and port to listen on (default: `:8080`).
*   `--version`: Show version information and exit.
*   `--watch-config`: Automatically reload the configuration file when it changes.

The project can also be run using Docker. See the `README.md` for more details.

## Development Conventions

*   **Dependency Management:** The project uses Go modules for dependency management. Dependencies are listed in the `go.mod` file.
*   **Testing:** The `Makefile` includes targets for running tests:
    *   `make test`: Runs the tests in short mode.
    *   `make test-all`: Runs all tests.
*   **Configuration:** The application is configured via a YAML file. The `config.example.yaml` file provides a comprehensive overview of the available configuration options.
*   **UI:** The web UI is a React application located in the `ui/` directory. It is built using `npm run build` and the output is embedded in the Go binary.
