CREATE MATERIALIZED VIEW mv_dmitriy_visits_and_eats AS
SELECT pizzeria.name AS pizzeria_name
FROM menu
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
INNER JOIN person_visits ON menu.pizzeria_id = person_visits.pizzeria_id
INNER JOIN person ON person.id = person_visits.person_id
WHERE person.name = 'Dmitriy'
AND price < 800
AND visit_date = '2022-01-08';