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
  - `server/`
    - `route/` : 라우팅 정의
      - `v1.go` : 그룹 라우팅 진입점 (`/`, `/v1/...`)
      - `v1/` : `/v1` 하위 라우트
    - `handler/` : 핸들러(현재는 라우트 파일 내에서 확장 가능)
  - 향후 도메인/기능별로 세분화하여 관리
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
- `GET /v1/health` : JSON 헬스체크 응답
