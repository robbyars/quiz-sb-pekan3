# quiz-sb-pekan3
## Book Management API
Ini adalah aplikasi API untuk manajemen buku yang dibuat menggunakan Gin Framework dan Go (Golang). API ini memungkinkan pengguna untuk mengelola data buku dan kategori buku. Fitur utama dari aplikasi ini mencakup CRUD untuk buku dan kategori, serta autentikasi dasar menggunakan Basic Auth.

## Kegunaan
Aplikasi ini memungkinkan pengguna untuk melakukan operasi berikut:

* CRUD untuk Kategori Buku: Menambah, mengubah, melihat, dan menghapus kategori buku.
* CRUD untuk Buku: Menambah, mengubah, melihat, dan menghapus buku.
* Basic Authentication: Menggunakan Basic Auth untuk memastikan hanya pengguna yang terautentikasi yang dapat mengakses API.
## Persyaratan
* Go: Pastikan Go terpasang di sistem Anda.
* PostgreSQL: Database yang digunakan adalah PostgreSQL. Anda harus menyiapkan koneksi database terlebih dahulu.
## Instalasi
### Clone Repository

Clone repositori ini ke mesin lokal Anda:

* git clone https://github.com/robbyars/quiz-sb-pekan3.git
* cd quiz-sb-pekan3

### Install Dependencies

Install dependencies menggunakan go get:
* https://github.com/gin-gonic/gin
* https://github.com/lib/pq
* https://github.com/rubenv/sql-migrate


### Konfigurasi Database

* Buat database di PostgreSQL dan sesuaikan konfigurasi di file config/.env untuk menyesuaikan kredensial dan URL koneksi Anda.

## Menjalankan Aplikasi

Jalankan server menggunakan perintah berikut:

* go run main.go
Server akan berjalan di port 8080.

Cara Menggunakan
API ini menggunakan Basic Authentication untuk mengakses endpoint yang membutuhkan otentikasi. Anda dapat mengirimkan header Authorization dengan format berikut:

* Authorization: Basic <base64(username:password)>
* Gunakan Postman untuk menguji endpoint API.

### List Path API yang Tersedia
 Kategori Buku (Categories)
* GET /api/categories
 
Menampilkan semua kategori buku.
Contoh Response:

json
{
    "result": [
        {
            "id": 10,
            "name": "Robot",
            "created_at": "2024-11-18T04:07:23.604054Z",
            "created_by": "admin",
            "modified_at": "2024-11-18T04:07:23.604639Z",
            "modified_by": ""
        },
        {
            "id": 9,
            "name": "Technology Car",
            "created_at": "2024-11-18T01:08:07.340256Z",
            "created_by": "",
            "modified_at": "2024-11-18T04:07:58.002738Z",
            "modified_by": "admin"
        },
        {
            "id": 11,
            "name": "Plant and Animals",
            "created_at": "2024-11-18T04:13:37.648385Z",
            "created_by": "admin",
            "modified_at": "2024-11-18T04:15:57.383116Z",
            "modified_by": "admin"
        }
    ]
}

* POST /api/categories

Menambahkan kategori baru.
Body Request:

json
{
    "name":"Economy"
}

* GET /api/categories/12

Menampilkan detail kategori berdasarkan ID.
Contoh Response:

json
{
    "id": 12,
    "name": "Economy",
    "created_at": "2024-11-18T04:56:41.408936Z",
    "created_by": "admin",
    "modified_at": "2024-11-18T04:56:41.410044Z",
    "modified_by": ""
}

* PUT /api/categories/12

Mengubah kategori berdasarkan ID.
Body Request:

json
{
    "name":"Economy Modern"
}

* DELETE /api/categories/11

Menghapus kategori berdasarkan ID.

* GET /api/categories/9/books

Menampilkan semua buku berdasarkan kategori ID.

json
{
    "result": [
        {
            "id": 3,
            "title": "1984",
            "description": "A dystopian novel by George Orwell that explores the dangers of totalitarianism and extreme political ideology.",
            "image_url": "https://example.com/images/1984.jpg",
            "release_year": 1995,
            "price": 1000000,
            "total_page": 200,
            "thickness": "tebal",
            "category_id": 9,
            "created_at": "2024-11-18T01:53:41.032331Z",
            "created_by": "",
            "modified_at": "2024-11-18T02:10:32.975939Z",
            "modified_by": ""
        },
        {
            "id": 4,
            "title": "Go Programming Language",
            "description": "A comprehensive guide to Go programming.",
            "image_url": "https://example.com/images/go-programming.jpg",
            "release_year": 2024,
            "price": 300000,
            "total_page": 400,
            "thickness": "tebal",
            "category_id": 9,
            "created_at": "2024-11-18T03:40:42.644881Z",
            "created_by": "",
            "modified_at": "0001-01-01T00:00:00Z",
            "modified_by": ""
        }
    ]
}

Buku (Books)
* GET /api/books

Menampilkan semua buku.

{
    "result": [
        {
            "id": 3,
            "title": "1984",
            "description": "A dystopian novel by George Orwell that explores the dangers of totalitarianism and extreme political ideology.",
            "image_url": "https://example.com/images/1984.jpg",
            "release_year": 1995,
            "price": 1000000,
            "total_page": 200,
            "thickness": "tebal",
            "category_id": 9,
            "created_at": "2024-11-18T01:53:41.032331Z",
            "created_by": "",
            "modified_at": "2024-11-18T02:10:32.975939Z",
            "modified_by": ""
        },
        {
            "id": 4,
            "title": "Go Programming Language",
            "description": "A comprehensive guide to Go programming.",
            "image_url": "https://example.com/images/go-programming.jpg",
            "release_year": 2024,
            "price": 300000,
            "total_page": 400,
            "thickness": "tebal",
            "category_id": 9,
            "created_at": "2024-11-18T03:40:42.644881Z",
            "created_by": "",
            "modified_at": "0001-01-01T00:00:00Z",
            "modified_by": ""
        },
        {
            "id": 5,
            "title": "Introduction to Web Development",
            "description": "A beginner's guide to web development, covering HTML, CSS, and JavaScript fundamentals.",
            "image_url": "https://example.com/images/web-development.jpg",
            "release_year": 2023,
            "price": 220000,
            "total_page": 120,
            "thickness": "tebal",
            "category_id": 10,
            "created_at": "2024-11-18T04:19:46.444888Z",
            "created_by": "admin",
            "modified_at": "2024-11-18T04:23:16.737301Z",
            "modified_by": "admin"
        }
    ]
}

* POST /api/books

Menambahkan buku baru.
Body Request:

json
{
  "title": "Mastering JavaScript",
  "description": "Dive deep into JavaScript, exploring both basic concepts and advanced techniques for building modern web applications.",
  "image_url": "https://example.com/images/mastering-javascript.jpg",
  "release_year": 2024,
  "price": 420000,
  "total_page": 500,
  "category_id": 12
}

* GET /api/books/6

Menampilkan detail buku berdasarkan ID.
Contoh Response:

json
{
    "id": 6,
    "title": "Mastering JavaScript",
    "description": "Dive deep into JavaScript, exploring both basic concepts and advanced techniques for building modern web applications.",
    "image_url": "https://example.com/images/mastering-javascript.jpg",
    "release_year": 2024,
    "price": 420000,
    "total_page": 500,
    "thickness": "tebal",
    "category_id": 12,
    "created_at": "2024-11-18T05:01:28.758375Z",
    "created_by": "admin",
    "modified_at": "0001-01-01T00:00:00Z",
    "modified_by": ""
}

* PUT /api/books/6

Mengubah data buku berdasarkan ID.
Body Request:

json
{
  "title": "Go Programming Language - Advanced Edition",
  "description": "Updated guide to advanced Go concepts.",
  "image_url": "https://example.com/images/mastering-javascript.jpg",
  "release_year": 2023,
  "price": 400000,
  "total_page": 500,
  "category_id": 12
}

* DELETE /api/books/5

Menghapus buku berdasarkan ID.

### Autentikasi
API ini menggunakan Basic Authentication untuk memastikan bahwa hanya pengguna yang terautentikasi yang dapat mengakses endpoint tertentu. Autentikasi ini dilakukan dengan mengirimkan username dan password dalam header 

Authorization menggunakan format:

Authorization: Basic <base64(username:password)>

Cara Mendapatkan base64(username:password)
Untuk menghasilkan string base64 untuk kredensial Anda (username
), Anda dapat menggunakan berbagai tool seperti Base64 Encode.

Contoh:

Username: admin

Password: admin

Format yang dikirimkan dalam header:

Authorization: Basic YWRtaW46YWRtaW4=

Error Handling
API ini mengembalikan kode status HTTP berikut dalam berbagai situasi:

* 200 OK: Sesuai dengan error message berhasil dari setiap fungsi.
* 400 Bad Request: Input yang diberikan tidak valid.
* 401 Unauthorized: Pengguna tidak terautentikasi atau kredensial tidak valid.
* 404 Not Found: Data yang diminta tidak ditemukan.
* 500 Internal Server Error: Terjadi kesalahan pada server.

## Contributing
Jika Anda ingin berkontribusi pada proyek ini, silakan fork repositori ini, buat perubahan, dan kirimkan pull request.