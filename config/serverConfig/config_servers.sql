drop table if exists config_servers; 
CREATE TABLE config_servers (
    id INT NOT NULL,
    server_id INT NOT NULL,
    server_name VARCHAR(255) NOT NULL,
    start_time TIMESTAMP
);
LOAD DATA LOCAL INFILE './config_servers.csv' 
INTO TABLE config_servers 
FIELDS TERMINATED BY ',' 
ENCLOSED BY '"'
LINES TERMINATED BY '\n'
IGNORE 1 ROWS;