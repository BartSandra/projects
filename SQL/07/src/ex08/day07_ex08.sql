SELECT address, pizzeria.name AS name, COUNT(*) AS count_of_orders
FROM person_order
INNER JOIN menu ON person_order.menu_id = menu.id
INNER JOIN pizzeria ON menu.pizzeria_id = pizzeria.id
INNER JOIN person ON person.id = person_order.person_id
GROUP BY address, pizzeria.name
ORDER BY address ASC, pizzeria.name ASC;