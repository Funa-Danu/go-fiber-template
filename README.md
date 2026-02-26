# go-fiber-template

Go 1.26 + Fiber v3 프로젝트 템플릿입니다.

## 요구사항
- [mise](https://mise.jdx.dev/) 설치
- Go 1.26

## 빠른 시작

```bash
cd go-fiber-template
mise install        # go 1.26 설치
mise run           # 서버 실행
mise lint          # 정적 검사
mise format        # 코드 포맷팅
```

## 기본 구조

- `main.go`: 기본 Fiber v3 서버 예시
- `go.mod`: Go 모듈 정의 + Fiber v3 의존성
- `mise.toml`: mise tool/task 구성
- `.gitignore`: 빌드 산출물 무시
