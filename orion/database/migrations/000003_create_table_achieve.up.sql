CREATE TABLE achievements (
  id int unsigned NOT NULL AUTO_INCREMENT,
  created_at bigint(20) NOT NULL,
  updated_at bigint(20) NOT NULL,
  deleted_at datetime,
  year int unsigned NOT NULL,
  message text NOT NULL,
  PRIMARY KEY (id)
) AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;
