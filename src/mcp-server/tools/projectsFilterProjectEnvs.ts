/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { projectsFilterProjectEnvs } from "../../funcs/projectsFilterProjectEnvs.js";
import { FilterProjectEnvsRequest$inboundSchema } from "../../models/filterprojectenvsop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: FilterProjectEnvsRequest$inboundSchema,
};

export const tool$projectsFilterProjectEnvs: ToolDefinition<typeof args> = {
  name: "projects-filter-project-envs",
  description: `Retrieve the environment variables of a project by id or name

Retrieve the environment variables for a given project by passing either the project \`id\` or \`name\` in the URL.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await projectsFilterProjectEnvs(
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
