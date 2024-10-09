package slicer

import (
	"reflect"
	"testing"
)

func TestValidEmail(t *testing.T) {
	inputEmail := "cmsoapp@gmail.com"

	got, err := EmailSlicer(inputEmail)
	want := UserAccount{
		"cmsoapp",
		"gmail.com",
	}
	assertNoError(t, err)
	assertUserAccount(t, got, want)
}

func TestInvalidEmail(t *testing.T) {
	tests := []struct {
		email string
		user  UserAccount
		want  error
	}{
		{"cmsoapp@", UserAccount{}, ErrMessageInvalidEmail},
		{"cmsoapp", UserAccount{}, ErrMessageInvalidEmail},
	}

	for _, tt := range tests {
		got, err := EmailSlicer(tt.email)
		assertError(t, err, tt.want)
		assertUserAccount(t, got, tt.user)
	}
}

func assertUserAccount(t testing.TB, got, want UserAccount) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error, none recieved")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("revieved error, none expected")
	}
}
