# GetProjectsErl

## Example Usage

```typescript
import { GetProjectsErl } from "@vercel/sdk/models/getprojectsop.js";

let value: GetProjectsErl = {
  algo: "fixed_window",
  window: 2518.73,
  limit: 4529.22,
  keys: [
    "<value 1>",
    "<value 2>",
    "<value 3>",
  ],
};
```

## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `algo`                                                 | [models.GetProjectsAlgo](../models/getprojectsalgo.md) | :heavy_check_mark:                                     | N/A                                                    |
| `window`                                               | *number*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `limit`                                                | *number*                                               | :heavy_check_mark:                                     | N/A                                                    |
| `keys`                                                 | *string*[]                                             | :heavy_check_mark:                                     | N/A                                                    |