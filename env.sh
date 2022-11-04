#!/usr/bin/env bash
#


if [[ "x${JAVA_HOME}" = "x" ]]; then
  echo "Please specify JAVA_HOME"
  return
fi

export CGO_CFLAGS="-I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin"
export CGO_LDFLAGS="-L${JAVA_HOME}/lib -L${JAVA_HOME}/lib/server -ljvm"

