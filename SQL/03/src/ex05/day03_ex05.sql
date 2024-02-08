SELECT pizzeria.name AS pizzeria_name
FROM person_visits
INNER JOIN person ON person.id = person_visits.person_id
INNER JOIN pizzeria ON pizzeria.id = person_visits.pizzeria_id
WHERE person.name IN ('Andrey')
EXCEPT
SELECT pizzeria.name
FROM person_order
INNER JOIN person ON person.id = person_order.person_id
INNER JOIN menu ON menu.id = person_order.menu_id
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
WHERE person.name IN ('Andrey')
ORDER BY pizzeria_name ASC;