#!/bin/bash

# git 설치
sudo dnf install git

# jenkins 사용자 확인
ps aux | grep jenkins

# 사용자를 jenkins 로 변환
sudo -u jenkins /bin/bash

# jenkins 홈 directory 에 .ssh directory 생성 후 이동
mkdir /var/lib/jenkins/.ssh
cd /var/lib/jenkins/.ssh

# ssh 키 생성
ssh-keygen -t ed25519 -f /var/lib/jenkins/.ssh/github_PARKNAMSU

# Github에 ssh 첫 연결 시 호스트 체크 skip 설정
sudo cat <<EOF >> /etc/ssh/ssh_config
StrictHostKeyChecking no
EOF

# 아래의 내용은 jenkins 계정으로 실행
cat <<EOF >> /var/lib/jenkins/.ssh/config
Host github.com
  Hostname ssh.github.com
  Port 443
EOF

ssh -T git@github.com