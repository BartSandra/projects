SELECT name
FROM person
INNER JOIN person_order ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
WHERE menu.pizza_name IN ('pepperoni pizza')
AND gender IN ('female')
INTERSECT
SELECT name
FROM person
INNER JOIN person_order ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
WHERE menu.pizza_name IN ('cheese pizza')
AND gender IN ('female')
ORDER BY name ASC;