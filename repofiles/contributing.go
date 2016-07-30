package repofiles

const contribFilename = "CONTRIBUTING.md"

const contribContents = `
# Contributing Guidelines

Want to hack on Kubernetes? Yay! [Check out our contributor guidelines!](
https://github.com/kubernetes/kubernetes/tree/master/CONTRIBUTING.md)
`

func Contributing() *RepoFile {
	return &RepoFile{
		Contents: contribContents,
		Name:     contribFilename,
	}
}
