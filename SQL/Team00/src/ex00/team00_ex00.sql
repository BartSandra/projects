CREATE TABLE nodes (
    point1 VARCHAR(1),
    point2 VARCHAR(1),
    cost INT
);

INSERT INTO nodes (point1, point2, cost)
VALUES ('a', 'b', 10),
       ('b', 'a', 10),
       ('a', 'c', 15),
       ('c', 'a', 15),
       ('a', 'd', 20),
       ('d', 'a', 20),
       ('b', 'c', 35),
       ('c', 'b', 35),
       ('b', 'd', 25),
       ('d', 'b', 25),
       ('c', 'd', 30),
       ('d', 'c', 30);

WITH travel AS
(WITH RECURSIVE r AS
(SELECT nodes.point1::bpchar AS tour, nodes.point1, nodes.point2, nodes.cost
FROM nodes
WHERE nodes.point1 = 'a'
UNION
SELECT r.tour || ',' || r.point2 AS tour, nodes.point1, nodes.point2, nodes.cost + r.cost
FROM nodes
INNER JOIN r ON nodes.point1 = r.point2
WHERE r.tour NOT LIKE '%' || r.point2 || '%')
SELECT *
FROM r)
(SELECT cost AS total_cost, '{' || tour || ',a}' AS tour 
FROM travel
WHERE point2 = 'a'
AND cost = (SELECT MIN(cost)
			FROM travel
			WHERE point2 = 'a'
			AND length(tour) = 7)
ORDER BY total_cost ASC, tour ASC);