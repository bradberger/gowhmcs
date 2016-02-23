package whmcs

import (
    "fmt"
    "time"
)

var (
    // DateFormat is the Go date format string that must match the WHMCS install's
    // date format to support WHMCS date fields which must be formatted according
    // to the system date/time format.00
    DateFormat       = "01/02/2006"
)

// Date implements WHMCS formatted JSON marshaler for time values that need to
// match the WHMCS system format.
type Date time.Time

// MarshalJSON changes the date to a format which WHMCS recognizes.
// @TODO This is for the AddTransaction momentarily, and must match the WHMCS
// date format of the local install because the WHMCS API is whack. We need to
// be able to change this as needed. Also, other API endpoints require other
// fixed date formats, so we have to take that into account down the road too.
func (d Date) MarshalJSON() ([]byte, error) {
	t := time.Time(d)
	f := t.Format(DateFormat)
	str := fmt.Sprintf(`"%s"`, f)
	return []byte(str), nil
}

// Time returns the original time.Time value of the date.
func (d Date) Time() time.Time {
	return time.Time(d)
}


// InvoiceDate is a time.Time value that gets marshaled into the correct
// format for the WHMCS invoice date fields.
type InvoiceDate time.Time

// MarshalJSON changes the date to a format which WHMCS recognizes for invoices.
func (d InvoiceDate) MarshalJSON()  ([]byte, error) {
    t := time.Time(d)
	f := t.Format("20060102")
	str := fmt.Sprintf(`"%s"`, f)
	return []byte(str), nil
}

// Time returns the original time.Time value of the date.
func (d InvoiceDate) Time() time.Time {
    return time.Time(d)
}
