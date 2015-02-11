/*

    Incomplete, only focuses on video


*/

package vast

import (
    "log"
    "encoding/xml"
)

type VAST struct {
    XMLName xml.Name `xml:"VAST"`
    Ad []Ad `xml:"Ad"`
 }

type Ad struct {
    XMLName xml.Name `xml:"Ad"`
    InLine []InLine  `xml:"InLine"`
    Id string `xml:"type,attr"`
}

type InLine struct {
    XMLName xml.Name `xml:"InLine"`
    Impressions []Impression `xml:"Impression"`
    Error string `xml:"Error"`
    Creatives []Creatives `xml:"Creatives"`
}

type Impression struct {
    XMLName xml.Name `xml:"Impression"`
    Impression string `xml:"Impression"`
    URL    string      `xml:",chardata"`
}

type Creatives struct {
    XMLName xml.Name `xml:"Creatives"`
    Creative []Creative `xml:"Creative"`
}

type Creative struct {
    XMLName xml.Name `xml:"Creative"`
    Linear []Linear `xml:"Linear"`
}

type Linear struct {
    XMLName xml.Name `xml:"Linear"`
    Duration string `xml:"Duration"`
    TrackingEvents []TrackingEvents `xml:"TrackingEvents"`
}

type TrackingEvents struct {
    XMLName xml.Name `xml:"TrackingEvents"`
    Events []Tracking `xml:"Tracking"`
}

type Tracking struct {
    XMLName     xml.Name    `xml:"Tracking"`
    Event       string      `xml:"type,attr"`
    Tracking    string      `xml:",chardata"`
}

func ParseBidXML(xmlString string) VAST {
    var vast VAST

    err := xml.Unmarshal([]byte(xmlString), &vast)
    if err != nil {
        log.Fatal(err)
        panic("Unable to return XML")
    } else {
        return vast
    }
}

func ParseBidXMLBytes(xmlString []byte) VAST {
    var vast VAST

    err := xml.Unmarshal(xmlString, &vast)
    if err != nil {
        log.Fatal(err)
        panic("Unable to return XML")
    } else {
        return vast
    }
}