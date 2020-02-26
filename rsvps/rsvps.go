package rsvps

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/pkg/errors"
)

// RSVP is a data model representation of RSVPs
type RSVP struct {
	ID         int32  `json:"id"`
	Email      string `json:"email"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Attending  string `json:"attending"`
	FoodChoice string `json:"food"`
	GuestName  string `json:"extraAttendees"`
	GuestFood  string `json:"guestFood"`
	Note       string `json:"message"`
}

// RegisterRoutes attaches all routes related to RSVPs to the HTTP mux
func RegisterRoutes(db *sql.DB) {
	fmt.Println("Registering RSVP routes")
	repository := newRepository(db)

	http.HandleFunc("/rsvps/new", createRSVPHandler(db, repository))
	http.HandleFunc("/rsvps", allRSVPsHandler(db, repository))

	fmt.Println("Finished registering RSVP routes")
}

func createRSVPHandler(db *sql.DB, rep repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("handling request to create rsvp")
		defer fmt.Println("finished handling request")
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error reading request body"))
		}

		rsvp := RSVP{}
		err = json.Unmarshal(body, &rsvp)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error unmarshalling JSON"))
		}

		err = validateRsvp(rsvp)
		if err != nil {
			fmt.Println(errors.Wrap(err, "error validating new rsvp"))
			w.WriteHeader(http.StatusUnprocessableEntity)
			w.Write([]byte(err.Error()))
			return
		}

		err = rep.createRSVP(rsvp)
		if err != nil {
			log.Fatal(errors.Wrap(err, "error inserting rsvp into database"))
		}

		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	}
}

func allRSVPsHandler(db *sql.DB, rep repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		rsvps, err := rep.allRSVPs(db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error querying for RSVPs"))
			return
		}

		j, err := json.Marshal(rsvps)
		if err != nil {
			log.Fatal(err)
		}

		json.NewEncoder(w).Encode(string(j))
	}
}

func validateRsvp(req RSVP) error {
	if req.Email == "" {
		return errors.New("email is required")
	} else if req.FirstName == "" {
		return errors.New("first name is required")
	} else if req.LastName == "" {
		return errors.New("last name is required")
	} else if req.FoodChoice == "" {
		return errors.New("food choice is required")
	}

	return nil
}
