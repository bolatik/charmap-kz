package cp1048

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestBasics(t *testing.T) {
	encoder := KZ1048.NewEncoder()
	s, e := encoder.String("Бұл қазақша сөйлем")
	if e != nil {
		t.Error(e)
	}

	ioutil.WriteFile("example.txt", []byte(s), os.ModePerm)
	// Декодировка в UTF-8
	f, e := os.Open("example.txt")
	if e != nil {
		t.Error(e)
	}
	defer f.Close()
	decoder := KZ1048.NewDecoder()
	reader := decoder.Reader(f)
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(b))
}
