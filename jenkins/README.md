# 자동 배포 Jenkins 서버 세팅

## Jenkins 처음 접속 시 password 확인

```
# 접속 domain http://[server ip or hostname]:8080
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```

## Git 설치 후 가이드

1. 공개 키 복사 후 github repo > settings > Deploy keys > Add deploy key에 추가
   - `cat /var/lib/jenkins/.ssh/{file name}`
2. Jenkins Main > Credentials > System > Globalcredentials > Add Credentials 진입해 secret key 등록
   - Kind : SSH Username with private key
   - Scope : Global (Jenkins, nodes, items, all child items, etc)
   - Username : github_global-{name}
   - private Key: `cat /var/lib/jenkins/.ssh/${file name}`

## Jenkins 서버에 배포 시 사용할 aws 계정 설정

```
sudo -u jenkins /bin/bash
mkdir /var/lib/jenkins/.aws
cat <<EOF >> /var/lib/jenkins/.aws/config
[profile {InstanceRole}]
output = json
region = {region}
role-arn = {InstanceRole ARN}
credential_source = Ec2InstanceMetadata
EOF
```
