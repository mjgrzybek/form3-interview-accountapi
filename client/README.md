# Design decisions
- `Client` contains http client, apiUrl and AccountApiService
- `AccountApi` is a field in `Client`, it uses `Client`'s http client and `ApiUrl`
- Requests timeout is supported and should be used as `context`
- Github Actions are used for CI

# Trade-offs
- e2e tests limited to capabilities of provided server
- `mockhttp` is used to mock non-standard (200, 404, ...) error codes
- some possible paths to errors are not covered with tests (coverage = 84%)

# How to test
## On github
Tests are executed as a [github action](https://github.com/mjgrzybek/form3-interview-accountapi/actions/workflows/tests.yml) on every push.

## Locally
Locally tests can be executed two ways:
1. Using `Makefile`
   - `$ make tests`
2. Using `act` to run github actions locally
   - Install [act](https://github.com/nektos/act)
   - `$ act` (in repo root)