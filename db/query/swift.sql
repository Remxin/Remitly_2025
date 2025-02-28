-- name: CreateSwiftData :exec
INSERT INTO swift_data (
    country_iso2_code,
    swift_code,
    code_type,
    bank_name,
    address,
    town_name,
    country_name,
    time_zone
) 
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8
);