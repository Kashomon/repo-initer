package repofiles

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

const ownersName = "OWNERS"

var ownersTmpl = `
assignees:{{range .}}
  - {{.}}{{end}}
`

// Owners creates an OWNERS YAML file that specifies who has the power to
// perform reviews and give submit approvals.
func Owners(owners []string) *RepoFile {
	return &RepoFile{
		Contents: templatize(buf, owners),
		Name:     ownersName,
	}
}
