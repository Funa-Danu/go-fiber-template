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
    - `mocks/` : `mockgen`으로 생성한 테스트 Mock
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
mise lint             # 정적 검사
mise format           # 코드 포맷팅
mise unit-test        # 유닛 테스트 (integration 제외)
mise integration-test  # 통합 테스트 (Testcontainers 실행)
mise mockgen           # valkey 패키지 mock 생성
mise test-all          # unit + integration 한 번에 실행
```

## Valkey 클라이언트

`client/valkey`는 다음을 제공합니다.

- `LoadConfigFromEnv()`
  - `VALKEY_ADDR`, `VALKEY_USERNAME`, `VALKEY_PASSWORD`, `VALKEY_DB` 읽기
- `New(ctx, cfg)`
  - 연결 유효성 검사(`PING`)까지 수행

## Mockgen 사용 예시

- `client/valkey/interface.go`의 `ValkeyClient`를 기준으로 Mock 파일을 생성/재생성합니다.

```bash
mise mockgen
```

Go 파일 내부에는 다음의 `go:generate` 지시어가 있어요.

```go
//go:generate mockgen -source=interface.go -destination=mocks/mock_valkey.go -package=mockclient
```

## 기본 엔드포인트

- `GET /` : "hello fiber3"
- `GET /v1/health` : JSON 헬스체크 응답

### 테스트 주의

- `mise integration-test`는 Docker가 있는 환경에서만 실행 가능합니다.
