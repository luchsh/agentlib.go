#!/usr/bin/env bash

#
# Copyright 2020 chuanshenglu@gmail.com
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

# Detect and set environment variable JAVA_HOME
function detect_java_home {
  if [[ "x${JAVA_HOME}" != "x" ]]; then return; fi
  JAVA=`which java 2>&1`
  if [[ $? -ne 0 ]] || [[ "x${JAVA}" = "x" ]]; then
    echo "Cannot find command 'java', pls install"
    exit 1
  fi

  # detect based on position of java executables
  JAVA=`readlink ${JAVA}`
  DIR=`dirname ${JAVA}`
  if [[ -f "${DIR}/../include/jni.h" ]]; then
    export JAVA_HOME=${DIR}
  fi

  # platform specific localtions
  if [[ "$(uname -s)" = "Darwin" ]]; then
    SYS_JAVA_DIR="/Library/Java/JavaVirtualMachines/"
    for D in `ls ${SYS_JAVA_DIR}`; do
      DIR="${SYS_JAVA_DIR}/${D}/Contents/Home"
      if [[ -f "${DIR}/bin/java" ]] && [[ -f "${DIR}/include/jni.h" ]]; then
        export JAVA_HOME=${DIR}
      fi
    done
  fi
}

detect_java_home
echo "Using JAVA_HOME=${JAVA_HOME}"

# do build in child process without contaminating current process' state
(
  export PATH=${JAVA_HOME}/bin:${PATH}
  export CGO_CFLAGS="-I ${JAVA_HOME}/include"
  # add platform specific flags
  case `uname -s` in
    Darwin)
    export CGO_CFLAGS="${CGO_CFLAGS} -I ${JAVA_HOME}/include/darwin"
    DL_POSTFIX="dylib"
    ;;
    Linux)
    DL_POSTFIX="so"
    ;;
    *)
    echo "Unkown OS, exiting"
    exit 1
    ;;
  esac

  cd src
  go build -v -buildmode=c-shared -o libjnigo.${DL_POSTFIX}


  if [[ $? -eq 0 ]]; then
    echo "Build successful"
    exit 0
  else
    exit 128
  fi
)
