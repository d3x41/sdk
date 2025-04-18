# GetProjectsBranchMatcher

## Example Usage

```typescript
import { GetProjectsBranchMatcher } from "@vercel/sdk/models/getprojectsop.js";

let value: GetProjectsBranchMatcher = {
  type: "endsWith",
  pattern: "<value>",
};
```

## Fields

| Field                                                                                                                                              | Type                                                                                                                                               | Required                                                                                                                                           | Description                                                                                                                                        |
| -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------- |
| `type`                                                                                                                                             | [models.GetProjectsProjectsResponse200ApplicationJSONResponseBodyType](../models/getprojectsprojectsresponse200applicationjsonresponsebodytype.md) | :heavy_check_mark:                                                                                                                                 | The type of matching to perform                                                                                                                    |
| `pattern`                                                                                                                                          | *string*                                                                                                                                           | :heavy_check_mark:                                                                                                                                 | The pattern to match against branch names                                                                                                          |