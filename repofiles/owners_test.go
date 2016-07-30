package repofiles

import (
	"strings"
	"testing"
)

func TestOwners(t *testing.T) {
	o := []string{"Ford", "Prefect", "Zaphod"}
	repofile := Owners(o)
	for _, s := range o {
		substr := "  - " + s
		if !strings.Contains(repofile.Contents, substr) {
			t.Fatalf("Wanted %v to contain %v", repofile.Contents, substr)
		}
	}
}
