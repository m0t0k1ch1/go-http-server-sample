CREATE DATABASE `go_http_server_sample` DEFAULT CHARACTER SET utf8mb4;
CREATE DATABASE `go_http_server_sample_test` DEFAULT CHARACTER SET utf8mb4;

USE `go_http_server_sample`;

CREATE TABLE `albums` (
  `ean` CHAR(13) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `artist` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`ean`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `albums` (`ean`, `title`, `artist`)
VALUES (4988002758807, 'Juice', 'iri'),
       (4988005553027, 'This Is The One', 'Utada'),
       (4988008803235, 'MADRUGADA / TIGER EYES', 'Jazztronik'),
       (4995879601242, 'GREEN', 'SEEDA'),
       (4997184881425, 'modal soul', 'Nujabes');
