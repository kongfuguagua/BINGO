CREATE TABLE `World`
(
  `ID` varchar(255) NOT NULL COMMENT 'id',
  `Namespace` varchar(255) NOT NULL COMMENT 'namespace',
  `DLName` varchar(255) NOT NULL COMMENT 'dlname',
  PRIMARY KEY(`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;