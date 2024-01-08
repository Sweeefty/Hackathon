-- insert.sql

-- Insert sample users
INSERT INTO accounts (email, password , role ,  campus_id) VALUES ('gurvan', 'gurvan' , 'admin', 1 );
INSERT INTO accounts (email, password , role , bde_id , campus_id) VALUES ('lenny', 'lenny' , 'admin' , 1 , 1);
INSERT INTO accounts (email, password , role ,  campus_id) VALUES ('rayen', 'rayen' , 'admin', 1);

-- Insert sample campus
INSERT INTO campus (name) VALUES ('Nantes');
INSERT INTO campus (name) VALUES ('Paris');
INSERT INTO campus (name) VALUES ('Montcuq');

-- Insert sample bde
INSERT INTO bde (name, campus_id) VALUES ('Peacocktail', 1);
INSERT INTO bde (name, campus_id) VALUES ('ParisBDE', 2);
INSERT INTO bde (name, campus_id) VALUES ('MontcuqBDE', 3);

-- Insert sample events
INSERT INTO events (name, description, date, price, bde_id) VALUES ('Soirée de fin année', '', '2024-05-23', 10, 1);

-- Insert sample products
INSERT INTO products (name, description, price) VALUES ('Place WareHouse', 'Une place pour aller au WareHouse pour la prochaine soirée sponsorisé par votre bde', 15);
INSERT INTO products (name, description, price) VALUES ('Pain au chocolat', 'Pour ceux qui ont laissé leur pc allumé', 1);

-- Insert sample inventory
INSERT INTO inventory (accounts_id, products_id) VALUES (1, 1);
INSERT INTO inventory (accounts_id, products_id) VALUES (2, 1);
INSERT INTO inventory (accounts_id, products_id) VALUES (3, 1);