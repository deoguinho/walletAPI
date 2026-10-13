# Wallet API

API REST em Go para representar carteiras digitais. O projeto propõe um serviço simples para criar carteiras associadas a usuários e movimentar saldo por depósitos, saques e transferências entre carteiras.

As regras atuais rejeitam valores menores ou iguais a zero, saques acima do saldo e transferências para a mesma carteira. Transferências usam uma transação PostgreSQL e são registradas na tabela de transações.

> Os valores são inteiros, mas o projeto ainda não define moeda nem se o inteiro representa unidades ou centavos. Defina essa convenção antes de integrar clientes reais.

## Tecnologias

- Go e net/http
- PostgreSQL com pgx/v5
- Swagger
- Docker Compose para o banco local

## Estrutura

- internal/domain: entidades e regras de carteira e transação
- internal/application: casos de uso
- internal/handler: endpoints HTTP
- internal/infra/database: conexão e repositórios PostgreSQL
- migrations: scripts SQL
- docs: especificação Swagger

## Requisitos

- Go compatível com a versão indicada em go.mod
- Docker e Docker Compose
- Cliente PostgreSQL (psql) para aplicar as migrations

## Executar localmente

Execute os comandos a partir da pasta walletAPI.

1. Inicie o banco:

   ```sh
   docker compose up -d postgres
   ```

2. Configure DATABASE_URL. O Compose configura o usuário e banco locais como wallet, com senha wallet, na porta 5432.

   PowerShell:

   ```powershell
   $env:DATABASE_URL = "postgres://wallet:wallet@localhost:5432/wallet?sslmode=disable"
   ```

   Bash:

   ```sh
   export DATABASE_URL="postgres://wallet:wallet@localhost:5432/wallet?sslmode=disable"
   ```

   A aplicação também tenta carregar um arquivo .env da pasta atual. Não versione credenciais reais.

3. Aplique as migrations na ordem:

   ```sh
   psql "$DATABASE_URL" -f migrations/001_create_wallets.up.sql
   psql "$DATABASE_URL" -f migrations/002_create_transactions.up.sql
   ```

   No PowerShell, use psql $env:DATABASE_URL -f migrations/001_create_wallets.up.sql e repita para a migration 002.

4. Inicie a API:

   ```sh
   go run .
   ```

O servidor escuta na porta 8080. O schema precisa ser criado antes da inicialização; a aplicação não executa migrations automaticamente.

## Endpoints

| Método | Caminho | Descrição |
| --- | --- | --- |
| POST | /wallets | Cria carteira com saldo inicial zero |
| POST | /wallets/{id}/deposit | Deposita valor na carteira |
| POST | /wallets/{id}/withdraw | Saca valor da carteira |
| POST | /wallets/transfer | Transfere entre carteiras |

### Criar carteira

```http
POST /wallets
Content-Type: application/json

{"user_id": 123}
```

### Depositar ou sacar

```http
POST /wallets/1/deposit
Content-Type: application/json

{"amount": 500}
```

Para sacar, use /wallets/1/withdraw com o mesmo corpo.

### Transferir

```http
POST /wallets/transfer
Content-Type: application/json

{"from_wallet_id": 1, "to_wallet_id": 2, "amount": 100}
```

## Swagger

Com a API em execução, acesse http://localhost:8080/swagger/index.html.

## Testes

Para executar os testes existentes:

```sh
go test ./...
```

Os testes atuais cobrem regras de domínio; ainda não há testes de integração dos handlers e repositórios PostgreSQL.

## Escopo e próximos passos

Este projeto é uma base para estudar operações de carteira. Antes de expô-lo em produção, faltam autenticação e autorização, consulta de saldo e histórico, idempotência para evitar operações duplicadas, controle de concorrência para depósitos e saques, constraints e chaves estrangeiras adequadas, timeouts HTTP e testes de integração.

As migrations atuais exigem as duas carteiras em cada registro de transação. Isso precisa ser revisto para registrar depósitos e saques no histórico.
