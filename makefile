setup:
	cd front
	npm install
	yarn install
	yarn add -D @types/react

watch-front:
	cd front
	yarn build
	yarn start
