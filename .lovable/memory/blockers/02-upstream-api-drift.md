# Blocker: Upstream API Drift (suspected)

Pre-existing drift against (now-removed, see Task M) `core-v9 v1.5.8`.
`core-v9` itself was removed from `go.mod`/`go.sum` because no `.go` file
imported it — but matching drift may still affect other packages.

## Suspect signatures

- `corestr.New.LinkedCollections`
- `converters.StringToIntegerWithDefault`
- `coredynamic.SliceItemsAsStringsAny`
- `errwrappers.NewEmpty`
- `errtype.InvalidValidate`
- `errnew.Type.Message`
- `errnew.NotFound.Simple`

## Known real bug

`errwrappers/Collection.go:1216` — a `sync.noCopy` violation flagged by `go vet`.
Outside agent scope without the failure log; needs targeted patch.

## Resolution path

1. Get build-errors.txt (see `01-build-errors-log.md`).
2. Confirm which signatures are actually referenced.
3. Patch call sites in-repo OR coordinate with the upstream library owner.
