# SCAP
A GO module of the Security Content Automation Protocol (SCAP) Specification

This project represents complete parsers of SCAP standard. Users are advised to use [OpenSCAP](https://www.open-scap.org/) for SCAP evaluation and for any advanced use of SCAP standard. This project is beneficial for quick manipulation of SCAP data within golang code.

## Examplary usage
```
	import "github.com/gocomply/scap/pkg/scap/scap_document"
	doc, err := scap_document.ReadDocumentFromFile("DOCUMENT.XML")
```
