package migrations1

import (
	"fmt"
	"database/sql"
)

func Up00001(tx *sql.Tx) error {
	fmt.Println("********** UP 00001 ************")
	_, err := tx.Exec(
		`CREATE TABLE moosey
		(
			id          int unsigned     NOT NULL AUTO_INCREMENT,
			created_at  datetime         NOT NULL,
			updated_at  datetime         NOT NULL,
			deleted_at  datetime,
			program_id  varchar(64)      NOT NULL UNIQUE,
			name        varchar(255)     NOT NULL,
			grade1      tinyint unsigned NOT NULL,
			grade2      tinyint unsigned NOT NULL,
			description text             NOT NULL,
			PRIMARY KEY (id)
		) AUTO_INCREMENT = 1
		  DEFAULT CHARSET = UTF8MB4;
		`)
	if err != nil {
		return err
	}
	return nil
}

func Down00001(tx *sql.Tx) error {
	fmt.Println("********** DOWN 00001 ************")
	_, err := tx.Exec("DROP TABLE IF EXISTS moosey;")
	if err != nil {
		return err
	}
	return nil
}