# jarpatcher
[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![Build][Build-Status-Image]][Build-Status-Url] [![Coverage][Coverage-Image]][Coverage-Url]

Replace OSGi bundles in an installation

patch -s *sourcedir* -t *targetdir*

Scans all bundles in *sourcedir* for `Bundle-SymbolicName` names, replacing any bundles with a matching bundle name in targetdir. This is useful for patching an installation with a new bundle for testing. There's no backup for the replaced jars, so make sure that you can revert your installation.

[License-Url]: http://opensource.org/licenses/MIT
[License-Image]: https://img.shields.io/npm/l/express.svg
[Build-Status-Url]: http://travis-ci.org/aricart/jarpatcher
[Build-Status-Image]: https://travis-ci.org/aricart/jarpatcher.svg?branch=master
[Release-Url]: https://github.com/aricart/jarpatcher/releases/tag/v0.7.2
[Release-image]: http://img.shields.io/badge/release-v0.7.2-1eb0fc.svg
[Coverage-Url]: https://coveralls.io/r/aricart/jarpatcher?branch=master
[Coverage-image]: https://img.shields.io/coveralls/aricart/jarpatcher.svg
[ReportCard-Url]: http://goreportcard.com/report/aricart/jarpatcher
[ReportCard-Image]: http://goreportcard.com/badge/aricart/jarpatcher
[github-release]: https://github.com/aricart/jarpatcher/releases/
