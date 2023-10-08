import "strings"
func isAnagram(s string, t string) bool {
    if  len(t) != len(s) {
        return false
    }

    result := true
    for i := 0; i < len(t); i++ {
        ch := string(t[i])
        if strings.Count(s, ch) != strings.Count(t, ch) {
            result = false
            break
        } 
    }
    return result
}