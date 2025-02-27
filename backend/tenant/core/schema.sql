CREATE TABLE tenants (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    tier TEXT CHECK (tier IN ('free', 'paid')) NOT NULL DEFAULT 'free',
    db_usage BIGINT DEFAULT 0, -- Stores database size in bytes
    db_name TEXT,  -- Only for paid tenants
    db_host TEXT,  -- Only for paid tenants
    created_at TIMESTAMP DEFAULT now()
);

CREATE TABLE free_tenant_users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    tenant_id UUID REFERENCES tenants(id) ON DELETE CASCADE,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT now()
);