CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "image_url" varchar NOT NULL,
  "name" varchar NOT NULL,
  "message" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);