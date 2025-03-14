-- name: GetCity :one
SELECT * FROM cities_with_used
WHERE id = ? LIMIT 1;

-- name: ListCities :many
SELECT * FROM cities_with_used
ORDER BY name;

-- name: CreateCity :exec
INSERT INTO cities (
	name
) VALUES (
	?
)
RETURNING *;

-- name: UpdateCity :exec
UPDATE cities
SET name = ?
WHERE cities.id = sqlc.arg(id);

-- name: DeleteCity :exec
DELETE FROM cities
WHERE id = ?;

-- name: GetHouse :one
SELECT * FROM houses_with_cities
WHERE id = ? LIMIT 1;

-- name: ListHouses :many
SELECT * FROM houses_with_cities
ORDER BY created_at DESC;

-- name: CreateHouse :execlastid
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
);

-- name: UpdateHouse :exec
UPDATE houses
SET
	updated_at = CURRENT_TIMESTAMP,
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
WHERE id = ?;

-- name: DeleteHouse :exec
DELETE FROM houses
WHERE id = ?;

-- name: GetPublicationURLs :many
SELECT * FROM publication_urls
WHERE house_id = ?
ORDER BY publication_date DESC;

-- name: CreatePublicationURL :exec
INSERT INTO publication_urls (
	house_id,
	url,
	publication_date
) VALUES (
	?, ?, ?
);

-- name: UpdatePublicationURL :exec
UPDATE publication_urls
SET
	url = ?,
	publication_date = ?
WHERE id = ?;

-- name: DeletePublicationURL :exec
DELETE FROM publication_urls
WHERE id = ?;

-- name: DeleteAllPublicationURLs :exec
DELETE FROM publication_urls
WHERE house_id = ?;
