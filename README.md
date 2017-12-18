## Build release and push to docker hub registry

```bash
make release -e NS=zion -e IMAGE_NAME=health-checker -e VERSION=latest
```
Where:
 - `NS` namespace in docker hub
 - `IMAGE_NAME` name of publishing image
 - `VERSION` version of release that will be published to registry


## To run service locally
```bash
docker pull postgres:9.6-alpine
docker run -d --rm --name health-checker-postgres --hostname postgres postgres:9.6-alpine
docker run --rm --link health-checker-postgres:postgres postgres:9.6-alpine \
	psql -h postgres -U postgres -c "CREATE DATABASE checker;"
docker run -d --rm --link health-checker-postgres:postgres zion/health-checker:latest /service \
	-database checker -host health-checker-postgres -user postgres -interval 10 -source /url_list.txt
```

## CLI to see statistics of resource availability
Detail statistic with specified url to resource:
```bash
docker run --rm --link health-checker-postgres:postgres zion/health-checker:latest /cli \
	-database checker -host health-checker-postgres -user postgres -url "https://google.com" -start 01-01-2017 -end 31-12-2017
```

General statistic of availability all tracked resources:
```bash
docker run --rm --link health-checker-postgres:postgres zion/health-checker:latest /cli \
	-database checker -host health-checker-postgres -user postgres -start 01-01-2017 -end 31-12-2017
```
