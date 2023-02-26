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
go run main.go
go build main.go
./main
```

## URL's
Utilizar as seguintes URL's para verificar o comportamento dependendo de onde o contexto transacional do banco é aberto / fechado:

http://localhost:8081/testTransactionalOutside/3000

http://localhost:8081/testTransactionalCorrect/3000

Memória inicial: 18572
Percentage of the requests served within a certain time (ms)
  50%      7
  66%      8
  75%     10
  80%     11
  90%     16
  95%     22
  98%     31
  99%     45
 100%    133 (longest request)
Memória: 25688
Percentage of the requests served within a certain time (ms)
  50%     13
  66%     20
  75%     27
  80%     32
  90%     54
  95%     88
  98%    151
  99%    215
 100%    552 (longest request)
Memória final: 29520
