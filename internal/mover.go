package reposort

import (
    "fmt"
    "os"
    "path/filepath"
    "time"
)

func SortRepositories(path string) error {
    files, err := os.ReadDir(path)
    if err != nil {
        return fmt.Errorf("error reading directory: %v", err)
    }

    languageDirs := make(map[string]string)

    done := make(chan bool)
    go func() {
        characters := []string{"|", "/", "-", "\\"}
        for {
            for _, char := range characters {
                fmt.Printf("\rSorting repositories %s", char)
                time.Sleep(100 * time.Millisecond)
            }
            select {
            case <-done:
                fmt.Print("\rSorting completed! \n")
                return
            default:
            }
        }
    }()

    for _, file := range files {
        if file.IsDir() {
            repoPath := filepath.Join(path, file.Name())
            language := detectMainLanguage(repoPath)
            languageDir := filepath.Join(path, language)
            if _, ok := languageDirs[language]; !ok {
                err := os.Mkdir(languageDir, os.ModePerm)
                if err != nil {
                    fmt.Printf("Error creating directory: %v\n", err)
                    continue
                }
                languageDirs[language] = languageDir
            }

            oldPath := repoPath
            newPath := filepath.Join(languageDirs[language], file.Name())
            err := os.Rename(oldPath, newPath)
            if err != nil {
                fmt.Printf("Error moving repository: %v\n", err)
            }
        }
    }

    done <- true
    return nil
}