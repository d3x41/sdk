/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { remap as remap$ } from "../lib/primitives.js";
import { safeParse } from "../lib/schemas.js";
import { ClosedEnum } from "../types/enums.js";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./sdkvalidationerror.js";

export type GetBypassIpRequest = {
  projectId: string;
  limit?: number | undefined;
  /**
   * Filter by source IP
   */
  sourceIp?: string | undefined;
  /**
   * Filter by domain
   */
  domain?: string | undefined;
  /**
   * Filter by project scoped rules
   */
  projectScope?: boolean | undefined;
  /**
   * Used for pagination. Retrieves results after the provided id
   */
  offset?: string | undefined;
  /**
   * The Team identifier to perform the request on behalf of.
   */
  teamId?: string | undefined;
  /**
   * The Team slug to perform the request on behalf of.
   */
  slug?: string | undefined;
};

export const ResponseBodyAction = {
  Block: "block",
  Bypass: "bypass",
} as const;
export type ResponseBodyAction = ClosedEnum<typeof ResponseBodyAction>;

export type GetBypassIpResponseBodyResult = {
  ownerId: string;
  id: string;
  domain: string;
  ip: string;
  action?: ResponseBodyAction | undefined;
  projectId?: string | undefined;
  isProjectRule?: boolean | undefined;
  note?: string | undefined;
  createdAt: string;
  actorId?: string | undefined;
  updatedAt: string;
  updatedAtHour: string;
  deletedAt?: string | undefined;
  expiresAt?: number | undefined;
};

export type GetBypassIpResponseBodyPagination = {
  ownerId: string;
  id: string;
};

export type GetBypassIpResponseBody2 = {
  result?: Array<GetBypassIpResponseBodyResult> | undefined;
  pagination?: GetBypassIpResponseBodyPagination | undefined;
};

export type ResponseBodyResult = {
  ownerId: string;
  id: string;
  domain: string;
  ip: string;
  projectId: string;
  isProjectRule: boolean;
};

export type GetBypassIpResponseBody1 = {
  result: Array<ResponseBodyResult>;
  pagination?: any | null | undefined;
};

export type GetBypassIpResponseBody =
  | GetBypassIpResponseBody1
  | GetBypassIpResponseBody2;

/** @internal */
export const GetBypassIpRequest$inboundSchema: z.ZodType<
  GetBypassIpRequest,
  z.ZodTypeDef,
  unknown
> = z.object({
  projectId: z.string(),
  limit: z.number().optional(),
  sourceIp: z.string().optional(),
  domain: z.string().optional(),
  projectScope: z.boolean().optional(),
  offset: z.string().optional(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/** @internal */
export type GetBypassIpRequest$Outbound = {
  projectId: string;
  limit?: number | undefined;
  sourceIp?: string | undefined;
  domain?: string | undefined;
  projectScope?: boolean | undefined;
  offset?: string | undefined;
  teamId?: string | undefined;
  slug?: string | undefined;
};

/** @internal */
export const GetBypassIpRequest$outboundSchema: z.ZodType<
  GetBypassIpRequest$Outbound,
  z.ZodTypeDef,
  GetBypassIpRequest
> = z.object({
  projectId: z.string(),
  limit: z.number().optional(),
  sourceIp: z.string().optional(),
  domain: z.string().optional(),
  projectScope: z.boolean().optional(),
  offset: z.string().optional(),
  teamId: z.string().optional(),
  slug: z.string().optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetBypassIpRequest$ {
  /** @deprecated use `GetBypassIpRequest$inboundSchema` instead. */
  export const inboundSchema = GetBypassIpRequest$inboundSchema;
  /** @deprecated use `GetBypassIpRequest$outboundSchema` instead. */
  export const outboundSchema = GetBypassIpRequest$outboundSchema;
  /** @deprecated use `GetBypassIpRequest$Outbound` instead. */
  export type Outbound = GetBypassIpRequest$Outbound;
}

export function getBypassIpRequestToJSON(
  getBypassIpRequest: GetBypassIpRequest,
): string {
  return JSON.stringify(
    GetBypassIpRequest$outboundSchema.parse(getBypassIpRequest),
  );
}

export function getBypassIpRequestFromJSON(
  jsonString: string,
): SafeParseResult<GetBypassIpRequest, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetBypassIpRequest$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetBypassIpRequest' from JSON`,
  );
}

/** @internal */
export const ResponseBodyAction$inboundSchema: z.ZodNativeEnum<
  typeof ResponseBodyAction
> = z.nativeEnum(ResponseBodyAction);

/** @internal */
export const ResponseBodyAction$outboundSchema: z.ZodNativeEnum<
  typeof ResponseBodyAction
> = ResponseBodyAction$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace ResponseBodyAction$ {
  /** @deprecated use `ResponseBodyAction$inboundSchema` instead. */
  export const inboundSchema = ResponseBodyAction$inboundSchema;
  /** @deprecated use `ResponseBodyAction$outboundSchema` instead. */
  export const outboundSchema = ResponseBodyAction$outboundSchema;
}

/** @internal */
export const GetBypassIpResponseBodyResult$inboundSchema: z.ZodType<
  GetBypassIpResponseBodyResult,
  z.ZodTypeDef,
  unknown
> = z.object({
  OwnerId: z.string(),
  Id: z.string(),
  Domain: z.string(),
  Ip: z.string(),
  Action: ResponseBodyAction$inboundSchema.optional(),
  ProjectId: z.string().optional(),
  IsProjectRule: z.boolean().optional(),
  Note: z.string().optional(),
  CreatedAt: z.string(),
  ActorId: z.string().optional(),
  UpdatedAt: z.string(),
  UpdatedAtHour: z.string(),
  DeletedAt: z.string().optional(),
  ExpiresAt: z.number().optional(),
}).transform((v) => {
  return remap$(v, {
    "OwnerId": "ownerId",
    "Id": "id",
    "Domain": "domain",
    "Ip": "ip",
    "Action": "action",
    "ProjectId": "projectId",
    "IsProjectRule": "isProjectRule",
    "Note": "note",
    "CreatedAt": "createdAt",
    "ActorId": "actorId",
    "UpdatedAt": "updatedAt",
    "UpdatedAtHour": "updatedAtHour",
    "DeletedAt": "deletedAt",
    "ExpiresAt": "expiresAt",
  });
});

/** @internal */
export type GetBypassIpResponseBodyResult$Outbound = {
  OwnerId: string;
  Id: string;
  Domain: string;
  Ip: string;
  Action?: string | undefined;
  ProjectId?: string | undefined;
  IsProjectRule?: boolean | undefined;
  Note?: string | undefined;
  CreatedAt: string;
  ActorId?: string | undefined;
  UpdatedAt: string;
  UpdatedAtHour: string;
  DeletedAt?: string | undefined;
  ExpiresAt?: number | undefined;
};

/** @internal */
export const GetBypassIpResponseBodyResult$outboundSchema: z.ZodType<
  GetBypassIpResponseBodyResult$Outbound,
  z.ZodTypeDef,
  GetBypassIpResponseBodyResult
> = z.object({
  ownerId: z.string(),
  id: z.string(),
  domain: z.string(),
  ip: z.string(),
  action: ResponseBodyAction$outboundSchema.optional(),
  projectId: z.string().optional(),
  isProjectRule: z.boolean().optional(),
  note: z.string().optional(),
  createdAt: z.string(),
  actorId: z.string().optional(),
  updatedAt: z.string(),
  updatedAtHour: z.string(),
  deletedAt: z.string().optional(),
  expiresAt: z.number().optional(),
}).transform((v) => {
  return remap$(v, {
    ownerId: "OwnerId",
    id: "Id",
    domain: "Domain",
    ip: "Ip",
    action: "Action",
    projectId: "ProjectId",
    isProjectRule: "IsProjectRule",
    note: "Note",
    createdAt: "CreatedAt",
    actorId: "ActorId",
    updatedAt: "UpdatedAt",
    updatedAtHour: "UpdatedAtHour",
    deletedAt: "DeletedAt",
    expiresAt: "ExpiresAt",
  });
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetBypassIpResponseBodyResult$ {
  /** @deprecated use `GetBypassIpResponseBodyResult$inboundSchema` instead. */
  export const inboundSchema = GetBypassIpResponseBodyResult$inboundSchema;
  /** @deprecated use `GetBypassIpResponseBodyResult$outboundSchema` instead. */
  export const outboundSchema = GetBypassIpResponseBodyResult$outboundSchema;
  /** @deprecated use `GetBypassIpResponseBodyResult$Outbound` instead. */
  export type Outbound = GetBypassIpResponseBodyResult$Outbound;
}

export function getBypassIpResponseBodyResultToJSON(
  getBypassIpResponseBodyResult: GetBypassIpResponseBodyResult,
): string {
  return JSON.stringify(
    GetBypassIpResponseBodyResult$outboundSchema.parse(
      getBypassIpResponseBodyResult,
    ),
  );
}

export function getBypassIpResponseBodyResultFromJSON(
  jsonString: string,
): SafeParseResult<GetBypassIpResponseBodyResult, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetBypassIpResponseBodyResult$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetBypassIpResponseBodyResult' from JSON`,
  );
}

/** @internal */
export const GetBypassIpResponseBodyPagination$inboundSchema: z.ZodType<
  GetBypassIpResponseBodyPagination,
  z.ZodTypeDef,
  unknown
> = z.object({
  OwnerId: z.string(),
  Id: z.string(),
}).transform((v) => {
  return remap$(v, {
    "OwnerId": "ownerId",
    "Id": "id",
  });
});

/** @internal */
export type GetBypassIpResponseBodyPagination$Outbound = {
  OwnerId: string;
  Id: string;
};

/** @internal */
export const GetBypassIpResponseBodyPagination$outboundSchema: z.ZodType<
  GetBypassIpResponseBodyPagination$Outbound,
  z.ZodTypeDef,
  GetBypassIpResponseBodyPagination
> = z.object({
  ownerId: z.string(),
  id: z.string(),
}).transform((v) => {
  return remap$(v, {
    ownerId: "OwnerId",
    id: "Id",
  });
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetBypassIpResponseBodyPagination$ {
  /** @deprecated use `GetBypassIpResponseBodyPagination$inboundSchema` instead. */
  export const inboundSchema = GetBypassIpResponseBodyPagination$inboundSchema;
  /** @deprecated use `GetBypassIpResponseBodyPagination$outboundSchema` instead. */
  export const outboundSchema =
    GetBypassIpResponseBodyPagination$outboundSchema;
  /** @deprecated use `GetBypassIpResponseBodyPagination$Outbound` instead. */
  export type Outbound = GetBypassIpResponseBodyPagination$Outbound;
}

export function getBypassIpResponseBodyPaginationToJSON(
  getBypassIpResponseBodyPagination: GetBypassIpResponseBodyPagination,
): string {
  return JSON.stringify(
    GetBypassIpResponseBodyPagination$outboundSchema.parse(
      getBypassIpResponseBodyPagination,
    ),
  );
}

export function getBypassIpResponseBodyPaginationFromJSON(
  jsonString: string,
): SafeParseResult<GetBypassIpResponseBodyPagination, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetBypassIpResponseBodyPagination$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetBypassIpResponseBodyPagination' from JSON`,
  );
}

/** @internal */
export const GetBypassIpResponseBody2$inboundSchema: z.ZodType<
  GetBypassIpResponseBody2,
  z.ZodTypeDef,
  unknown
> = z.object({
  result: z.array(z.lazy(() => GetBypassIpResponseBodyResult$inboundSchema))
    .optional(),
  pagination: z.lazy(() => GetBypassIpResponseBodyPagination$inboundSchema)
    .optional(),
});

/** @internal */
export type GetBypassIpResponseBody2$Outbound = {
  result?: Array<GetBypassIpResponseBodyResult$Outbound> | undefined;
  pagination?: GetBypassIpResponseBodyPagination$Outbound | undefined;
};

/** @internal */
export const GetBypassIpResponseBody2$outboundSchema: z.ZodType<
  GetBypassIpResponseBody2$Outbound,
  z.ZodTypeDef,
  GetBypassIpResponseBody2
> = z.object({
  result: z.array(z.lazy(() => GetBypassIpResponseBodyResult$outboundSchema))
    .optional(),
  pagination: z.lazy(() => GetBypassIpResponseBodyPagination$outboundSchema)
    .optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetBypassIpResponseBody2$ {
  /** @deprecated use `GetBypassIpResponseBody2$inboundSchema` instead. */
  export const inboundSchema = GetBypassIpResponseBody2$inboundSchema;
  /** @deprecated use `GetBypassIpResponseBody2$outboundSchema` instead. */
  export const outboundSchema = GetBypassIpResponseBody2$outboundSchema;
  /** @deprecated use `GetBypassIpResponseBody2$Outbound` instead. */
  export type Outbound = GetBypassIpResponseBody2$Outbound;
}

export function getBypassIpResponseBody2ToJSON(
  getBypassIpResponseBody2: GetBypassIpResponseBody2,
): string {
  return JSON.stringify(
    GetBypassIpResponseBody2$outboundSchema.parse(getBypassIpResponseBody2),
  );
}

export function getBypassIpResponseBody2FromJSON(
  jsonString: string,
): SafeParseResult<GetBypassIpResponseBody2, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetBypassIpResponseBody2$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetBypassIpResponseBody2' from JSON`,
  );
}

/** @internal */
export const ResponseBodyResult$inboundSchema: z.ZodType<
  ResponseBodyResult,
  z.ZodTypeDef,
  unknown
> = z.object({
  OwnerId: z.string(),
  Id: z.string(),
  Domain: z.string(),
  Ip: z.string(),
  ProjectId: z.string(),
  IsProjectRule: z.boolean(),
}).transform((v) => {
  return remap$(v, {
    "OwnerId": "ownerId",
    "Id": "id",
    "Domain": "domain",
    "Ip": "ip",
    "ProjectId": "projectId",
    "IsProjectRule": "isProjectRule",
  });
});

/** @internal */
export type ResponseBodyResult$Outbound = {
  OwnerId: string;
  Id: string;
  Domain: string;
  Ip: string;
  ProjectId: string;
  IsProjectRule: boolean;
};

/** @internal */
export const ResponseBodyResult$outboundSchema: z.ZodType<
  ResponseBodyResult$Outbound,
  z.ZodTypeDef,
  ResponseBodyResult
> = z.object({
  ownerId: z.string(),
  id: z.string(),
  domain: z.string(),
  ip: z.string(),
  projectId: z.string(),
  isProjectRule: z.boolean(),
}).transform((v) => {
  return remap$(v, {
    ownerId: "OwnerId",
    id: "Id",
    domain: "Domain",
    ip: "Ip",
    projectId: "ProjectId",
    isProjectRule: "IsProjectRule",
  });
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace ResponseBodyResult$ {
  /** @deprecated use `ResponseBodyResult$inboundSchema` instead. */
  export const inboundSchema = ResponseBodyResult$inboundSchema;
  /** @deprecated use `ResponseBodyResult$outboundSchema` instead. */
  export const outboundSchema = ResponseBodyResult$outboundSchema;
  /** @deprecated use `ResponseBodyResult$Outbound` instead. */
  export type Outbound = ResponseBodyResult$Outbound;
}

export function responseBodyResultToJSON(
  responseBodyResult: ResponseBodyResult,
): string {
  return JSON.stringify(
    ResponseBodyResult$outboundSchema.parse(responseBodyResult),
  );
}

export function responseBodyResultFromJSON(
  jsonString: string,
): SafeParseResult<ResponseBodyResult, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => ResponseBodyResult$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'ResponseBodyResult' from JSON`,
  );
}

/** @internal */
export const GetBypassIpResponseBody1$inboundSchema: z.ZodType<
  GetBypassIpResponseBody1,
  z.ZodTypeDef,
  unknown
> = z.object({
  result: z.array(z.lazy(() => ResponseBodyResult$inboundSchema)),
  pagination: z.nullable(z.any()).optional(),
});

/** @internal */
export type GetBypassIpResponseBody1$Outbound = {
  result: Array<ResponseBodyResult$Outbound>;
  pagination?: any | null | undefined;
};

/** @internal */
export const GetBypassIpResponseBody1$outboundSchema: z.ZodType<
  GetBypassIpResponseBody1$Outbound,
  z.ZodTypeDef,
  GetBypassIpResponseBody1
> = z.object({
  result: z.array(z.lazy(() => ResponseBodyResult$outboundSchema)),
  pagination: z.nullable(z.any()).optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetBypassIpResponseBody1$ {
  /** @deprecated use `GetBypassIpResponseBody1$inboundSchema` instead. */
  export const inboundSchema = GetBypassIpResponseBody1$inboundSchema;
  /** @deprecated use `GetBypassIpResponseBody1$outboundSchema` instead. */
  export const outboundSchema = GetBypassIpResponseBody1$outboundSchema;
  /** @deprecated use `GetBypassIpResponseBody1$Outbound` instead. */
  export type Outbound = GetBypassIpResponseBody1$Outbound;
}

export function getBypassIpResponseBody1ToJSON(
  getBypassIpResponseBody1: GetBypassIpResponseBody1,
): string {
  return JSON.stringify(
    GetBypassIpResponseBody1$outboundSchema.parse(getBypassIpResponseBody1),
  );
}

export function getBypassIpResponseBody1FromJSON(
  jsonString: string,
): SafeParseResult<GetBypassIpResponseBody1, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetBypassIpResponseBody1$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetBypassIpResponseBody1' from JSON`,
  );
}

/** @internal */
export const GetBypassIpResponseBody$inboundSchema: z.ZodType<
  GetBypassIpResponseBody,
  z.ZodTypeDef,
  unknown
> = z.union([
  z.lazy(() => GetBypassIpResponseBody1$inboundSchema),
  z.lazy(() => GetBypassIpResponseBody2$inboundSchema),
]);

/** @internal */
export type GetBypassIpResponseBody$Outbound =
  | GetBypassIpResponseBody1$Outbound
  | GetBypassIpResponseBody2$Outbound;

/** @internal */
export const GetBypassIpResponseBody$outboundSchema: z.ZodType<
  GetBypassIpResponseBody$Outbound,
  z.ZodTypeDef,
  GetBypassIpResponseBody
> = z.union([
  z.lazy(() => GetBypassIpResponseBody1$outboundSchema),
  z.lazy(() => GetBypassIpResponseBody2$outboundSchema),
]);

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace GetBypassIpResponseBody$ {
  /** @deprecated use `GetBypassIpResponseBody$inboundSchema` instead. */
  export const inboundSchema = GetBypassIpResponseBody$inboundSchema;
  /** @deprecated use `GetBypassIpResponseBody$outboundSchema` instead. */
  export const outboundSchema = GetBypassIpResponseBody$outboundSchema;
  /** @deprecated use `GetBypassIpResponseBody$Outbound` instead. */
  export type Outbound = GetBypassIpResponseBody$Outbound;
}

export function getBypassIpResponseBodyToJSON(
  getBypassIpResponseBody: GetBypassIpResponseBody,
): string {
  return JSON.stringify(
    GetBypassIpResponseBody$outboundSchema.parse(getBypassIpResponseBody),
  );
}

export function getBypassIpResponseBodyFromJSON(
  jsonString: string,
): SafeParseResult<GetBypassIpResponseBody, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => GetBypassIpResponseBody$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'GetBypassIpResponseBody' from JSON`,
  );
}
