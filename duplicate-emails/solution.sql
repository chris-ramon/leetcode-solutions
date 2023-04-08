-- @lc code=start
SELECT email AS 'Email'
FROM Person
GROUP BY email
HAVING COUNT(email) > 1;
-- @lc code=end

/*
TestCases
{"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[1, "a@b.com"], [2, "c@d.com"], [3, "a@b.com"]]}}\n{"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[1, "a@b.com"], [2, "a@d.com"], [3, "a@b.com"]]}}\n{"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[1, "a@b.com"], [2, "a@d.com"], [3, "a@b.com"],[4, "aa@b.com"],[5, "aaa@b.com"]]}}
*/