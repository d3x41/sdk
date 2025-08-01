/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { expect, test } from "vitest";
import { Vercel } from "../index.js";
import { GetDeploymentEventsDeploymentsResponseBody } from "../models/getdeploymenteventsop.js";
import { createTestHTTPClient } from "./testclient.js";

test("Deployments Get Deployment Events", async () => {
  const testHttpClient = createTestHTTPClient("getDeploymentEvents");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.getDeploymentEvents({
    idOrUrl: "dpl_5WJWYSyB7BpgTj3EuwF37WMRBXBtPQ2iTMJHJBJyRfd",
    direction: "backward",
    follow: 1,
    limit: 100,
    name: "bld_cotnkcr76",
    since: 1540095775941,
    until: 1540106318643,
    statusCode: "5xx",
    delimiter: 1,
    builds: 1,
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result as GetDeploymentEventsDeploymentsResponseBody[]).toEqual([
    {
      created: 9364.53,
      date: 65.47,
      deploymentId: "<id>",
      id: "<id>",
      info: {
        type: "<value>",
        name: "<value>",
      },
      serial: "<value>",
      type: "command",
    },
    {
      type: "stderr",
      created: 2829.12,
      payload: {
        deploymentId: "<id>",
        id: "<id>",
        date: 7964.23,
        serial: "<value>",
      },
    },
    {
      type: "stderr",
      created: 2829.12,
      payload: {
        deploymentId: "<id>",
        id: "<id>",
        date: 7964.23,
        serial: "<value>",
      },
    },
  ]);
});

test("Deployments Update Integration Deployment Action", async () => {
  const testHttpClient = createTestHTTPClient(
    "update-integration-deployment-action",
  );

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  await vercel.deployments.updateIntegrationDeploymentAction({
    deploymentId: "<id>",
    integrationConfigurationId: "<id>",
    resourceId: "<id>",
    action: "<value>",
  });
});

test("Deployments Get Deployment", async () => {
  const testHttpClient = createTestHTTPClient("getDeployment");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.getDeployment({
    idOrUrl: "dpl_89qyp1cskzkLrVicDaZoDbjyHuDJ",
    withGitRepoInfo: "true",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    aliasAssigned: true,
    bootedAt: 2363.88,
    buildingAt: 7418.28,
    buildSkipped: false,
    creator: {
      uid: "<id>",
    },
    public: false,
    status: "BUILDING",
    id: "<id>",
    name: "<value>",
    type: "LAMBDAS",
    createdAt: 548.4,
    readyState: "ERROR",
    meta: {
      "key": "<value>",
      "key1": "<value>",
    },
    regions: [
      "<value 1>",
      "<value 2>",
      "<value 3>",
    ],
    url: "https://jagged-pacemaker.biz/",
    version: 5784.12,
  });
});

test("Deployments Create Deployment", async () => {
  const testHttpClient = createTestHTTPClient("createDeployment");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.createDeployment({
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
    requestBody: {
      deploymentId: "dpl_2qn7PZrx89yxY34vEZPD31Y9XVj6",
      files: [
        {
          file: "folder/file.js",
        },
      ],
      gitMetadata: {
        remoteUrl: "https://github.com/vercel/next.js",
        commitAuthorName: "kyliau",
        commitMessage:
          "add method to measure Interaction to Next Paint (INP) (#36490)",
        commitRef: "main",
        commitSha: "dc36199b2234c6586ebe05ec94078a895c707e29",
        dirty: true,
      },
      gitSource: {
        org: "vercel",
        ref: "main",
        repo: "next.js",
        sha: "a1b2c3d4e5f6g7h8i9j0k1l2m3n4o5p6q7r8s9t0",
        type: "github",
      },
      meta: {
        "foo": "bar",
      },
      name: "my-instant-deployment",
      project: "my-deployment-project",
      projectSettings: {
        buildCommand: "next build",
        installCommand: "pnpm install",
      },
      target: "production",
    },
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    build: {
      env: [
        "<value 1>",
        "<value 2>",
        "<value 3>",
      ],
    },
    env: [],
    inspectorUrl: null,
    isInConcurrentBuildsQueue: false,
    isInSystemBuildsQueue: false,
    projectSettings: {},
    aliasAssigned: false,
    bootedAt: 3619.08,
    buildingAt: 1664.36,
    buildSkipped: false,
    creator: {
      uid: "<id>",
    },
    public: false,
    status: "CANCELED",
    id: "<id>",
    name: "<value>",
    createdAt: 179.28,
    type: "LAMBDAS",
    version: 8477.72,
    meta: {
      "key": "<value>",
    },
    readyState: "BUILDING",
    regions: [],
    url: "https://unwritten-viability.org",
    projectId: "<id>",
    ownerId: "<id>",
    plan: "enterprise",
    routes: [],
    createdIn: "<value>",
  });
});

test("Deployments Cancel Deployment", async () => {
  const testHttpClient = createTestHTTPClient("cancelDeployment");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.cancelDeployment({
    id: "dpl_5WJWYSyB7BpgTj3EuwF37WMRBXBtPQ2iTMJHJBJyRfd",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    build: {
      env: [
        "<value 1>",
        "<value 2>",
      ],
    },
    env: [],
    inspectorUrl: "https://grown-gymnast.net",
    isInConcurrentBuildsQueue: false,
    isInSystemBuildsQueue: false,
    projectSettings: {},
    aliasAssigned: false,
    bootedAt: 9923.4,
    buildingAt: 4182.97,
    buildSkipped: true,
    creator: {
      uid: "<id>",
    },
    public: false,
    status: "CANCELED",
    id: "<id>",
    createdAt: 4076.8,
    name: "<value>",
    meta: {},
    readyState: "READY",
    regions: [],
    type: "LAMBDAS",
    url: "https://medium-tribe.org/",
    version: 7316.38,
    createdIn: "<value>",
    ownerId: "<id>",
    plan: "hobby",
    projectId: "<id>",
    routes: [
      {
        src: "<value>",
      },
      {
        src: "<value>",
        continue: true,
        middleware: 1635.94,
      },
    ],
  });
});

test("Deployments Upload File", async () => {
  const testHttpClient = createTestHTTPClient("uploadFile");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.uploadFile({
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
});

test("Deployments List Deployment Files", async () => {
  const testHttpClient = createTestHTTPClient("listDeploymentFiles");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.listDeploymentFiles({
    id: "<id>",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual([
    {
      name: "my-file.json",
      type: "file",
      uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
      children: [
        {
          name: "my-file.json",
          type: "file",
          uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
          contentType: "application/json",
          mode: 7131.4,
        },
        {
          name: "my-file.json",
          type: "file",
          uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
          contentType: "application/json",
          mode: 7131.4,
        },
        {
          name: "my-file.json",
          type: "file",
          uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
          contentType: "application/json",
          mode: 7131.4,
        },
      ],
      contentType: "application/json",
      mode: 6417.18,
    },
    {
      name: "my-file.json",
      type: "file",
      uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
      children: [
        {
          name: "my-file.json",
          type: "file",
          uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
          contentType: "application/json",
          mode: 7131.4,
        },
        {
          name: "my-file.json",
          type: "file",
          uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
          contentType: "application/json",
          mode: 7131.4,
        },
        {
          name: "my-file.json",
          type: "file",
          uid: "2d4aad419917f15b1146e9e03ddc9bb31747e4d0",
          contentType: "application/json",
          mode: 7131.4,
        },
      ],
      contentType: "application/json",
      mode: 6417.18,
    },
  ]);
});

test("Deployments Get Deployment File Contents", async () => {
  const testHttpClient = createTestHTTPClient("getDeploymentFileContents");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  await vercel.deployments.getDeploymentFileContents({
    id: "<id>",
    fileId: "<id>",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
});

test("Deployments Get Deployments", async () => {
  const testHttpClient = createTestHTTPClient("getDeployments");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.getDeployments({
    app: "docs",
    from: 1612948664566,
    limit: 10,
    projectId: "QmXGTs7mvAMMC7WW5ebrM33qKG32QK3h4vmQMjmY",
    target: "production",
    to: 1612948664566,
    users: "kr1PsOIzqEL5Xg6M4VZcZosf,K4amb7K9dAt5R2vBJWF32bmY",
    since: 1540095775941,
    until: 1540095775951,
    state: "BUILDING,READY",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    pagination: {
      count: 20,
      next: 1540095775951,
      prev: 1540095775951,
    },
    deployments: [],
  });
});

test("Deployments Delete Deployment", async () => {
  const testHttpClient = createTestHTTPClient("deleteDeployment");

  const vercel = new Vercel({
    serverURL: process.env["TEST_SERVER_URL"] ?? "http://localhost:18080",
    httpClient: testHttpClient,
    bearerToken: "<YOUR_BEARER_TOKEN_HERE>",
  });

  const result = await vercel.deployments.deleteDeployment({
    id: "dpl_5WJWYSyB7BpgTj3EuwF37WMRBXBtPQ2iTMJHJBJyRfd",
    url: "https://files-orcin-xi.vercel.app/",
    teamId: "team_1a2b3c4d5e6f7g8h9i0j1k2l",
    slug: "my-team-url-slug",
  });
  expect(result).toBeDefined();
  expect(result).toEqual({
    uid: "dpl_5WJWYSyB7BpgTj3EuwF37WMRBXBtPQ2iTMJHJBJyRfd",
    state: "DELETED",
  });
});
