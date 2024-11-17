-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE categories (
    id SERIAL PRIMARY KEY,                      -- id auto-increment (primary key)
    name VARCHAR(255) NOT NULL,                  -- Nama kategori, tidak boleh kosong
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp saat data dibuat
    created_by VARCHAR(255),                     -- User yang membuat data
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp saat data terakhir dimodifikasi
    modified_by VARCHAR(255)                     -- User yang memodifikasi data terakhir
);


CREATE TABLE books (
    id SERIAL PRIMARY KEY,                      -- id auto-increment (primary key)
    title VARCHAR(255) NOT NULL,                 -- Judul buku, tidak boleh kosong
    description VARCHAR(255),                           -- Deskripsi buku
    image_url VARCHAR(255),                     -- URL gambar buku
    release_year INTEGER,                       -- Tahun rilis buku
    price INTEGER,                              -- Harga buku (misalnya dalam sen)
    total_page INTEGER,                         -- Jumlah halaman
    thickness VARCHAR(255),                      -- Ketebalan buku (misalnya dalam cm)
    category_id INTEGER,                        -- ID kategori (relasi dengan tabel kategori)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp saat data dibuat
    created_by VARCHAR(255),                    -- User yang membuat data
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp saat data terakhir dimodifikasi
    modified_by VARCHAR(255),                   -- User yang memodifikasi data terakhir
    CONSTRAINT fk_category
        FOREIGN KEY (category_id) REFERENCES categories(id) -- Relasi dengan tabel kategori (jika ada)
);

-- Index untuk mempercepat pencarian berdasarkan kategori
CREATE INDEX idx_category_id ON books(category_id);

-- Index untuk mempercepat pencarian berdasarkan nama
CREATE INDEX idx_name ON categories(name);


CREATE TABLE users (
    id SERIAL PRIMARY KEY,                      -- id auto-increment (primary key)
    username VARCHAR(255) UNIQUE NOT NULL,       -- Username, harus unik dan tidak boleh kosong
    password VARCHAR(255) NOT NULL,              -- Password, tidak boleh kosong
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp saat data dibuat
    created_by VARCHAR(255),                     -- User yang membuat data
    modified_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- Timestamp saat data terakhir dimodifikasi
    modified_by VARCHAR(255)                     -- User yang memodifikasi data terakhir
);

-- Index untuk mempercepat pencarian berdasarkan username
CREATE INDEX idx_username ON users(username);

INSERT INTO users (username, password, created_by, modified_by)
VALUES
    ('admin', 'admin', 'system', 'system');


-- +migrate StatementEnd