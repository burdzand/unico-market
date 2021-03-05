## Instrução

file: DEINFO_AB_FEIRASLIVRES_2014.csv


> Running Local

```bash
Local postgres connection: 
export DB_CONNECTION=host=postgres port=5432 user=unico-dev password=123456 dbname=pismo sslmode=disable  

You need a Postgresql running, it can use docker-compose inside deployents folder:  docker-composer up postgres  

>> go run cmd/unico/main.go

```

> Running Makefile

```bash
make up
make build
make rebuild
make down

```


> Running local Docker compose

```bash
docker-composer up

```

> Testing

```bash
CGO_ENABLED=0 go test -v ./...

```

- [x] HTTP Create Market
- [x] HTTP Update Market
- [x] HTTP Delete Market
- [ ] HTTP Get Market
- [ ] HTTP Save SVC Market
- [x] Func Create Market
- [x] Func Update Market
- [x] Func Remove Market
- [x] Func Create SubPrefecture
- [x] Func Create District


>  DockerFile

```bash
docker build -t unico.io/market:1.0 -f build/Dockerfile .  
docker run -p 3000:3000 -e DB_CONNECTION="host=postgres port=5432 user=unico-dev password=123456 dbname=unico sslmode=disable" unico.io/market:1.0  
```

> Postman Json

```bash
 scripts/postman/Unique.postman_collection
```

> Kubernetes Deploy

```bash
deployents/kube/develop/secrets.yaml
deployents/kube/develop/deployment.yaml
deployents/kube/develop/hpa.yaml
```

> Documentation
Browse to http://localhost:3000/swagger/index.html

<img src="/docs/swagger.png" alt="Unico Market Swagger"/>

