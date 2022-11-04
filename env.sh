#!/usr/bin/env bash
#
# source env.sh before running tests or compilation


if [[ "x${JAVA_HOME}" = "x" ]]; then
  echo "Please specify JAVA_HOME"
  return
fi

export CGO_CFLAGS="-I${JAVA_HOME}/include -I${JAVA_HOME}/include/darwin"
export CGO_LDFLAGS="-L${JAVA_HOME}/lib/server -ljvm"

OS=`uname -s`
if [[ "x${OS}" = "Darwin" ]]; then
  export DYLD_LIBRARY_PATH=${JAVA_HOME}/lib/server:${DYLD_LIBRARY_PATH}
elif [[ "x${OS}" = "Linux" ]]; then
  export LD_LIBRARY_PATH=${JAVA_HOME}/lib/server:${DYLD_LIBRARY_PATH}
fi
