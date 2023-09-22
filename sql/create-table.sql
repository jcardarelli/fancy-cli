CREATE TABLE IF NOT EXISTS restaurants (
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT,
    address TEXT,
    stars INTEGER
);

-- sample data
INSERT INTO
    restaurants(name, address, stars)
VALUES
(
        'The French Laundry',
        '6640 Washington St, Yountville, CA 94599',
        3
    );

-- SELECT * FROM restaurants;
-- 1 | The French Laundry | 6640 Washington St, Yountville, CA 94599 | 3
