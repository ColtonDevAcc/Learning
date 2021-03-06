module myapp

go 1.18

replace github.com/VooDooStack/Voo => ../Voo

require (
	github.com/CloudyKit/jet/v6 v6.1.0
	github.com/VooDooStack/Voo v0.0.0-00010101000000-000000000000
	github.com/go-chi/chi/v5 v5.0.7
)

require github.com/go-sql-driver/mysql v1.6.0 // indirect

require (
	github.com/CloudyKit/fastprinter v0.0.0-20200109182630-33d98a066a53 // indirect
	github.com/alexedwards/scs/v2 v2.5.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgconn v1.12.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.11.0 // indirect
	github.com/jackc/pgx/v4 v4.16.1 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
	github.com/lib/pq v1.10.5 // indirect
	github.com/upper/db/v4 v4.5.2
	golang.org/x/crypto v0.0.0-20220507011949-2cf3adece122
	golang.org/x/text v0.3.7 // indirect
)
