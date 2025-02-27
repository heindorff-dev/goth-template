compose:
	@docker compose -f docker-compose.yml up --build

run:
	@air

css:
	@npx tailwindcss -i ./view/css/tailwind.css -o ./view/public/styles.css --watch

css-nowatch:
	@npx tailwindcss -i ./view/css/tailwind.css -o ./view/public/styles.css