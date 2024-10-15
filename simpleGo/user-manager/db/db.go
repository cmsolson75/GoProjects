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
	Add(userRecord []string) error
	Delete(email string)
	ViewAll() error
	ViewEmail(email string)
	Search(email string) ([]string, bool, error)
}

type CSV struct {
	file     string
	data     [][]string
	emailMap map[string][]string
}

func (d *CSV) ViewAll() error {
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)

	for _, r := range d.data {
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
	header := d.data[0]
	w := tabwriter.NewWriter(os.Stdout, 5, 0, 3, ' ', 0)
	fmt.Fprintln(w, header[0], "\t", header[1], "\t", header[2])

	fmt.Fprintln(w, record[0], "\t", record[1], "\t", record[2])

	w.Flush()
	return nil
}

func (d *CSV) Search(email string) ([]string, bool, error) {
	if record, exists := d.emailMap[email]; exists {
		return record, true, nil
	}
	return []string{}, false, nil
}

func sliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func (d *CSV) Delete(email string) error {
	// Check if email exists in the map
	record, exists := d.emailMap[email]
	if !exists {
		return errors.New("email not in db")
	}

	// find index of record
	var index int = -1
	for i, rec := range d.data {
		if sliceEqual(rec, record) {
			index = i
			break
		}
	}

	// If the record was found remove in place
	if index != -1 {
		d.data = append(d.data[:index], d.data[index+1:]...) // remove
	}

	// Remove from emailMap
	delete(d.emailMap, email)

	return d.Write()
}

func (d *CSV) Read() error {
	f, err := os.Open(d.file)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	d.data, err = reader.ReadAll()
	if err != nil {
		return err
	}

	d.emailMap = make(map[string][]string)
	for _, record := range d.data[1:] {
		if len(record) >= 3 {
			email := fmt.Sprintf("%s@%s", record[1], record[2])
			d.emailMap[email] = record
		}
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
	record := d.data[1:]
	n := len(record)
	newID := "0"
	if n > 0 {
		i, err := strconv.Atoi(record[n-1][0])
		if err != nil {
			return err
		}
		newID = strconv.Itoa(i + 1)
	}
	newRecord := []string{newID, userName, userDomain}
	d.data = append(d.data, newRecord)

	d.emailMap[userEmail] = newRecord
	return nil

}

func (d *CSV) Write() error {
	f, err := os.Create(d.file)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	w.WriteAll(d.data)
	return nil
}

func CSVInit(filename string) (*CSV, error) {
	db := &CSV{file: filename}
	err := db.Read()
	if err != nil {
		return &CSV{}, err
	}
	return db, nil
}
