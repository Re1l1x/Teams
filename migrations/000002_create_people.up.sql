CREATE TABLE people
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    birth_year INT NOT NULL,
    group_id UUID NOT NULL REFERENCES groups(id)
        ON DELETE CASCADE
);
