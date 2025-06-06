/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { logDrainsDeleteConfigurableLogDrain } from "../../funcs/logDrainsDeleteConfigurableLogDrain.js";
import { DeleteConfigurableLogDrainRequest$inboundSchema } from "../../models/deleteconfigurablelogdrainop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: DeleteConfigurableLogDrainRequest$inboundSchema,
};

export const tool$logDrainsDeleteConfigurableLogDrain: ToolDefinition<
  typeof args
> = {
  name: "log-drains-delete-configurable-log-drain",
  description: `Deletes a Configurable Log Drain

Deletes a Configurable Log Drain. This endpoint must be called with a team AccessToken (integration OAuth2 clients are not allowed). Only log drains owned by the authenticated team can be deleted.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await logDrainsDeleteConfigurableLogDrain(
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
