import click
import sys
import subprocess
import re
import json


def cmd_array(cmd):
    return cmd.split(" ")


def execute(cmd):
    p = subprocess.run(cmd_array(cmd), stdout=subprocess.PIPE)
    return p.stdout.decode("utf-8").strip().rstrip("\n")


def perror(msg):
    print(f"\033[91;1m{msg}\033[0m", file=sys.stderr)


def git_origin():
    return execute("git remote show")


def git_repo_url():
    origin = git_origin()
    return execute(f"git remote get-url {origin}")


def git_server():
    r = re.compile(r"(.+://.+?):.+")
    repo_url = git_repo_url()
    m = re.match(r, repo_url)
    if m is None:
        perror("error: unable to parse repo url")
        raise SystemExit

    return m.group(1).strip()


def git_upstream():
    s = execute("git rev-parse --abbrev-ref @{u}")
    r = re.compile(r"(.+?)/(.+)")
    m = re.match(r, s)
    if m is None:
        perror("error: unable to parse upstream")
        raise SystemExit

    origin, upstream = m.group(1), m.group(2)
    return upstream


def git_head():
    return execute("git rev-parse HEAD")


def git_branch():
    return execute("git rev-parse --abbrev-ref HEAD")


def query(commit):
    server = git_server()
    res = execute(
        f"ssh {server} gerrit query {commit} --current-patch-set --format JSON"
    )
    return res.splitlines()[0]


@click.group()
def cli():
    pass


@cli.command()
def push():
    """
    Amend the HEAD commit and push.
    """
    origin = git_origin()
    upstream = git_upstream()
    execute("git commit --amend --no-edit -s")
    execute(f"git push {origin} HEAD:refs/for/{upstream}")


@cli.command()
@click.argument("commits", nargs=-1)
def download(commits: tuple[str, ...]):
    """
    Cherrypick the specified patch from remote to current branch.
    """
    repo_url = git_repo_url()
    for commit in commits:
        resp = json.loads(query(commit))
        ref = resp["currentPatchSet"]["ref"]
        execute(f"git fetch {repo_url} {ref}")
        execute("git cherry-pick FETCH_HEAD")


@cli.command()
@click.argument("branch")
def branch(branch: str):
    """
    Checkout a new branch based on and tracking upstream.
    """
    origin = git_origin()
    upstream = git_upstream()
    execute(f"git checkout {upstream}")
    execute(f"git checkout -b {branch}")
    execute(f"git branch -u {origin}/{upstream}")


@cli.command()
@click.argument("branches", nargs=-1)
def pick(branches: tuple[str, ...]):
    """
    Cherrypick HEAD commit to the specified branches on remote server.
    """
    head = git_head()
    origin = git_origin()
    cur_branch = git_branch()
    for branch in branches:
        tmp_branch = f"pick/{head}"
        execute(f"git checkout {branch}")
        execute(f"git checkout -b {tmp_branch}")
        execute(f"git cherry-pick {head}")
        execute(f"git push {origin} HEAD:refs/for/{branch}")
        execute(f"git checkout {cur_branch}")
        execute(f"git branch -D {tmp_branch}")
