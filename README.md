[![CircleCI](https://circleci.com/gh/boriska70/bmt/tree/master.svg?style=svg)](https://circleci.com/gh/boriska70/bmt/tree/master)

# bmt


### Build
 - docker build -t boriska70/bmt-builder -f Dockerfile.build .
 - docker run -it --name=bmt-builder boriska70/bmt-builder
 - mkdir .dist
 - docker cp  bmt-builder:/go/src/github.com/boriska70/bmt/.dist/bmt ./.dist/
 - docker build -t boriska70/bmt .

### Run in docker
  - Assuming that we link to elasticsearch running as another docker named es and queries.yml is sitting in the current folder:
  `docker run --rm -w="/usr/bin" -v .:/usr/share --link es:elasticsearch /usr/bin/bmt -ies=http://elasticsearch:9200 -oes=http://elasticsearch:9200 -queries=/usr/share/queries.yml`
  - Elasticsearch can be started as
  `docker run -d --name es -p 9200:9200 -p 9300:9300 elasticsearch:5.1.1 elasticsearch -Enetwork.host=0.0.0.0 -Enetwork.bind_host=0.0.0.0 -Ecluster.name=elasticlaster -Enode.name=$(hostname)`
  - Kibana run: docker run --link es:elasticsearch -d -p5601:5601 --name kibana kibana
  - NOTE: run ```sudo sysctl vm.max_map_count=262144``` before starting Elasticsearch

### Run in docker-compose
  - docker-compose up
  - NOTE: run ```sudo sysctl vm.max_map_count=262144``` before starting compose

