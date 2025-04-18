# GetAliasDeployment

A map with the deployment ID, URL and metadata

## Example Usage

```typescript
import { GetAliasDeployment } from "@vercel/sdk/models/getaliasop.js";

let value: GetAliasDeployment = {
  id: "dpl_5m8CQaRBm3FnWRW1od3wKTpaECPx",
  url: "my-instant-deployment-3ij3cxz9qr.now.sh",
  meta: "{}",
};
```

## Fields

| Field                                   | Type                                    | Required                                | Description                             | Example                                 |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `id`                                    | *string*                                | :heavy_check_mark:                      | The deployment unique identifier        | dpl_5m8CQaRBm3FnWRW1od3wKTpaECPx        |
| `url`                                   | *string*                                | :heavy_check_mark:                      | The deployment unique URL               | my-instant-deployment-3ij3cxz9qr.now.sh |
| `meta`                                  | *string*                                | :heavy_minus_sign:                      | The deployment metadata                 | {}                                      |