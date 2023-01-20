#### Langkah Penggunaan
`git clone https://github.com/fakhrads/Final-Project---Sanbercode-Go-Batch-41`

`cd golang-api`

`go mod tidy`

`go run main.go`

#### Endpoint User
|  TYPE | ENDPOINT  | USAGE  | PARAMS  |
| ------------ | ------------ | ------------ | ------------ |
|  POST | /api/v1/login  |  Melakukan login pengguna untuk mendapatkan JWT token | username, password  |
|  POST |  /api/v1/register | Melakukan register pengguna  | username, password  |
|  GET  | /api/v1/user  |  Mengambil seluruh data pengguna | -  |
|  GET  | /api/v1/user/:id  |  Mengambil data pengguna berdasarkan ID | -  |
| PUT  | /api/v1/user/:id  | Merubah data pengguna  | username, password  |
| DELETE  |  /api/v1/user/:id | Menghapus data pengguna  | -  |


#### Endpoint Product
|  TYPE | ENDPOINT  | USAGE  | PARAMS  |
| ------------ | ------------ | ------------ | ------------ |
|  POST | /api/v1/product  |  Mengunggah data produk |    product_name, product_description, product_price, product_stock  |
|  GET  | /api/v1/product  |  Mengambil seluruh data produk | -  |
|  GET  | /api/v1/product/:id  |  Mengambil data produk berdasarkan ID | -  |
| PUT  | /api/v1/product/:id  | Merubah data produk  | product_name, product_description, product_price, product_stock  |
| DELETE  |  /api/v1/product/:id | Menghapus data produk  | -  |

#### Endpoint Order
|  TYPE | ENDPOINT  | USAGE  | PARAMS  |
| ------------ | ------------ | ------------ | ------------ |
|  POST | /api/v1/order  |  Melakukan order produk |    users_id, products_id, total_bought  |
|  GET  | /api/v1/order  |  Mengambil seluruh data order | -  |
|  GET  | /api/v1/order/:id  |  Mengambil data order berdasarkan ID | -  |
