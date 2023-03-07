CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "first_name" varchar NOT NULL,
  "last_name" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "daily_goals" (
  "id" bigserial PRIMARY KEY,
  "discription" varchar NOT NULL,
  "completed" boolean NOT NULL,
  "user_id" int,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "weekly_goals" (
  "id" bigserial PRIMARY KEY,
  "discription" varchar NOT NULL,
  "completed" boolean NOT NULL,
  "user_id" int,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "monthly_goals" (
  "id" bigserial PRIMARY KEY,
  "discription" varchar NOT NULL,
  "completed" boolean NOT NULL,
  "user_id" int,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE INDEX ON "users" ("first_name");

CREATE INDEX ON "daily_goals" ("user_id");

CREATE INDEX ON "weekly_goals" ("user_id");

CREATE INDEX ON "monthly_goals" ("user_id");

ALTER TABLE "daily_goals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "weekly_goals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "monthly_goals" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");