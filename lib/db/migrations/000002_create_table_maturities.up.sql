CREATE TABLE maturities (
    id SERIAL PRIMARY KEY,
    month INTEGER CHECK (month BETWEEN 1 AND 12) NOT NULL,
    month_name VARCHAR(3) GENERATED ALWAYS AS (
        CASE 
            WHEN month = 1 THEN 'JAN'
            WHEN month = 2 THEN 'FEB'
            WHEN month = 3 THEN 'MAR'
            WHEN month = 4 THEN 'APR'
            WHEN month = 5 THEN 'MAY'
            WHEN month = 6 THEN 'JUN'
            WHEN month = 7 THEN 'JUL'
            WHEN month = 8 THEN 'AUG'
            WHEN month = 9 THEN 'SEP'
            WHEN month = 10 THEN 'OCT'
            WHEN month = 11 THEN 'NOV'
            WHEN month = 12 THEN 'DEC'
        END
    ) STORED,
    year INTEGER CHECK (year >= 2024) NOT NULL,
    UNIQUE (month, year)
);
