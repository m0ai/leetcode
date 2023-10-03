
import (
    "fmt"
    "strings"
)

func canConstruct(ransomNote string, magazine string) bool {
    result := false
    for _, ch := range magazine {
        i := strings.Index(ransomNote, string(ch))
        if i > -1 {
            // remove index obj
            fmt.Println(ransomNote)
            ss := []string{ransomNote[:i], ransomNote[i+1:]}
            ransomNote = strings.Join(ss, "")
        }

        if len(ransomNote) == 0 {
            result = true
        }
    }
    return result
}