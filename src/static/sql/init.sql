CREATE DATABASE IF NOT EXISTS julius;

CREATE TABLE events (
        id INT NOT NULL AUTO_INCREMENT, 
        name varchar(255) NOT NULL, 
        host varchar(255) NOT NULL, 
        start TIMESTAMP NOT NULL, 
        end TIMESTAMP NOT NULL,
        dc varchar(255),
        theme varchar(255),
        price FLOAT,
        signup varchar(255),
        PRIMARY KEY (id)
        );


