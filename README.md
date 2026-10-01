# Go Practice

Notes and exercises from working through the Go basics — packages, variables, types,
functions, and flow control. Every file is a small, self-contained lesson with the
explanation written in comments above the code.

## Running the examples

Run these from the project root (the folder that contains `go.mod`):

| Topic | Command |
| --- | --- |
| Packages, imports, exported names | `go run .` |
| Functions and return values | `go run ./package_variables_functions` |
| Variables | `go run ./package_variables_functions/variables` |
| Types, conversion, constants | `go run ./package_variables_functions/types` |
| for, if, switch | `go run ./flow_controls` |

Each folder has a `main.go` that calls every lesson in that folder and prints a header
before each one, so the output is easy to match back to the file it came from:

```
--- for_1.go: the basic for loop ---
45
--- for_3.go: for as a while loop ---
1024
```

### Running just one example

Open that folder's `main.go` and comment out the calls you don't want:

```go
func main() {
	fmt.Println("--- for_1.go: the basic for loop ---")
	forBasic()

	// fmt.Println("--- for_2.go: init and post are optional ---")
	// forOptionalInitPost()
}
```

### Building instead of running

```powershell
go build ./...      # compile everything, report errors, keep no binaries
go vet ./...        # catch likely mistakes the compiler allows
go fmt ./...        # reformat to standard Go style
```

## Project layout

```
practice/
├── go.mod                            module definition
├── main.go                           packages + entry point for the root folder
├── imports.go                        import syntax
├── exportedNames.go                  why exported names start with a capital letter
├── package_variables_functions/
│   ├── main.go                       entry point for this folder
│   ├── functions.go                  parameters, multiple return values
│   ├── namedReturnValue.go           named return values and naked return
│   ├── variables/
│   │   ├── main.go                   entry point for this folder
│   │   ├── variables.go              var keyword, zero values
│   │   ├── variables2.go             omitting the type when you initialize
│   │   └── variables3.go             short declaration with :=
│   └── types/
│       ├── main.go                   entry point for this folder
│       ├── basicTypes.go             Go's basic types
│       ├── typeInference.go          type inference
│       ├── typeConversion.go         converting between types
│       └── constants.go              constants
└── flow_controls/
    ├── main.go                       entry point for this folder
    ├── for_1.go … for_4.go           the for loop in its four forms
    ├── if_1.go … if_3.go             if, if with a short statement, if/else
    └── switch_1.go … switch_3.go     switch
```

## What is `go.mod`?

`go.mod` is the file that marks the root of a **module** — Go's unit of a project. It's
created once with `go mod init <name>`, and this project's looks like:

```
module practice

go 1.27.1
```

Three lines, three jobs:

- `module practice` names the project. Import paths are built from it, so the variables
  folder is importable as `practice/package_variables_functions/variables`.
- `go 1.27.1` records the language version to compile against.
- A `require` block would list third-party dependencies and their exact versions. This
  project has none yet, since everything here comes from the standard library.

### Why this project needs one

Without a `go.mod`, Go can't resolve a folder into a package, so `go run ./flow_controls`
fails with *"go.mod file not found in the current directory or any parent directory"*. You'd
be stuck naming every file by hand. The editor's Go extension also degrades without a
module and starts reporting unreliable errors.

If you ever add a dependency, `go mod tidy` adds it to `go.mod` and records its checksum
in a `go.sum` file. Commit both.

## How Go organizes code, and the mistakes it causes

These two rules explain almost every confusing error in a project like this one.

**One folder is one package.** Go compiles a whole directory at once, not one file at a
time, and all files in it share a single namespace. Subfolders are separate packages —
`flow_controls` and `types` have nothing to do with each other.

**A package has exactly one `main`.** That's the entry point, and there can only be one
per folder. This is why each folder here has a single `main.go` and the lesson files
expose named functions like `forBasic()` and `switchCase()` instead.

The practical consequence: two files in the same folder can't both declare something with
the same name. That applies to functions, package-level variables, everything. Where that
happened here, one side was renamed — `if_2.go` has `powShort` because `if_3.go` already
has `pow`, and `namedReturnValue.go` has `addNamed` because `functions.go` already has `add`.
Package-level variables are shared the other way around: `variables3.go` uses the `i` and
`j` declared over in `variables2.go`, because one package means one set of them.

### Errors you'll likely hit

**`main redeclared in this block`** — two files in the same folder both declare `func main()`.
Rename one and call it from the folder's `main.go`.

**`function main is undeclared in the main package`** — you ran a file that has no `main`
function. Run the whole folder instead, or check you're pointing at the right place.

**`package command-line-arguments is not a main package`** — the file says something like
`package variables` instead of `package main`. Only `package main` can be executed.

**`undefined: someFunction`** — the function lives in a different folder, which means a
different package. Either move it, or import its package by its full module path.

## Requirements

Go 1.27.1 or newer, and the Go extension for the editor (which runs `gopls` and reports
these errors inline as you type).
