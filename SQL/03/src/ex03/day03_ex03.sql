(SELECT pizzeria.name AS pizzeria_name
FROM person 
INNER JOIN person_visits ON person.id = person_visits.person_id
INNER JOIN pizzeria ON pizzeria.id = person_visits.pizzeria_id
WHERE person.gender IN ('female')
EXCEPT ALL
SELECT pizzeria.name
FROM person
INNER JOIN person_visits ON person.id = person_visits.person_id
INNER JOIN pizzeria ON pizzeria.id = person_visits.pizzeria_id
WHERE person.gender IN ('male'))
UNION ALL
(SELECT pizzeria.name
FROM person
INNER JOIN person_visits ON person.id = person_visits.person_id
INNER JOIN pizzeria ON pizzeria.id = person_visits.pizzeria_id
WHERE person.gender IN ('male')
EXCEPT ALL
SELECT pizzeria.name
FROM person 
INNER JOIN person_visits ON person.id = person_visits.person_id
INNER JOIN pizzeria ON pizzeria.id = person_visits.pizzeria_id
WHERE person.gender IN ('female'))
ORDER BY pizzeria_name;