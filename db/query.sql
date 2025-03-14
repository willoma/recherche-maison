-- name: GetCity :one
SELECT * FROM cities
WHERE id = ? LIMIT 1;

-- name: ListCities :many
SELECT * FROM cities
ORDER BY name;

-- name: CreateCity :one
INSERT INTO cities (
  name
) VALUES (
  ?
)
RETURNING *;

-- name: UpdateCity :one
UPDATE cities
SET name = ?
WHERE id = ?
RETURNING *;

-- name: DeleteCity :exec
DELETE FROM cities
WHERE id = ?;

-- name: IsCityUsedByHouses :one
SELECT EXISTS(
  SELECT 1 FROM houses WHERE city_id = ?
) AS is_used;

-- name: GetHouse :one
SELECT * FROM houses
WHERE id = ? LIMIT 1;

-- name: ListHouses :many
SELECT * FROM houses
ORDER BY created_at DESC;

-- name: CreateHouse :one
INSERT INTO houses (
  title,
  city_id,
  address,
  price,
  surface,
  rooms,
  bedrooms,
  bathrooms,
  floors,
  construction_year,
  house_type,
  land_surface,
  has_garage,
  outdoor_parking_spaces,
  main_photo,
  notes
) VALUES (
  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?
)
RETURNING *;

-- name: UpdateHouse :one
UPDATE houses
SET updated_at = CURRENT_TIMESTAMP,
    title = ?,
    city_id = ?,
    address = ?,
    price = ?,
    surface = ?,
    rooms = ?,
    bedrooms = ?,
    bathrooms = ?,
    floors = ?,
    construction_year = ?,
    house_type = ?,
    land_surface = ?,
    has_garage = ?,
    outdoor_parking_spaces = ?,
    main_photo = ?,
    notes = ?
WHERE id = ?
RETURNING *;

-- name: DeleteHouse :exec
DELETE FROM houses
WHERE id = ?;

-- name: GetPublicationURLs :many
SELECT * FROM publication_urls
WHERE house_id = ?
ORDER BY is_main DESC, publication_date DESC;

-- name: GetMainPublicationURL :one
SELECT * FROM publication_urls
WHERE house_id = ? AND is_main = true
LIMIT 1;

-- name: CreatePublicationURL :one
INSERT INTO publication_urls (
  house_id,
  url,
  publication_date,
  is_main
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: UpdatePublicationURL :one
UPDATE publication_urls
SET url = ?,
    publication_date = ?,
    is_main = ?
WHERE id = ?
RETURNING *;

-- name: DeletePublicationURL :exec
DELETE FROM publication_urls
WHERE id = ?;

-- name: DeleteAllPublicationURLs :exec
DELETE FROM publication_urls
WHERE house_id = ?;
