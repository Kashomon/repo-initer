package repofiles

const readmeTmpl = `
# {{.Name}}

{{.Description}}

## Ownership

This repository is owned by Kubernetes organization on Github, which is in turn
managed by the [Cloud Native Foundation or CNF](LINK ME).

This repository is sponsored and maintained by the following SIG or Special
Interest Group: {{.SIG}}.

You can check out the primary maintainers via the [OWNERS.yaml](OWNERS.yaml) file,
which specifies how reviews get routed and who has the power to approve merges.

## Code-of-conduct

All that contribute to and participate agree to abide by the [Kubernetes Code of
Conduct](LINK ME)

## Contributing

## LICENSE

This repository uses the Apache 2.0 License, and can be [found here](LICENSE).
`

func Readme(c *Config) *RepoFile {
	return &RepoFile{
		Contents: templatize(readmeTmpl, c),
		Name:     "README.md",
	}
}
