if __name__ == '__main__':
  from os import sys, path
  p=path.dirname(path.dirname(path.abspath(__file__)))
  sys.path.append(p)

import click
import six
from generate_build.output import Out
from generate_build import __version__ as gb_version

@click.command()
@click.option('-v', '--verbose', count=True)
def run(verbose):
  """
  CLI for yet another build generator for C++
  """
  out = Out(verbose)
  out.write("gb %s" % gb_version, bold=True)


if __name__ == '__main__':
  run()
