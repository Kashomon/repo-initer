package repofiles

const releaseFilename = "RELEASE"

const releaseContents = "TODO:add contents here"

// Release createsa a RELEASE file
func Release() *RepoFile {
	return &RepoFile{
		Contents: releaseContents,
		Name:     releaseFilename,
	}
}
