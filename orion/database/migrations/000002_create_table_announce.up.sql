CREATE TABLE announcements (
  id int unsigned NOT NULL AUTO_INCREMENT,
  created_at bigint(20) NOT NULL,
  updated_at bigint(20) NOT NULL,
  deleted_at datetime,
  posted_at bigint(20) NOT NULL,
  author varchar(255) NOT NULL,
  message text NOT NULL,
  PRIMARY KEY (id)
) AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;
