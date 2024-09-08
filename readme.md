# Gin_golang_REST_Api

Gin_golang_REST_Api adalah aplikasi backend sederhana menggunakan framework Gin untuk Go. Aplikasi ini mencakup rute otentikasi dan produk, serta middleware untuk autentikasi dan otorisasi.

## Prerequisites

Sebelum memulai, pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/) (versi 1.23.1)
- [Air](https://github.com/air-verse/air) (untuk hot reloading, opsional)

## Setup

1. **Clone Repository**

   Clone repositori ini ke mesin lokal Anda:

   ```bash
   git clone https://github.com/LuthfiSad/Gin_golang_REST_Api.git
   ```

2. **Masuk ke Direktori Proyek**

   ```bash
   cd Gin_golang_REST_Api
   ```

3. **Instal Dependensi**

   Instal dependensi Go yang diperlukan:

   ```bash
   go mod tidy
   ```

4. **Buat File `.env`**

   Salin file `.env.example` ke `.env` dan sesuaikan sesuai kebutuhan Anda:

   ```bash
   cp .env.example .env
   ```

   Pastikan Anda mengisi variabel yang diperlukan dalam file `.env`, seperti `PORT`, `DB`, `JWT_SECRET` dan `BASE_URL`.

5. **Inisialisasi Database**

   Pastikan Anda telah mengonfigurasi database Anda dan melakukan migrasi jika diperlukan.

## Menjalankan Aplikasi

1. **Dengan Air (Hot Reloading)**

   Jika Anda menggunakan `air` untuk hot reloading, jalankan:

   ```bash
   air
   ```

   Pastikan Anda telah menginstal `air` secara global dengan:

   ```bash
   go install github.com/air-verse/air@latest
   ```

2. **Tanpa Air**

   Untuk menjalankan aplikasi tanpa hot reloading, gunakan:

   ```bash
   go run main.go
   ```

## Rute API

### Autentikasi

- **Login**
  - `POST /login`
  - Mengautentikasi pengguna dan mengembalikan token JWT.

- **Register**
  - `POST /register`
  - Mendaftar pengguna baru.

- **Profile**
  - `GET /profile`
  - Mengambil informasi profil pengguna (memerlukan otentikasi).

- **Users**
  - `GET /users`
  - Mengambil daftar pengguna (memerlukan otentikasi dan hak admin).

### Produk

- **Get Products**
  - `GET /product/`
  - Mengambil daftar produk.

- **Get Product**
  - `GET /product/:id`
  - Mengambil detail produk berdasarkan ID.

- **Create Product**
  - `POST /product/`
  - Menambahkan produk baru.

- **Update Product**
  - `PUT /product/:id`
  - Memperbarui produk berdasarkan ID.

- **Delete Product**
  - `DELETE /product/:id`
  - Menghapus produk berdasarkan ID.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan fork repositori ini dan buat pull request dengan perubahan Anda.