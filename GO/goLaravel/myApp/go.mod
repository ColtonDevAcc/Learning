module myapp

go 1.18

replace github.com/VooDooStack/Voo => ../Voo

require github.com/VooDooStack/Voo v0.0.0-00010101000000-000000000000

require (
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
