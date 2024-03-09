package reposort

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDetectMainLanguage(t *testing.T) {
	tempDir := t.TempDir()

	createRepo(t, tempDir, "repo1", []string{"main.go", "utils.go"})
	createRepo(t, tempDir, "repo2", []string{"main.py", "utils.py"})
	createRepo(t, tempDir, "repo3", []string{"main.js", "utils.js"})
	createRepo(t, tempDir, "repo4", []string{"main.rs", "utils.rs"})
	createRepo(t, tempDir, "repo5", []string{"main.txt", "utils.txt"})

	testCases := []struct {
		repoPath string
		expected string
	}{
		{filepath.Join(tempDir, "repo1"), "Go"},
		{filepath.Join(tempDir, "repo2"), "Python"},
		{filepath.Join(tempDir, "repo3"), "JavaScript"},
		{filepath.Join(tempDir, "repo4"), "Rust"},
		{filepath.Join(tempDir, "repo5"), "Other"},
	}

	for _, tc := range testCases {
		result := detectMainLanguage(tc.repoPath)
		if result != tc.expected {
			t.Errorf("detectMainLanguage(%q) = %q; want %q", tc.repoPath, result, tc.expected)
		}
	}
}

func TestGetLanguageByExtension(t *testing.T) {
	testCases := []struct {
		ext      string
		expected string
	}{
		{"go", "Go"},
		{"py", "Python"},
		{"js", "JavaScript"},
		{"rs", "Rust"},
		{"txt", ""},
	}

	for _, tc := range testCases {
		result := getLanguageByExtension(tc.ext)
		if result != tc.expected {
			t.Errorf("getLanguageByExtension(%q) = %q; want %q", tc.ext, result, tc.expected)
		}
	}
}

func createRepo(t *testing.T, parentDir, repoName string, files []string) {
	repoPath := filepath.Join(parentDir, repoName)
	err := os.Mkdir(repoPath, os.ModePerm)
	if err != nil {
		t.Fatalf("Failed to create repository directory: %v", err)
	}

	for _, file := range files {
		filePath := filepath.Join(repoPath, file)
		_, err := os.Create(filePath)
		if err != nil {
			t.Fatalf("Failed to create file %q: %v", file, err)
		}
	}
}
