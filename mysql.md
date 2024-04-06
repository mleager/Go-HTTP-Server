# MySQL Setup for Inventory Products

1. Install MySQL

2. Sign into MySQL
`$ mysql -u root -p`
* login and set the password

3. Create DB 'inventory'
`mysql$ create database inventory;`

4. Switch to the DB 'inventory'
`mysql$ use inventory;`

5. Create 'products' Table
```
mysql$ CREATE TABLE products(
    id INT NOT NULL AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    quantity INT,
    price FLOAT(10,7),
    PRIMARY KEY(id)
    );
```

6. Describe the Table
`mysql$ desc products;`

7. Add Values to the Table
`mysql$ insert into products values(1,"chair", 100, 200.00);`
`mysql$ insert into products values(2,"desk", 800, 600.00;`

8. Get Table info
`mysql$ SELECT * FROM products;`

9. Install MySQL Go Module
`go get -u github.com/go-sql-driver/mysql`
* works indirectly with builtin `sql` module
(hence the `_` before module in `mysql.go`)
