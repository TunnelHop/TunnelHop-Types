/* eslint-disable */
/**
 * This file was automatically generated by json-schema-to-typescript.
 * DO NOT MODIFY IT BY HAND. Instead, modify the source JSONSchema file,
 * and run json-schema-to-typescript to regenerate this file.
 */

export interface Application {
  id?: number;
  user_email: string;
  protocol: string;
  internal_port: number;
  require_auth?: boolean;
  require_encryption?: boolean;
  tcp_port?: number;
  is_active?: boolean;
  application_name?: string;
  [k: string]: unknown;
}
