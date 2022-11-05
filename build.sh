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

source env.sh

detect_java_home
echo "Using JAVA_HOME=${JAVA_HOME}"

# do build in child process without contaminating current process' state
(
  # add platform specific flags
  case ${OS} in
    darwin)
    DL_POSTFIX="dylib"
    ;;
    linux)
    DL_POSTFIX="so"
    ;;
    *)
    echo "Unkown OS, exiting"
    exit 1
    ;;
  esac

  go build -v -buildmode=c-shared -o libjnigo.${DL_POSTFIX}

  if [[ $? -eq 0 ]]; then
    echo "Build successful"
    exit 0
  else
    exit 128
  fi
)
