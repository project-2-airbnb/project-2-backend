# Proyek 2: Backend Airbnb

Ini adalah layanan backend untuk Proyek 2: Airbnb. Dibangun menggunakan Go dan framework Echo, layanan ini menyediakan API untuk manajemen pengguna, daftar homestay, pemesanan, dan ulasan.

## Fitur

- Registrasi dan Autentikasi Pengguna
- Operasi CRUD untuk data Pengguna
- Melihat Daftar dan Detail Homestay
- Memesan Homestay
- Manajemen Host (Tambah, Perbarui, Hapus homestay)
- Memberikan Ulasan Homestay

## Memulai

### Prasyarat

- Go 1.16 atau lebih baru
- MySQL
- Docker (opsional, untuk kontainerisasi)

### Instalasi

1. Klon repositori:

   ```sh
   git clone https://github.com/project-2-airbnb/project-2-backend.git
   cd project-2-backend

   ```

2. Salin .env.example menjadi .env dan konfigurasi variabel lingkungan:
   cp .env.example .env

3. Install dependensi:
   go mod download

4. Jalankan aplikasi
   go run main.go

### Menggunakan Docker

Untuk menjalankan aplikasi menggunakan Docker, Anda dapat menggunakan Dockerfile yang sudah disediakan.

1. Build image Docker:
   docker build -t project-2-backend .

2. Jalankan kontainer Docker:
   docker run -p 8080:8080 --env-file .env project-2-backend

## Struktur Direktori

    app/: Berisi kode aplikasi utama, termasuk konfigurasi, basis data, migrasi, dan rute.

    features/: Berisi kode khusus fitur.

    utils/: Berisi fungsi utilitas.

    .github/workflows/: Berisi file workflow GitHub Actions.

    documentation/: Berisi dokumentasi proyek.

    .env.example: Contoh file variabel lingkungan.

    Dockerfile: File konfigurasi Docker.

    main.go: Titik masuk aplikasi.

## Endpoint API

### Pengguna

    POST /register: Registrasi pengguna baru

    POST /login: Login pengguna

    GET /users: Mendapatkan daftar pengguna

    GET /users/:id: Mendapatkan detail pengguna

    PUT /users/:id: Memperbarui detail pengguna

    DELETE /users/:id: Menghapus pengguna

## Kontribusi

    Kontribusi sangat diterima! Silakan buka isu atau kirim pull request.

## Lisensi

    Proyek ini dilisensikan di bawah Lisensi MIT.
