<!--next-version-placeholder-->

## v0.3.5 (2024-11-25)

### Bug Fixes

- Update pyproject.toml for semantic release configuration and dependency versions
  ([`0ec0f04`](https://github.com/entelecheia/entelecheia/commit/0ec0f049404ea1b2b04ca8273077f63573005f28))

- Update release workflow to trigger on push and streamline release steps
  ([`08d5d64`](https://github.com/entelecheia/entelecheia/commit/08d5d644166ab456254516d11620e8b2c592b4bf))

### Chores

- **deps**: Bump codecov/codecov-action from 4 to 5
  ([`3d8a4b8`](https://github.com/entelecheia/entelecheia/commit/3d8a4b819fde0422b9aaf674340eac0a5067a2df))

Bumps [codecov/codecov-action](https://github.com/codecov/codecov-action) from 4 to 5. - [Release
  notes](https://github.com/codecov/codecov-action/releases) -
  [Changelog](https://github.com/codecov/codecov-action/blob/main/CHANGELOG.md) -
  [Commits](https://github.com/codecov/codecov-action/compare/v4...v5)

--- updated-dependencies: - dependency-name: codecov/codecov-action dependency-type:
  direct:production update-type: version-update:semver-major ...

Signed-off-by: dependabot[bot] <support@github.com>

- **deps**: Bump hyfi from 1.36.4 to 1.37.0
  ([`dd4dd60`](https://github.com/entelecheia/entelecheia/commit/dd4dd60ef3ee56023e4752ebbf103d616213c384))

Bumps [hyfi](https://github.com/entelecheia/hyfi) from 1.36.4 to 1.37.0. - [Release
  notes](https://github.com/entelecheia/hyfi/releases) -
  [Changelog](https://github.com/entelecheia/hyfi/blob/main/CHANGELOG.md) -
  [Commits](https://github.com/entelecheia/hyfi/compare/v1.36.4...v1.37.0)

--- updated-dependencies: - dependency-name: hyfi dependency-type: direct:production update-type:
  version-update:semver-minor ...

Signed-off-by: dependabot[bot] <support@github.com>

- **deps**: Bump jupyter-book from 1.0.2 to 1.0.3
  ([`82f7a55`](https://github.com/entelecheia/entelecheia/commit/82f7a55d86111f6c439f4ca359beab782be85234))

Bumps [jupyter-book](https://github.com/executablebooks/jupyter-book) from 1.0.2 to 1.0.3. -
  [Release notes](https://github.com/executablebooks/jupyter-book/releases) -
  [Changelog](https://github.com/jupyter-book/jupyter-book/blob/main/CHANGELOG.md) -
  [Commits](https://github.com/executablebooks/jupyter-book/compare/v1.0.2...v1.0.3)

--- updated-dependencies: - dependency-name: jupyter-book dependency-type: direct:production
  update-type: version-update:semver-patch ...

Signed-off-by: dependabot[bot] <support@github.com>

### Documentation

- Update personal introduction and research focus in README.md
  ([`fb7536c`](https://github.com/entelecheia/entelecheia/commit/fb7536c285aa3ae975a9cb3eb112d7e4511a4f82))


## v0.3.4 (2024-09-14)

### Fix

* **docs:** Update documentation URLs to youngjoon-lee.com ([`f70c640`](https://github.com/entelecheia/entelecheia/commit/f70c6403cdc854d8875b266c0b4aa33e9320c95c))
* **docs:** Update repository URL in README.md ([`ea0e505`](https://github.com/entelecheia/entelecheia/commit/ea0e50504ca9c53ee207b793d6d3077fb9145aa4))

## v0.3.3 (2024-07-20)

### Fix

* **dependencies:** Update multiple dependencies in poetry.lock ([`7099ee4`](https://github.com/entelecheia/entelecheia/commit/7099ee480790bb9238029fa5d4c5e15ec2f9f96e))

## v0.3.2 (2024-04-21)

### Fix

* **docs:** Update GitHub Stats link in README.md ([`1ced68b`](https://github.com/entelecheia/entelecheia/commit/1ced68b2a47c9cc8cb0c3653d114d1e1e2f5fba8))

## v0.3.1 (2024-03-25)

### Fix

* **entelecheia:** Remove unused function open_homepage and update about yaml configuration ([`0f34a67`](https://github.com/entelecheia/entelecheia/commit/0f34a6701c1e60896da953c30474883597c9c840))
* **dependencies:** Update hyfi version to 1.36.3 ([`8bcc3dd`](https://github.com/entelecheia/entelecheia/commit/8bcc3dd782bd619a7fd106dad5c7a29cdf82b8b7))

## v0.3.0 (2024-03-23)

### Feature

* **entelecheia:** Add target field in about page manifest ([`6d53342`](https://github.com/entelecheia/entelecheia/commit/6d533425bfe2e8cae91571881ccc35f04be92e61))
* **init:** Add open_homepage function ([`f02a063`](https://github.com/entelecheia/entelecheia/commit/f02a06358f56cee7503f313b1901c6cd46f0f72e))

### Fix

* **pyproject:** Upgrade Python version to ^3.9, hyfi to ^1.36.1 ([`75b158f`](https://github.com/entelecheia/entelecheia/commit/75b158f0ddfb9d0bcfa9c6c16021ac55bdb89a72))

## v0.2.16 (2023-07-30)
### Fix
* **dependencies:** Upgrade hyfi to 1.12.5 ([`91e13e3`](https://github.com/entelecheia/entelecheia/commit/91e13e3ff6537e61cc4c5985d67603e4cf87f23e))

## v0.2.15 (2023-07-26)
### Fix
* **entelecheia:** Change initialization methods for HyFI and logging ([`79fd6eb`](https://github.com/entelecheia/entelecheia/commit/79fd6eb8df7117a4f9e0abf31a0c2c718c2f6012))

## v0.2.14 (2023-07-23)
### Fix
* **entelecheia:** Change initialization of HyFI object and its logger ([`1c236ec`](https://github.com/entelecheia/entelecheia/commit/1c236eca021b13649554ff9065cc75935e5088b1))
* **entelecheia:** Replace hydra_main with hyfi_main ([`4bfaa38`](https://github.com/entelecheia/entelecheia/commit/4bfaa380e974ffcc03c4f00fcef49f88a594497b))
* **dependencies:** Update hyfi to 1.8.1, limit black version to 23.3.0 ([`1632eb9`](https://github.com/entelecheia/entelecheia/commit/1632eb975f1f31e54ca435a1b8214acedde27711))
* **conf/about:** Rename __init__.yaml to entelecheia.yaml ([`e37a295`](https://github.com/entelecheia/entelecheia/commit/e37a29565ca8b105d82c8a485f716f7e6fc6de66))

## v0.2.13 (2023-07-22)
### Fix
* **entelecheia:** Add new config.yaml file ([`ec7704e`](https://github.com/entelecheia/entelecheia/commit/ec7704e067c3114394a390710840c9b599a55409))
* **dependencies:** Upgrade hyfi to 1.6.1 ([`62b995c`](https://github.com/entelecheia/entelecheia/commit/62b995c784e50daa6112ab40bdeaeea4162b85ce))

## v0.2.12 (2023-07-22)
### Fix
* **dependencies:** Upgrade hyfi to 1.6.0 ([`f43ce55`](https://github.com/entelecheia/entelecheia/commit/f43ce55c04468ce6ae20f1c6eb28988f9bb4f8ca))

## v0.2.11 (2023-07-20)
### Fix
* **HyFI:** Initialize logger, modify get_version docstring ([`6b89139`](https://github.com/entelecheia/entelecheia/commit/6b8913993e5ad96f94912f131b7352b3dad11c00))
* **dependencies:** Upgrade hyfi to 1.3.1 ([`1be70c3`](https://github.com/entelecheia/entelecheia/commit/1be70c39c0f86f3303c8a40bf3204fafd9aca542))

## v0.2.10 (2023-07-18)
### Fix
* **dependencies:** Upgrade hyfi to 1.2.14, remove test files ([`32ce923`](https://github.com/entelecheia/entelecheia/commit/32ce9234b034a35fd65ac03f584c3640ca334071))

## v0.2.9 (2023-07-17)
### Fix
* **dependencies:** Upgrade hyfi to 1.2.13 ([`a4f107b`](https://github.com/entelecheia/entelecheia/commit/a4f107bae69b0e83626ddc885bd0ecbd73080c6c))

## v0.2.8 (2023-07-17)
### Fix
* **dependencies:** Upgrade hyfi to 1.2.11 ([`e3a71d5`](https://github.com/entelecheia/entelecheia/commit/e3a71d532fe8f952db851030282e573e5bb9cfd6))

## v0.2.7 (2023-07-16)
### Fix
* **dependencies:** Upgrade hyfi to 1.2.10 ([`210281e`](https://github.com/entelecheia/entelecheia/commit/210281e7a789b5f6843e8ff888ac77e1857aeed1))

## v0.2.6 (2023-07-04)
### Fix
* **dependencies:** Upgrade hyfi to 1.0.2 ([`e6a95da`](https://github.com/entelecheia/entelecheia/commit/e6a95da6a9826302e8045e30f57cf67ee330482b))

## v0.2.5 (2023-07-04)
### Fix
* **dependencies:** Upgrade hyfi to 1.0.0 ([`688ac62`](https://github.com/entelecheia/entelecheia/commit/688ac6268a09da40eb073a7b53435af490635e18))

## v0.2.4 (2023-06-27)
### Fix
* **dependencies:** Upgrade hyfi to 0.15.0 ([`1bdcaa3`](https://github.com/entelecheia/entelecheia/commit/1bdcaa3d182ea69c8653aa55f5739f8a0524dcf5))

## v0.2.3 (2023-06-21)
### Fix
* **template:** Apply the latest template ([`ff9a73e`](https://github.com/entelecheia/entelecheia/commit/ff9a73e932ce704ad91f2fe8a7e488e5d57092cd))

## v0.2.2 (2023-06-14)
### Fix
* **hyfi-conf:** Update hyfi configs to the latest ([`22ba27c`](https://github.com/entelecheia/entelecheia/commit/22ba27c223a4040d8186ceefffcb0048f1bc3c7e))

## v0.2.1 (2023-06-06)
### Fix
* Update the template ([`44c4de9`](https://github.com/entelecheia/entelecheia/commit/44c4de946ed16210b6b55da1d3450dc2dba2842a))

## v0.2.0 (2023-05-05)
### Feature
* Apply updated template ([`0fc6434`](https://github.com/entelecheia/entelecheia/commit/0fc6434aa5d3ca49ecb888e614070e8abb9f3501))

## v0.1.4 (2023-04-26)
### Fix
* Apply updated template ([`2679928`](https://github.com/entelecheia/entelecheia/commit/2679928d0a8702546fd491b94e16bca3f8fea363))
