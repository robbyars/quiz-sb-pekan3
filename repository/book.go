package repository

import (
	"database/sql"
	"fmt"
	"quiz-sb-pekan3/structs"
)

func GetAllBook(db *sql.DB) (result []structs.Book, err error) {
	sql := "SELECT * FROM books"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var book structs.Book

		err = rows.Scan(&book.ID, &book.Title, &book.Description, &book.Image_url, &book.Release_year, &book.Price, &book.Total_page, &book.Thickness, &book.Category_id, &book.Created_at, &book.Created_by, &book.Modified_at, &book.Modified_by)
		if err != nil {
			return
		}

		result = append(result, book)
	}

	return
}

func InsertBook(db *sql.DB, book structs.Book) (err error) {
	// Cek apakah kategori dengan ID tersebut ada
	sqlCheck := "SELECT id FROM categories WHERE id = $1"
	var id int
	err = db.QueryRow(sqlCheck, book.Category_id).Scan(&id)

	if err == sql.ErrNoRows {
		// Jika tidak ada kategori yang ditemukan, kembalikan error kustom
		return fmt.Errorf("Kategori dengan ID %d tidak ditemukan. Periksa kembali daftar Kategori", book.Category_id)
	} else if err != nil {
		// Jika terjadi error lainnya, kembalikan error tersebut
		return fmt.Errorf("gagal memeriksa kategori: %w", err)
	}

	sql := "INSERT INTO books(title,description,image_url,release_year,price,total_page,thickness,category_id,created_at,created_by,modified_at,modified_by) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)"

	errs := db.QueryRow(sql, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness, book.Category_id, book.Created_at, book.Created_by, book.Modified_at, book.Modified_by)

	return errs.Err()
}

func UpdateBook(db *sql.DB, book structs.Book) (err error) {
	// Cek apakah kategori dengan ID tersebut ada
	sqlCheckCat := "SELECT id FROM categories WHERE id = $1"
	var idcat int
	err = db.QueryRow(sqlCheckCat, book.Category_id).Scan(&idcat)

	if err == sql.ErrNoRows {
		// Jika tidak ada kategori yang ditemukan, kembalikan error kustom
		return fmt.Errorf("Kategori dengan ID %d tidak ditemukan. Periksa kembali daftar Kategori", book.Category_id)
	} else if err != nil {
		// Jika terjadi error lainnya, kembalikan error tersebut
		return fmt.Errorf("gagal memeriksa kategori: %w", err)
	}
	// Cek apakah kategori dengan ID tersebut ada
	sqlCheck := "SELECT id FROM books WHERE id = $1"
	var id int
	err = db.QueryRow(sqlCheck, book.ID).Scan(&id)

	if err == sql.ErrNoRows {
		// Jika tidak ada kategori yang ditemukan, kembalikan error kustom
		return fmt.Errorf("Buku dengan ID %d tidak ditemukan", book.ID)
	} else if err != nil {
		// Jika terjadi error lainnya, kembalikan error tersebut
		return fmt.Errorf("gagal memeriksa buku: %w", err)
	}

	sqlUp := "UPDATE books SET title =$1,description=$2,image_url=$3,release_year=$4,price=$5,total_page=$6,thickness=$7,category_id=$8, modified_by = $9, modified_at=$10 WHERE id = $11"

	errs := db.QueryRow(sqlUp, book.Title, book.Description, book.Image_url, book.Release_year, book.Price, book.Total_page, book.Thickness, book.Category_id, book.Modified_by, book.Modified_at, book.ID)

	return errs.Err()
}

func DeleteBook(db *sql.DB, book structs.Book) (err error) {
	// Cek apakah kategori dengan ID tersebut ada
	sqlCheck := "SELECT id FROM books WHERE id = $1"
	var id int
	err = db.QueryRow(sqlCheck, book.ID).Scan(&id)

	if err == sql.ErrNoRows {
		// Jika tidak ada kategori yang ditemukan, kembalikan error kustom
		return fmt.Errorf("Buku dengan ID %d tidak ditemukan", book.ID)
	} else if err != nil {
		// Jika terjadi error lainnya, kembalikan error tersebut
		return fmt.Errorf("Gagal memeriksa buku: %w", err)
	}

	// Jika kategori ditemukan, lanjutkan untuk menghapus kategori
	sqlDelete := "DELETE FROM books WHERE id = $1"
	_, err = db.Exec(sqlDelete, book.ID)
	if err != nil {
		return fmt.Errorf("Gagal menghapus buku: %w", err)
	}

	// Jika penghapusan berhasil, tidak ada error yang dikembalikan
	return nil
}

func GetDetailBook(db *sql.DB, book *structs.Book) (err error) {
	sqlquery := "SELECT * FROM books WHERE id = $1"

	err = db.QueryRow(sqlquery, book.ID).Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.Image_url,
		&book.Release_year,
		&book.Price,
		&book.Total_page,
		&book.Thickness,
		&book.Category_id,
		&book.Created_at,
		&book.Created_by,
		&book.Modified_at,
		&book.Modified_by,
	)

	if err == sql.ErrNoRows {
		return fmt.Errorf("Buku dengan ID %d tidak ditemukan", book.ID)
	}

	if err != nil {
		return fmt.Errorf("Gagal mengambil data buku: %w", err)
	}

	return
}
