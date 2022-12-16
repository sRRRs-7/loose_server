-- +migrate Up
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "sex" varchar NOT NULL,
  "data_of_birth" varchar NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "codes" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "code" text UNIQUE NOT NULL,
  "img" bytea NOT NULL,
  "description" text NOT NULL,
  "performance" varchar NOT NULL,
  "star" bigint[] NOT NULL,
  "tags" varchar[] NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL,
  "access" bigint NOT NULL DEFAULT 0
);

CREATE TABLE "collection" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint NOT NULL,
  "code_id" bigint NOT NULL
);

CREATE TABLE "media" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "contents" text NOT NULL,
  "img" bytea NOT NULL,
  "created_at" timestamp NOT NULL,
  "updated_at" timestamp NOT NULL
);

CREATE TABLE "adminuser" (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "password" varchar UNIQUE NOT NULL,
  "created_at" timestamp NOT NULL
);

ALTER TABLE "codes" ADD FOREIGN KEY ("username") REFERENCES "users" ("username");

ALTER TABLE "collection" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "collection" ADD FOREIGN KEY ("code_id") REFERENCES "codes" ("id");

-- gin index
CREATE INDEX ON "codes" USING GIN ("tags");

-- +migrate Down
DROP TABLE IF EXISTS "users" cascade;
DROP TABLE IF EXISTS "stars" cascade;
DROP TABLE IF EXISTS "codes" cascade;
DROP TABLE IF EXISTS "tags" cascade;
DROP TABLE IF EXISTS "language" cascade;
DROP TABLE IF EXISTS "carts" cascade;
DROP TABLE IF EXISTS "media" cascade;
DROP TABLE IF EXISTS "adminuser" cascade;
