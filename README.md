# Prerequisite
- If you need to be able to have conversation list on website you need to create line account and replace message.secret and message.accessToken with you own config in api/config/config.json
- You need some service that will proxy you localhost service to public site such as ngrok to let line service hook back to you in path `/hook/line`

# How to run
## With docker-compose
- Install docker
- Run command `docker-compose up`

## With manual
- Install golang, nodejs and make
- Run command `make start-api` and `make start-web`
