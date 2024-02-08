-- Session #1 then session #2
BEGIN;

-- Session #1
UPDATE pizzeria SET rating = 0.1 WHERE id = 1;

-- Session #2
UPDATE pizzeria SET rating = 0.2 WHERE id = 2;

-- Session #1
UPDATE pizzeria SET rating = 0.3 WHERE id = 2;

-- Session #2
UPDATE pizzeria SET rating = 0.4 WHERE id = 1;

-- Session #1 then session #2
COMMIT;