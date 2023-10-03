
import (
    "fmt"
    "strings"
)


func canConstruct(ransomNote string, magazine string) bool {
    // runtime 7ms , memory 3.85mb
    for _, v := range ransomNote {
        if strings.Count(ransomNote, string(v)) > strings.Count(magazine, string(v))  {
            return false
        }
    }
    return true
}
func oldCanConstruct(ransomNote string, magazine string) bool {
    // runtime 12ms, mem 6.84mb
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