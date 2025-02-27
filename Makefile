# Build project and DB in docker containers. Run project with Air.
compose:
	@docker compose -f docker-compose.yml up --build

# Run project with Air.
build:
	@air

# Generate *_templ files for frontend.
templ:
	@templ generate

# Generate CSS and watch for changes.
css:
	@npx tailwindcss -i ./view/css/tailwind.css -o ./view/public/styles.css --watch

# Generate CSS.
css-nowatch:
	@npx tailwindcss -i ./view/css/tailwind.css -o ./view/public/styles.css

# Generate database models.
sqlc:
	@(cd database ; sqlc generate)

