insert into currency values (100, 'EUR', 0.85, '2022-01-01 13:29');
insert into currency values (100, 'EUR', 0.79, '2022-01-08 13:29');

SELECT
    COALESCE("user".name, 'not defined') AS name,
    COALESCE("user".lastname, 'not defined') AS lastname,
    currency2.name AS currency_name,
    CASE
        WHEN (SELECT rate_to_usd
            FROM currency
            WHERE currency.name = currency2.name AND balance.updated >= currency.updated
            ORDER BY currency.updated DESC
            LIMIT 1) IS NOT NULL THEN
            CAST((SELECT rate_to_usd
                    FROM currency
                    WHERE currency.name = currency2.name AND balance.updated >= currency.updated
                    ORDER BY currency.updated DESC
                    LIMIT 1) * balance.money AS REAL)
        ELSE CAST((SELECT rate_to_usd
                    FROM currency
                    WHERE currency.name = currency2.name AND balance.updated < currency.updated
                    ORDER BY currency.updated
                    LIMIT 1) * balance.money AS REAL)
    END AS currency_in_usd
FROM balance
LEFT JOIN "user" ON "user".id = balance.user_id
FULL JOIN (SELECT DISTINCT id, name FROM currency) AS currency2 ON currency2.id = balance.currency_id
WHERE currency2.name IS NOT NULL
ORDER BY name DESC, lastname ASC, currency_name ASC;