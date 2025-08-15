flowchart LR
  subgraph Edge
    WAF[WAF / Rate Limit]
    APIGW[API Gateway]
    BFF_Admin[BFF Admin]
    BFF_App[BFF Pública]
  end

  subgraph Identity
    AS[Auth Server (OIDC/OAuth2.1)]
    IDP[(DB Identidades)]
    RT[(Redis: refresh/blacklist)]
  end

  subgraph Core
    USvc[User Service]
    RSvc[Role Service (RBAC)]
    PSvc[Policy Service (OPA)]
    AuditIn[Audit Ingestor]
    AIdx[(OpenSearch/Elastic)]
    AStore[(S3 u otro storage)]
  end

  subgraph Infra
    BUS[(Kafka/Redpanda)]
    CFG[Secrets/Config]
    OTEL[Logs/Metrics/Traces]
  end

  Client[Clientes / Admin]

  Client -->|HTTPS| WAF --> APIGW
  APIGW --> BFF_App
  APIGW --> BFF_Admin

  BFF_App -->|/oauth2/*| AS
  BFF_App -->|/users/*| USvc
  BFF_App -->|/authz/check| PSvc

  BFF_Admin -->|/roles/*| RSvc
  BFF_Admin -->|/users/*| USvc
  BFF_Admin -->|/policies/*| PSvc
  BFF_Admin -->|/audit/search| AIdx

  AS --> IDP
  AS --> RT
  AS --> BUS

  USvc --> BUS
  RSvc --> BUS
  PSvc --> BUS

  BUS --> AuditIn
  AuditIn --> AIdx
  AuditIn --> AStore

  APIGW --> OTEL
  AS --> OTEL
  USvc --> OTEL
  RSvc --> OTEL
  PSvc --> OTEL
  AuditIn --> OTEL

  CFG -.-> AS
  CFG -.-> USvc
  CFG -.-> RSvc
  CFG -.-> PSvc
  CFG -.-> AuditIn
