CREATE TABLE sessions (
	id int unsigned NOT NULL AUTO_INCREMENT,
	created_at datetime NOT NULL,
	updated_at datetime NOT NULL,
	deleted_at datetime,
	class_id varchar(64) NOT NULL,
	starts_at datetime NOT NULL,
	ends_at datetime NOT NULL,
	canceled boolean NOT NULL DEFAULT false,
	notes text NOT NULL,
	PRIMARY KEY (id),
	FOREIGN KEY (class_id) REFERENCES classes(class_id)
) AUTO_INCREMENT=1 DEFAULT CHARSET=UTF8MB4;