interface Database {
  host: string;
  version: string;
  postgresEngine?: string;
  postgres_engine?: string;
  releaseChannel?: string;
  release_channel?: string;
}

export interface Project {
  id: string;
  ref: string;
  organization_id: string;
  organizationSlug?: string;
  organization_slug?: string;
  name: string;
  region: string;
  status: string;
  database: Database;
  createdAt?: string;
  created_at?: string;
}

export interface OrgDetail {
  id: string;
  name: string;
  plan: string;
  opt_in_tags: string[];
  allowed_release_channels: string[];
}

export interface CPUUsage {
  load_avg_15m: number;
  usage_percent?: number;
  per_core_seconds?: Record<string, number>;
}

export interface MemoryUsage {
  swap_total_bytes: number;
  page_tables_bytes: number;
  slab_bytes: number;
  committed_as_bytes: number;
  dirty_bytes: number;
  shmem_bytes: number;
}

export interface DiskUsage {
  device: string;
  io_time_weighted_seconds: number;
  read_time_seconds: number;
  discard_time_seconds?: number;
  filesystem_type?: string;
}

export interface PgBouncerInfo {
  version: string;
  max_client_connections: number;
  server_active_connections: number;
  server_login_connections: number;
  cached_dns_names: number;
}

export interface SystemUsage {
  project_ref: string;
  cpu: CPUUsage;
  memory: MemoryUsage;
  disk: DiskUsage[];
  pgbouncer: PgBouncerInfo;
  timestamp: number;
}

export interface AnalyticsCount {
  timestamp: string;
  total_auth_requests: number;
  total_realtime_requests: number;
  total_rest_requests: number;
  total_storage_requests: number;
}

export interface LogMetadataRequest {
  method?: string;
  path?: string;
  host?: string;
  cf?: Array<{
    city?: string;
    country?: string;
    asOrganization?: string;
  }>;
}

export interface LogMetadataResponse {
  status_code?: number;
  origin_time?: number;
  sb_gateway_version?: string;
}

export interface LogItem {
  id: string;
  event_message: string;
  timestamp: number;
  metadata?: Array<{
    request?: LogMetadataRequest[];
    response?: LogMetadataResponse[];
  }>;
}

export interface StatusHistoryDay {
  date: string;
  formattedDate: string;
  status: 'operational' | 'degraded' | 'outage';
  uptimePercentage: number;
  incidents: string[];
}
