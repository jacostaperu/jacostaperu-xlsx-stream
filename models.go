package xlsx_stream

import "encoding/xml"

type StringItem struct {
	T string `xml:"t"`
}

type Sheets struct {
	XMLName xml.Name `xml:"sheets"`
	Text    string   `xml:",chardata"`
	Sheet   []struct {
		Text    string `xml:",chardata"`
		Name    string `xml:"name,attr"`
		SheetId string `xml:"sheetId,attr"`
		State   string `xml:"state,attr"`
		ID      string `xml:"id,attr"`
	} `xml:"sheet"`
}

type Row struct {
	C []struct {
		T string `xml:"t,attr"`
		V string `xml:"v"`
	} `xml:"c"`
}
