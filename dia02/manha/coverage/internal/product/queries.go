package product

const (
	GetAll  = "SELECT p.id, p.name, p.price, p.quantity, p.type FROM products AS p"
	GetById = "SELECT p.id, p.name, p.price, p.quantity, p.type FROM products AS p WHERE p.id = ?"
	Create  = "INSERT INTO products (name, price, quantity, type) VALUES (?, ?, ?, ?)"
	Delete  = "DELETE FROM products WHERE id=?"
)

/*
	CREATE TABLE PRODUCTS(
		id INT PRIMARY KEY AUTO_INCREMENT,
		name VARCHAR(50) NOT NULL,
		type VARCHAR(50) NOT NULL,
		quantity INT DEFAULT 0,
		price DECIMAL(12,2) DEFAULT 0,
	)
*/
