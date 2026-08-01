interface Database {
  host: string;
  version: string;
  postgresEngine: string;  
  releaseChannel: string;
}

export interface Project {
  id: string;
  ref: string;
  organizationId: string;
  organizationSlug: string;
  name: string;
  region: string;
  status: string;
  database: Database;
  createdAt: string;
}