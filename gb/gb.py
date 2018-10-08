import click
import six
from output import Out

@click.command()
@click.option('-v', '--verbose', count=True)
def run(verbose):
    """
    CLI for yet another build generator for C++
    """
    out = Out(verbose)
    out.write("gb %s" % "0.0.1", bold=True)    
    

if __name__ == '__main__':
    from os import sys, path
    sys.path.append(path.dirname(path.dirname(path.abspath(__file__))))
    run()
