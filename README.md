LabAssistant!
================================================================================

A Go library for carefully refactoring critical paths and code.
It's a lightweight port of Github's [Scientist](https://github.com/github/scientist)
and with inspiration from the python port [Laboratory](https://github.com/joealcorn/laboratory).

Why?
--------------------------------------------------------------------------------

See GitHub's blog post — http://githubengineering.com/scientist/

Status
--------------------------------------------------------------------------------

Project has been shut down, if anyone want to take it over please go see [issue #2](https://github.com/lmas/labassistant/issues/2).

~~The library is currently in an alpha stage and under development.~~

~~Main features have been implemented, along with basic test coverage. It haven't
been tested in production yet, so beware.~~

Installation
--------------------------------------------------------------------------------

`go get github.com/lmas/labassistant`.

Usage
--------------------------------------------------------------------------------

For now, please see the `example.go` file for how to use the library.

Contribution
--------------------------------------------------------------------------------

Any and all contributions are welcome. Just make sure to run `go fmt` and
`go test` on any code in a new pull request.

License
--------------------------------------------------------------------------------

MIT License, see the LICENSE file.

TODO
--------------------------------------------------------------------------------

Library:
- Ditch the reflect magic and make a code gen tool instead?
- Remove the panic'ing code and return errors instead.
- Default publish functions? One simple CLI output and one web page graph.

Misc:
- Show feature list.
- More documentation and usage examples (especially in this file).
