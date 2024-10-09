# Project: User Manager


## Why?

I want to learn the CSV package.


## Notes

### CSV DOCS

Just like in my tests, using error messages that are standardized it the norm. 

exp
```
var (
	ErrBareQuote  = errors.New("bare \" in non-quoted-field")
	ErrQuote      = errors.New("extraneous or missing \" in quoted-field")
	ErrFieldCount = errors.New("wrong number of fields")

	// Deprecated: ErrTrailingComma is no longer used.
	ErrTrailingComma = errors.New("extra delimiter at end of line")
)
```

type Reader
- This uses the rune; what is this?
- This reads records from a CSV-encoded file
- Returned by NewReader
- This reader is for reading a CSV and converting it to a string
- You can change the delimiter

This library is made up of Reader & Writer
- Both handle what they say they do.

The writer looks like it is doing the following
- Writing data to a buffer
- You flush the buffer with Writer.Flush.
- This dumps the data to the io.Writer
- Check for error.
    - encorages to use internal error w.Error()


MOCH Inputs
```
in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	// this is moching the reader
	r := csv.NewReader(strings.NewReader(in))
	records, err := r.ReadAll()
```

What i need to learn how to do 
- How do you write to a file in Go?
- `w := csv.NewWriter(file)` the file var is where the writer buffer will write its data too.



TUTORIAL: [Twilio Tutorial on CSV in GO](https://www.twilio.com/en-us/blog/read-write-csv-file-go)

