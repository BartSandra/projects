(SELECT pizzeria.name AS pizzeria_name
FROM person 
INNER JOIN person_order ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
WHERE person.gender IN ('female')
EXCEPT
SELECT pizzeria.name
FROM person
INNER JOIN person_order ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
WHERE person.gender IN ('male'))
UNION
(SELECT pizzeria.name
FROM person
INNER JOIN person_order ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
WHERE person.gender IN ('female')
EXCEPT
SELECT pizzeria.name
FROM person 
INNER JOIN person_order ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
WHERE person.gender IN ('male'))
ORDER BY pizzeria_name;