DELETE
FROM Person
WHERE id NOT IN (
    SELECT unique_emails.id
    FROM (
             SELECT MIN(id) as id
             FROM Person
             GROUP BY(email)
         ) unique_emails
);


/*
TestCases

{"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[1, "john@example.com"], [2, "bob@example.com"], [3, "john@example.com"]]}}

{"headers": {"Person": ["id", "email"]}, "rows": {"Person": [[3, "john@example.com"], [2, "bob@example.com"], [1, "john@example.com"]]}}

*/
