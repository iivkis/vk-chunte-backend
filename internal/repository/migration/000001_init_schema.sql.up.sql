CREATE TABLE "users" (
  "id" SERIAL PRIMARY KEY,
  "vk_id" bigint,
  "name" varchar,
  "age" int,
  "created_at" timestamp DEFAULT (now())
);
