#!/usr/bin/env bash

echo $(curl -XPOST -d '{"author": "Mickey Mause","timestamp": "2017-03-10T14:00:00Z","tweet": "Hello from Mickey"}' http://localhost:9200/tweets/1)
echo $(curl -XPOST -d '{"author": "Mini Mause","timestamp": "2017-03-10T16:00:00Z","tweet": "Would you like some coffee?"}' http://localhost:9200/tweets/1)
echo $(curl -XPOST -d '{"author": "Mini Mause","timestamp": "2017-03-10T16:30:00Z","tweet": "Would you like some more coffee?"}' http://localhost:9200/tweets/1)
echo $(curl -XPOST -d '{"author": "Mickey Mause","timestamp": "2017-03-10T18:00:00Z","tweet": "Hello from Mickey"}' http://localhost:9200/tweets/1)
