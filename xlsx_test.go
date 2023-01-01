package xlsx_stream

import (
	"archive/zip"
	"io"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	expected := []string{
		"name,pais",
		"hola s1,cl",
		"hola s2,pe",
		"hola s3,es",
	}
	f, err := zip.OpenReader("testdata/xlsx_test.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := NewReader(f)

	res := []string{}
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		res = append(res, strings.Join(row, ","))
	}

	for k, v := range expected {
		if res[k] != v {
			t.Fatalf("unexpected results, got: %q, want: %q", res[k], v)
		}
	}
}

func TestReaderWorksheet2(t *testing.T) {
	expected := []string{
		"code,product",
		"1,perro",
		"2,gato",
		"3.1,loro",
	}
	f, err := zip.OpenReader("testdata/xlsx_test.xlsx")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	r := NewReader(f)
	r.Worksheet = "Sheet2"

	res := []string{}
	for {
		row, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatal(err)
		}
		res = append(res, strings.Join(row, ","))
	}

	for k, v := range expected {
		if res[k] != v {
			t.Fatalf("unexpected results, got: %q, want: %q", res[k], v)
		}
	}
}

func BenchmarkRead(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		f, err := zip.OpenReader("testdata/xlsx_test.xlsx")
		if err != nil {
			b.Fatal(err)
		}
		defer f.Close()
		r := NewReader(f)

		for {
			_, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				b.Fatal(err)
			}
		}
	}
}
