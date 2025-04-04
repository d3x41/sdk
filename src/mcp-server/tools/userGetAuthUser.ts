/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { userGetAuthUser } from "../../funcs/userGetAuthUser.js";
import { formatResult, ToolDefinition } from "../tools.js";

export const tool$userGetAuthUser: ToolDefinition = {
  name: "user-get-auth-user",
  description: `Get the User

Retrieves information related to the currently authenticated User.`,
  tool: async (client, ctx) => {
    const [result, apiCall] = await userGetAuthUser(
      client,
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
