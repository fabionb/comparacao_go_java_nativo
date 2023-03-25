# Comparação de memória em aplicações go / spring boot (e nativo) / quarkus (e nativo)

Este repositório faz o comparativo de uma aplicação simples que apenas persiste uma entidade no banco de dados para verificar o uso de memória.

## Requisitos
Rodar um MySQL com o seguinte comando:
```
docker run -d --name mysql --network host -e MYSQL_ROOT_PASSWORD=root -d mysql:8
```

## Execução
Para rodar o código, executar o seguinte comando:
```
./memory.sh
```

O script acima precisa receber um parâmetro que é qual o tipo de aplicação ele irá rodar, podendo ser um dos seguintes:
```
springboot
springboot_nativo
quarkus
quarkus_nativo
go
```

Enquanto o script executa, você pode simular uma carga com a ferramenta *ab* da seguinte maneira:

```
ab -c 30 -n 3000 http://localhost:8081/
```

Após parar o script memory.sh, será gerada a imagem *memory.png* com as informações de uso de memória com o passar do tempo.
