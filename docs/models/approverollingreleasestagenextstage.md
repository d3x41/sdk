# ApproveRollingReleaseStageNextStage

The next stage to be activated, null if not in ACTIVE state

## Example Usage

```typescript
import { ApproveRollingReleaseStageNextStage } from "@vercel/sdk/models/approverollingreleasestageop.js";

let value: ApproveRollingReleaseStageNextStage = {
  index: 2,
  isFinalStage: false,
  targetPercentage: 60,
  requireApproval: true,
  duration: null,
};
```

## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              | Example                                                                                  |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `index`                                                                                  | *number*                                                                                 | :heavy_check_mark:                                                                       | The zero-based index of the stage                                                        | 0                                                                                        |
| `isFinalStage`                                                                           | *boolean*                                                                                | :heavy_check_mark:                                                                       | Whether or not this stage is the final stage (targetPercentage === 100)                  | false                                                                                    |
| `targetPercentage`                                                                       | *number*                                                                                 | :heavy_check_mark:                                                                       | The percentage of traffic to serve to the canary deployment (0-100)                      | 25                                                                                       |
| `requireApproval`                                                                        | *boolean*                                                                                | :heavy_check_mark:                                                                       | Whether or not this stage requires manual approval to proceed                            |                                                                                          |
| `duration`                                                                               | *number*                                                                                 | :heavy_check_mark:                                                                       | Duration in seconds for automatic advancement, null for manual stages or the final stage | <nil>                                                                                    |