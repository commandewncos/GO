USE people;

CREATE TABLE names(
    id INT NOT NULL AUTO_INCREMENT
    ,name VARCHAR(100)
    ,age INT


    , PRIMARY KEY (id)
);

-- JUst in init
INSERT INTO names(name, age) VALUES 
('Wellington Nascimento Costa', 36),
('Wilson Nascimento Costa', 35),
('Washington Nascimento Costa', 31);