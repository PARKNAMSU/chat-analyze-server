# chat-system

## 프로젝트 설명

- 채팅 시스템을 제공해주는 서비스

## 프로젝트 구성

- [broker server (web socket)](https://github.com/PARKNAMSU/chat-analyze-server/tree/main/broker-server)
- [api serverless](https://github.com/PARKNAMSU/chat-analyze-server/tree/main/api-serverless)
- [kafka](https://github.com/PARKNAMSU/chat-analyze-server/tree/main/kafka)
- [client]()

## 프로젝트 아키텍처

<img width="727" alt="Image" src="https://github.com/user-attachments/assets/6199ba51-9d13-4bc0-aa23-0af8b09aa99e" />

### 아키텍처 설명

1. API 비즈니스 로직을 처리할 Lambda 함수
2. API 게이트웨이를 통해 Client 와 Lambda 함수 간 통신
3. 서비스의 데이터 저장소
4. 메세지 Publish, Poll 이벤트 처리를 위한 Kafka 서버.
5. Client 의 Kafka 구독 처리를 위한 중계 서버
6. 서비스를 이용하는 클라이언트

## 사용 stack
