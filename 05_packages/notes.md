- every go file must be apart of a package
- logically, if making a separate package with utility functions
  you would want to ensure that the funcs and var names are generic
- different packages must be in own directory
  - convention to name directory as package name
- in go, capitalized func names are exported
  - this is why we call fmt.Print()
  - so basically we don't have an export keyword like in Ts, but the casing is what exports

- go mod file stores packages like package.json
- `go get` same as `pnpm install`
- install packages with go get link
