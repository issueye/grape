export CGO_ENABLED=0
export GOARCH=amd64
# export GOARCH=arm64
export GOOS=linux

go mod tidy

go mod vendor

TAG="Beta"
BRANCH=$(git symbolic-ref --short -q HEAD)
COMMIT=$(git rev-parse --verify HEAD)
NOW=$(date '+%FT%T%z')

VERSION="v0.1.2-${TAG}"
APPNAME="grape-${VERSION}"
DESCRIPTION="版本管理服务"

go build -o bin/${APPNAME} -tags=ui -ldflags "-X demo/build.AppName=Demo \
-X github.com/issueye/grape/internal/initialize.Branch=${BRANCH} \
-X github.com/issueye/grape/internal/initialize.Commit=${COMMIT} \
-X github.com/issueye/grape/internal/initialize.Date=${NOW} \
-X github.com/issueye/grape/internal/initialize.AppName=${DESCRIPTION} \
-X github.com/issueye/grape/internal/initialize.Version=${VERSION}" .