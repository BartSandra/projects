SELECT pizza_name, menu.price, pizzeria.name AS pizzeria_name, person_visits.visit_date
FROM menu
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
INNER JOIN person_visits ON pizzeria.id = person_visits.pizzeria_id
INNER JOIN person ON person.id = person_visits.person_id
WHERE person.name IN ('Kate')
AND price >= 800
AND price <= 1000
ORDER BY pizza_name, menu.price, pizzeria_name;