DROP TABLE IF EXISTS events;

CREATE TABLE events (
    id INT NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    host varchar(255) NOT NULL,
    location varchar(255),
    start TIMESTAMP NOT NULL,
    end TIMESTAMP NOT NULL,
    dc varchar(255),
    theme varchar(255),
    price FLOAT,
    signup varchar(255),

    PRIMARY KEY (id)
);
