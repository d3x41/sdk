/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { checksGetCheck } from "../../funcs/checksGetCheck.js";
import { GetCheckRequest$inboundSchema } from "../../models/getcheckop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: GetCheckRequest$inboundSchema,
};

export const tool$checksGetCheck: ToolDefinition<typeof args> = {
  name: "checks-get-check",
  description: `Get a single check

Return a detailed response for a single check.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await checksGetCheck(
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
