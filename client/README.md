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