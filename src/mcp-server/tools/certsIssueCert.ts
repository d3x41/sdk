/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { certsIssueCert } from "../../funcs/certsIssueCert.js";
import { IssueCertRequest$inboundSchema } from "../../models/issuecertop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: IssueCertRequest$inboundSchema,
};

export const tool$certsIssueCert: ToolDefinition<typeof args> = {
  name: "certs-issue-cert",
  description: `Issue a new cert

Issue a new cert`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await certsIssueCert(
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
