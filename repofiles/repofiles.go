// Package repofiles creates the actual repo files that should be created
// in a specific repository.
package repofiles

import "io"

type Config struct {
	// The name of the
	Name        string
	Description string
	GithubPath  string
	Owners      []string
	Description string
	SIG         string
	License     string
}

// RepoFile represents a file that will be initialized into the repository.
type RepoFile struct {
	Contents string
	Name     string
}

// Templatize is a simple helpers which fills out a template in-memory, returning
// the raw string. If any error occurs, it panics.
func templatize(rawTpml string, data interface{}) string {
	tmpl, terr := template.New("owners").Parse(strings.TrimSpace(rawTpml))
	if terr != nil {
		panic(fmt.Sprintf("Programming Error! Template didn't parse. %v", terr))
	}

	buf := new(bytes.Buffer)
	err := tmpl.Execute(buf, data)
	if err != nil {
		// This shouldn't error, but we might as well blow up if it does.
		panic(fmt.Sprintf("Programming Err! Couldn't parse template. %v", err))
	}
	return buf.String()
}

// TODO(Kashomon): Add writer support
