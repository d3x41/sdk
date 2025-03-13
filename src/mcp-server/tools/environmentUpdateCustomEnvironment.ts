/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { environmentUpdateCustomEnvironment } from "../../funcs/environmentUpdateCustomEnvironment.js";
import { UpdateCustomEnvironmentRequest$inboundSchema } from "../../models/updatecustomenvironmentop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: UpdateCustomEnvironmentRequest$inboundSchema,
};

export const tool$environmentUpdateCustomEnvironment: ToolDefinition<
  typeof args
> = {
  name: "environment-update-custom-environment",
  description: `Update a custom environment

Update a custom environment for the project. Must not be named 'Production' or 'Preview'.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await environmentUpdateCustomEnvironment(
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

    const value = result.value;

    return formatResult(value, apiCall);
  },
};
