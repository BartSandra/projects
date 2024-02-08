SELECT pizzeria.name AS pizzeria_name
FROM menu
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
INNER JOIN person_visits ON menu.pizzeria_id = person_visits.pizzeria_id
INNER JOIN person ON person.id = person_visits.person_id
WHERE person.name IN ('Dmitriy')
	AND visit_date = '2022-01-08'
	AND price < 800;