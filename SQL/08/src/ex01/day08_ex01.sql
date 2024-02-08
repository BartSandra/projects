-- Session #1 then session #2
SHOW TRANSACTION ISOLATION LEVEL;

-- Session #1 then session #2
BEGIN;

-- Session #1 then session #2
SELECT * FROM pizzeria WHERE name = 'Pizza Hut';

-- Session #1 
UPDATE pizzeria SET rating = 4 WHERE name = 'Pizza Hut';

-- Session #2 
UPDATE pizzeria SET rating = 3.6 WHERE name = 'Pizza Hut';

-- Session #1 then session #2
COMMIT;

-- Session #1 then session #2
SELECT * FROM pizzeria WHERE name = 'Pizza Hut';
