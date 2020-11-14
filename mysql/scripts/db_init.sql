CREATE DATABASE IF NOT EXISTS fullstackAsociacion;
CREATE USER IF NOT EXISTS 'dev'@'localhost' IDENTIFIED BY 'passdev';
GRANT ALL PRIVILEGES ON *.* TO 'dev'@'localhost';
