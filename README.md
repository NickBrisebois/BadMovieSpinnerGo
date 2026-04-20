# Bad Movie Spinner (Go Edition)

Largely Go-based rewrite of my BadMovieSpinner originally done in Typescript


## Configuration

Configuration for both API and Spinner are done via environment variables:

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


### Spinner required variables:
```
# Server details for connecting to backend API
SERVER_HOST=localhost
SERVER_PORT=8080
```
