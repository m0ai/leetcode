func maxProfit(prices []int) int {
    profit := 0
    minPrice := prices[0]

    for i := 1; i < len(prices); i += 1 {

        if prices[i] < minPrice {
            minPrice = prices[i]
        } else if profit < prices[i] - minPrice{
            profit = prices[i] - minPrice
        }
    }
    return profit

    // myProfit := 0 
    // for i, j := 0, 1; i < len(prices); i, j = i+1, i+2 {
    //     for j < len(prices) {
    //         if prices[j] - prices[i] > myProfit{
    //             myProfit = prices[j] - prices[i]
    //         }
    //         j += 1
    //     }
    // }
    // return myProfit
}