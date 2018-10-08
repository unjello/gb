import click

_verbosity_levels = {
  'info': 1,
  'debug': 2,
  'trace': 3
}

class Out:
  def __init__(self, verbosity):
    self.verbosity = verbosity

  def _can(self, level):
    return self.verbosity >= _verbosity_levels[level]
  
  def write(self, data, **style):
    click.secho(data, **style)

  def info(self, data, **style):
    if self._can('info'):
      click.secho("info ", fg='green', nl=False)
      click.secho(data, **style)

  def debug(self, data, **style):
    if self._can('debug'):
      click.secho("debg ", fg='green', nl=False)
      click.secho(data, **style)

  def trace(self, data, **style):
    if self._can('trace'):
      click.secho("trac ", fg='white', dim=True, nl=False)
      click.secho(data, **style)

  def warning(self, data, **style):
    click.secho("warn ", fg='yellow', nl=False)
    click.secho(data, **style)
