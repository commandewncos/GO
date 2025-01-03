USE names;


CREATE TABLE IF NOT EXISTS informations(
    id INT NOT NULL AUTO_INCREMENT
    
    ,name VARCHAR(150) NOT NULL
    
    ,age INT NOT NULL
    
    ,PRIMARY KEY (id)
);


INSERT INTO informations(name, age)
VALUES
    ('Wilson Nascimento Costa', 35),
    ('Jessina Nascimento Costa', 31);