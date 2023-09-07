postgresinit:
	docker run --name chat-room -p 5432:5432 -e POSTGRES_USER=dmilyano -e POSTGRES_PASSWORD=qwerty -e POSTGRES_DB=chatroomdb -d postgres:13.3
.PHONY: postgresinit