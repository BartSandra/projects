SELECT select1.name, coalesce(select1.count, 0) + coalesce(select2.count, 0) AS total_count
FROM (SELECT pizzeria.name, COUNT(*) AS count
	FROM person_visits
	INNER JOIN pizzeria ON pizzeria.id = person_visits.pizzeria_id
	GROUP BY pizzeria.name) AS select1 
FULL JOIN (SELECT pizzeria.name, COUNT(*) AS count
FROM person_order
INNER JOIN menu ON person_order.menu_id = menu.id
INNER JOIN pizzeria ON menu.pizzeria_id = pizzeria.id
GROUP BY pizzeria.name) AS select2 ON select1.name = select2.name
ORDER BY total_count DESC, name ASC;