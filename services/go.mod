module github.com/satheeshds/shortly/services

go 1.18

replace github.com/satheeshds/shortly/interfaces => ../interfaces

replace github.com/satheeshds/shortly/mock => ../mock

require (
	github.com/golang/mock v1.6.0
	github.com/satheeshds/shortly/interfaces v0.0.0-00010101000000-000000000000
	github.com/satheeshds/shortly/mock v0.0.0-00010101000000-000000000000
)
