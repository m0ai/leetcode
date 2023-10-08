import "strings"
func isAnagram(s string, t string) bool {
    if  len(t) != len(s) {
        return false
    }

    count := make([]int, 26) // lowerase alphabet count 
    for _, ch := range s {
        count[int(ch - 'a')]++
    }

    for _, ch := range t {
        i := int(ch - 'a') 
        count[i]--
        if count[i] < 0 {
            return false
        }
    }
    return true
}