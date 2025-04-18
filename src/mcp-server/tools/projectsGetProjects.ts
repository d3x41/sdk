/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { projectsGetProjects } from "../../funcs/projectsGetProjects.js";
import { GetProjectsRequest$inboundSchema } from "../../models/getprojectsop.js";
import { formatResult, ToolDefinition } from "../tools.js";

const args = {
  request: GetProjectsRequest$inboundSchema,
};

export const tool$projectsGetProjects: ToolDefinition<typeof args> = {
  name: "projects-get-projects",
  description: `Retrieve a list of projects

Allows to retrieve the list of projects of the authenticated user or team. The list will be paginated and the provided query parameters allow filtering the returned projects.`,
  args,
  tool: async (client, args, ctx) => {
    const [result, apiCall] = await projectsGetProjects(
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
