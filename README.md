[![CircleCI](https://circleci.com/gh/boriska70/bmt/tree/master.svg?style=svg)](https://circleci.com/gh/boriska70/bmt/tree/master)

# bmt


### Build
 - docker build -t boriska70/bmt-builder -f Dockerfile.build .
 - docker run -it --name=bmt-builder boriska70/bmt-builder
 - mkdir .dist
 - docker cp  bmt-builder:/go/src/github.com/boriska70/bmt/.dist/bmt ./.dist/
 - docker build -t boriska70/bmt .

### Run in docker
  - Assuming that we link to elasticsearch running as another docker named es:
  `docker run --rm --name bmt --log-driver=json-file -w /usr/bin--link es:es boriska70/bmt -ies=http://somehost:9200 -oes=http://es:9200`
  - Elasticsearch can be started as
  `docker run -d --name es -p 9200:9200 -p 9300:9300 elasticsearch elasticsearch -Des.network.host=0.0.0.0 -Des.network.bind_host=0.0.0.0 -Des.cluster.name=elasticlaster -Des.node.name=$(hostname)`
  - Kibana run: docker run --link es:elasticsearch -d -p5601:5601 --name kibana kibana
  - NOTE: run ```sudo sysctl vm.max_map_count=262144``` before starting Elasticsearch

### Run in docker-compose
  - docker-compose up
  - NOTE: run ```sudo sysctl vm.max_map_count=262144``` before starting compose

