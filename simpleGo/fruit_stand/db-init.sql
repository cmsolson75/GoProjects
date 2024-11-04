CREATE TABLE IF NOT EXISTS "produce" (
    "Name" TEXT NOT NULL UNIQUE,
    "Amount" REAL,
    PRIMARY KEY("name")
);