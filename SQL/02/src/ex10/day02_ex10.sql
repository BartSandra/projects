SELECT person1.name, person2.name, person1.address AS common_address
FROM person person1
INNER JOIN person person2 ON person1.id > person2.id
AND person1.address = person2.address
ORDER BY person1.name, person2.name, person1.address;