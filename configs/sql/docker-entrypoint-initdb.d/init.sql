CREATE DATABASE `go_http_server_sample` DEFAULT CHARACTER SET utf8mb4;
CREATE DATABASE `go_http_server_sample_test` DEFAULT CHARACTER SET utf8mb4;

USE `go_http_server_sample`;

CREATE TABLE `albums` (
  `ean` CHAR(13) NOT NULL,
  `title` VARCHAR(255) NOT NULL,
  `artist` VARCHAR(255) NOT NULL,
  PRIMARY KEY (`ean`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
