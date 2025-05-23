/*
 * Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.
 */

import { accessGroupsCreateAccessGroup } from "../funcs/accessGroupsCreateAccessGroup.js";
import { accessGroupsCreateAccessGroupProject } from "../funcs/accessGroupsCreateAccessGroupProject.js";
import { accessGroupsDeleteAccessGroup } from "../funcs/accessGroupsDeleteAccessGroup.js";
import { accessGroupsDeleteAccessGroupProject } from "../funcs/accessGroupsDeleteAccessGroupProject.js";
import { accessGroupsListAccessGroupMembers } from "../funcs/accessGroupsListAccessGroupMembers.js";
import { accessGroupsListAccessGroupProjects } from "../funcs/accessGroupsListAccessGroupProjects.js";
import { accessGroupsListAccessGroups } from "../funcs/accessGroupsListAccessGroups.js";
import { accessGroupsReadAccessGroup } from "../funcs/accessGroupsReadAccessGroup.js";
import { accessGroupsReadAccessGroupProject } from "../funcs/accessGroupsReadAccessGroupProject.js";
import { accessGroupsUpdateAccessGroup } from "../funcs/accessGroupsUpdateAccessGroup.js";
import { accessGroupsUpdateAccessGroupProject } from "../funcs/accessGroupsUpdateAccessGroupProject.js";
import { ClientSDK, RequestOptions } from "../lib/sdks.js";
import {
  CreateAccessGroupRequest,
  CreateAccessGroupResponseBody,
} from "../models/createaccessgroupop.js";
import {
  CreateAccessGroupProjectRequest,
  CreateAccessGroupProjectResponseBody,
} from "../models/createaccessgroupprojectop.js";
import { DeleteAccessGroupRequest } from "../models/deleteaccessgroupop.js";
import { DeleteAccessGroupProjectRequest } from "../models/deleteaccessgroupprojectop.js";
import {
  ListAccessGroupMembersRequest,
  ListAccessGroupMembersResponseBody,
} from "../models/listaccessgroupmembersop.js";
import {
  ListAccessGroupProjectsRequest,
  ListAccessGroupProjectsResponseBody,
} from "../models/listaccessgroupprojectsop.js";
import {
  ListAccessGroupsRequest,
  ListAccessGroupsResponseBody,
} from "../models/listaccessgroupsop.js";
import {
  ReadAccessGroupRequest,
  ReadAccessGroupResponseBody,
} from "../models/readaccessgroupop.js";
import {
  ReadAccessGroupProjectRequest,
  ReadAccessGroupProjectResponseBody,
} from "../models/readaccessgroupprojectop.js";
import {
  UpdateAccessGroupRequest,
  UpdateAccessGroupResponseBody,
} from "../models/updateaccessgroupop.js";
import {
  UpdateAccessGroupProjectRequest,
  UpdateAccessGroupProjectResponseBody,
} from "../models/updateaccessgroupprojectop.js";
import { unwrapAsync } from "../types/fp.js";

export class AccessGroups extends ClientSDK {
  /**
   * Reads an access group
   *
   * @remarks
   * Allows to read an access group
   */
  async readAccessGroup(
    request: ReadAccessGroupRequest,
    options?: RequestOptions,
  ): Promise<ReadAccessGroupResponseBody> {
    return unwrapAsync(accessGroupsReadAccessGroup(
      this,
      request,
      options,
    ));
  }

  /**
   * Update an access group
   *
   * @remarks
   * Allows to update an access group metadata
   */
  async updateAccessGroup(
    request: UpdateAccessGroupRequest,
    options?: RequestOptions,
  ): Promise<UpdateAccessGroupResponseBody> {
    return unwrapAsync(accessGroupsUpdateAccessGroup(
      this,
      request,
      options,
    ));
  }

  /**
   * Deletes an access group
   *
   * @remarks
   * Allows to delete an access group
   */
  async deleteAccessGroup(
    request: DeleteAccessGroupRequest,
    options?: RequestOptions,
  ): Promise<void> {
    return unwrapAsync(accessGroupsDeleteAccessGroup(
      this,
      request,
      options,
    ));
  }

  /**
   * List members of an access group
   *
   * @remarks
   * List members of an access group
   */
  async listAccessGroupMembers(
    request: ListAccessGroupMembersRequest,
    options?: RequestOptions,
  ): Promise<ListAccessGroupMembersResponseBody> {
    return unwrapAsync(accessGroupsListAccessGroupMembers(
      this,
      request,
      options,
    ));
  }

  /**
   * List access groups for a team, project or member
   *
   * @remarks
   * List access groups
   */
  async listAccessGroups(
    request: ListAccessGroupsRequest,
    options?: RequestOptions,
  ): Promise<ListAccessGroupsResponseBody> {
    return unwrapAsync(accessGroupsListAccessGroups(
      this,
      request,
      options,
    ));
  }

  /**
   * Creates an access group
   *
   * @remarks
   * Allows to create an access group
   */
  async createAccessGroup(
    request: CreateAccessGroupRequest,
    options?: RequestOptions,
  ): Promise<CreateAccessGroupResponseBody> {
    return unwrapAsync(accessGroupsCreateAccessGroup(
      this,
      request,
      options,
    ));
  }

  /**
   * List projects of an access group
   *
   * @remarks
   * List projects of an access group
   */
  async listAccessGroupProjects(
    request: ListAccessGroupProjectsRequest,
    options?: RequestOptions,
  ): Promise<ListAccessGroupProjectsResponseBody> {
    return unwrapAsync(accessGroupsListAccessGroupProjects(
      this,
      request,
      options,
    ));
  }

  /**
   * Create an access group project
   *
   * @remarks
   * Allows creation of an access group project
   */
  async createAccessGroupProject(
    request: CreateAccessGroupProjectRequest,
    options?: RequestOptions,
  ): Promise<CreateAccessGroupProjectResponseBody> {
    return unwrapAsync(accessGroupsCreateAccessGroupProject(
      this,
      request,
      options,
    ));
  }

  /**
   * Reads an access group project
   *
   * @remarks
   * Allows reading an access group project
   */
  async readAccessGroupProject(
    request: ReadAccessGroupProjectRequest,
    options?: RequestOptions,
  ): Promise<ReadAccessGroupProjectResponseBody> {
    return unwrapAsync(accessGroupsReadAccessGroupProject(
      this,
      request,
      options,
    ));
  }

  /**
   * Update an access group project
   *
   * @remarks
   * Allows update of an access group project
   */
  async updateAccessGroupProject(
    request: UpdateAccessGroupProjectRequest,
    options?: RequestOptions,
  ): Promise<UpdateAccessGroupProjectResponseBody> {
    return unwrapAsync(accessGroupsUpdateAccessGroupProject(
      this,
      request,
      options,
    ));
  }

  /**
   * Delete an access group project
   *
   * @remarks
   * Allows deletion of an access group project
   */
  async deleteAccessGroupProject(
    request: DeleteAccessGroupProjectRequest,
    options?: RequestOptions,
  ): Promise<void> {
    return unwrapAsync(accessGroupsDeleteAccessGroupProject(
      this,
      request,
      options,
    ));
  }
}
