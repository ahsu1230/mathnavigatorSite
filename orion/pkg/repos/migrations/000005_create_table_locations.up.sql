CREATE TABLE locations (
	id int unsigned NOT NULL AUTO_INCREMENT,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	deleted_at datetime,
	loc_id varchar(64) NOT NULL UNIQUE,
	street varchar(255) NOT NULL,
	city varchar(64) NOT NULL,
	state varchar(64) NOT NULL,
	zipcode varchar(64) NOT NULL,
	room varchar(64),
	PRIMARY KEY (id)
) AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;