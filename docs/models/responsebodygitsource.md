# ResponseBodyGitSource


## Supported Types

### `models.GetDeploymentGitSource1`

```typescript
const value: models.GetDeploymentGitSource1 = {
  type: "github",
  repoId: "<id>",
};
```

### `models.GetDeploymentGitSource2`

```typescript
const value: models.GetDeploymentGitSource2 = {
  type: "github",
  org: "<value>",
  repo: "<value>",
};
```

### `models.GetDeploymentGitSource3`

```typescript
const value: models.GetDeploymentGitSource3 = {
  type: "gitlab",
  projectId: "<id>",
};
```

### `models.GetDeploymentGitSource4`

```typescript
const value: models.GetDeploymentGitSource4 = {
  type: "bitbucket",
  repoUuid: "<id>",
};
```

### `models.GetDeploymentGitSource5`

```typescript
const value: models.GetDeploymentGitSource5 = {
  type: "bitbucket",
  owner: "<value>",
  slug: "<value>",
};
```

### `models.GetDeploymentGitSource6`

```typescript
const value: models.GetDeploymentGitSource6 = {
  type: "custom",
  ref: "<value>",
  sha: "<value>",
  gitUrl: "https://tired-baseboard.biz/",
};
```

### `models.GetDeploymentGitSource7`

```typescript
const value: models.GetDeploymentGitSource7 = {
  type: "github",
  ref: "<value>",
  sha: "<value>",
  repoId: 1210.01,
};
```

### `models.GetDeploymentGitSource8`

```typescript
const value: models.GetDeploymentGitSource8 = {
  type: "gitlab",
  ref: "<value>",
  sha: "<value>",
  projectId: 1758.71,
};
```

### `models.GetDeploymentGitSource9`

```typescript
const value: models.GetDeploymentGitSource9 = {
  type: "bitbucket",
  ref: "<value>",
  sha: "<value>",
  workspaceUuid: "<id>",
  repoUuid: "<id>",
};
```

