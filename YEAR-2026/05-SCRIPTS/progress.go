go
// YEAR-2026/05-SCRIPTS/progress.go
package main

import (
    "fmt"
    "io/ioutil"
    "strings"
)

func main() {
    data, _ := ioutil.ReadFile("../02-CHECKLIST.md")
    content := string(data)
    
    total := strings.Count(content, "- [ ]") + strings.Count(content, "- [x]")
    done := strings.Count(content, "- [x]")
    
    percent := 0
    if total > 0 {
        percent = int(float64(done) / float64(total) * 100)
    }
    
    fmt.Printf("✅ Прогресс: %d%% (%d/%d задач)\n", percent, done, total)
}