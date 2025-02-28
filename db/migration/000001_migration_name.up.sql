CREATE TABLE swift_data (
    id SERIAL PRIMARY KEY,
    country_iso2_code CHAR(2) NOT NULL,
    swift_code VARCHAR(11) NOT NULL,
    code_type VARCHAR(5) NOT NULL,
    bank_name VARCHAR(150) NOT NULL,
    address VARCHAR(255) NOT NULL,
    town_name VARCHAR(100) NOT NULL,
    country_name VARCHAR(50) NOT NULL,
    time_zone VARCHAR(40) NOT NULL
);

CREATE INDEX idx_iso2_code ON swift_data(country_iso2_code);
