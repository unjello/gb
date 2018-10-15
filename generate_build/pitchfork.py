import re
import click
from os import path


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
  _check_if_build_folder_is_ignored(project_root, out)

def generate_build(project_root, out=None):
  _run_checks(project_root, out)
