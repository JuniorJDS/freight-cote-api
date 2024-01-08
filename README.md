# freight cote- API

<p align="center">
  <img src="https://img.shields.io/badge/Golang-v1.21.5-blue"/>
  <img src="https://github.com/JuniorJDS/freight-cote-api/actions/workflows/integration-tests.yml/badge.svg">
  <img src="https://github.com/JuniorJDS/freight-cote-api/actions/workflows/golangci-lint.yml/badge.svg">
</p>

API responsável por fazer a integração entre um E-commerce fictício e um API de cálculo de Frete. Sua principal finalidade é permitir que o E-commerce obtenha informações precisas sobre os custos de transporte associados aos produtos vendidos, possibilitando uma experiência de compra mais transparente para os clientes.

Código escrito na linguagem <a href="https://go.dev/" target="_blank">Golang</a> na versão 1.21.5, mais detalhes são descritos abaixo:

## Como Executar:

É possível executar a API, tal como seus testes, através do Docker ou no ambiente local, também foi criado um arquivo `Make` para facilitar algumas execuções. Além disso, para as execuções fora do Docker, atente-se para o arquivo `.env` que deve estar na pasta raiz do projeto. Caso deseje debugar os testes, também deve ter um arquivo `.env` na pasta `/tests/integration`.

### Comandos Make:

- `make run-local-environment`: Inicia a API através do Docker, juntamente com os outros serviços necessários;
- `make run`: Inicia, unicamente, a API;
- `make integration-tests`: Inicia um ambiente no Docker e roda todos os testes de integração;
- `make mongo-local`: Inicia uma instância do mongo-db através do Docker, muito útil para o debug do código ou dos testes;

### Docker:

Para rodar utilizando o Docker, basta subir o arquivo `docker-compose-run-local.yaml` com o docker-compose:

```
docker-compose -f docker-compose-run-local.yaml build
docker-compose -f docker-compose-run-local.yaml up
```

ou utilizando o comando `make`:

```
make run-local-environment
```

### Ambiente Local:

Para executar a API em um abiente local é necessário um MongoDB, além de instalar os pacotes necessários. Sendo assim, os comandos necessários para instalação e execução são:

```
go get -v
go run main.go
```

## Swagger

A documentação do Swagger se encontra em `/docs`. A partir dela é possível ver as rotas geradas e uma curta explicação.

Ao atualizar _qualquer_ componente da documentação, deve-se executar o comando:

```
swag init --parseDependency --parseInternal
```

\*Os argumentos adicionais resolvem tipos externos, como `fiber.Map`.

Caso deseje saber mais sobre as anotações disponíveis, acesse [este link](https://github.com/swaggo/swag#declarative-comments-format).
