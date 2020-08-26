# SCAP ![Build CI](https://github.com/gocomply/scap/workflows/Build%20CI/badge.svg) [![Gitter](https://badges.gitter.im/GoComply/community.svg)](https://gitter.im/GoComply/community?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)
A GO module of the Security Content Automation Protocol (SCAP) Specification

This project represents complete parsers of SCAP standard. Parsers are generated automatically by [xsd2go](https://github.com/GoComply/xsd2go) project.

Users are advised to use [OpenSCAP](https://www.open-scap.org/) for SCAP evaluation and for any advanced use of SCAP standard. This project is beneficial for quick manipulation of SCAP data within golang code.

## Examplary usage
```
	import "github.com/gocomply/scap/pkg/scap/scap_document"
	doc, err := scap_document.ReadDocumentFromFile("DOCUMENT.XML")
```

## Developer information
How to re-generate SCAP parsers with the latest [xsd2go](https://github.com/GoComply/xsd2go).
```
make all
```
