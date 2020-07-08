# Write your MySQL query statement below
SELECT
    IFNULL(
        (SELECT DISTINCT Salary
        FROM Employee ORDER BY Salary DESC
        LIMIT 1 OFFSET 1), NULL) AS SecondHighestSalary;

# Testcases:
# {"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100], [2, 200], [3, 300]]}}
# {"headers": {"Employee": ["Id", "Salary"]}, "rows": {"Employee": [[1, 100]]}}