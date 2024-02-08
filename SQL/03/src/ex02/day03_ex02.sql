SELECT pizza_name, menu.price, pizzeria.name AS pizzeria_name
FROM (
	SELECT id AS menu_id
	FROM menu
	EXCEPT
	SELECT menu_id
	FROM person_order) AS one
INNER JOIN menu ON one.menu_id = menu.id
INNER JOIN pizzeria ON pizzeria.id = menu.pizzeria_id
ORDER BY pizza_name, menu.price;