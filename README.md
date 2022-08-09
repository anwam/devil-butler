# go workspaces

new capability of Go, workspaces.
empower Go multi-modules repo with built-in `go work` commands

## usage

`go work sync` syncs the workspace's build list back to the workspace's modules
`go work use <module path>` uses module in workspaces - `OPTION` can use `-r` to recursively uses modules
