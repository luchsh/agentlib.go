#!/usr/bin/env bash
#
# source env.sh before running tests or compilation

OS=`uname -s | tr '[[:upper:]]' '[[:lower:]]'`
if [[ "x${OS}" = "xdarwin" ]]; then
  export DYLD_LIBRARY_PATH=${JAVA_HOME}/lib/server:${DYLD_LIBRARY_PATH}
elif [[ "x${OS}" = "xlinux" ]]; then
  # detect JAVA_HOME
  if [[ "x${JAVA_HOME}" = "x" ]]; then
    which java > /dev/null 2>&1 ||  {
      echo "cannot find command java, aborting"
      exit 1
    }
    # Use a Java snippet to detect
cat > /tmp/EchoJavaHome.java <<EOF
public class EchoJavaHome {
  public static void main(String[] args) {
    System.out.println(System.getProperty("java.home"));
  }
}
EOF
    javac /tmp/EchoJavaHome.java
    export JAVA_HOME="$(java -cp /tmp EchoJavaHome)"
  fi
  export LD_LIBRARY_PATH=${JAVA_HOME}/lib/server:${DYLD_LIBRARY_PATH}
fi

if [[ "x${JAVA_HOME}" = "x" ]] || [[ ! -d "${JAVA_HOME}" ]]; then
  echo "Canot find a valid JAVA_HOME, aborting..."
  exit 128
fi

export PATH=${JAVA_HOME}/bin:${PATH}
if [[ -d "${JAVA_HOME}/jre" ]]; then
  export PATH=${JAVA_HOME}/jre/bin:${PATH}
  export CGO_LDFLAGS="-L${JAVA_HOME}/lib/amd64/server -ljvm"
else
  export CGO_LDFLAGS="-L${JAVA_HOME}/lib/server -ljvm"
fi

export CGO_CFLAGS="-I${JAVA_HOME}/include -I${JAVA_HOME}/include/${OS}"
