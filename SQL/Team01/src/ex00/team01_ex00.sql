WITH
cte_last_updated_currency
	AS
	(
		SELECT
			last_updated_c.id,
			last_updated_c.name,
			c.rate_to_usd,
			last_updated_c.updated
		FROM
			(SELECT
			 	id,
				name,
				MAX(updated) AS updated
			 FROM
				currency
			 GROUP BY
			 	id,
				name) AS last_updated_c
		LEFT JOIN
			currency AS c
			ON
				c.id = last_updated_c.id AND
				c.updated = last_updated_c.updated
	),
cte_base_table
	AS
	(
		SELECT
			b.user_id,
			b.type,
			SUM(money) AS volume,
			c.id AS currency_id
		FROM
				balance AS b
			FULL JOIN
				cte_last_updated_currency AS c
				ON
					c.id = b.currency_id
		GROUP BY
			b.type,
			c.id,
			b.user_id
	)

SELECT
	COALESCE(u.name, 'not defined') AS name,
	COALESCE(u.lastname, 'not defined') AS lastname,
	base_table.type,
	base_table.volume,
	COALESCE(c.name, 'not defined') AS currency_name,
	COALESCE(c.rate_to_usd, 1) AS last_rate_to_usd,
	base_table.volume * COALESCE(c.rate_to_usd, 1) AS total_volume_in_usd
FROM
		cte_base_table AS base_table
	FULL JOIN
		"user" AS u
		ON
			u.id = base_table.user_id
	FULL JOIN
		cte_last_updated_currency AS c
		ON
			c.id = base_table.currency_id
ORDER BY
	name DESC,
	lastname,
	type
