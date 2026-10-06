# walletAPI
Repositório para o projeto wallet
┌─────────────────────────┐
│        HTTP / API       │
│                         │
│ Gin + Handlers          │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│       Application       │
│                         │
│ Casos de uso            │
│                         │
│ CreateWallet            │
│ Deposit                 │
│ Transfer                │
│ GetBalance              │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│         Domain          │
│                         │
│ Wallet                  │
│ Transaction             │
│ Regras de negócio       │
└────────────┬────────────┘
             │
             ▼
┌─────────────────────────┐
│      Infrastructure     │
│                         │
│ PostgreSQL              │
│ Redis                   │
│ etc.                    │
└─────────────────────────┘
