package rsvps

import (
	"database/sql"

	"github.com/pkg/errors"
)

type repository struct {
	DB *sql.DB
}

func newRepository(db *sql.DB) repository {
	return repository{DB: db}
}

func (r repository) createRSVP(rsvp RSVP) error {
	res, err := r.DB.Exec(`
	INSERT INTO rsvp (email, first_name, last_name, attending, food_choice, guest_name, note)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
	`, rsvp.Email, rsvp.FirstName, rsvp.LastName, rsvp.Attending, rsvp.FoodChoice, rsvp.GuestName, rsvp.Note)
	if err != nil {
		return errors.Wrap(err, "error inserting rsvp into database")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "error checking rows affected")
	}

	if rows == 0 {
		return errors.Wrap(err, "no rows inserted")
	}

	return nil
}

func (r repository) allRSVPs(db *sql.DB) ([]RSVP, error) {
	rows, err := db.Query("SELECT id, email, first_name, last_name, attending, food_choice, guest_name, note FROM rsvp;")
	if err != nil {
		return []RSVP{}, errors.Wrap(err, "error retrieving rsvps from database")
	}
	defer rows.Close()

	var rsvps []RSVP
	for rows.Next() {
		var (
			id         int32
			email      string
			firstName  string
			lastName   string
			attending  bool
			foodChoice string
			guestName  string
			note       string
		)

		err := rows.Scan(&id, &email, &firstName, &lastName, &attending, &foodChoice, &guestName, &note)
		if err != nil {
			return []RSVP{}, errors.Wrap(err, "error scanning rows")
		}

		rsvp := RSVP{
			ID:         id,
			Email:      email,
			FirstName:  firstName,
			LastName:   lastName,
			Attending:  attending,
			FoodChoice: foodChoice,
			GuestName:  guestName,
			Note:       note,
		}

		rsvps = append(rsvps, rsvp)
	}

	return rsvps, nil
}