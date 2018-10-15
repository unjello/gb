import os
from pitchfork import generate_build

if __name__ == '__main__':
  p=os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
  os.sys.path.append(p)

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
  working_dir = os.getcwd()
  generate_build(working_dir, out)



if __name__ == '__main__':
  run()
