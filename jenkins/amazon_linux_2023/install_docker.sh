#!/bin/bash

sudo dnf install docker
sudo systemctl start docker
sudo systemctl enable docker.service
sudo usermod -a -G docker jenkins