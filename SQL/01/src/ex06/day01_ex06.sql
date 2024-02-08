SELECT action_date, person.name AS person_name
FROM (SELECT order_date AS action_date, person_id
FROM person_order
INTERSECT ALL
SELECT visit_date, person_id
FROM person_visits) AS po
INNER JOIN person ON po.person_id = person.id
ORDER BY action_date ASC, person_name DESC;