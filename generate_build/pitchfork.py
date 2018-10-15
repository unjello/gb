import re
import click
from os import path

def _check_if_src_folder_exists(project_root, out=None):
  src_path=path.join(project_root, "src")
  if not (path.exists(src_path) and path.isdir(src_path)):
    if out:
      out.warning("Source folder {} not found.".format(click.style("src", fg="cyan")))
    return False
  return True

def _check_if_build_folder_is_ignored(project_root, out=None):
  gitignore_path=path.join(project_root, ".gitignore")
  if not path.exists(gitignore_path):
    if out:
      out.warning("No {} in project root. You should create one.".format(click.style(".gitignore", fg="cyan")))
    return False
  with open(gitignore_path, 'r') as gitignore_file:
    content = gitignore_file.read()
    m = re.search('^build/?$', content, flags=re.MULTILINE)
    if not m:
      if out:
        out.warning("Your build folder {} should be ignored by git. Add it to {}.".format(click.style("build/", fg="cyan"), click.style(".gitignore", fg="cyan")))
      return False
  return True

def _run_checks(project_root, out=None):
  checks = [_check_if_src_folder_exists, _check_if_build_folder_is_ignored]
  for c in checks:
    if not c(project_root, out):
      return False
  return True

def generate_build(project_root, out=None):
  _run_checks(project_root, out)
