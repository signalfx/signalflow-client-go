# Release Process

1. Add and push a signed tag:

   ```sh
   TAG='v<version>'
   COMMIT='<commit-sha>'
   git tag -s -m $TAG $TAG $COMMIT
   git push upstream $TAG
   ```

1. Create a GitHib Release named `v<version>` with `v<version>` tag.

   The release description should include all the release notes
   from the [`CHANGELOG.md`](CHANGELOG.md) for this release.
