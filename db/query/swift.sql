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
    RIGHT(@swift_code, 3) = 'XXX'
    AND swift_code LIKE CONCAT(LEFT(@swift_code, 8), '%');

-- name: GetDetailsCountry :many
SELECT
	country_iso2_code
	, country_name
	,address
	,bank_name
	,swift_code
    ,CASE 
        WHEN RIGHT(swift_code, 3) = 'XXX' THEN 'PARENT' 
        ELSE 'CHILD' 
    END AS parent
FROM swift_data
WHERE country_iso2_code = @country_iso2_code;


-- name: AddNewSwiftCode :one
INSERT INTO swift_data (
    address
    ,bank_name
    ,country_iso2_code
    ,country_name
    ,swift_code
    ,code_type
    ,town_name
    ,time_zone
) VALUES (
    @address
    ,@bank_name
    ,UPPER(@country_iso2_code)
    ,UPPER(@country_name)
    ,@swift_code
    ,'BIC11'
    ,' '
    ,' '
) RETURNING *;

-- name: DeleteSwiftCode :one
DELETE FROM swift_data
WHERE
    swift_code = @swift_code
RETURNING *;
