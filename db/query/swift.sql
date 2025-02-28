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

-- name: GetDetailsSwift :many
SELECT country_iso2_code, swift_code, code_type, bank_name, address, town_name, country_name, time_zone, 
    CASE 
        WHEN RIGHT(@swift_code, 3) = 'XXX' THEN 'PARENT' 
        ELSE 'CHILD' 
    END AS parent
FROM swift_data
WHERE swift_code = @swift_code

UNION ALL

SELECT country_iso2_code, swift_code, code_type, bank_name, address, town_name, country_name, time_zone, 'CHILD' AS parent
FROM swift_data
WHERE 
    RIGHT(@swift_code, 3) = 'XXX' -- Sprawdza, czy kod kończy się na XXX
    AND swift_code LIKE CONCAT(LEFT(@swift_code, 8), '%');