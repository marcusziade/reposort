package reposort

import (
	"os"
	"path/filepath"
	"strings"
)

func detectMainLanguage(repoPath string) string {
	languageCount := make(map[string]int)

	filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			ext := strings.TrimPrefix(filepath.Ext(path), ".")
			language := getLanguageByExtension(ext)

			if language != "" {
				languageCount[language]++
			}
		}

		return nil
	})

	var mainLanguage string
	var maxCount int

	for language, count := range languageCount {
		if count > maxCount {
			maxCount = count
			mainLanguage = language
		}
	}

	if mainLanguage == "" {
		mainLanguage = "Other"
	}

	return mainLanguage
}

func getLanguageByExtension(ext string) string {
	extensions := map[string]string{
		"swift": "Swift",
		"go":    "Go",
		"rs":    "Rust",
		"zig":   "Zig",
		"js":    "JavaScript",
		"ts":    "TypeScript",
		"py":    "Python",
		"rb":    "Ruby",
		"java":  "Java",
		"cpp":   "C++",
		"c":     "C",
	}

	return extensions[ext]
}
