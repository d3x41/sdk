# UpdateProjectProjectsDeploymentType

Specify if the Vercel Authentication (SSO Protection) will apply to every Deployment Target or just Preview

## Example Usage

```typescript
import { UpdateProjectProjectsDeploymentType } from "@vercel/sdk/models/updateprojectop.js";

let value: UpdateProjectProjectsDeploymentType =
  "prod_deployment_urls_and_all_previews";
```

## Values

```typescript
"all" | "preview" | "prod_deployment_urls_and_all_previews" | "all_except_custom_domains"
```