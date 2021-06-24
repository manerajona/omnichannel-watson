module main

go 1.16

require (
	github.com/olivere/elastic/v7 v7.0.22
	monitor v1.0.0
)

replace monitor => ./../../framework/monitor
