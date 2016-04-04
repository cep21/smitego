#!/bin/bash
set -ex

CIRCLEUTIL_TAG="v1.39"
DEFAULT_GOLANG_VERSION="1.6"
GO_TESTED_VERSIONS="1.4.3 1.5.1 1.6"
IMPORT_PATH="github.com/$CIRCLE_PROJECT_USERNAME/$CIRCLE_PROJECT_REPONAME"

export GOROOT="$HOME/go_circle"
export GOPATH="$HOME/.go_circle"
export GOPATH_INTO="$HOME/lints"
PATH="$GOROOT/bin:$GOPATH/bin:$GOPATH_INTO:$PATH"

GO_COMPILER_PATH="$HOME/gover"
SRC_PATH="$GOPATH/src/$IMPORT_PATH"

# Cache phase of circleci
function do_cache() {
  [ ! -d "$HOME/circleutil" ] && git clone https://github.com/signalfx/circleutil.git "$HOME/circleutil"
  (
    cd "$HOME/circleutil"
    git fetch -a -v
    git fetch --tags
    git reset --hard $CIRCLEUTIL_TAG
  )
  . "$HOME/circleutil/scripts/common.sh"
  mkdir -p "$GO_COMPILER_PATH"
  install_all_go_versions "$GO_COMPILER_PATH"
  install_go_version "$GO_COMPILER_PATH" "$DEFAULT_GOLANG_VERSION"
  mkdir -p "$GOPATH_INTO"
  install_circletasker "$GOPATH_INTO"
  versioned_goget "github.com/cep21/gobuild:v1.5"
  copy_local_to_path "$SRC_PATH"
}


# Test phase of circleci
function do_test() {
  . "$HOME/circleutil/scripts/common.sh"
  # The go get install only works on go 1.5+ (no -f parameter)
  install_go_version "$GO_COMPILER_PATH" "$DEFAULT_GOLANG_VERSION"
  gobuild
  for GO_VERSION in $GO_TESTED_VERSIONS; do
  install_go_version "$GO_COMPILER_PATH" "$GO_VERSION"
    go test -race -timeout 3s ./...
    go test -timeout 15s -race ./...
  done
}

# Deploy phase of circleci
function do_deploy() {
  echo "No deploy phase for library"
}

function do_all() {
  do_cache
  do_test
  do_deploy
}

case "$1" in
  cache)
    do_cache
    ;;
  test)
    do_test
    ;;
  deploy)
    do_deploy
    ;;
  all)
    do_all
    ;;
  *)
  echo "Usage: $0 {cache|test|deploy|all}"
    exit 1
    ;;
esac
