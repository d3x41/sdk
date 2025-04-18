/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { domainsGetDomainConfig } from "../../funcs/domainsGetDomainConfig.js";
import { GetDomainConfigRequest$inboundSchema } from "../../models/getdomainconfigop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: GetDomainConfigRequest$inboundSchema,
};

export const tool$domainsGetDomainConfig: ToolDefinition<typeof args> = {
  name: "domains-get-domain-config",
  description: `Get a Domain's configuration

Get a Domain's configuration.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await domainsGetDomainConfig(
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
