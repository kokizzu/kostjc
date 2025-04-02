#!/usr/bin/env bash

SERVER=root@benalu.dev
./no_restart.sh &&
ssh $SERVER sudo bash /home/kostjc/reload_services.sh

