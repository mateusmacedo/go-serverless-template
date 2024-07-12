# Go Serverless Template

Este repositório é um template para criar aplicações serverless utilizando Go. Ele inclui configurações essenciais e estruturas para facilitar o desenvolvimento de projetos serverless com AWS Lambda.

## Estrutura do Projeto

- `cmd/`: Contém os comandos principais da aplicação.
- `docker/`: Configurações para containerização com Docker.
- `serverless.yml`: Configuração do Serverless Framework.
- `Makefile`: Comandos para build e deploy.
- `go.mod` e `go.sum`: Gerenciamento de dependências Go.
- `codecove.sh`: Script para integração contínua.

## Pré-requisitos

- Go 1.18+
- Node.js e npm
- Serverless Framework

## Instalação

Clone o repositório:
```sh
git clone https://github.com/mateusmacedo/go-serverless-template.git
cd go-serverless-template
```

Instale as dependências:
```sh
npm install -G serverless
npm install
```

## Uso

### Desenvolvimento Local

Para rodar a aplicação localmente:
```sh
make up
make deploy
```

### Deploy

Para fazer o deploy na AWS:
```sh
make deploy STAGE=dev
make deploy STAGE=prod
```

## Contribuição

Contribuições são bem-vindas! Por favor, abra um issue ou envie um pull request.

## Licença

Este projeto está licenciado sob a licença MIT.

## Autor

- [Mateus Macedo](https://github.com/mateusmacedo)