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
    address TEXT,
    price INTEGER NOT NULL,
    surface INTEGER NOT NULL,
    rooms INTEGER,
    bedrooms INTEGER,
    bathrooms INTEGER,
    floors INTEGER,
    construction_year INTEGER,
    house_type TEXT NOT NULL, -- 'maison' or 'appartement'
    land_surface INTEGER,
    has_garage BOOLEAN,
    outdoor_parking_spaces INTEGER,
    main_photo TEXT, -- filename of the main photo
    notes TEXT
);

CREATE TABLE IF NOT EXISTS publication_urls (
    id INTEGER PRIMARY KEY,
    house_id INTEGER NOT NULL REFERENCES houses(id) ON DELETE CASCADE,
    url TEXT NOT NULL,
    publication_date DATE NOT NULL,
    is_main BOOLEAN NOT NULL
);
