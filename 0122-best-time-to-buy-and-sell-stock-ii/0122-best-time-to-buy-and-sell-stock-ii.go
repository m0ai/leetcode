func maxProfit(prices []int) int {
    profit := 0 
    boughtPrice := -1
    hasStock := false
    fmt.Println(prices ,len(prices))
    for day, price := range prices {
        // when end of days
        if hasStock && boughtPrice < price  {
            profit += price
            boughtPrice = -1
            hasStock = false
            fmt.Println("sell", day, price, profit)
        }

        if day == len(prices)-1 && !hasStock {
            continue
        }

        if  !hasStock && price < prices[day+1] {
            profit -= price
            boughtPrice = price
            hasStock = true
            fmt.Println("buy", day, price, profit)
        }
    }
    return profit
}
 
