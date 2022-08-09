# go workspaces

New features from `Go 1.18`, `workspaces`, improve Go multi-modules workflow with built-in `go work` commands.

## usage

- `go work sync` syncs the workspace's build list back to the workspace's modules.
- `go work use <module path>` uses module in workspaces - `OPTION` can use `-r` to recursively uses modules.
