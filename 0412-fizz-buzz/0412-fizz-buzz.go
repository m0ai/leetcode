import "strconv"
func fizzBuzz(n int) []string { 
  var result []string;
  
  for i := 1; i <= n; i++ {
    var num string 

    if i % 3 == 0 {
      num = "Fizz"
    } else if i % 5 == 0 {
      num = "Buzz"
    } else {
      num = strconv.Itoa(i)
    }

    if i % 3 == 0 && i % 5 == 0 {
      num = "FizzBuzz"
    }

    result = append(result, num)
  }  

  return result
}