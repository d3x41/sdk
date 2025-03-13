/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { artifactsRecordEvents } from "../../funcs/artifactsRecordEvents.js";
import { RecordEventsRequest$inboundSchema } from "../../models/recordeventsop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: RecordEventsRequest$inboundSchema,
};

export const tool$artifactsRecordEvents: ToolDefinition<typeof args> = {
  name: "artifacts-record-events",
  description: `Record an artifacts cache usage event

Records an artifacts cache usage event. The body of this request is an array of cache usage events. The supported event types are \`HIT\` and \`MISS\`. The source is either \`LOCAL\` the cache event was on the users filesystem cache or \`REMOTE\` if the cache event is for a remote cache. When the event is a \`HIT\` the request also accepts a number \`duration\` which is the time taken to generate the artifact in the cache.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await artifactsRecordEvents(
      client,
      args.request,
      { fetchOptions: { signal: ctx.signal } },
    ).$inspect();

    if (!result.ok) {
      return {
        content: [{ type: "text", text: result.error.message }],
        isError: true,
      };
    }

    return formatResult(void 0, apiCall);
  },
};
