package reposort

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSortRepositories(t *testing.T) {
	tempDir := t.TempDir()

	createRepo(t, tempDir, "repo1", []string{"main.go", "utils.go"})
	createRepo(t, tempDir, "repo2", []string{"main.py", "utils.py"})
	createRepo(t, tempDir, "repo3", []string{"main.js", "utils.js"})
	createRepo(t, tempDir, "repo4", []string{"main.rs", "utils.rs"})
	createRepo(t, tempDir, "repo5", []string{"main.txt", "utils.txt"})

	err := SortRepositories(tempDir)
	if err != nil {
		t.Fatalf("SortRepositories failed: %v", err)
	}

	expectedDirs := []string{"Go", "Python", "JavaScript", "Rust", "Other"}
	for _, dir := range expectedDirs {
		dirPath := filepath.Join(tempDir, dir)
		if _, err := os.Stat(dirPath); os.IsNotExist(err) {
			t.Errorf("Directory %q not created", dir)
		}
	}

	expectedRepoMoves := []struct {
		repoName string
		langDir  string
	}{
		{"repo1", "Go"},
		{"repo2", "Python"},
		{"repo3", "JavaScript"},
		{"repo4", "Rust"},
		{"repo5", "Other"},
	}

	for _, move := range expectedRepoMoves {
		repoPath := filepath.Join(tempDir, move.langDir, move.repoName)
		if _, err := os.Stat(repoPath); os.IsNotExist(err) {
			t.Errorf("Repository %q not moved to %q directory", move.repoName, move.langDir)
		}
	}
}
