/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { remap as remap$ } from "../lib/primitives.js";
import { safeParse } from "../lib/schemas.js";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./sdkvalidationerror.js";

export type DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody =
  {
    /**
     * Delete Environment Variables that are not assigned to any environments.
     */
    deleteUnassignedEnvironmentVariables?: boolean | undefined;
  };

export type DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest =
  {
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
    requestBody?:
      | DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody
      | undefined;
  };

export type DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody =
  {};

/** @internal */
export const DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$inboundSchema:
  z.ZodType<
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody,
    z.ZodTypeDef,
    unknown
  > = z.object({
    deleteUnassignedEnvironmentVariables: z.boolean().optional(),
  });

/** @internal */
export type DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$Outbound =
  {
    deleteUnassignedEnvironmentVariables?: boolean | undefined;
  };

/** @internal */
export const DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$outboundSchema:
  z.ZodType<
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$Outbound,
    z.ZodTypeDef,
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody
  > = z.object({
    deleteUnassignedEnvironmentVariables: z.boolean().optional(),
  });

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$ {
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$inboundSchema` instead. */
  export const inboundSchema =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$inboundSchema;
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$outboundSchema` instead. */
  export const outboundSchema =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$outboundSchema;
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$Outbound` instead. */
  export type Outbound =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$Outbound;
}

export function deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBodyToJSON(
  deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody:
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody,
): string {
  return JSON.stringify(
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$outboundSchema
      .parse(
        deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody,
      ),
  );
}

export function deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBodyFromJSON(
  jsonString: string,
): SafeParseResult<
  DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody,
  SDKValidationError
> {
  return safeParse(
    jsonString,
    (x) =>
      DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$inboundSchema
        .parse(JSON.parse(x)),
    `Failed to parse 'DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody' from JSON`,
  );
}

/** @internal */
export const DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$inboundSchema:
  z.ZodType<
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest,
    z.ZodTypeDef,
    unknown
  > = z.object({
    idOrName: z.string(),
    environmentSlugOrId: z.string(),
    teamId: z.string().optional(),
    slug: z.string().optional(),
    RequestBody: z.lazy(() =>
      DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$inboundSchema
    ).optional(),
  }).transform((v) => {
    return remap$(v, {
      "RequestBody": "requestBody",
    });
  });

/** @internal */
export type DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$Outbound =
  {
    idOrName: string;
    environmentSlugOrId: string;
    teamId?: string | undefined;
    slug?: string | undefined;
    RequestBody?:
      | DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$Outbound
      | undefined;
  };

/** @internal */
export const DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$outboundSchema:
  z.ZodType<
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$Outbound,
    z.ZodTypeDef,
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest
  > = z.object({
    idOrName: z.string(),
    environmentSlugOrId: z.string(),
    teamId: z.string().optional(),
    slug: z.string().optional(),
    requestBody: z.lazy(() =>
      DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestBody$outboundSchema
    ).optional(),
  }).transform((v) => {
    return remap$(v, {
      requestBody: "RequestBody",
    });
  });

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$ {
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$inboundSchema` instead. */
  export const inboundSchema =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$inboundSchema;
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$outboundSchema` instead. */
  export const outboundSchema =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$outboundSchema;
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$Outbound` instead. */
  export type Outbound =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$Outbound;
}

export function deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestToJSON(
  deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest:
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest,
): string {
  return JSON.stringify(
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$outboundSchema
      .parse(
        deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest,
      ),
  );
}

export function deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequestFromJSON(
  jsonString: string,
): SafeParseResult<
  DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest,
  SDKValidationError
> {
  return safeParse(
    jsonString,
    (x) =>
      DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest$inboundSchema
        .parse(JSON.parse(x)),
    `Failed to parse 'DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdRequest' from JSON`,
  );
}

/** @internal */
export const DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$inboundSchema:
  z.ZodType<
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody,
    z.ZodTypeDef,
    unknown
  > = z.object({});

/** @internal */
export type DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$Outbound =
  {};

/** @internal */
export const DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$outboundSchema:
  z.ZodType<
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$Outbound,
    z.ZodTypeDef,
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody
  > = z.object({});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$ {
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$inboundSchema` instead. */
  export const inboundSchema =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$inboundSchema;
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$outboundSchema` instead. */
  export const outboundSchema =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$outboundSchema;
  /** @deprecated use `DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$Outbound` instead. */
  export type Outbound =
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$Outbound;
}

export function deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBodyToJSON(
  deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody:
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody,
): string {
  return JSON.stringify(
    DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$outboundSchema
      .parse(
        deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody,
      ),
  );
}

export function deleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBodyFromJSON(
  jsonString: string,
): SafeParseResult<
  DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody,
  SDKValidationError
> {
  return safeParse(
    jsonString,
    (x) =>
      DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody$inboundSchema
        .parse(JSON.parse(x)),
    `Failed to parse 'DeleteV9ProjectsIdOrNameCustomEnvironmentsEnvironmentSlugOrIdResponseBody' from JSON`,
  );
}
