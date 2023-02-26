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
./gradlew quarkusDev
./gradlew build -Dquarkus.package.type=native
./build/quarkus_nativo-unspecified-runner
```

## URL's
Utilizar as seguintes URL's para verificar o comportamento dependendo de onde o contexto transacional do banco é aberto / fechado:

http://localhost:8081/testTransactionalOutside/3000

http://localhost:8081/testTransactionalCorrect/3000

Memória inicial: 73212
Percentage of the requests served within a certain time (ms)
  50%      9
  66%     11
  75%     12
  80%     13
  90%     16
  95%     18
  98%     22
  99%     24
 100%     35 (longest request)
Memória: 138092
Percentage of the requests served within a certain time (ms)
  50%     36
  66%     39
  75%     41
  80%     43
  90%     52
  95%     71
  98%     83
  99%     97
 100%    190 (longest request)
Memória final: 299200