run:
	@docker compose -f docker-compose.yml up --build

css:
	@npx tailwindcss -i ./view/css/tailwind.css -o ./view/public/styles.css --watch