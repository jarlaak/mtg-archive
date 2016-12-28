package main

import (
	"database/sql"
)

func Up_card_infos(txn *sql.Tx) error {
	_, err := txn.Exec(`CREATE TABLE card_infos(
                        id bigserial NOT NULL PRIMARY KEY,
                        card_name varchar(256) NOT NULL,
                        created_at timestamp with time zone,
                        modified_at timestamp with time zone
                        );`)

	return err
}

func Up_card_infos_card_name_index(txn *sql.Tx) error {
	_, err := txn.Exec(`CREATE UNIQUE INDEX card_infos_card_name_index
                        ON card_infos
                        USING btree
                        (lower(card_name));`)

	return err
}

func Up_update_cards(txn *sql.Tx) error {
	_, err := txn.Exec(`ALTER TABLE cards DROP card_name;
                        ALTER TABLE cards ADD card_info_id bigserial REFERENCES card_infos (id);
    CREATE INDEX cards_card_info_id
                        ON cards
                        USING btree
                        (id,card_info_id);`)
	return err
}

// Up is executed when this migration is applied
func Up_2(txn *sql.Tx) error {
	err := Up_card_infos(txn)
	if err != nil {
		return err
	}

	err = Up_card_infos_card_name_index(txn)
	if err != nil {
		return err
	}

	err = Up_update_cards(txn)
	if err != nil {
		return err
	}

	return nil
}

// Down is executed when this migration is rolled back
func Down_2(txn *sql.Tx) error {
	_, err := txn.Exec("DROP INDEX cards_card_info_id;")
	if err != nil {
		return err
	}

	_, err = txn.Exec("ALTER TABLE deck_items DROP card_info_id;")
	if err != nil {
		return err
	}

	_, err = txn.Exec("ALTER TABLE deck_items ADD card_info_id;")
	if err != nil {
		return err
	}

	_, err = txn.Exec("DROP INDEX card_infos_card_name_index;")
	if err != nil {
		return err
	}

	_, err = txn.Exec("DROP TABLE card_infos;")
	if err != nil {
		return err
	}

	return nil
}
