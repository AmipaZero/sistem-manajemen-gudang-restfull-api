# 📦 Sistem Manajemen Gudang

Aplikasi manajemen gudang sederhana untuk manajemen gudang. 
Aplikasi ini menggunakan **RESTful API** dengan mekanisme autentikasi **JWT (JSON Web Token)**.  
Setiap pengguna wajib login untuk mendapatkan token, dan token tersebut harus disertakan pada setiap request ke endpoint API.

---
## ⚡ Fitur

### 🔐 Autentikasi
- **Register** pengguna baru  
- **Login** dengan username & password, dan masuk aplikasi
- **Logout** untuk menghapus token aktif dan keluar dari aplikasi

### 👥 Role & Hak Akses
- **Admin**
  - Memiliki akses penuh ke semua fitur sistem  
  - Dapat mengelola produk, inbound, outbound  
  - Memiliki **hak khusus untuk mengakses laporan gudang**
- **Staff**
  - Hanya bisa mengelola produk, inbound, dan outbound  
  - **Tidak memiliki akses** ke fitur laporan  

### 📦 Produk
- Menambahkan produk baru ke dalam gudang  
- Menambahkan stok produk yang sudah ada  
- Mendukung lebih dari satu produk dalam satu transaksi inbound  

### 📥 Inbounds (Produk Masuk)
- Menambahkan stok produk ke gudang  
- Menyimpan riwayat inbound setiap produk  
- Mendukung lebih dari satu produk dalam satu transaksi inbound  

### 📤 Outbounds (Produk Keluar)
- Mengurangi stok produk dari gudang  
- Menyimpan riwayat outbound setiap produk  
- Mendukung lebih dari satu produk dalam satu transaksi outbound  
---
## 🛠️ Teknologi yang Digunakan
- **Golang** (Backend API)  
- **Gin Framework** (HTTP Web Framework)  
- **MySQL** (Database)  
- **JWT** (JSON Web Token untuk autentikasi)  

