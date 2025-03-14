CREATE TABLE IF NOT EXISTS cities (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS houses (
    id INTEGER PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    title TEXT NOT NULL,
    city_id INTEGER NOT NULL REFERENCES cities(id) ON DELETE RESTRICT,
    address TEXT NOT NULL DEFAULT '',
    price INTEGER NOT NULL,
    surface INTEGER NOT NULL,
    rooms INTEGER NOT NULL,
    bedrooms INTEGER NOT NULL,
    bathrooms INTEGER NOT NULL,
    floors INTEGER NOT NULL,
    construction_year INTEGER NOT NULL DEFAULT 0,
    house_type TEXT NOT NULL, -- 'maison' or 'appartement'
    land_surface INTEGER NOT NULL DEFAULT 0,
    has_garage BOOLEAN NOT NULL DEFAULT FALSE,
    outdoor_parking_spaces INTEGER NOT NULL DEFAULT 0,
    main_photo TEXT NOT NULL DEFAULT '', -- filename of the main photo
    notes TEXT NOT NULL DEFAULT ''
);

CREATE TABLE IF NOT EXISTS publication_urls (
    id INTEGER PRIMARY KEY,
    house_id INTEGER NOT NULL REFERENCES houses(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    publication_date DATE NOT NULL
);

CREATE VIEW IF NOT EXISTS cities_with_used
AS SELECT cities.*, CAST(EXISTS (SELECT 1 FROM houses WHERE houses.city_id = cities.id) AS BOOLEAN) AS is_used
FROM cities;

CREATE VIEW IF NOT EXISTS houses_with_cities
AS SELECT houses.*, cities.name AS city_name
FROM houses JOIN cities ON houses.city_id = cities.id;
