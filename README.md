# Locadora de Carros

API de uma locadora de Carros e portal para visualização do status usando Clean Architecture

## Compilação
```shell
  make
```

## Executar testes
```shell
  make test
```


## Usando a API

### Cadastrar Carro

```shell
curl --location --request POST 'http://localhost:3001/v1/carro' \
--header 'Authorization: Basic dXN1YXJpbzpzZW5oYQ==' \
--header 'Content-Type: application/json' \
--data-raw '{
    "placa": "QWE9876",
    "modelo": "Celta",
    "cor": "Branco",
    "renavan": 45678,
    "hodometro": 0
}'
```

### Listar Carros

```shell
curl --location --request GET 'http://localhost:3001/v1/carro' \
--header 'Authorization: Basic dXN1YXJpbzpzZW5oYQ=='
```

### Buscar por Placa

```shell
curl --location --request GET 'http://localhost:3001/v1/carro?query=ABC' \
--header 'Authorization: Basic dXN1YXJpbzpzZW5oYQ=='
```

### Buscar por ID 

```shell
curl --location --request GET 'http://localhost:3001/v1/carro/5167f297-393a-438f-9eb4-831bd521c0e2' \
--header 'Authorization: Basic dXN1YXJpbzpzZW5oYQ=='
```

### Excluir Carro

```shell
curl --location --request DELETE 'http://localhost:3001/v1/carro/5167f297-393a-438f-9eb4-831bd521c0e2' \
--header 'Authorization: Basic dXN1YXJpbzpzZW5oYQ=='
```

## Interface de Visualização

Acessar http://localhost:3001/login com credenciais "usuario" e "senha"