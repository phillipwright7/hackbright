CREATE TABLE "users" (
  "user_id" serial PRIMARY KEY,
  "username" varchar,
  "password" varchar,
  "email" varchar
);

CREATE TABLE "movies" (
  "movie_id" serial PRIMARY KEY,
  "title" varchar,
  "overview" varchar,
  "release_date" date,
  "poster_url" varchar
);

CREATE TABLE "ratings" (
  "rating_id" serial PRIMARY KEY,
  "movie_id" int,
  "user_id" int,
  "score" int
);

ALTER TABLE "ratings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("movie_id");