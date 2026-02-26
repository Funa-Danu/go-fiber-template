# go-fiber-template

Go 1.26 + Fiber v3 템플릿 프로젝트입니다.

## 요구사항
- [mise](https://mise.jdx.dev/) 설치
- Go 1.26

## 폴더 구조

- `cmd/`
  - `server/main.go` : 서버 시작점 (entrypoint)
- `client/`
  - `db/` : DB 연결/쿼리 클라이언트
  - `valkey/` : Valkey(또는 Redis 계열) 클라이언트/유틸
- `internal/`
  - `handler/` : HTTP 핸들러
  - `router/` : 라우팅 설정
- `pkg/`
  - `utils/` : 유틸리티(외부 모듈로 공유 가능한 코드)

## 빠른 시작

```bash
cd go-fiber-template
mise trust .
mise run           # 서버 실행
mise lint          # 정적 검사
mise format        # 코드 포맷팅
```

## 기본 엔드포인트

- `GET /` : "hello fiber3"
