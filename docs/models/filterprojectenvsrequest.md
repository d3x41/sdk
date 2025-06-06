# FilterProjectEnvsRequest

## Example Usage

```typescript
import { FilterProjectEnvsRequest } from "@vercel/sdk/models/filterprojectenvsop.js";

let value: FilterProjectEnvsRequest = {
  idOrName: "prj_XLKmu1DyR1eY7zq8UgeRKbA7yVLA",
  gitBranch: "feature-1",
  decrypt: "true",
  source: "vercel-cli:pull",
  customEnvironmentId: "env_123abc4567",
  customEnvironmentSlug: "my-custom-env",
  teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
  slug: "my-team-url-slug",
};
```

## Fields

| Field                                                                                                   | Type                                                                                                    | Required                                                                                                | Description                                                                                             | Example                                                                                                 |
| ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------- |
| `idOrName`                                                                                              | *string*                                                                                                | :heavy_check_mark:                                                                                      | The unique project identifier or the project name                                                       | prj_XLKmu1DyR1eY7zq8UgeRKbA7yVLA                                                                        |
| `gitBranch`                                                                                             | *string*                                                                                                | :heavy_minus_sign:                                                                                      | If defined, the git branch of the environment variable to filter the results (must have target=preview) | feature-1                                                                                               |
| `decrypt`                                                                                               | [models.Decrypt](../models/decrypt.md)                                                                  | :heavy_minus_sign:                                                                                      | If true, the environment variable value will be decrypted                                               | true                                                                                                    |
| `source`                                                                                                | *string*                                                                                                | :heavy_minus_sign:                                                                                      | The source that is calling the endpoint.                                                                | vercel-cli:pull                                                                                         |
| `customEnvironmentId`                                                                                   | *string*                                                                                                | :heavy_minus_sign:                                                                                      | The unique custom environment identifier within the project                                             | env_123abc4567                                                                                          |
| `customEnvironmentSlug`                                                                                 | *string*                                                                                                | :heavy_minus_sign:                                                                                      | The custom environment slug (name) within the project                                                   | my-custom-env                                                                                           |
| `teamId`                                                                                                | *string*                                                                                                | :heavy_minus_sign:                                                                                      | The Team identifier to perform the request on behalf of.                                                | team_1a2b3c4d5e6f7g8h9i0j1k2l                                                                           |
| `slug`                                                                                                  | *string*                                                                                                | :heavy_minus_sign:                                                                                      | The Team slug to perform the request on behalf of.                                                      | my-team-url-slug                                                                                        |