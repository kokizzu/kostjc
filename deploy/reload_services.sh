#!/usr/bin/env bash

mv /home/kostjc/.env.production /home/kostjc/.env &&
cp /home/kostjc/kostjc.service /lib/systemd/system/ &&
systemctl daemon-reload &&
systemctl enable kostjc &&
systemctl restart kostjc
