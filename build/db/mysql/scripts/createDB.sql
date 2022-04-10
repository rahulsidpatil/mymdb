CREATE DATABASE IF NOT EXISTS mDB;

CREATE USER IF NOT EXISTS 'docker'@'%' IDENTIFIED BY 'docker';
GRANT ALL PRIVILEGES ON *.* TO 'docker'@'%' WITH GRANT OPTION;

CREATE TABLE IF NOT EXISTS mDB.movies (
Id VARCHAR(512) NOT NULL PRIMARY KEY,
Title  VARCHAR(512),
Year INT,
Rated FLOAT,
Genre VARCHAR(512)
);
INSERT INTO mDB.movies (Id, Title, Year, Rated, Genre) VALUES 
('tt2884018','Macbeth',2015, 6.6,'Drama, History, War'),
('tt10095582','The Tragedy of Macbeth',2021, 7.4,'Drama, History, War'),
('tt4291600','Lady Macbeth',2016, 7.8,'Drama, History, War'),
('tt0067372','Macbeth',1971, 7.2,'Drama, History, War'),
('tt0040558','Macbeth',1948, 8.1,'Drama, History, War');