# TODO
- [ ] http client with context
- [x] input validation
- [x] add `defer` where needed to close readers etc.
# Testing
## unit
- [ ] unit tests full coverage
## e2e
### basic
- [ ] success scenarios
- [ ] failing scenarios
### performance ❓
- flood of operations
  - [ ] same type
  - [ ] mixed (create, fetch, delete)
## fuzzy ❓
- [ ] learn about it and try to apply

# Others
- [ ] separate structs for different methods to avoid sending multiple empty fields
  - check if it affects perf - less complicated code is preferred if perf is simliar