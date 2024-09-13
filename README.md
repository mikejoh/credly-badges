# credly-badges

[![goreleaser](https://github.com/mikejoh/credly-badges/actions/workflows/goreleaser.yml/badge.svg)](https://github.com/mikejoh/credly-badges/actions/workflows/goreleaser.yml)

A GitHub action to display your [Credly](https://info.credly.com/) badges in your personal GitHub profile README!

Inspired by this [repo](https://github.com/pemtajo/badge-readme).

_Note that this GitHub action only supports fetching badges from Credly._

## How-to

1. Add the following in your `github.com/<owner>/<repo>/README.md` add the following:
```
<!--START_BADGES:badges-->
<!--END_BADGES:badges-->
```
2. In your repository create a workflows directory:
```
mkdir -p .github/workflows
```
3. Create a file called e.g. `credly-badges.yaml` with the following contents:
```
name: "Credly Badges"

on:
  schedule:
    - cron: "0 0 * * *"

jobs:
  credly-badges:
    name: "Update README with Credly badges"
    runs-on: ubuntu-latest

    steps:
      - name: Update
        uses: mikejoh/credly-badges@main
        with:
          CREDLY_USERNAME: <Your Credly username>
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
          COMMIT_MESSAGE: "Update Credly badges!"
```
_Note that you might want to replace `@main` and pin to a specific version, see the [Releases](https://github.com/mikejoh/credly-badges/releases) page for available released versions._

If you want to try it out, without waiting for the trigger to be scheduled, you can add another trigger e.g.:
```
...
on:
  push:
...
```
And push a commit to your profile repository, in the `Actions` tab of your repository you shall now see that it has triggered.

## Test locally

1. Build:
```
go build -o credly-badges .
```
2. Create a GitHub token.
3. Run:
```
export INPUT_GITHUB_TOKEN="token"
./credly-badges -credly-username <username> -gh-username <GitHub username> -gh-token $INPUT_GITHUB_TOKEN
```
