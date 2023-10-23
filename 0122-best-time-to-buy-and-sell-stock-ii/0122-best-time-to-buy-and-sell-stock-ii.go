func maxProfit(prices []int) int {
    profit := 0 
    hasStock := false
    boughtPrice := 0
    for day, price := range prices {
        // when end of days
        if hasStock && boughtPrice < price  {
            profit += price
            hasStock = false
        }

        if day == len(prices)-1 && !hasStock {
            continue
        }

        if  !hasStock && price < prices[day+1] {
            profit -= price
            boughtPrice = price
            hasStock = true
        }
    }
    return profit
}
 
