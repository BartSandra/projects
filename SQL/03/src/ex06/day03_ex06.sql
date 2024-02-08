SELECT menu1.pizza_name, pizzeria1.name AS pizzeria_name_1, pizzeria2.name AS pizzeria_name_2, menu1.price
FROM menu menu1
INNER JOIN pizzeria ON pizzeria.id = menu1.pizzeria_id
INNER JOIN menu menu2 ON menu1.id <> menu2.id
AND menu1.price = menu2.price
AND menu1.pizza_name = menu2.pizza_name
AND menu1.pizzeria_id > menu2.pizzeria_id
INNER JOIN pizzeria pizzeria1 ON pizzeria1.id = menu1.pizzeria_id
INNER JOIN pizzeria pizzeria2 ON pizzeria2.id = menu2.pizzeria_id;