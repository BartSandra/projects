#!/bin/bash

sudo apt install -y nginx
sudo bash metrics.sh
sudo cp ./nginx.conf /etc/nginx/nginx.conf
sudo nginx -s reload
sudo cp metrics.html /usr/share/nginx/html/metrics.html
sudo cp ./prometheus.yml /etc/prometheus/prometheus.yml
sudo systemctl restart prometheus.service