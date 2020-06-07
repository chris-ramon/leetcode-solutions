/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    var min = null;
    var maxProfit = 0;
    for(var i = 0; i < prices.length; i++) {
        var p = prices[i];
        min = min === null ? p : Math.min(min, p);
        var profit = p - min;
        maxProfit = Math.max(maxProfit, profit);
    }
    return maxProfit;
};