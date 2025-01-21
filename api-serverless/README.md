# chat-api-serverless

## 개발 중 테스트

```
# sam build or invoke 오류 시 아래 명령어 입력하여 aws public ecr 및 aws public 계정으로 docker 로그인
aws ecr-public get-login-password --region us-east-1 | docker login --username AWS --password-stdin public.ecr.aws

1. sam build
2. sam local start-api --port {port}
```
