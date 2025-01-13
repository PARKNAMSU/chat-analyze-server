# chat-analyze-server
## 프로젝트 설명
* 상호간의 채팅 대화를 바탕으로 상대의 기분 및 상태를 분석해 주는 애플리케이션의 채팅 서버

## 프로젝트 세팅
### 종속성 설치
```
go mod tidy
```
### .env 파일 세팅
```
# sample

PORT=
SERVER_API_KEY=""
ENVIROMENT=""

MYSQL_SLAVE_USER=""
MYSQL_SLAVE_PASSWORD=""
MYSQL_SLAVE_HOST=""
MYSQL_SLAVE_DATABASE=""

MYSQL_MASTER_USER=""
MYSQL_MASTER_PASSWORD=""
MYSQL_MASTER_HOST=""
MYSQL_MASTER_DATABASE=""

PG_SLAVE_USER=""
PG_SLAVE_PASSWORD=""
PG_SLAVE_HOST=""
PG_SLAVE_DATABASE=""

PG_MASTER_USER=""
PG_MASTER_PASSWORD=""
PG_MASTER_HOST=""
PG_MASTER_DATABASE=""
```

## 프로젝트 아키텍처

## 사용 stack