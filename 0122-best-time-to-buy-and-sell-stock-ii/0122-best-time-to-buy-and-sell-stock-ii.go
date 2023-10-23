func maxProfit(prices []int) int {
    profit := 0 
    for i := 0; i < len(prices)-1; i ++ {
        if (prices[i] < prices[i+1]){
            profit += prices[i+1]-prices[i]

        }
    }
    return profit

    // for day, price := range prices {
    //     // when end of days
    //     // sell
    //     if hasStock && boughtPrice < price  {
    //         profit += price
    //         boughtPrice = -1
    //         hasStock = false
    //     }

    //     if day == len(prices)-1 && !hasStock {
    //         continue
    //     }

    //     // buy
    //     if !hasStock && price < prices[day+1] {
    //         profit -= price
    //         boughtPrice = price
    //         hasStock = true
    //     }
    // }
    // return profit
}
 
