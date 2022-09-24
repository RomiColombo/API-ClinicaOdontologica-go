CREATE SCHEMA IF NOT EXISTS `clinicaodontologica` ;
USE `clinicaodontologica`;

-- CREATE TABLES
CREATE TABLE `clinicaodontologica`.`odontologos` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(45) NOT NULL,
  `apellido` VARCHAR(45) NOT NULL,
  `matricula` INT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `matricula_UNIQUE` (`matricula` ASC) VISIBLE);
  
  
  CREATE TABLE `clinicaodontologica`.`pacientes` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `nombre` VARCHAR(45) NOT NULL,
  `apellido` VARCHAR(45) NOT NULL,
  `domicilio` VARCHAR(45) NOT NULL,
  `dni` INT NOT NULL,
  `fechaAlta` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `dni_UNIQUE` (`dni` ASC) VISIBLE);
  
-- DROP TABLE pacientes

-- INSERT ODONTOLOGOS
INSERT INTO odontologos (nombre, apellido, matricula) VALUES ("Emmanuel", "Forster", 1234);
INSERT INTO odontologos (nombre, apellido, matricula) VALUES ("Rodolfo", "Macias", 5678);
INSERT INTO odontologos (nombre, apellido, matricula) VALUES ("Susana", "Alfonso", 9123);
 
SELECT * FROM odontologos;