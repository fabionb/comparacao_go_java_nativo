# Chamadas remotas em contexto transacional de banco de dados
Este repositório foi utilizado para a escrita do artigo sobre como chamadas remotas em contexto transacional de banco de dados impactam o tempo de resposta / escalabilidade da aplicação.

O artigo pode ser encontrado na seguinte URL: https://www.fnbrandao.com.br/blog/2021/10/chamadas-remotas-em-contexto-transacional-de-banco-de-dados/

## Requisitos
Rodar um MySQL com o seguinte comando:
```
docker run -d --name mysql --network host -e MYSQL_ROOT_PASSWORD=root -d mysql:8
```

## Execução
Para rodar o código, executar o seguinte comando:
```
./gradlew bootRun

export JAVA_HOME=/usr/lib/jvm/java-17-openjdk/
export GRAALVM_HOME=/usr/lib/jvm/java-17-graalvm/
./gradlew bootBuildImage
./build/native/nativeCompile/springboot_nativo
```

## URL's
Utilizar as seguintes URL's para verificar o comportamento dependendo de onde o contexto transacional do banco é aberto / fechado:

http://localhost:8081/testTransactionalOutside/3000

http://localhost:8081/testTransactionalCorrect/3000

Memória inicial: 165388
Percentage of the requests served within a certain time (ms)
  50%      9
  66%     10
  75%     12
  80%     13
  90%     16
  95%     19
  98%     27
  99%     28
 100%     38 (longest request)
Memória: 204864
Percentage of the requests served within a certain time (ms)
  50%     28
  66%     30
  75%     33
  80%     38
  90%     53
  95%     60
  98%     78
  99%     85
 100%    169 (longest request)
Memória final: 331288