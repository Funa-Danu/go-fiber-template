# go-fiber-template

Go 1.26 + Fiber v3 템플릿 프로젝트입니다.

## 요구사항
- [mise](https://mise.jdx.dev/) 설치
- Go 1.26

## 폴더 구조

- `cmd/`
  - `server/main.go` : 서버 시작점 (entrypoint)
- `client/`
  - `valkey/` : Valkey(또는 Redis 계열) 클라이언트/유틸
    - `command_*` 테스트 + 통합 테스트 예시 제공
  - `pgx/` : PostgreSQL(PGX) 클라이언트/유틸
- `internal/`
  - `server/`
    - `route/` : 라우팅 정의
      - `v1.go` : 라우팅 그룹 엔트리 (`/`, `/v1/...`)
      - `v1/` : `/v1` 하위 라우트
    - `handler/` : 라우트 핸들러
- `pkg/`
  - `utils/` : 유틸리티(외부 모듈로 공유 가능한 코드)

## 빠른 시작

```bash
cd go-fiber-template
mise trust .

mise run-server       # 서버 실행
mise build-server     # 서버 바이너리 빌드 (출력: ./bin/server)
mise generate         # go:generate 실행 (mock/stub 생성 등)
mise lint             # 정적 검사
mise format           # 코드 포맷팅
mise unit-test        # 유닛 테스트 (integration 제외)
mise integration-test  # 통합 테스트 (현재 통합 대상은 valkey 테스트를 기본으로 구성)
mise test-all          # unit + integration 한 번에 실행
```

## 환경 변수 주입

실행 전 아래 환경변수값이 주입됩니다. 기본값은 `.env.example` 참조

- `SERVICE_NAME` (default: `go-fiber-template`)
- `ENV` (default: `local`)
- `PORT` (default: `3000`)
- `VALKEY_ADDR` (default: `localhost:6379`)
- `VALKEY_DB` (default: `0`)
- `PG_HOST` (default: `localhost`)
- `PG_PORT` (default: `5432`)
- `PG_USER` (default: `postgres`)
- `PG_PASSWORD` (default: `postgres`)
- `PG_DATABASE` (default: `postgres`)

`cmd/server/main.go`는 이 값을 `internal/config`에서 로드해서 라우팅/리스닝 포트 및 의존성 클라이언트를 구성합니다.

## 환경 변수 예시

로컬 실행 시 `.env.example` 기준으로 값을 채워서 사용할 수 있습니다.

```bash
cp .env.example .env
```

## 클라이언트 실행 환경

- Valkey

```bash
docker compose -f client-docker-compose.yaml up -d
```

```bash
# 중지
docker compose -f client-docker-compose.yaml down
```

- PostgreSQL(PGX)

```bash
docker compose -f client-pgx-docker-compose.yaml up -d
```

```bash
# 중지
docker compose -f client-pgx-docker-compose.yaml down
```

## 테스트 구조

- `client/valkey`는 실제 서비스 연동 코드와 테스트 템플릿을 제공합니다.
- `client/pgx`는 PostgreSQL 연동 코드와 테스트 템플릿(테스트컨테이너 연동 포함)을 제공합니다.
- 단위 테스트: `*_test.go`
- 통합 테스트: `*_integration_test.go` + `//go:build integration`
- 통합 테스트는 `TestMain`에서 Testcontainers 기반으로 컨테이너를 띄운 뒤 종료 처리

## 기본 엔드포인트

- `GET /` : "hello fiber3"
- `GET /v1/health` : JSON 헬스체크 응답

## 테스트 주의

- `mise integration-test`는 Docker/컨테이너 환경이 필요할 수 있습니다.
