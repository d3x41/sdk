/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { remap as remap$ } from "../lib/primitives.js";
import { safeParse } from "../lib/schemas.js";
import { ClosedEnum } from "../types/enums.js";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./sdkvalidationerror.js";

/**
 * Type of matcher. One of \"equals\", \"startsWith\", or \"endsWith\".
 */
export const UpdateCustomEnvironmentType = {
  Equals: "equals",
  StartsWith: "startsWith",
  EndsWith: "endsWith",
} as const;
/**
 * Type of matcher. One of \"equals\", \"startsWith\", or \"endsWith\".
 */
export type UpdateCustomEnvironmentType = ClosedEnum<
  typeof UpdateCustomEnvironmentType
>;

/**
 * How we want to determine a matching branch. This is optional.
 */
export type UpdateCustomEnvironmentBranchMatcher = {
  /**
   * Type of matcher. One of \"equals\", \"startsWith\", or \"endsWith\".
   */
  type: UpdateCustomEnvironmentType;
  /**
   * Git branch name or portion thereof.
   */
  pattern: string;
};

export type UpdateCustomEnvironmentRequestBody = {
  /**
   * The slug of the custom environment.
   */
  slug?: string | undefined;
  /**
   * Description of the custom environment. This is optional.
   */
  description?: string | undefined;
  /**
   * How we want to determine a matching branch. This is optional.
   */
  branchMatcher?: UpdateCustomEnvironmentBranchMatcher | null | undefined;
};

export type UpdateCustomEnvironmentRequest = {
  /**
   * The unique project identifier or the project name
   */
  idOrName: string;
  /**
   * The unique custom environment identifier within the project
   */
  environmentSlugOrId: string;
  /**
   * The Team identifier to perform the request on behalf of.
   */
  teamId?: string | undefined;
  /**
   * The Team slug to perform the request on behalf of.
   */
  slug?: string | undefined;
  requestBody?: UpdateCustomEnvironmentRequestBody | undefined;
};

/**
 * The type of environment (production, preview, or development)
 */
export const UpdateCustomEnvironmentEnvironmentType = {
  Production: "production",
  Preview: "preview",
  Development: "development",
} as const;
/**
 * The type of environment (production, preview, or development)
 */
export type UpdateCustomEnvironmentEnvironmentType = ClosedEnum<
  typeof UpdateCustomEnvironmentEnvironmentType
>;

/**
 * The type of matching to perform
 */
export const UpdateCustomEnvironmentEnvironmentResponseType = {
  EndsWith: "endsWith",
  StartsWith: "startsWith",
  Equals: "equals",
} as const;
/**
 * The type of matching to perform
 */
export type UpdateCustomEnvironmentEnvironmentResponseType = ClosedEnum<
  typeof UpdateCustomEnvironmentEnvironmentResponseType
>;

/**
 * Configuration for matching git branches to this environment
 */
export type UpdateCustomEnvironmentEnvironmentBranchMatcher = {
  /**
   * The type of matching to perform
   */
  type: UpdateCustomEnvironmentEnvironmentResponseType;
  /**
   * The pattern to match against branch names
   */
  pattern: string;
};

/**
 * A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
 */
export type UpdateCustomEnvironmentVerification = {
  type: string;
  domain: string;
  value: string;
  reason: string;
};

/**
 * List of domains associated with this environment
 */
export type UpdateCustomEnvironmentDomains = {
  name: string;
  apexName: string;
  projectId: string;
  redirect?: string | null | undefined;
  redirectStatusCode?: number | null | undefined;
  gitBranch?: string | null | undefined;
  customEnvironmentId?: string | null | undefined;
  updatedAt?: number | undefined;
  createdAt?: number | undefined;
  /**
   * `true` if the domain is verified for use with the project. If `false` it will not be used as an alias on this project until the challenge in `verification` is completed.
   */
  verified: boolean;
  /**
   * A list of verification challenges, one of which must be completed to verify the domain for use on the project. After the challenge is complete `POST /projects/:idOrName/domains/:domain/verify` to verify the domain. Possible challenges: - If `verification.type = TXT` the `verification.domain` will be checked for a TXT record matching `verification.value`.
   */
  verification?: Array<UpdateCustomEnvironmentVerification> | undefined;
};

/**
 * Internal representation of a custom environment with all required properties
 */
export type UpdateCustomEnvironmentResponseBody = {
  /**
   * Unique identifier for the custom environment (format: env_*)
   */
  id: string;
  /**
   * URL-friendly name of the environment
   */
  slug: string;
  /**
   * The type of environment (production, preview, or development)
   */
  type: UpdateCustomEnvironmentEnvironmentType;
  /**
   * Optional description of the environment's purpose
   */
  description?: string | undefined;
  /**
   * Configuration for matching git branches to this environment
   */
  branchMatcher?: UpdateCustomEnvironmentEnvironmentBranchMatcher | undefined;
  /**
   * List of domains associated with this environment
   */
  domains?: Array<UpdateCustomEnvironmentDomains> | undefined;
  /**
   * List of aliases for the current deployment
   */
  currentDeploymentAliases?: Array<string> | undefined;
  /**
   * Timestamp when the environment was created
   */
  createdAt: number;
  /**
   * Timestamp when the environment was last updated
   */
  updatedAt: number;
};

/** @internal */
export const UpdateCustomEnvironmentType$inboundSchema: z.ZodNativeEnum<
  typeof UpdateCustomEnvironmentType
> = z.nativeEnum(UpdateCustomEnvironmentType);

/** @internal */
export const UpdateCustomEnvironmentType$outboundSchema: z.ZodNativeEnum<
  typeof UpdateCustomEnvironmentType
> = UpdateCustomEnvironmentType$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentType$ {
  /** @deprecated use `UpdateCustomEnvironmentType$inboundSchema` instead. */
  export const inboundSchema = UpdateCustomEnvironmentType$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentType$outboundSchema` instead. */
  export const outboundSchema = UpdateCustomEnvironmentType$outboundSchema;
}

/** @internal */
export const UpdateCustomEnvironmentBranchMatcher$inboundSchema: z.ZodType<
  UpdateCustomEnvironmentBranchMatcher,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: UpdateCustomEnvironmentType$inboundSchema,
  pattern: z.string(),
});

/** @internal */
export type UpdateCustomEnvironmentBranchMatcher$Outbound = {
  type: string;
  pattern: string;
};

/** @internal */
export const UpdateCustomEnvironmentBranchMatcher$outboundSchema: z.ZodType<
  UpdateCustomEnvironmentBranchMatcher$Outbound,
  z.ZodTypeDef,
  UpdateCustomEnvironmentBranchMatcher
> = z.object({
  type: UpdateCustomEnvironmentType$outboundSchema,
  pattern: z.string(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentBranchMatcher$ {
  /** @deprecated use `UpdateCustomEnvironmentBranchMatcher$inboundSchema` instead. */
  export const inboundSchema =
    UpdateCustomEnvironmentBranchMatcher$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentBranchMatcher$outboundSchema` instead. */
  export const outboundSchema =
    UpdateCustomEnvironmentBranchMatcher$outboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentBranchMatcher$Outbound` instead. */
  export type Outbound = UpdateCustomEnvironmentBranchMatcher$Outbound;
}

export function updateCustomEnvironmentBranchMatcherToJSON(
  updateCustomEnvironmentBranchMatcher: UpdateCustomEnvironmentBranchMatcher,
): string {
  return JSON.stringify(
    UpdateCustomEnvironmentBranchMatcher$outboundSchema.parse(
      updateCustomEnvironmentBranchMatcher,
    ),
  );
}

export function updateCustomEnvironmentBranchMatcherFromJSON(
  jsonString: string,
): SafeParseResult<UpdateCustomEnvironmentBranchMatcher, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) =>
      UpdateCustomEnvironmentBranchMatcher$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'UpdateCustomEnvironmentBranchMatcher' from JSON`,
  );
}

/** @internal */
export const UpdateCustomEnvironmentRequestBody$inboundSchema: z.ZodType<
  UpdateCustomEnvironmentRequestBody,
  z.ZodTypeDef,
  unknown
> = z.object({
  slug: z.string().optional(),
  description: z.string().optional(),
  branchMatcher: z.nullable(
    z.lazy(() => UpdateCustomEnvironmentBranchMatcher$inboundSchema),
  ).optional(),
});

/** @internal */
export type UpdateCustomEnvironmentRequestBody$Outbound = {
  slug?: string | undefined;
  description?: string | undefined;
  branchMatcher?:
    | UpdateCustomEnvironmentBranchMatcher$Outbound
    | null
    | undefined;
};

/** @internal */
export const UpdateCustomEnvironmentRequestBody$outboundSchema: z.ZodType<
  UpdateCustomEnvironmentRequestBody$Outbound,
  z.ZodTypeDef,
  UpdateCustomEnvironmentRequestBody
> = z.object({
  slug: z.string().optional(),
  description: z.string().optional(),
  branchMatcher: z.nullable(
    z.lazy(() => UpdateCustomEnvironmentBranchMatcher$outboundSchema),
  ).optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentRequestBody$ {
  /** @deprecated use `UpdateCustomEnvironmentRequestBody$inboundSchema` instead. */
  export const inboundSchema = UpdateCustomEnvironmentRequestBody$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentRequestBody$outboundSchema` instead. */
  export const outboundSchema =
    UpdateCustomEnvironmentRequestBody$outboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentRequestBody$Outbound` instead. */
  export type Outbound = UpdateCustomEnvironmentRequestBody$Outbound;
}

export function updateCustomEnvironmentRequestBodyToJSON(
  updateCustomEnvironmentRequestBody: UpdateCustomEnvironmentRequestBody,
): string {
  return JSON.stringify(
    UpdateCustomEnvironmentRequestBody$outboundSchema.parse(
      updateCustomEnvironmentRequestBody,
    ),
  );
}

export function updateCustomEnvironmentRequestBodyFromJSON(
  jsonString: string,
): SafeParseResult<UpdateCustomEnvironmentRequestBody, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) =>
      UpdateCustomEnvironmentRequestBody$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'UpdateCustomEnvironmentRequestBody' from JSON`,
  );
}

/** @internal */
export const UpdateCustomEnvironmentRequest$inboundSchema: z.ZodType<
  UpdateCustomEnvironmentRequest,
  z.ZodTypeDef,
  unknown
> = z.object({
  idOrName: z.string(),
  environmentSlugOrId: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
  RequestBody: z.lazy(() => UpdateCustomEnvironmentRequestBody$inboundSchema)
    .optional(),
}).transform((v) => {
  return remap$(v, {
    "RequestBody": "requestBody",
  });
});

/** @internal */
export type UpdateCustomEnvironmentRequest$Outbound = {
  idOrName: string;
  environmentSlugOrId: string;
  teamId?: string | undefined;
  slug?: string | undefined;
  RequestBody?: UpdateCustomEnvironmentRequestBody$Outbound | undefined;
};

/** @internal */
export const UpdateCustomEnvironmentRequest$outboundSchema: z.ZodType<
  UpdateCustomEnvironmentRequest$Outbound,
  z.ZodTypeDef,
  UpdateCustomEnvironmentRequest
> = z.object({
  idOrName: z.string(),
  environmentSlugOrId: z.string(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
  requestBody: z.lazy(() => UpdateCustomEnvironmentRequestBody$outboundSchema)
    .optional(),
}).transform((v) => {
  return remap$(v, {
    requestBody: "RequestBody",
  });
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentRequest$ {
  /** @deprecated use `UpdateCustomEnvironmentRequest$inboundSchema` instead. */
  export const inboundSchema = UpdateCustomEnvironmentRequest$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentRequest$outboundSchema` instead. */
  export const outboundSchema = UpdateCustomEnvironmentRequest$outboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentRequest$Outbound` instead. */
  export type Outbound = UpdateCustomEnvironmentRequest$Outbound;
}

export function updateCustomEnvironmentRequestToJSON(
  updateCustomEnvironmentRequest: UpdateCustomEnvironmentRequest,
): string {
  return JSON.stringify(
    UpdateCustomEnvironmentRequest$outboundSchema.parse(
      updateCustomEnvironmentRequest,
    ),
  );
}

export function updateCustomEnvironmentRequestFromJSON(
  jsonString: string,
): SafeParseResult<UpdateCustomEnvironmentRequest, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => UpdateCustomEnvironmentRequest$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'UpdateCustomEnvironmentRequest' from JSON`,
  );
}

/** @internal */
export const UpdateCustomEnvironmentEnvironmentType$inboundSchema:
  z.ZodNativeEnum<typeof UpdateCustomEnvironmentEnvironmentType> = z.nativeEnum(
    UpdateCustomEnvironmentEnvironmentType,
  );

/** @internal */
export const UpdateCustomEnvironmentEnvironmentType$outboundSchema:
  z.ZodNativeEnum<typeof UpdateCustomEnvironmentEnvironmentType> =
    UpdateCustomEnvironmentEnvironmentType$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentEnvironmentType$ {
  /** @deprecated use `UpdateCustomEnvironmentEnvironmentType$inboundSchema` instead. */
  export const inboundSchema =
    UpdateCustomEnvironmentEnvironmentType$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentEnvironmentType$outboundSchema` instead. */
  export const outboundSchema =
    UpdateCustomEnvironmentEnvironmentType$outboundSchema;
}

/** @internal */
export const UpdateCustomEnvironmentEnvironmentResponseType$inboundSchema:
  z.ZodNativeEnum<typeof UpdateCustomEnvironmentEnvironmentResponseType> = z
    .nativeEnum(UpdateCustomEnvironmentEnvironmentResponseType);

/** @internal */
export const UpdateCustomEnvironmentEnvironmentResponseType$outboundSchema:
  z.ZodNativeEnum<typeof UpdateCustomEnvironmentEnvironmentResponseType> =
    UpdateCustomEnvironmentEnvironmentResponseType$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentEnvironmentResponseType$ {
  /** @deprecated use `UpdateCustomEnvironmentEnvironmentResponseType$inboundSchema` instead. */
  export const inboundSchema =
    UpdateCustomEnvironmentEnvironmentResponseType$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentEnvironmentResponseType$outboundSchema` instead. */
  export const outboundSchema =
    UpdateCustomEnvironmentEnvironmentResponseType$outboundSchema;
}

/** @internal */
export const UpdateCustomEnvironmentEnvironmentBranchMatcher$inboundSchema:
  z.ZodType<
    UpdateCustomEnvironmentEnvironmentBranchMatcher,
    z.ZodTypeDef,
    unknown
  > = z.object({
    type: UpdateCustomEnvironmentEnvironmentResponseType$inboundSchema,
    pattern: z.string(),
  });

/** @internal */
export type UpdateCustomEnvironmentEnvironmentBranchMatcher$Outbound = {
  type: string;
  pattern: string;
};

/** @internal */
export const UpdateCustomEnvironmentEnvironmentBranchMatcher$outboundSchema:
  z.ZodType<
    UpdateCustomEnvironmentEnvironmentBranchMatcher$Outbound,
    z.ZodTypeDef,
    UpdateCustomEnvironmentEnvironmentBranchMatcher
  > = z.object({
    type: UpdateCustomEnvironmentEnvironmentResponseType$outboundSchema,
    pattern: z.string(),
  });

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentEnvironmentBranchMatcher$ {
  /** @deprecated use `UpdateCustomEnvironmentEnvironmentBranchMatcher$inboundSchema` instead. */
  export const inboundSchema =
    UpdateCustomEnvironmentEnvironmentBranchMatcher$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentEnvironmentBranchMatcher$outboundSchema` instead. */
  export const outboundSchema =
    UpdateCustomEnvironmentEnvironmentBranchMatcher$outboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentEnvironmentBranchMatcher$Outbound` instead. */
  export type Outbound =
    UpdateCustomEnvironmentEnvironmentBranchMatcher$Outbound;
}

export function updateCustomEnvironmentEnvironmentBranchMatcherToJSON(
  updateCustomEnvironmentEnvironmentBranchMatcher:
    UpdateCustomEnvironmentEnvironmentBranchMatcher,
): string {
  return JSON.stringify(
    UpdateCustomEnvironmentEnvironmentBranchMatcher$outboundSchema.parse(
      updateCustomEnvironmentEnvironmentBranchMatcher,
    ),
  );
}

export function updateCustomEnvironmentEnvironmentBranchMatcherFromJSON(
  jsonString: string,
): SafeParseResult<
  UpdateCustomEnvironmentEnvironmentBranchMatcher,
  SDKValidationError
> {
  return safeParse(
    jsonString,
    (x) =>
      UpdateCustomEnvironmentEnvironmentBranchMatcher$inboundSchema.parse(
        JSON.parse(x),
      ),
    `Failed to parse 'UpdateCustomEnvironmentEnvironmentBranchMatcher' from JSON`,
  );
}

/** @internal */
export const UpdateCustomEnvironmentVerification$inboundSchema: z.ZodType<
  UpdateCustomEnvironmentVerification,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: z.string(),
  domain: z.string(),
  value: z.string(),
  reason: z.string(),
});

/** @internal */
export type UpdateCustomEnvironmentVerification$Outbound = {
  type: string;
  domain: string;
  value: string;
  reason: string;
};

/** @internal */
export const UpdateCustomEnvironmentVerification$outboundSchema: z.ZodType<
  UpdateCustomEnvironmentVerification$Outbound,
  z.ZodTypeDef,
  UpdateCustomEnvironmentVerification
> = z.object({
  type: z.string(),
  domain: z.string(),
  value: z.string(),
  reason: z.string(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentVerification$ {
  /** @deprecated use `UpdateCustomEnvironmentVerification$inboundSchema` instead. */
  export const inboundSchema =
    UpdateCustomEnvironmentVerification$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentVerification$outboundSchema` instead. */
  export const outboundSchema =
    UpdateCustomEnvironmentVerification$outboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentVerification$Outbound` instead. */
  export type Outbound = UpdateCustomEnvironmentVerification$Outbound;
}

export function updateCustomEnvironmentVerificationToJSON(
  updateCustomEnvironmentVerification: UpdateCustomEnvironmentVerification,
): string {
  return JSON.stringify(
    UpdateCustomEnvironmentVerification$outboundSchema.parse(
      updateCustomEnvironmentVerification,
    ),
  );
}

export function updateCustomEnvironmentVerificationFromJSON(
  jsonString: string,
): SafeParseResult<UpdateCustomEnvironmentVerification, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) =>
      UpdateCustomEnvironmentVerification$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'UpdateCustomEnvironmentVerification' from JSON`,
  );
}

/** @internal */
export const UpdateCustomEnvironmentDomains$inboundSchema: z.ZodType<
  UpdateCustomEnvironmentDomains,
  z.ZodTypeDef,
  unknown
> = z.object({
  name: z.string(),
  apexName: z.string(),
  projectId: z.string(),
  redirect: z.nullable(z.string()).optional(),
  redirectStatusCode: z.nullable(z.number()).optional(),
  gitBranch: z.nullable(z.string()).optional(),
  customEnvironmentId: z.nullable(z.string()).optional(),
  updatedAt: z.number().optional(),
  createdAt: z.number().optional(),
  verified: z.boolean(),
  verification: z.array(
    z.lazy(() => UpdateCustomEnvironmentVerification$inboundSchema),
  ).optional(),
});

/** @internal */
export type UpdateCustomEnvironmentDomains$Outbound = {
  name: string;
  apexName: string;
  projectId: string;
  redirect?: string | null | undefined;
  redirectStatusCode?: number | null | undefined;
  gitBranch?: string | null | undefined;
  customEnvironmentId?: string | null | undefined;
  updatedAt?: number | undefined;
  createdAt?: number | undefined;
  verified: boolean;
  verification?:
    | Array<UpdateCustomEnvironmentVerification$Outbound>
    | undefined;
};

/** @internal */
export const UpdateCustomEnvironmentDomains$outboundSchema: z.ZodType<
  UpdateCustomEnvironmentDomains$Outbound,
  z.ZodTypeDef,
  UpdateCustomEnvironmentDomains
> = z.object({
  name: z.string(),
  apexName: z.string(),
  projectId: z.string(),
  redirect: z.nullable(z.string()).optional(),
  redirectStatusCode: z.nullable(z.number()).optional(),
  gitBranch: z.nullable(z.string()).optional(),
  customEnvironmentId: z.nullable(z.string()).optional(),
  updatedAt: z.number().optional(),
  createdAt: z.number().optional(),
  verified: z.boolean(),
  verification: z.array(
    z.lazy(() => UpdateCustomEnvironmentVerification$outboundSchema),
  ).optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentDomains$ {
  /** @deprecated use `UpdateCustomEnvironmentDomains$inboundSchema` instead. */
  export const inboundSchema = UpdateCustomEnvironmentDomains$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentDomains$outboundSchema` instead. */
  export const outboundSchema = UpdateCustomEnvironmentDomains$outboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentDomains$Outbound` instead. */
  export type Outbound = UpdateCustomEnvironmentDomains$Outbound;
}

export function updateCustomEnvironmentDomainsToJSON(
  updateCustomEnvironmentDomains: UpdateCustomEnvironmentDomains,
): string {
  return JSON.stringify(
    UpdateCustomEnvironmentDomains$outboundSchema.parse(
      updateCustomEnvironmentDomains,
    ),
  );
}

export function updateCustomEnvironmentDomainsFromJSON(
  jsonString: string,
): SafeParseResult<UpdateCustomEnvironmentDomains, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => UpdateCustomEnvironmentDomains$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'UpdateCustomEnvironmentDomains' from JSON`,
  );
}

/** @internal */
export const UpdateCustomEnvironmentResponseBody$inboundSchema: z.ZodType<
  UpdateCustomEnvironmentResponseBody,
  z.ZodTypeDef,
  unknown
> = z.object({
  id: z.string(),
  slug: z.string(),
  type: UpdateCustomEnvironmentEnvironmentType$inboundSchema,
  description: z.string().optional(),
  branchMatcher: z.lazy(() =>
    UpdateCustomEnvironmentEnvironmentBranchMatcher$inboundSchema
  ).optional(),
  domains: z.array(z.lazy(() => UpdateCustomEnvironmentDomains$inboundSchema))
    .optional(),
  currentDeploymentAliases: z.array(z.string()).optional(),
  createdAt: z.number(),
  updatedAt: z.number(),
});

/** @internal */
export type UpdateCustomEnvironmentResponseBody$Outbound = {
  id: string;
  slug: string;
  type: string;
  description?: string | undefined;
  branchMatcher?:
    | UpdateCustomEnvironmentEnvironmentBranchMatcher$Outbound
    | undefined;
  domains?: Array<UpdateCustomEnvironmentDomains$Outbound> | undefined;
  currentDeploymentAliases?: Array<string> | undefined;
  createdAt: number;
  updatedAt: number;
};

/** @internal */
export const UpdateCustomEnvironmentResponseBody$outboundSchema: z.ZodType<
  UpdateCustomEnvironmentResponseBody$Outbound,
  z.ZodTypeDef,
  UpdateCustomEnvironmentResponseBody
> = z.object({
  id: z.string(),
  slug: z.string(),
  type: UpdateCustomEnvironmentEnvironmentType$outboundSchema,
  description: z.string().optional(),
  branchMatcher: z.lazy(() =>
    UpdateCustomEnvironmentEnvironmentBranchMatcher$outboundSchema
  ).optional(),
  domains: z.array(z.lazy(() => UpdateCustomEnvironmentDomains$outboundSchema))
    .optional(),
  currentDeploymentAliases: z.array(z.string()).optional(),
  createdAt: z.number(),
  updatedAt: z.number(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace UpdateCustomEnvironmentResponseBody$ {
  /** @deprecated use `UpdateCustomEnvironmentResponseBody$inboundSchema` instead. */
  export const inboundSchema =
    UpdateCustomEnvironmentResponseBody$inboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentResponseBody$outboundSchema` instead. */
  export const outboundSchema =
    UpdateCustomEnvironmentResponseBody$outboundSchema;
  /** @deprecated use `UpdateCustomEnvironmentResponseBody$Outbound` instead. */
  export type Outbound = UpdateCustomEnvironmentResponseBody$Outbound;
}

export function updateCustomEnvironmentResponseBodyToJSON(
  updateCustomEnvironmentResponseBody: UpdateCustomEnvironmentResponseBody,
): string {
  return JSON.stringify(
    UpdateCustomEnvironmentResponseBody$outboundSchema.parse(
      updateCustomEnvironmentResponseBody,
    ),
  );
}

export function updateCustomEnvironmentResponseBodyFromJSON(
  jsonString: string,
): SafeParseResult<UpdateCustomEnvironmentResponseBody, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) =>
      UpdateCustomEnvironmentResponseBody$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'UpdateCustomEnvironmentResponseBody' from JSON`,
  );
}
