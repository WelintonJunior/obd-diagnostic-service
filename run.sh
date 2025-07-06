#!/bin/bash

./bin/api_server initApiServer &
./bin/api_server initCli &
wait
