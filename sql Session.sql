  CREATE TABLE `members` ( 
    `id` int( 11 ) NOT NULL AUTO_INCREMENT ,
   `username` varchar( 30 ) NOT NULL ,
   `password` varchar( 128 ) NOT NULL ,
    PRIMARY KEY ( `id` ) ,
    UNIQUE KEY `username` ( `username` ) ) ENGINE = MYISAM DEFAULT CHARSET = utf8;

  CREATE TABLE `logs` (
      `id` int(11) NOT NULL AUTO_INCREMENT,
      `action` varchar(15),
      `reminder` timestamp NULL DEFAULT NULL,
      INDEX `par_id` (`id`) ,
      `user` varchar (30) NOT NULL ,
      FOREIGN KEY (`user`) REFERENCES members(username) ON DELETE CASCADE
  ) ENGINE = MYISAM DEFAULT CHARSET = utf8;
