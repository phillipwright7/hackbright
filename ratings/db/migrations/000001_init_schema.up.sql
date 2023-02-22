CREATE TABLE "users" (
  "user_id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "movies" (
  "movie_id" serial PRIMARY KEY,
  "title" varchar NOT NULL,
  "overview" varchar NOT NULL,
  "release_date" date NOT NULL,
  "poster_url" varchar NOT NULL
);

CREATE TABLE "ratings" (
  "rating_id" serial PRIMARY KEY,
  "movie_id" int NOT NULL,
  "user_id" int NOT NULL,
  "score" int NOT NULL
);

ALTER TABLE "ratings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("movie_id");