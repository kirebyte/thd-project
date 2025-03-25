CREATE TABLE IF NOT EXISTS cars (
    id TEXT PRIMARY KEY,
    make TEXT NOT NULL,
    model TEXT NOT NULL,
    package TEXT NOT NULL,
    color TEXT NOT NULL,
    year INTEGER NOT NULL,
    category TEXT NOT NULL,
    mileage INTEGER NOT NULL, -- miles
    price INTEGER NOT NULL    -- cents
);
