
package main

import (
	"database/sql"
)

func Up_users(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE TABLE users(
                        id bigserial NOT NULL PRIMARY KEY,
                        username varchar(256) NOT NULL,
                        pw_salt text,
                        pw_hash text,
                        created_at timestamp with time zone,
                        modified_at timestamp with time zone,
                        seen_at timestamp with time zone
                        );`)

    return err
}

func Up_cards(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE TABLE cards(
                        id bigserial NOT NULL PRIMARY KEY,
                        card_name varchar(256) NOT NULL,
                        multiverse_id varchar(10) UNIQUE,
                        created_at timestamp with time zone,
                        modified_at timestamp with time zone
                        );`)

    return err
}

func Up_cards_card_name_index(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE INDEX cards_card_name_index
                        ON cards
                        USING btree (card_name);`)

    return err
}

func Up_cards_multiverse_id_index(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE INDEX cards_multiverse_id_index
                        ON cards
                        USING btree
                        (multiverse_id NULLS LAST);`)

    return err
}

func Up_decks(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE TABLE decks(
                        id bigserial NOT NULL PRIMARY KEY,
                        deck_name varchar(256) NOT NULL,
                        description text,
                        user_id bigserial REFERENCES users (id) ON DELETE CASCADE,
                        created_at timestamp with time zone,
                        modified_at timestamp with time zone
                        );`)
    return err
}

func Up_decks_deck_name_index(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE INDEX decks_deck_name_index
                        ON decks
                        USING btree (deck_name);`)

    return err
}

func Up_deck_items(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE TABLE deck_items(
                        id bigserial NOT NULL PRIMARY KEY,
                        deck_id bigserial REFERENCES decks (id) ON DELETE CASCADE,
                        user_id bigserial REFERENCES users (id) ON DELETE CASCADE,
                        card_id bigserial REFERENCES cards (id),
                        created_at timestamp with time zone,
                        modified_at timestamp with time zone
                        );`)
    return err
}

func Up_deck_items_index(txn *sql.Tx) error {
    _, err := txn.Exec(`CREATE INDEX deck_items_index
                        ON deck_items
                        USING btree
                        (user_id,card_id,deck_id);`)

    return err
}

// Up is executed when this migration is applied
func Up_1(txn *sql.Tx) error {
    err1 := Up_users(txn)
    if err1 != nil {
        return err1
    }

    err2 := Up_cards(txn)
    if err2 != nil {
        return err2
    }

    err3 := Up_cards_card_name_index(txn)
    if err3 != nil {
        return err3
    }

    err4 := Up_cards_multiverse_id_index(txn)
    if err4 != nil {
        return err4
    }

    err5 := Up_decks(txn)
    if err5 != nil {
        return err5
    }

    err6 := Up_decks_deck_name_index(txn)
    if err6 != nil {
        return err6
    }

    err7 := Up_deck_items(txn)
    if err7 != nil {
        return err7
    }


    err8 := Up_deck_items_index(txn)
    if err8 != nil {
        return err8
    }

    return nil
}

// Down is executed when this migration is rolled back
func Down_1(txn *sql.Tx) error {
    _, err8 := txn.Exec("DROP INDEX deck_items_index;")
    if err8 != nil {
        return err8
    }

    _, err7 := txn.Exec("DROP TABLE deck_items;")
    if err7 != nil {
        return err7
    }

    _, err6 := txn.Exec("DROP INDEX decks_deck_name_index;")
    if err6 != nil {
        return err6
    }

    _, err5 := txn.Exec("DROP TABLE decks;")
    if err5 != nil {
        return err5
    }

    _, err4 := txn.Exec("DROP INDEX cards_multiverse_id_index;")
    if err4 != nil {
        return err4
    }

    _, err3 := txn.Exec("DROP INDEX cards_card_name_index;")
    if err3 != nil {
        return err3
    }

    _, err2 := txn.Exec("DROP TABLE cards;")
    if err2 != nil {
        return err2
    }

    _, err1 := txn.Exec("DROP TABLE users;")
    if err1 != nil {
        return err1
    }
    
    return nil
}
