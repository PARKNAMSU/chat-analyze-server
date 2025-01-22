#!/bin/bash

# docker 설치
sudo dnf install docker
sudo systemctl start docker
sudo systemctl enable docker.service

# docker 그룹에 jenkins 계정 추가
sudo usermod -a -G docker jenkins