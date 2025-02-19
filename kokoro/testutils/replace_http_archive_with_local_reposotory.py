#!/usr/bin/env python3

# Copyright 2022 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
r"""Utility that replaces http_archive with local_repository in WORKSPACE files.

Usage:
  ./kokoro/testutils/replace_http_archive_with_local_reposotory.py \
    -f <workspace directory> \
    -t <tink local base path>

For examples:
  ./kokoro/testutils/replace_http_archive_with_local_reposotory.py \
    -f "cc/WORKSPACE" \
    -t "../../tink"
"""

import argparse
import textwrap


def _replace_http_archive_with_local_repository(workspace_content: str,
                                                tink_base_path: str) -> None:
  """Replaces http_archive with local_repository in workspace_content.

  Args:
    workspace_content: Content of the WORKSPACE file to modify.
    tink_base_path: Path to the local Tink folder.

  Returns:
    The modified WORKSPACE file content.
  """
  # Remove loading of http_archive.
  http_archive_load = textwrap.dedent("""\
      load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
      """)
  workspace_content = workspace_content.replace(http_archive_load, '')

  # Tink Base.
  tink_base_before = textwrap.dedent("""\
      http_archive(
          name = "tink_base",
          urls = ["https://github.com/google/tink/archive/master.zip"],
          strip_prefix = "tink-master/",
      )""")
  tink_base_after = textwrap.dedent("""\
      local_repository(
          name = "tink_base",
          path = "{}",
      )""".format(tink_base_path))
  workspace_content = workspace_content.replace(tink_base_before,
                                                tink_base_after)

  # Tink C++.
  tink_cc_before = textwrap.dedent("""\
      http_archive(
          name = "tink_cc",
          urls = ["https://github.com/google/tink/archive/master.zip"],
          strip_prefix = "tink-master/cc",
      )""")
  tink_cc_after = textwrap.dedent("""\
      local_repository(
          name = "tink_cc",
          path = "{}/cc",
      )""".format(tink_base_path))
  workspace_content = workspace_content.replace(tink_cc_before, tink_cc_after)

  return workspace_content


def main():
  parser = argparse.ArgumentParser(
      description='Replaces http_archive rules with local_repository rules')
  parser.add_argument('--workspace_file_path', '-f', required=True)
  parser.add_argument('--tink_base_path', '-t', required=True)
  args = parser.parse_args()

  with open(args.workspace_file_path, 'r') as workspace_file:
    content = workspace_file.read()
    content = _replace_http_archive_with_local_repository(
        content, args.tink_base_path)
  with open(args.workspace_file_path, 'w') as workspace_file:
    workspace_file.write(content)


if __name__ == '__main__':
  main()
