from os import path


def _check_if_build_folder_is_ignored(project_root, out=None):
  gitignore_path=path.join(project_root, ".gitignore")
  if not path.exists(gitignore_path):
    if out:
      out.warning("No .gitignore in project root. You should probably create one.")
      return False
  return True

def _run_checks(project_root, out=None):
  _check_if_build_folder_is_ignored(project_root, out)

def generate_build(project_root, out=None):
  _run_checks(project_root, out)
