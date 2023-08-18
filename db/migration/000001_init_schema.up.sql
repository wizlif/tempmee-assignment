CREATE TABLE users (
  id integer PRIMARY KEY,
  email varchar UNIQUE NOT NULL,
  password varchar NOT NULL,
  created_at timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp
);

CREATE TABLE books (
  id integer PRIMARY KEY,
  title varchar NOT NULL,
  author varchar NOT NULL,
  price integer NOT NULL,
  page_count integer NOT NULL DEFAULT 0,
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp
);

CREATE TABLE orders (
  id integer PRIMARY KEY,
  price integer NOT NULL,
  book_id integer NOT NULL,
  user_id integer NOT NULL,
  status varchar NOT NULL DEFAULT 'PENDING',
  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp,
  FOREIGN KEY(user_id) REFERENCES users(id),
  FOREIGN KEY (book_id) REFERENCES books(id)
);
