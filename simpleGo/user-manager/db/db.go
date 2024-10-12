package db

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type DB interface {
	Write() error
	Read() error
	Add(userData []string) error
	Delete(id int)
	ViewAll()
	ViewEmail()
	Search(email string)
}

type CSV struct {
	File string
	Data [][]string
}

func (d *CSV) ViewAll() error {
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)

	for _, r := range d.Data {
		fmt.Fprintln(w, r[0], "\t", r[1], "\t", r[2])
	}

	w.Flush()

	return nil
}

func (d *CSV) ViewEmail(email string) error {
	record, present, err := d.Search(email)
	if err != nil {
		return err
	}

	if !present {
		return errors.New("email not in db")
	}
	header := d.Data[0]
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)
	fmt.Fprintln(w, header[0], "\t", header[1], "\t", header[2])

	fmt.Fprintln(w, record[0], "\t", record[1], "\t", record[2])

	w.Flush()
	return nil
}

func (d *CSV) Search(email string) ([]string, bool, error) {
	// I want to setup a hash map in the CSV struct to make
	// Lookup faster
	// I need to do some kind of map or something
	// Its okay we are in a compiled language but still.
	for _, record := range d.Data {
		userName := record[1]
		userDomain := record[2]
		idxEmail := fmt.Sprintf("%s@%s", userName, userDomain)
		if idxEmail == email {
			return record, true, nil
		}
	}
	return []string{}, false, nil
}

func (d *CSV) Delete(email string) error {
	// fail fast
	_, present, _ := d.Search(email)
	if !present {
		return errors.New("email not in db")
	}

	var data [][]string
	data = append(data, d.Data[0])
	for _, record := range d.Data[1:] {
		userName := record[1]
		userDomain := record[2]
		userEmail := fmt.Sprintf("%s@%s", userName, userDomain)
		if userEmail == email {
			continue
		}
		data = append(data, record)
	}

	d.Data = data

	d.Write()

	return nil
}

func (d *CSV) Read() error {
	f, err := os.Open(d.File)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	d.Data, err = reader.ReadAll()
	if err != nil {
		return err
	}
	return nil
}

func (d *CSV) Add(userRecord []string) error {
	if len(userRecord) != 2 {
		return errors.New("invalid user record")
	}
	userName := userRecord[0]
	userDomain := userRecord[1]
	userEmail := fmt.Sprintf("%s@%s", userName, userDomain)
	_, present, _ := d.Search(userEmail)
	if present {
		return errors.New("email already in db")
	}

	// remove headers for processing
	record := d.Data[1:]
	n := len(record)
	if n == 0 {
		d.Data = append(d.Data, []string{"0", userName, userDomain})
		return nil
	}
	i, err := strconv.Atoi(record[n-1][0])
	if err != nil {
		return err
	}
	d.Data = append(d.Data, []string{strconv.Itoa(i + 1), userName, userDomain})
	return nil

}

func (d *CSV) Write() error {
	f, err := os.Create(d.File)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	w.WriteAll(d.Data)
	return nil
}

func CSVInit(filename string) (CSV, error) {
	db := CSV{File: filename}
	err := db.Read()
	if err != nil {
		return CSV{}, err
	}
	return db, nil
}
