CREATE TABLE products (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        slug TEXT NOT NULL UNIQUE,
        detail TEXT NOT NULL,
        category_id INTEGER NOT NULL,
        FOREIGN KEY (category_id)
                REFERENCES categories (id)
);

CREATE TABLE categories (
        id INTEGER PRIMARY KEY,
        name TEXT NOT NULL,
        slug TEXT NOT NULL UNIQUE
);

CREATE TABLE images (
        id INTEGER PRIMARY KEY,
        filename TEXT NOT NULL,
        product_id INTEGER NOT NULL,
        FOREIGN KEY (product_id)
                REFERENCES products (id)
);

