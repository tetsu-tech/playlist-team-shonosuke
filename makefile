setup:
	cd front
	npm install
	yarn install
	yarn add react-router-dom
	yarn add -D @types/react

watch-front:
	cd front
	yarn build
	yarn start
