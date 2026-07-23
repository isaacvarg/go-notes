## packages

- main is reserved package name
  - tells go this is entry point to app building
- package name can be anything
  - think of it it kind of like name spaces 
- go requires at least 1 package
- can have many package
- packages can be split into multiple files

## modules

- modules consists of multiple packages
- simply a go 'project'
- `go mod init` turns project into a module
- go can build a module into executable

## function

- `func main() {}` ran on app start
- project can only have one main

## conventions

- snake_case for filenames
- kebab-case for modules (think url)
- camelCase for vars
