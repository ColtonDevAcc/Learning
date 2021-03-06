module myapp

go 1.17

replace github.com/tsawler/celeritas => ../celeritas

require (
	github.com/CloudyKit/jet/v6 v6.1.0
	github.com/go-chi/chi/v5 v5.0.4
	github.com/tsawler/celeritas v0.0.0
)

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/alexedwards/scs/v2 v2.4.0 // indirect
	github.com/joho/godotenv v1.3.0 // indirect
)
