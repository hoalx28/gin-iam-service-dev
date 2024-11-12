install:
	go install github.com/air-verse/air@latest
env:
	sed "s/=.*/=/" .env > .env.example
dev:
	air -d
