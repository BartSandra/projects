-- Session #1 then session #2
BEGIN;

-- Session #1
SELECT SUM(rating) FROM pizzeria;

-- Session #2 
UPDATE pizzeria SET rating = 1 WHERE name = 'Pizza Hut';

COMMIT;

-- Session #1
SELECT SUM(rating) FROM pizzeria;

COMMIT;

-- Session #1 then session #2
SELECT SUM(rating) FROM pizzeria;
