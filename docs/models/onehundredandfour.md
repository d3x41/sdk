# OneHundredAndFour

The payload of the event, if requested.

## Example Usage

```typescript
import { OneHundredAndFour } from "@vercel/sdk/models/userevent.js";

let value: OneHundredAndFour = {
  slug: "<value>",
  teamId: "<id>",
  by: "<value>",
};
```

## Fields

| Field                                    | Type                                     | Required                                 | Description                              |
| ---------------------------------------- | ---------------------------------------- | ---------------------------------------- | ---------------------------------------- |
| `slug`                                   | *string*                                 | :heavy_check_mark:                       | N/A                                      |
| `teamId`                                 | *string*                                 | :heavy_check_mark:                       | N/A                                      |
| `by`                                     | *string*                                 | :heavy_check_mark:                       | N/A                                      |
| `reasons`                                | [models.Reasons](../models/reasons.md)[] | :heavy_minus_sign:                       | N/A                                      |