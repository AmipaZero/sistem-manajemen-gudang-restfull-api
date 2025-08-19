# 📦 Sistem Manajemen Gudang

Aplikasi manajemen gudang sederhana untuk manajemen gudang. 
Aplikasi ini menggunakan **RESTful API** dengan mekanisme autentikasi **JWT (JSON Web Token)**.  
Setiap pengguna wajib login untuk mendapatkan token, dan token tersebut harus disertakan pada setiap request ke endpoint API.

---
## ⚡ Fitur

### 🔐 Autentikasi
- **Register** pengguna baru  
- **Login** dengan username & password  
- **Logout** untuk menghapus token aktif  

### 👥 Role & Hak Akses
- **Admin**
  - Memiliki akses penuh ke semua fitur sistem  
  - Dapat mengelola barang, kategori, inbound, outbound  
  - Memiliki **hak khusus untuk mengakses laporan gudang** (stok, transaksi masuk/keluar, histori per barang)  
- **Staff**
  - Hanya bisa mengelola barang, kategori, inbound, dan outbound  
  - **Tidak memiliki akses** ke fitur laporan  

### 📦 Manajemen Barang
- Tambah, ubah, hapus, dan lihat daftar barang  
- Barang memiliki stok, kategori, dan informasi detail  

### 🏷️ Manajemen Kategori
- Tambah, ubah, hapus, dan lihat daftar kategori  
- Barang wajib terkait dengan kategori tertentu

### 📦 Produk
- Menambahkan produk baru ke dalam gudang  
- Menambahkan stok produk yang sudah ada 
- Mendukung lebih dari satu produk dalam satu transaksi inbound  

### 📥 Inbounds (Barang Masuk)
- Menambahkan stok barang ke gudang  
- Menyimpan riwayat inbound setiap barang  
- Mendukung lebih dari satu barang dalam satu transaksi inbound  

### 📤 Outbounds (Barang Keluar)
- Mengurangi stok barang dari gudang  
- Menyimpan riwayat outbound setiap barang  
- Mendukung lebih dari satu barang dalam satu transaksi outbound  
---
## 🛠️ Teknologi yang Digunakan
- **Golang** (Backend API)  
- **Gin Framework** (HTTP Web Framework)  
- **MySQL** (Database)  
- **JWT** (JSON Web Token untuk autentikasi)  

