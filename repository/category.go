package repository

import (
	"database/sql"
	"fmt"
	"quiz-sb-pekan3/structs"
)

func GetAllCategory(db *sql.DB) (result []structs.Category, err error) {
	sql := "SELECT * FROM categories"

	rows, err := db.Query(sql)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var cat structs.Category

		err = rows.Scan(&cat.ID, &cat.Name, &cat.Created_at, &cat.Created_by, &cat.Modified_at, &cat.Modified_by)
		if err != nil {
			return
		}

		result = append(result, cat)
	}

	return
}

func InsertCategory(db *sql.DB, cat structs.Category) (err error) {
	sql := "INSERT INTO categories(name,created_at,created_by,modified_by) VALUES ($1,$2,$3,$4)"

	errs := db.QueryRow(sql, cat.Name, cat.Created_at, cat.Created_by, cat.Modified_by)

	return errs.Err()
}

func UpdateCategory(db *sql.DB, cat structs.Category) (err error) {
	// Cek apakah kategori dengan ID tersebut ada
	sqlCheck := "SELECT id FROM categories WHERE id = $1"
	var id int
	err = db.QueryRow(sqlCheck, cat.ID).Scan(&id)

	if err == sql.ErrNoRows {
		// Jika tidak ada kategori yang ditemukan, kembalikan error kustom
		return fmt.Errorf("kategori dengan ID %d tidak ditemukan", cat.ID)
	} else if err != nil {
		// Jika terjadi error lainnya, kembalikan error tersebut
		return fmt.Errorf("gagal memeriksa kategori: %w", err)
	}

	sqlUp := "UPDATE categories SET name = $1, modified_by = $2, modified_at=$3 WHERE id = $4"

	errs := db.QueryRow(sqlUp, cat.Name, cat.Modified_by, cat.Modified_at, cat.ID)

	return errs.Err()
}

func DeleteCategory(db *sql.DB, cat structs.Category) (err error) {
	// Cek apakah kategori dengan ID tersebut ada
	sqlCheck := "SELECT id FROM categories WHERE id = $1"
	var id int
	err = db.QueryRow(sqlCheck, cat.ID).Scan(&id)

	if err == sql.ErrNoRows {
		// Jika tidak ada kategori yang ditemukan, kembalikan error kustom
		return fmt.Errorf("Kategori dengan ID %d tidak ditemukan", cat.ID)
	} else if err != nil {
		// Jika terjadi error lainnya, kembalikan error tersebut
		return fmt.Errorf("Gagal memeriksa kategori: %w", err)
	}

	// Jika kategori ditemukan, lanjutkan untuk menghapus kategori
	sqlDelete := "DELETE FROM categories WHERE id = $1"
	_, err = db.Exec(sqlDelete, cat.ID)
	if err != nil {
		return fmt.Errorf("Gagal menghapus kategori: %w", err)
	}

	// Jika penghapusan berhasil, tidak ada error yang dikembalikan
	return nil
}

func GetDetailCategory(db *sql.DB, cat *structs.Category) (err error) {
	sqlquery := "SELECT * FROM categories WHERE id = $1"

	err = db.QueryRow(sqlquery, cat.ID).Scan(
		&cat.ID,
		&cat.Name,
		&cat.Created_at,
		&cat.Created_by,
		&cat.Modified_at,
		&cat.Modified_by,
	)

	if err == sql.ErrNoRows {
		return fmt.Errorf("Kategori dengan ID %d tidak ditemukan", cat.ID)
	}

	if err != nil {
		return fmt.Errorf("Gagal mengambil data kategori: %w", err)
	}

	return
}

func GetDetailBookbyCategory(db *sql.DB, book *structs.Book) (result []structs.Book, err error) {
	sqlquery := "SELECT * FROM books WHERE category_id = $1"

	rows, err := db.Query(sqlquery, book.Category_id)
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
