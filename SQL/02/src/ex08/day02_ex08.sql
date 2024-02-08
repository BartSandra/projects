SELECT name
FROM person
INNER JOIN person_order ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
WHERE gender IN ('male')
	AND address IN ('Moscow', 'Samara')
	AND pizza_name IN ('pepperoni pizza', 'mushroom pizza')
ORDER BY name DESC;