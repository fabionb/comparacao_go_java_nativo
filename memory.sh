#!/bin/bash
set -e

function cleanup {
  kill $pid
  gnuplot memory.tpl
}

rm memory.log 2> /dev/null || echo ''

case $1 in
  springboot)
    cd springboot_nativo
    ./gradlew bootJar
    java -jar build/libs/springboot_nativo.jar &
    ;;

  springboot_nativo)
    cd springboot_nativo
    ./gradlew bootBuildImage
    ./build/native/nativeCompile/springboot_nativo &
    ;;

  quarkus)
    cd quarkus_nativo
    ./gradlew build
    java -jar build/quarkus-app/quarkus-run.jar &
    ;;

  quarkus_nativo)
    cd quarkus_nativo
    ./gradlew build -Dquarkus.package.type=native
    ./build/quarkus_nativo-unspecified-runner &
    ;;

  go)
    cd go_gin_gorm
    go build main.go
    ./main &
    ;;

  *)
    echo "Passar ao script um dos valores: springboot, springboot_nativo, quarkus, quarkus_nativo ou go"
    exit -1
    ;;
esac

pid=$!
trap cleanup EXIT
cd ..

echo "Monitorando PID $pid"
while true; do
  ps -p $pid -o 'rss=' >> memory.log
  sleep 0.01
done
