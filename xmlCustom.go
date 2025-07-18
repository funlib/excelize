// Copyright 2016 - 2025 The excelize Authors. All rights reserved. Use of
// this source code is governed by a BSD-style license that can be found in
// the LICENSE file.
//
// Package excelize providing a set of functions that allow you to write to and
// read from XLAM / XLSM / XLSX / XLTM / XLTX files. Supports reading and
// writing spreadsheet documents generated by Microsoft Excel™ 2007 and later.
// Supports complex components by high compatibility, and provided streaming
// API for generating or reading data from a worksheet with huge amounts of
// data. This library needs Go version 1.23 or later.

package excelize

import "encoding/xml"

// xlsxCustomProperties directly maps the element for the custom file properties
// part, that represents additional information. The information can be used as
// metadata for XML.
type xlsxCustomProperties struct {
	XMLName  xml.Name       `xml:"http://schemas.openxmlformats.org/officeDocument/2006/custom-properties Properties"`
	Vt       string         `xml:"xmlns:vt,attr"`
	Property []xlsxProperty `xml:"property"`
}

// xlsxProperty directly maps the element specifies a single custom file
// property. Custom file property type is defined through child elements in the
// File Properties Variant Type namespace. Custom file property value can be set
// by setting the appropriate Variant Type child element value.
type xlsxProperty struct {
	XMLName    xml.Name `xml:"property"`
	FmtID      string   `xml:"fmtid,attr"`
	PID        int      `xml:"pid,attr"`
	Name       string   `xml:"name,attr,omitempty"`
	LinkTarget string   `xml:"linkTarget,attr,omitempty"`
	Vector     *string  `xml:"vt:vector"`
	Array      *string  `xml:"vt:array"`
	Blob       *string  `xml:"vt:blob"`
	Oblob      *string  `xml:"vt:oblob"`
	Empty      *string  `xml:"vt:empty"`
	Null       *string  `xml:"vt:null"`
	I1         *int8    `xml:"vt:i1"`
	I2         *int16   `xml:"vt:i2"`
	I4         *int32   `xml:"vt:i4"`
	I8         *int64   `xml:"vt:i8"`
	Int        *int     `xml:"vt:int"`
	Ui1        *uint8   `xml:"vt:ui1"`
	Ui2        *uint16  `xml:"vt:ui2"`
	Ui4        *uint32  `xml:"vt:ui4"`
	Ui8        *uint64  `xml:"vt:ui8"`
	Uint       *uint    `xml:"vt:uint"`
	R4         *float32 `xml:"vt:r4"`
	R8         *float64 `xml:"vt:r8"`
	Decimal    *string  `xml:"vt:decimal"`
	Lpstr      *string  `xml:"vt:lpstr"`
	Lpwstr     *string  `xml:"vt:lpwstr"`
	Bstr       *string  `xml:"vt:bstr"`
	Date       *string  `xml:"vt:date"`
	FileTime   *string  `xml:"vt:filetime"`
	Bool       *bool    `xml:"vt:bool"`
	Cy         *string  `xml:"vt:cy"`
	Error      *string  `xml:"vt:error"`
	Stream     *string  `xml:"vt:stream"`
	Ostream    *string  `xml:"vt:ostream"`
	Storage    *string  `xml:"vt:storage"`
	Ostorage   *string  `xml:"vt:ostorage"`
	Vstream    *string  `xml:"vt:vstream"`
	ClsID      *string  `xml:"vt:clsid"`
}

// decodeCustomProperties specifies to an OOXML document custom properties.
// decodeCustomProperties just for deserialization.
type decodeCustomProperties struct {
	XMLName  xml.Name         `xml:"http://schemas.openxmlformats.org/officeDocument/2006/custom-properties Properties"`
	Vt       string           `xml:"xmlns:vt,attr"`
	Property []decodeProperty `xml:"property"`
}

// decodeProperty specifies to an OOXML document custom property. decodeProperty
// just for deserialization.
type decodeProperty struct {
	XMLName    xml.Name `xml:"property"`
	FmtID      string   `xml:"fmtid,attr"`
	PID        int      `xml:"pid,attr"`
	Name       string   `xml:"name,attr,omitempty"`
	LinkTarget string   `xml:"linkTarget,attr,omitempty"`
	Vector     *string  `xml:"vector"`
	Array      *string  `xml:"array"`
	Blob       *string  `xml:"blob"`
	Oblob      *string  `xml:"oblob"`
	Empty      *string  `xml:"empty"`
	Null       *string  `xml:"null"`
	I1         *int8    `xml:"i1"`
	I2         *int16   `xml:"i2"`
	I4         *int32   `xml:"i4"`
	I8         *int64   `xml:"i8"`
	Int        *int     `xml:"int"`
	Ui1        *uint8   `xml:"ui1"`
	Ui2        *uint16  `xml:"ui2"`
	Ui4        *uint32  `xml:"ui4"`
	Ui8        *uint64  `xml:"ui8"`
	Uint       *uint    `xml:"uint"`
	R4         *float32 `xml:"r4"`
	R8         *float64 `xml:"r8"`
	Decimal    *string  `xml:"decimal"`
	Lpstr      *string  `xml:"lpstr"`
	Lpwstr     *string  `xml:"lpwstr"`
	Bstr       *string  `xml:"bstr"`
	Date       *string  `xml:"date"`
	FileTime   *string  `xml:"filetime"`
	Bool       *bool    `xml:"bool"`
	Cy         *string  `xml:"cy"`
	Error      *string  `xml:"error"`
	Stream     *string  `xml:"stream"`
	Ostream    *string  `xml:"ostream"`
	Storage    *string  `xml:"storage"`
	Ostorage   *string  `xml:"ostorage"`
	Vstream    *string  `xml:"vstream"`
	ClsID      *string  `xml:"clsid"`
}

// CustomProperty directly maps the custom property of the workbook. The value
// date type may be one of the following: int32, float64, string, bool,
// time.Time, or nil.
type CustomProperty struct {
	Name  string
	Value interface{}
}
