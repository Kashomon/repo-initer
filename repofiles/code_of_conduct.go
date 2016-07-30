package repofiles

// TODO(Kashomon): Is this link accurate?
const k8sCodeOfConduct = `
Contributors to this repo abide by by the [Kubernetes Comunity Code of Conduct](https://github.com/kubernetes/kubernetes/tree/master/code-of-conduct.md). Check it out!
`

// CodeOfConduct creates a code-of-conduct markdown file which references the
// Kubernetes code-of-conduct file.
func CodeOfConduct() *RepoFile {
	return &RepoFile{
		Contents: k8sCodeOfConduct,
		Name:     "code-of-conduct.md",
	}
}
