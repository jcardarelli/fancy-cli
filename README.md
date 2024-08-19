# fancy-cli
CLI to keep track of your favorite fancy restaurants in a SQLite database

![image](https://github.com/jcardarelli/fancy-cli/assets/1383816/46d9618f-755b-44a7-a669-7dc38b0ce203)


Features:
* View all restaurants
* Add new restaurant
* Delete restaurant


## Installation
1. `$ ./install.sh`
## Data
All data is stored in the file fancy-cli.db

## Testing
* CI on GitHub Actions:
  * Run golangci-lint
  * Build docker image
* Run GitHub Actions locally with `$ make local-github-actions`
  * This requires `act`
  * Install with `$ brew install act` on MacOS

## Development
1. `$ python3 -m pip install pre-commit`
1. `$ pre-commit`
