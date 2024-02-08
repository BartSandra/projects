SET ENABLE_SEQSCAN = OFF;

EXPLAIN ANALYZE SELECT menu.pizza_name AS pizza_name, pizzeria.name AS pizzeria_name
FROM menu
INNER JOIN pizzeria ON menu.pizzeria_id = pizzeria.id;