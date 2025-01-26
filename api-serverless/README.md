# chat-api-serverless

## 프로젝트 패키지 구조
<img width="799" alt="Image" src="https://github.com/user-attachments/assets/1fa551c2-e012-4657-a66f-e62e20b34aaf" />

## 사용 스택

## 보안 정책

## 주요 기능

## 로컬 환경 테스트

```
# sam build or invoke 오류 시 아래 명령어 입력하여 aws public ecr 및 aws public 계정으로 docker 로그인
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

1. sam build
2. sam local start-api --port {port}
```
