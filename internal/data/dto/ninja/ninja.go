package ninja

import "encoding/xml"

type NinjaDTO struct {
	XMLName xml.Name `xml:"ninja2"`
	Content string   `xml:",chardata"`
}
