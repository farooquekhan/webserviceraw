module todo

replace farooque.in/WebServicesRaw/storage => ./storage

replace farooque.in/WebServicesRaw/webutils => ./webutils

replace farooque.in/WebServicesRaw/webservice => ./webservice

go 1.15

require (
	farooque.in/WebServicesRaw/storage v0.0.0-00010101000000-000000000000 // indirect
	farooque.in/WebServicesRaw/webservice v0.0.0-00010101000000-000000000000
	farooque.in/WebServicesRaw/webutils v0.0.0-00010101000000-000000000000 // indirect
)
