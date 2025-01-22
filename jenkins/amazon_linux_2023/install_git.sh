#!/bin/bash

sudo dnf install git
ps aux | grep jenkins
sudo -u jenkins /bin/bash

mkdir /var/lib/jenkins/.ssh
cd /var/lib/jenkins/.ssh
ssh-keygen -t ed25519 -f /var/lib/jenkins/.ssh/github_PARKNAMSU

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