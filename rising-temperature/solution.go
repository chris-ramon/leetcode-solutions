-- @lc code=start
SELECT w1.id as 'id'
FROM Weather AS w1
INNER JOIN Weather AS w2
    ON DATEDIFF(w1.recordDate, w2.recordDate) = 1
    AND w1.temperature > w2.temperature
-- @lc code=end

/*
TestCases
{"headers": {"Weather": ["id", "recordDate", "temperature"]}, "rows": {"Weather": [[1, "2015-01-01", 10], [2, "2015-01-02", 25], [3, "2015-01-03", 20], [4, "2015-01-04", 30]]}}
{"headers": {"Weather": ["id", "recordDate", "temperature"]}, "rows": {"Weather": [[2, "2015-01-02", 25], [3, "2015-01-03", 20], [4, "2015-01-04", 30]]}}
*/