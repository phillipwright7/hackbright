CREATE TABLE "cars" (
    car_id SERIAL PRIMARY KEY,
    make VARCHAR(25) NOT NULL,
    model VARCHAR(35) NOT NULL,
    price DECIMAL(11,2) NOT NULL,
    transmission VARCHAR(4) NOT NULL,
    trim_level VARCHAR(15) NOT NULL,
    color VARCHAR(20) NOT NULL
);

CREATE TABLE "owners" (
    owner_id SERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name VARCHAR(30) NOT NULL,
    phone_number INTEGER NOT NULL,
    email VARCHAR(30) NOT NULL
);

CREATE TABLE "sales" (
    sale_id SERIAL PRIMARY KEY,
    car_id INTEGER,
    owner_id INTEGER
);

ALTER TABLE "sales" ADD FOREIGN KEY ("car_id") REFERENCES "cars" ("car_id");

ALTER TABLE "sales" ADD FOREIGN KEY ("owner_id") REFERENCES "owners" ("owner_id");