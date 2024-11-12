env:
	sed "s/=.*/=/" .env > .env.example
dev:
	air -d
