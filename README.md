# Shopping Cart

### Technology Stack
Project ini menggunakan technology stack
* PostgreSQL (main db)
* Redis (cached)
* Elasticsearch (optional: used for monitoring logging)

### Feature
Adapun fitur-fitur yang tersedia pada project ini :
* User
	- [x] Login
	- [x] Register
	- [ ] Forgot Password
* Product
	- [x] List Product
	- [x] Product Detail
* Cart
	- [x] Add to cart
	- [x] Add quantity product
	- [ ] Remove Product
	- [x] Decrease quantity product
	- [x] Show Cart
* Checkout
	- [x] List available Courrier
	- [x] List Available Payment Method
	- [x] Finish Checkout (convert Cart to Order)

### Design Pattern
####  1. Programming on Interface
* Mempunyai 4 layer : usecase (bisnis logic), model (data structure), controller (handler endpoint), dan repository (query layer). Masing-masing layer hanya bisa meng-akses melalui interface.
#### 2. Injection using Factory Method Pattern
* Injeksi ini digunakan untuk menyambungkan dari satu package ke package lain melalui interface.
#### 3.  Minimize Dependency