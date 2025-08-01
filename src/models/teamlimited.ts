/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import * as z from "zod";
import { safeParse } from "../lib/schemas.js";
import { ClosedEnum } from "../types/enums.js";
import { Result as SafeParseResult } from "../types/fp.js";
import { SDKValidationError } from "./sdkvalidationerror.js";

export const LimitedBy = {
  Scope: "scope",
  Mfa: "mfa",
} as const;
export type LimitedBy = ClosedEnum<typeof LimitedBy>;

/**
 * Information for the SAML Single Sign-On configuration.
 */
export type TeamLimitedConnection = {
  /**
   * The Identity Provider "type", for example Okta.
   */
  type: string;
  /**
   * Current status of the connection.
   */
  status: string;
  /**
   * Current state of the connection.
   */
  state: string;
  /**
   * Timestamp (in milliseconds) of when the configuration was connected.
   */
  connectedAt: number;
  /**
   * Timestamp (in milliseconds) of when the last webhook event was received from WorkOS.
   */
  lastReceivedWebhookEvent?: number | undefined;
};

/**
 * Information for the Directory Sync configuration.
 */
export type TeamLimitedDirectory = {
  /**
   * The Identity Provider "type", for example Okta.
   */
  type: string;
  /**
   * Current state of the connection.
   */
  state: string;
  /**
   * Timestamp (in milliseconds) of when the configuration was connected.
   */
  connectedAt: number;
  /**
   * Timestamp (in milliseconds) of when the last webhook event was received from WorkOS.
   */
  lastReceivedWebhookEvent?: number | undefined;
};

/**
 * When "Single Sign-On (SAML)" is configured, this object contains information that allows the client-side to identify whether or not this Team has SAML enforced.
 */
export type TeamLimitedSaml = {
  /**
   * Information for the SAML Single Sign-On configuration.
   */
  connection?: TeamLimitedConnection | undefined;
  /**
   * Information for the Directory Sync configuration.
   */
  directory?: TeamLimitedDirectory | undefined;
  /**
   * When `true`, interactions with the Team **must** be done with an authentication token that has been authenticated with the Team's SAML Single Sign-On provider.
   */
  enforced: boolean;
};

export type TeamLimitedEntitlements = {
  entitlement: string;
};

export const TeamLimitedRole = {
  Owner: "OWNER",
  Member: "MEMBER",
  Developer: "DEVELOPER",
  Security: "SECURITY",
  Billing: "BILLING",
  Viewer: "VIEWER",
  Contributor: "CONTRIBUTOR",
} as const;
export type TeamLimitedRole = ClosedEnum<typeof TeamLimitedRole>;

export const TeamLimitedTeamRoles = {
  Owner: "OWNER",
  Member: "MEMBER",
  Developer: "DEVELOPER",
  Security: "SECURITY",
  Billing: "BILLING",
  Viewer: "VIEWER",
  Contributor: "CONTRIBUTOR",
} as const;
export type TeamLimitedTeamRoles = ClosedEnum<typeof TeamLimitedTeamRoles>;

export const TeamLimitedTeamPermissions = {
  CreateProject: "CreateProject",
  FullProductionDeployment: "FullProductionDeployment",
  UsageViewer: "UsageViewer",
  EnvVariableManager: "EnvVariableManager",
  EnvironmentManager: "EnvironmentManager",
} as const;
export type TeamLimitedTeamPermissions = ClosedEnum<
  typeof TeamLimitedTeamPermissions
>;

export const TeamLimitedOrigin = {
  Link: "link",
  Saml: "saml",
  Mail: "mail",
  Import: "import",
  Teams: "teams",
  Github: "github",
  Gitlab: "gitlab",
  Bitbucket: "bitbucket",
  Dsync: "dsync",
  Feedback: "feedback",
  OrganizationTeams: "organization-teams",
} as const;
export type TeamLimitedOrigin = ClosedEnum<typeof TeamLimitedOrigin>;

export type TeamLimitedGitUserId = string | number;

export type TeamLimitedJoinedFrom = {
  origin: TeamLimitedOrigin;
  commitId?: string | undefined;
  repoId?: string | undefined;
  repoPath?: string | undefined;
  gitUserId?: string | number | undefined;
  gitUserLogin?: string | undefined;
  ssoUserId?: string | undefined;
  ssoConnectedAt?: number | undefined;
  idpUserId?: string | undefined;
  dsyncUserId?: string | undefined;
  dsyncConnectedAt?: number | undefined;
};

/**
 * The membership of the authenticated User in relation to the Team.
 */
export type TeamLimitedMembership = {
  uid?: string | undefined;
  entitlements?: Array<TeamLimitedEntitlements> | undefined;
  teamId?: string | undefined;
  confirmed: boolean;
  confirmedAt: number;
  accessRequestedAt?: number | undefined;
  role: TeamLimitedRole;
  teamRoles?: Array<TeamLimitedTeamRoles> | undefined;
  teamPermissions?: Array<TeamLimitedTeamPermissions> | undefined;
  createdAt: number;
  created: number;
  joinedFrom?: TeamLimitedJoinedFrom | undefined;
};

/**
 * A limited form of data representing a Team, due to the authentication token missing privileges to read the full Team data.
 */
export type TeamLimited = {
  /**
   * Property indicating that this Team data contains only limited information, due to the authentication token missing privileges to read the full Team data or due to team having MFA enforced and the user not having MFA enabled. Re-login with the Team's configured SAML Single Sign-On provider in order to upgrade the authentication token with the necessary privileges.
   */
  limited: boolean;
  limitedBy: Array<LimitedBy>;
  /**
   * When "Single Sign-On (SAML)" is configured, this object contains information that allows the client-side to identify whether or not this Team has SAML enforced.
   */
  saml?: TeamLimitedSaml | undefined;
  /**
   * The Team's unique identifier.
   */
  id: string;
  /**
   * The Team's slug, which is unique across the Vercel platform.
   */
  slug: string;
  /**
   * Name associated with the Team account, or `null` if none has been provided.
   */
  name: string | null;
  /**
   * The ID of the file used as avatar for this Team.
   */
  avatar: string | null;
  /**
   * The membership of the authenticated User in relation to the Team.
   */
  membership: TeamLimitedMembership;
  /**
   * UNIX timestamp (in milliseconds) when the Team was created.
   */
  createdAt: number;
};

/** @internal */
export const LimitedBy$inboundSchema: z.ZodNativeEnum<typeof LimitedBy> = z
  .nativeEnum(LimitedBy);

/** @internal */
export const LimitedBy$outboundSchema: z.ZodNativeEnum<typeof LimitedBy> =
  LimitedBy$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace LimitedBy$ {
  /** @deprecated use `LimitedBy$inboundSchema` instead. */
  export const inboundSchema = LimitedBy$inboundSchema;
  /** @deprecated use `LimitedBy$outboundSchema` instead. */
  export const outboundSchema = LimitedBy$outboundSchema;
}

/** @internal */
export const TeamLimitedConnection$inboundSchema: z.ZodType<
  TeamLimitedConnection,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: z.string(),
  status: z.string(),
  state: z.string(),
  connectedAt: z.number(),
  lastReceivedWebhookEvent: z.number().optional(),
});

/** @internal */
export type TeamLimitedConnection$Outbound = {
  type: string;
  status: string;
  state: string;
  connectedAt: number;
  lastReceivedWebhookEvent?: number | undefined;
};

/** @internal */
export const TeamLimitedConnection$outboundSchema: z.ZodType<
  TeamLimitedConnection$Outbound,
  z.ZodTypeDef,
  TeamLimitedConnection
> = z.object({
  type: z.string(),
  status: z.string(),
  state: z.string(),
  connectedAt: z.number(),
  lastReceivedWebhookEvent: z.number().optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedConnection$ {
  /** @deprecated use `TeamLimitedConnection$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedConnection$inboundSchema;
  /** @deprecated use `TeamLimitedConnection$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedConnection$outboundSchema;
  /** @deprecated use `TeamLimitedConnection$Outbound` instead. */
  export type Outbound = TeamLimitedConnection$Outbound;
}

export function teamLimitedConnectionToJSON(
  teamLimitedConnection: TeamLimitedConnection,
): string {
  return JSON.stringify(
    TeamLimitedConnection$outboundSchema.parse(teamLimitedConnection),
  );
}

export function teamLimitedConnectionFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimitedConnection, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimitedConnection$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimitedConnection' from JSON`,
  );
}

/** @internal */
export const TeamLimitedDirectory$inboundSchema: z.ZodType<
  TeamLimitedDirectory,
  z.ZodTypeDef,
  unknown
> = z.object({
  type: z.string(),
  state: z.string(),
  connectedAt: z.number(),
  lastReceivedWebhookEvent: z.number().optional(),
});

/** @internal */
export type TeamLimitedDirectory$Outbound = {
  type: string;
  state: string;
  connectedAt: number;
  lastReceivedWebhookEvent?: number | undefined;
};

/** @internal */
export const TeamLimitedDirectory$outboundSchema: z.ZodType<
  TeamLimitedDirectory$Outbound,
  z.ZodTypeDef,
  TeamLimitedDirectory
> = z.object({
  type: z.string(),
  state: z.string(),
  connectedAt: z.number(),
  lastReceivedWebhookEvent: z.number().optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedDirectory$ {
  /** @deprecated use `TeamLimitedDirectory$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedDirectory$inboundSchema;
  /** @deprecated use `TeamLimitedDirectory$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedDirectory$outboundSchema;
  /** @deprecated use `TeamLimitedDirectory$Outbound` instead. */
  export type Outbound = TeamLimitedDirectory$Outbound;
}

export function teamLimitedDirectoryToJSON(
  teamLimitedDirectory: TeamLimitedDirectory,
): string {
  return JSON.stringify(
    TeamLimitedDirectory$outboundSchema.parse(teamLimitedDirectory),
  );
}

export function teamLimitedDirectoryFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimitedDirectory, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimitedDirectory$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimitedDirectory' from JSON`,
  );
}

/** @internal */
export const TeamLimitedSaml$inboundSchema: z.ZodType<
  TeamLimitedSaml,
  z.ZodTypeDef,
  unknown
> = z.object({
  connection: z.lazy(() => TeamLimitedConnection$inboundSchema).optional(),
  directory: z.lazy(() => TeamLimitedDirectory$inboundSchema).optional(),
  enforced: z.boolean(),
});

/** @internal */
export type TeamLimitedSaml$Outbound = {
  connection?: TeamLimitedConnection$Outbound | undefined;
  directory?: TeamLimitedDirectory$Outbound | undefined;
  enforced: boolean;
};

/** @internal */
export const TeamLimitedSaml$outboundSchema: z.ZodType<
  TeamLimitedSaml$Outbound,
  z.ZodTypeDef,
  TeamLimitedSaml
> = z.object({
  connection: z.lazy(() => TeamLimitedConnection$outboundSchema).optional(),
  directory: z.lazy(() => TeamLimitedDirectory$outboundSchema).optional(),
  enforced: z.boolean(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedSaml$ {
  /** @deprecated use `TeamLimitedSaml$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedSaml$inboundSchema;
  /** @deprecated use `TeamLimitedSaml$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedSaml$outboundSchema;
  /** @deprecated use `TeamLimitedSaml$Outbound` instead. */
  export type Outbound = TeamLimitedSaml$Outbound;
}

export function teamLimitedSamlToJSON(
  teamLimitedSaml: TeamLimitedSaml,
): string {
  return JSON.stringify(TeamLimitedSaml$outboundSchema.parse(teamLimitedSaml));
}

export function teamLimitedSamlFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimitedSaml, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimitedSaml$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimitedSaml' from JSON`,
  );
}

/** @internal */
export const TeamLimitedEntitlements$inboundSchema: z.ZodType<
  TeamLimitedEntitlements,
  z.ZodTypeDef,
  unknown
> = z.object({
  entitlement: z.string(),
});

/** @internal */
export type TeamLimitedEntitlements$Outbound = {
  entitlement: string;
};

/** @internal */
export const TeamLimitedEntitlements$outboundSchema: z.ZodType<
  TeamLimitedEntitlements$Outbound,
  z.ZodTypeDef,
  TeamLimitedEntitlements
> = z.object({
  entitlement: z.string(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedEntitlements$ {
  /** @deprecated use `TeamLimitedEntitlements$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedEntitlements$inboundSchema;
  /** @deprecated use `TeamLimitedEntitlements$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedEntitlements$outboundSchema;
  /** @deprecated use `TeamLimitedEntitlements$Outbound` instead. */
  export type Outbound = TeamLimitedEntitlements$Outbound;
}

export function teamLimitedEntitlementsToJSON(
  teamLimitedEntitlements: TeamLimitedEntitlements,
): string {
  return JSON.stringify(
    TeamLimitedEntitlements$outboundSchema.parse(teamLimitedEntitlements),
  );
}

export function teamLimitedEntitlementsFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimitedEntitlements, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimitedEntitlements$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimitedEntitlements' from JSON`,
  );
}

/** @internal */
export const TeamLimitedRole$inboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedRole
> = z.nativeEnum(TeamLimitedRole);

/** @internal */
export const TeamLimitedRole$outboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedRole
> = TeamLimitedRole$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedRole$ {
  /** @deprecated use `TeamLimitedRole$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedRole$inboundSchema;
  /** @deprecated use `TeamLimitedRole$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedRole$outboundSchema;
}

/** @internal */
export const TeamLimitedTeamRoles$inboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedTeamRoles
> = z.nativeEnum(TeamLimitedTeamRoles);

/** @internal */
export const TeamLimitedTeamRoles$outboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedTeamRoles
> = TeamLimitedTeamRoles$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedTeamRoles$ {
  /** @deprecated use `TeamLimitedTeamRoles$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedTeamRoles$inboundSchema;
  /** @deprecated use `TeamLimitedTeamRoles$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedTeamRoles$outboundSchema;
}

/** @internal */
export const TeamLimitedTeamPermissions$inboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedTeamPermissions
> = z.nativeEnum(TeamLimitedTeamPermissions);

/** @internal */
export const TeamLimitedTeamPermissions$outboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedTeamPermissions
> = TeamLimitedTeamPermissions$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedTeamPermissions$ {
  /** @deprecated use `TeamLimitedTeamPermissions$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedTeamPermissions$inboundSchema;
  /** @deprecated use `TeamLimitedTeamPermissions$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedTeamPermissions$outboundSchema;
}

/** @internal */
export const TeamLimitedOrigin$inboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedOrigin
> = z.nativeEnum(TeamLimitedOrigin);

/** @internal */
export const TeamLimitedOrigin$outboundSchema: z.ZodNativeEnum<
  typeof TeamLimitedOrigin
> = TeamLimitedOrigin$inboundSchema;

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedOrigin$ {
  /** @deprecated use `TeamLimitedOrigin$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedOrigin$inboundSchema;
  /** @deprecated use `TeamLimitedOrigin$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedOrigin$outboundSchema;
}

/** @internal */
export const TeamLimitedGitUserId$inboundSchema: z.ZodType<
  TeamLimitedGitUserId,
  z.ZodTypeDef,
  unknown
> = z.union([z.string(), z.number()]);

/** @internal */
export type TeamLimitedGitUserId$Outbound = string | number;

/** @internal */
export const TeamLimitedGitUserId$outboundSchema: z.ZodType<
  TeamLimitedGitUserId$Outbound,
  z.ZodTypeDef,
  TeamLimitedGitUserId
> = z.union([z.string(), z.number()]);

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedGitUserId$ {
  /** @deprecated use `TeamLimitedGitUserId$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedGitUserId$inboundSchema;
  /** @deprecated use `TeamLimitedGitUserId$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedGitUserId$outboundSchema;
  /** @deprecated use `TeamLimitedGitUserId$Outbound` instead. */
  export type Outbound = TeamLimitedGitUserId$Outbound;
}

export function teamLimitedGitUserIdToJSON(
  teamLimitedGitUserId: TeamLimitedGitUserId,
): string {
  return JSON.stringify(
    TeamLimitedGitUserId$outboundSchema.parse(teamLimitedGitUserId),
  );
}

export function teamLimitedGitUserIdFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimitedGitUserId, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimitedGitUserId$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimitedGitUserId' from JSON`,
  );
}

/** @internal */
export const TeamLimitedJoinedFrom$inboundSchema: z.ZodType<
  TeamLimitedJoinedFrom,
  z.ZodTypeDef,
  unknown
> = z.object({
  origin: TeamLimitedOrigin$inboundSchema,
  commitId: z.string().optional(),
  repoId: z.string().optional(),
  repoPath: z.string().optional(),
  gitUserId: z.union([z.string(), z.number()]).optional(),
  gitUserLogin: z.string().optional(),
  ssoUserId: z.string().optional(),
  ssoConnectedAt: z.number().optional(),
  idpUserId: z.string().optional(),
  dsyncUserId: z.string().optional(),
  dsyncConnectedAt: z.number().optional(),
});

/** @internal */
export type TeamLimitedJoinedFrom$Outbound = {
  origin: string;
  commitId?: string | undefined;
  repoId?: string | undefined;
  repoPath?: string | undefined;
  gitUserId?: string | number | undefined;
  gitUserLogin?: string | undefined;
  ssoUserId?: string | undefined;
  ssoConnectedAt?: number | undefined;
  idpUserId?: string | undefined;
  dsyncUserId?: string | undefined;
  dsyncConnectedAt?: number | undefined;
};

/** @internal */
export const TeamLimitedJoinedFrom$outboundSchema: z.ZodType<
  TeamLimitedJoinedFrom$Outbound,
  z.ZodTypeDef,
  TeamLimitedJoinedFrom
> = z.object({
  origin: TeamLimitedOrigin$outboundSchema,
  commitId: z.string().optional(),
  repoId: z.string().optional(),
  repoPath: z.string().optional(),
  gitUserId: z.union([z.string(), z.number()]).optional(),
  gitUserLogin: z.string().optional(),
  ssoUserId: z.string().optional(),
  ssoConnectedAt: z.number().optional(),
  idpUserId: z.string().optional(),
  dsyncUserId: z.string().optional(),
  dsyncConnectedAt: z.number().optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedJoinedFrom$ {
  /** @deprecated use `TeamLimitedJoinedFrom$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedJoinedFrom$inboundSchema;
  /** @deprecated use `TeamLimitedJoinedFrom$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedJoinedFrom$outboundSchema;
  /** @deprecated use `TeamLimitedJoinedFrom$Outbound` instead. */
  export type Outbound = TeamLimitedJoinedFrom$Outbound;
}

export function teamLimitedJoinedFromToJSON(
  teamLimitedJoinedFrom: TeamLimitedJoinedFrom,
): string {
  return JSON.stringify(
    TeamLimitedJoinedFrom$outboundSchema.parse(teamLimitedJoinedFrom),
  );
}

export function teamLimitedJoinedFromFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimitedJoinedFrom, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimitedJoinedFrom$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimitedJoinedFrom' from JSON`,
  );
}

/** @internal */
export const TeamLimitedMembership$inboundSchema: z.ZodType<
  TeamLimitedMembership,
  z.ZodTypeDef,
  unknown
> = z.object({
  uid: z.string().optional(),
  entitlements: z.array(z.lazy(() => TeamLimitedEntitlements$inboundSchema))
    .optional(),
  teamId: z.string().optional(),
  confirmed: z.boolean(),
  confirmedAt: z.number(),
  accessRequestedAt: z.number().optional(),
  role: TeamLimitedRole$inboundSchema,
  teamRoles: z.array(TeamLimitedTeamRoles$inboundSchema).optional(),
  teamPermissions: z.array(TeamLimitedTeamPermissions$inboundSchema).optional(),
  createdAt: z.number(),
  created: z.number(),
  joinedFrom: z.lazy(() => TeamLimitedJoinedFrom$inboundSchema).optional(),
});

/** @internal */
export type TeamLimitedMembership$Outbound = {
  uid?: string | undefined;
  entitlements?: Array<TeamLimitedEntitlements$Outbound> | undefined;
  teamId?: string | undefined;
  confirmed: boolean;
  confirmedAt: number;
  accessRequestedAt?: number | undefined;
  role: string;
  teamRoles?: Array<string> | undefined;
  teamPermissions?: Array<string> | undefined;
  createdAt: number;
  created: number;
  joinedFrom?: TeamLimitedJoinedFrom$Outbound | undefined;
};

/** @internal */
export const TeamLimitedMembership$outboundSchema: z.ZodType<
  TeamLimitedMembership$Outbound,
  z.ZodTypeDef,
  TeamLimitedMembership
> = z.object({
  uid: z.string().optional(),
  entitlements: z.array(z.lazy(() => TeamLimitedEntitlements$outboundSchema))
    .optional(),
  teamId: z.string().optional(),
  confirmed: z.boolean(),
  confirmedAt: z.number(),
  accessRequestedAt: z.number().optional(),
  role: TeamLimitedRole$outboundSchema,
  teamRoles: z.array(TeamLimitedTeamRoles$outboundSchema).optional(),
  teamPermissions: z.array(TeamLimitedTeamPermissions$outboundSchema)
    .optional(),
  createdAt: z.number(),
  created: z.number(),
  joinedFrom: z.lazy(() => TeamLimitedJoinedFrom$outboundSchema).optional(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimitedMembership$ {
  /** @deprecated use `TeamLimitedMembership$inboundSchema` instead. */
  export const inboundSchema = TeamLimitedMembership$inboundSchema;
  /** @deprecated use `TeamLimitedMembership$outboundSchema` instead. */
  export const outboundSchema = TeamLimitedMembership$outboundSchema;
  /** @deprecated use `TeamLimitedMembership$Outbound` instead. */
  export type Outbound = TeamLimitedMembership$Outbound;
}

export function teamLimitedMembershipToJSON(
  teamLimitedMembership: TeamLimitedMembership,
): string {
  return JSON.stringify(
    TeamLimitedMembership$outboundSchema.parse(teamLimitedMembership),
  );
}

export function teamLimitedMembershipFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimitedMembership, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimitedMembership$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimitedMembership' from JSON`,
  );
}

/** @internal */
export const TeamLimited$inboundSchema: z.ZodType<
  TeamLimited,
  z.ZodTypeDef,
  unknown
> = z.object({
  limited: z.boolean(),
  limitedBy: z.array(LimitedBy$inboundSchema),
  saml: z.lazy(() => TeamLimitedSaml$inboundSchema).optional(),
  id: z.string(),
  slug: z.string(),
  name: z.nullable(z.string()),
  avatar: z.nullable(z.string()),
  membership: z.lazy(() => TeamLimitedMembership$inboundSchema),
  createdAt: z.number(),
});

/** @internal */
export type TeamLimited$Outbound = {
  limited: boolean;
  limitedBy: Array<string>;
  saml?: TeamLimitedSaml$Outbound | undefined;
  id: string;
  slug: string;
  name: string | null;
  avatar: string | null;
  membership: TeamLimitedMembership$Outbound;
  createdAt: number;
};

/** @internal */
export const TeamLimited$outboundSchema: z.ZodType<
  TeamLimited$Outbound,
  z.ZodTypeDef,
  TeamLimited
> = z.object({
  limited: z.boolean(),
  limitedBy: z.array(LimitedBy$outboundSchema),
  saml: z.lazy(() => TeamLimitedSaml$outboundSchema).optional(),
  id: z.string(),
  slug: z.string(),
  name: z.nullable(z.string()),
  avatar: z.nullable(z.string()),
  membership: z.lazy(() => TeamLimitedMembership$outboundSchema),
  createdAt: z.number(),
});

/**
 * @internal
 * @deprecated This namespace will be removed in future versions. Use schemas and types that are exported directly from this module.
 */
export namespace TeamLimited$ {
  /** @deprecated use `TeamLimited$inboundSchema` instead. */
  export const inboundSchema = TeamLimited$inboundSchema;
  /** @deprecated use `TeamLimited$outboundSchema` instead. */
  export const outboundSchema = TeamLimited$outboundSchema;
  /** @deprecated use `TeamLimited$Outbound` instead. */
  export type Outbound = TeamLimited$Outbound;
}

export function teamLimitedToJSON(teamLimited: TeamLimited): string {
  return JSON.stringify(TeamLimited$outboundSchema.parse(teamLimited));
}

export function teamLimitedFromJSON(
  jsonString: string,
): SafeParseResult<TeamLimited, SDKValidationError> {
  return safeParse(
    jsonString,
    (x) => TeamLimited$inboundSchema.parse(JSON.parse(x)),
    `Failed to parse 'TeamLimited' from JSON`,
  );
}
