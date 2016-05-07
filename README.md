# jarpatcher
[![License][License-Image]][License-Url] [![ReportCard][ReportCard-Image]][ReportCard-Url] [![Build][Build-Status-Image]][Build-Status-Url] [![Release][Release-Image]][Release-Url] [![Coverage][Coverage-Image]][Coverage-Url]

Replace OSGi bundles in an installation

patch -s *sourcedir* -t *targetdir*

Scans all bundles in *sourcedir* for `Bundle-SymbolicName` names, replacing any bundles with a matching bundle name in targetdir. This is useful for patching an installation with a new bundle for testing. There's no backup for the replaced jars, so make sure that you can revert your installation.
