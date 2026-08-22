interface Database {
  host: string;
  version: string;
  postgresEngine: string;  
  releaseChannel: string;
}

export interface Project {
  id: string;
  ref: string;
  organization_id: string;
  organizationSlug: string;
  name: string;
  region: string;
  status: string;
  database: Database;
  createdAt: string;
}


export interface OrgDetail {
  id : string 
  Name : string 
  plan : string
  opt_in_tags : string[]
  allowed_release_channels : string[]
}

