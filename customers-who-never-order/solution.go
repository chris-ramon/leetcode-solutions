-- @lc code=start
SELECT name AS 'Customers'
FROM Customers
LEFT JOIN Orders ON Orders.customerId = Customers.id
GROUP BY Customers.id
HAVING COUNT(Orders.id) = 0;
-- @lc code=end

/*
TestCases
{"headers": {"Customers": ["id", "name"], "Orders": ["id", "customerId"]}, "rows": {"Customers": [[1, "Joe"], [2, "Henry"], [3, "Sam"], [4, "Max"]], "Orders": [[1, 3], [2, 1]]}}
{"headers": {"Customers": ["id", "name"], "Orders": ["id", "customerId"]}, "rows": {"Customers": [[1, "James"], [2, "James"]], "Orders": []}}
*/