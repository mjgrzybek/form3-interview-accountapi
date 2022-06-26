# TODO
- [x] http ~~client~~ requests with context
- [x] ~~input validation~~
- [x] add `defer` where needed to close readers etc.
# Testing
## unit
- [x] unit tests full coverage ~ 81%
## e2e
- [ ] packt.io ?
### basic
- [x] success scenarios
- [x] failing scenarios
### ~~performance~~
Client's side performance should not be an issue until proven.
## fuzzy ❓
- [ ] learn about it and try to apply

# Others
- [ ] separate structs for different methods to avoid sending multiple empty fields
  - check if it affects perf - less complicated code is preferred if perf is similar