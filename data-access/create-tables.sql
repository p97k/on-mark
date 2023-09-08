DROP TABLE IF EXISTS products;
CREATE TABLE products (
    id         INT AUTO_INCREMENT NOT NULL,
    name      VARCHAR(128) NOT NULL,
    description     VARCHAR(255) NOT NULL,
    price      DECIMAL(5,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
);

INSERT INTO products
(name, description, price)
VALUES
    ('Espresso', 'Strong coffee without milk', 2.25),
    ('Latte', 'Milk + Espresso', 3.65)