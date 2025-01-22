#!/bin/bash

sudo amazon-linux-extras install docker
sudo service docker start
sudo systemctl enable docker.service

sudo usermod -a -G docker jenkins