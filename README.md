
# Orders System — Clean Architecture

Implementação de um sistema de listagem de Orders utilizando **Clean Architecture**, com suporte simultâneo a três interfaces de comunicação: REST, gRPC e GraphQL.

O principal objetivo do projeto é demonstrar como um único caso de uso (*Use Case*) pode ser disponibilizado por diferentes protocolos de comunicação, mantendo as regras de negócio independentes dos mecanismos de transporte.

## Arquitetura

A aplicação utiliza Clean Architecture para separar as responsabilidades e garantir baixo acoplamento entre as camadas.

O caso de uso `ListOrders` é compartilhado pelas três interfaces:

- **REST API:** exposição dos dados por meio de endpoints HTTP.
- **gRPC:** comunicação utilizando Protocol Buffers.
- **GraphQL:** consulta flexível dos dados por meio de queries.

Essa abordagem permite adicionar ou modificar interfaces de comunicação sem duplicar as regras de negócio.

## Executando o projeto

### Pré-requisitos

- Docker
- Docker Compose (plugin `docker compose`)

Não é necessário instalar Go, MySQL ou RabbitMQ localmente. Toda a infraestrutura é provisionada automaticamente pelo Docker Compose.

### Configuração de variáveis de ambiente

Antes de iniciar a aplicação, crie o arquivo `.env` a partir do arquivo de exemplo:

```bash
cp cmd/ordersystem/.env.example cmd/ordersystem/.env
```

Revise o arquivo `cmd/ordersystem/.env` e ajuste as variáveis de ambiente conforme necessário.

O arquivo de exemplo contém as configurações necessárias para executar a aplicação localmente utilizando Docker Compose.

> **Nota:** Não versione o arquivo `.env` com credenciais ou informações sensíveis. Utilize o `.env.example` como referência para compartilhar as configurações necessárias.

### Inicialização

Na raiz do projeto, execute:

```bash
docker compose up -d
```

O Docker Compose será responsável por:

1. Construir a imagem da aplicação Go.
2. Inicializar o banco de dados MySQL.
3. Inicializar o RabbitMQ.
4. Executar as migrations do banco de dados.
5. Iniciar a aplicação após a disponibilidade dos serviços necessários.

Para acompanhar os logs:

```bash
docker compose logs -f
```

Para encerrar os serviços:

```bash
docker compose down
```

**Observação:** os dados do MySQL são persistidos em volume local, portanto, reiniciar os containers não implica perder os dados armazenados.

---

## Interfaces de comunicação

Após a inicialização, as três interfaces estarão disponíveis simultaneamente.


### 0. Criando Orders

O projeto disponibiliza um script para realizar a carga inicial de Orders automaticamente, utilizando a REST API.

#### Executando a carga de dados

Com a aplicação em execução, execute o script na raiz do projeto:

```bash
bash scripts/seed_orders.sh
```

O script realizará requisições `POST` para o endpoint `/order`, cadastrando cinco Orders de exemplo.

**Endpoint utilizado:**

```http
POST http://localhost:8000/order
Content-Type: application/json
```

#### Criando uma Order manualmente

Caso deseje cadastrar uma Order individualmente, utilize a requisição disponibilizada no arquivo [`api/list_order.http`](./api/list_order.http).

Abra o arquivo no VS Code e execute a requisição `POST /order` utilizando a extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client).

**Exemplo de payload:**

```json
{
  "id": "0002",
  "price": 100.5,
  "tax": 0.5
}
```

Também é possível realizar o cadastro utilizando cURL:

```bash
curl -X POST http://localhost:8000/order \
  -H "Content-Type: application/json" \
  -d '{
    "id": "0002",
    "price": 100.5,
    "tax": 0.5
  }'
```


### 1. REST API

A listagem de Orders pode ser acessada pelo endpoint:

```http
GET http://localhost:8000/order
```

**Utilizando cURL:**

```bash
curl http://localhost:8000/order
```

**Utilizando VS Code:**

O projeto também disponibiliza uma requisição HTTP pronta em [`api/list_order.http`](./api/list_order.http).

Para executá-la, utilize a extensão [REST Client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client) do VS Code.


### 2. gRPC

A aplicação disponibiliza um servidor gRPC na porta `50051`.

Para realizar chamadas, é possível utilizar o [Evans](https://github.com/ktr0731/evans), um cliente interativo para serviços gRPC.

Execute o comando abaixo na raiz do projeto:

```bash
docker run --rm -it \
  --network ordersystem-network \
  -v "$(pwd):/mount:ro" \
  ghcr.io/ktr0731/evans:latest \
  --path /mount/internal/infra/grpc/protofiles \
  --proto order.proto \
  --host app \
  --port 50051 \
  repl
```

O comando inicia um terminal interativo do Evans, utilizando a rede interna do Docker Compose para acessar o servidor gRPC.

**Executando a chamada:**

Dentro do terminal do Evans, execute:

```text
call ListOrders
```

O cliente apresentará os parâmetros necessários, caso existam, e retornará a resposta do serviço.

### 3. GraphQL

O servidor GraphQL está disponível em:

[http://localhost:8080](http://localhost:8080)

Acesse o endereço pelo navegador e execute a seguinte query:

```graphql
query List {
  ListOrders {
    id
    Price
    Tax
    FinalPrice
  }
}
```

A query retorna as Orders com seus respectivos identificadores, preços, impostos e valores finais.

---

## Serviços e portas

| Serviço | Porta | Descrição |
|---|---:|---|
| REST API | 8000 | Endpoint HTTP para listagem de Orders |
| gRPC | 50051 | Servidor de comunicação gRPC |
| GraphQL | 8080 | Endpoint para execução de queries |
| MySQL | 3306 | Banco de dados relacional |
| RabbitMQ | 5672 | Message Broker |
| RabbitMQ Management | 15672 | Interface administrativa do RabbitMQ |

## Tecnologias utilizadas

- **Go:** linguagem de programação da aplicação.
- **Clean Architecture:** organização e separação das responsabilidades.
- **MySQL:** persistência de dados.
- **RabbitMQ:** mensageria.
- **REST:** interface HTTP.
- **gRPC / Protocol Buffers:** comunicação RPC.
- **GraphQL:** consultas flexíveis de dados.
- **Docker / Docker Compose:** containerização e orquestração dos serviços.

---

## Considerações finais

O projeto demonstra a aplicação de Clean Architecture em um cenário com múltiplas interfaces de comunicação, permitindo que diferentes mecanismos de transporte compartilhem o mesmo caso de uso e as mesmas regras de negócio.

Toda a infraestrutura necessária para avaliação é provisionada pelo Docker Compose, simplificando a execução e permitindo testar as três interfaces sem configurações adicionais no ambiente local.
