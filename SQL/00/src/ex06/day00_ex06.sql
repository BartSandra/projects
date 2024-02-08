SELECT
  (SELECT name FROM person WHERE person.id = person_order.person_id) AS Name,
  (SELECT name = 'Denis' FROM person WHERE person.id = person_order.person_id) AS Check_name
FROM person_order
WHERE menu_id IN (13, 14, 18)
  AND order_date = '2022-01-07';