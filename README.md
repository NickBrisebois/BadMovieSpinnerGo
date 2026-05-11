# Bad Movie Spinner (Go Edition)

A purely Go implementation of a (bad) movie spinner with data pulled from Google Sheets and TheMovieDatabase (TMDB). The frontend spinner application can be exported to either Linux/Desktop or web when built to target WASM

There are three main components to the project:

  - An API that handles pulling and caching movie data from Google Sheets
  - An Ebitengine-based project that contains the whole spinner that can be compiled to both Linux as a native application or WASM to be served on a website
  - A web server that hosts the WASM build along with some wrapper HTML

Each component has its own build target in the Makefile

## Configuration

Configuration for both API and Spinner is primarily done via environment variables, however, for web binaries configuration is done by injecting values during build

### API required variables:
```
SERVER_HOST=localhost
SERVER_PORT=8080

# gcpserviceaccountkey.json needed for accessing private google sheets
GOOGLE_SERVICE_ACCOUNT_KEY_PATH=./gcpserviceaccountkey.json
GOOGLE_SHEET_ID=1UASDBDSFSDFSDF_etcetcetc
GOOGLE_SCOPES=https://www.googleapis.com/auth/spreadsheets,https://www.googleapis.com/auth/drive.file

TMDB_ACCESS_TOKEN=[...]  # get here: https://www.themoviedb.org/settings/api

# Where movie posters are cached to on the API side
IMAGE_CACHE_DIR=cache
```


### Spinner (Desktop/Linux) required variables:
```
# Server details for connecting to backend API
SERVER_HOST=localhost
SERVER_PORT=8080
```

### Spinner (Web/WASM) required variables:
```
# Host and Port details for hosting frontend web server containing WASM spinner
WEB_HOST=localhost
WEB_PORT=8000
```

Note: Web/WASM spinner requires the API information to be injected at build time with the build flags:

`-tags="js wasm" -ldflags="-X 'main.APIHost=$(APP_API_HOST)' -X 'main.APIPort=$(APP_API_PORT)'"`

This is automated with `make build-web` or `make build-wasm` where it defaults to `badmovie2.api.acid1.xyz:443` or can be overriden by setting `APP_API_HOST` and `APP_API_PORT`


## Makefile help
```Makefile
Usage: make [target]
Targets:
build-api                      Build the spinner's backend API
build-linux                    Build the spinner as a Linux binary
build-wasm                     Build the spinner WASM binary
build-web                      Build fully self-contained front-end web server binary with WASM spinner
clean                          Clean up builds and reset to a clean state
debug                          Run the spinner with live reload as a linux binary through delve (see `.air-spinner.toml` for debugger connection details)
format                         Format codebase
gen-docs                       Generate API documentation
help                           Print this help
install-api                    Copy API binary to install target directory
install-linux                  Copy Linux binary to install target directory
install-web                    Copy web binary to install target directory
run-api                        Run the spinner API with live reload with delve debugger
run                            Run the spinner linux binary with live reload but no debugger
run-web                        Run the full web-based frontend
```
