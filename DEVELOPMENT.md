## Development
Requires Go 1.13+.
```
$ make build
$ make test
```

### Releases
This project is using [goreleaser](https://goreleaser.com). GitHub release creation is automated using GitHub
actions. New releases are automatically created when new tags are pushed to the repo.
```
$ TAG=v0.0.2 make tag
```

How to manually create a release without relying on GitHub actions.
```
$ TAG=v0.0.2 make tag
$ export GITHUB_TOKEN=xxx
$ make clean && make release
```