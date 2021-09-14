docker build -t go-agan-tryout-api .
docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network
docker run --rm -d --name dev-agan-tryout-api --network dev-network -p 3000:3000 go-agan-tryout-api