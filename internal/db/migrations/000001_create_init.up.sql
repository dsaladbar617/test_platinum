CREATE TABLE "sheets" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "short_name" varchar NOT NULL,
    "templates" varchar
);

CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "name" VARCHAR(255) NOT NULL,
    "man_number" INTEGER,
    "firebase_uuid" VARCHAR NOT NULL UNIQUE,
    "picture" VARCHAR,
    "email" VARCHAR(255) NOT NULL
);

CREATE TABLE "user_roles" (
    "id" bigserial PRIMARY KEY,
    "user_id" INTEGER REFERENCES users(id) on delete cascade NOT NULL,
    "role_name" VARCHAR(255) NOT NULL,
    "sheet_id" INTEGER REFERENCES sheets(id) on delete cascade NOT NULL
);

CREATE TABLE "fields" (
    "id" bigserial PRIMARY KEY,
    "type" VARCHAR(255) NOT NULL,
    "name" VARCHAR(255) NOT NULL,
    "archived" BOOLEAN NOT NULL,
    "favorite" BOOLEAN NOT NULL,
    "sheet_id" INTEGER REFERENCES sheets(id) on delete cascade NOT NULL
);

CREATE TABLE "entries" (
  "id" bigserial PRIMARY KEY,
  "sheet_id" INTEGER REFERENCES sheets(id) on delete cascade NOT NULL,
  "archived" BOOLEAN NOT NULL
);

CREATE TABLE "values" (
  "id" bigserial PRIMARY KEY,
  "value" TEXT NOT NULL,
  "entry_id" INTEGER REFERENCES entries(id) on delete cascade NOT NULL,
  "field_id" INTEGER REFERENCES fields(id) on delete cascade NOT NULL,
  "checked" BOOLEAN NOT NULL DEFAULT FALSE
);
