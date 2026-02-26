# go-fiber-template

Go 1.26 + Fiber v3 템플릿 프로젝트입니다.

## 요구사항
- [mise](https://mise.jdx.dev/) 설치
- Go 1.26

## 폴더 구조

- `cmd/`
  - `server/main.go` : 서버 시작점 (entrypoint)
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
mise integration-test  # 통합 테스트 (현재 테스트 범위는 내부 라우팅/서버 위주)
mise test-all          # unit + integration 한 번에 실행
```

## 기본 엔드포인트

- `GET /` : "hello fiber3"
- `GET /v1/health` : JSON 헬스체크 응답

## 테스트 주의

- `mise integration-test`는 현재 환경의 통합 대상에 따라 Docker 등이 필요한 테스트는 생략될 수 있습니다.
