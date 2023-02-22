CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar NOT NULL
);

CREATE TABLE "movies" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "overview" varchar NOT NULL,
  "release_date" varchar NOT NULL,
  "poster_url" varchar NOT NULL
);

CREATE TABLE "ratings" (
  "id" bigserial PRIMARY KEY,
  "score" int NOT NULL,
  "movie_id" int,
  "user_id" int
);

ALTER TABLE "ratings" ADD FOREIGN KEY ("movie_id") REFERENCES "movies" ("id");

ALTER TABLE "ratings" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");