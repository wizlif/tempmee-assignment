-- SQL dump generated using DBML (dbml-lang.org)
-- Database: SQLIte
-- Generated at: 2023-08-18T12:20:04.470Z

PRAGMA foreign_keys = ON;

CREATE TABLE "users" (
  "id" integer PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" integer PRIMARY KEY,
  "title" varchar NOT NULL,
  "author" varchar NOT NULL,
  "price" integer NOT NULL,
  "page_count" integer,
  "created_at" timestamp,
  "updated_at" timestamp
);

CREATE TABLE "orders" (
  "id" integer PRIMARY KEY,
  "price" integer NOT NULL,
  "book_id" integer NOT NULL,
  "user_id" integer NOT NULL,
  "status" varchar DEFAULT 'PENDING',
  "created_at" timestamp,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY (book_id) REFERENCES books(id)
);
