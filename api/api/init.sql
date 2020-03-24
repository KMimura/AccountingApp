CREATE TABLE IF NOT EXISTS test
(
    id INT(20) AUTO_INCREMENT PRIMARY KEY,
    amount INT(20) NOT NULL,
    t_date DATE NOT NULL,
    t_type CHAR(25),    
    if_earning BOOLEAN
);