# Final Project - Sanbercode Go Batch 41 Nando

Link project yang sudah di deploy : https://fp-sanbercode-go-batch-41-production.up.railway.app/. <br>
Link PPT : https://docs.google.com/presentation/d/118UTXRavY_5kMjCv2JWDA-4cgwqYJR8q/edit?usp=sharing&ouid=102141334842329430375&rtpof=true&sd=true

Aplikasi ini dapat digunakan oleh user baik sudah login ataupun belum login. Berikut adalah penjelasan dari tiap route yang ada di aplikasi ini.

### Route for all users
Route ini bisa diakses oleh siapapun (tidak perlu login untuk mengakses route ini). Berikut adalah daftar routenya.
- /menu
  - GET (/) : Pada route ini akan ditampilkan semua menu yang ada di dalam database dimana menu tersebut tidak dihapus oleh admin.
  - GET (/search?menu_name) : Menampilkan semua menu makanan yang tidak dihapus berdasarkan namanya.
  
### Route for admin only
Route ini hanya bisa diakses oleh user yang login sebagai admin saja. Berikut adalah daftar routenya.
- /admin
  - /order
    - GET  (/) : Menampilkan semua order customer yang telah terdaftar di database.
    - POST (/) : Menambahkan order baru untuk customer.
    - GET  (/detail?order_id) : Menampilkan detail order customer dari parameter yang diberikan.
  - /menu
    - GET (/) : Menampilkan semua menu baik yang dihapus maupun tidak.
    - POST (/) : Menambahkan menu baru.
  - /menu?menu_id
    - PUT (/) : Mengupdate data dari menu  id yang dikirimkan.
    - DELETE (/) : Menghapus menu dari menu ide yang dikirimkan. Delete bersifat soft delete.
  
### Route for customer only
Route ini hanya bisa diakses oleh user yang login sebagai customer saja. Berikut adalah daftar routenya.
- /customer
  - /order
    - GET (/) : Menampilkan semua orderan yang ada di dalam keranjang customer.
    - POST (/insert) : Menambahkan orderan baru ke dalam keranjang customer.
    - PUT (/delete) : Mengupdate / menghapus orderan dari keranjang customer.
    - POST (/clear) : Menghapus semua orderan yang ada di dalam keranjang.
  - /checkout
    - POST (/) : Mengcheckout semua orderan yang ada di dalam keranjang customer.
- /checkout
  - GET (/) : Menampilkan semua detail transaksi dari customer.