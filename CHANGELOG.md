# Changelog

## v0.0.3 - Adding developer console and splitting application.js
- Split all the libraries contained in public/js/application.js into separate files for each library. 
- Add a developer console to have more interaction and options during the run time and a tool for searching errors in the future
- Updating the documentation for the code and the project itself

## v0.0.2 - Code cleanup and optimization
- Adding, updating and extending the documentation in source code and the README
- Completing setup of repo with actions, codeQL, dependabot, automatic releasing, ...
- Clean up dependencies, use originals instead of a fork
- Fix warnings given by dependabot and codeQL 
- Improve logging and the beauty of the output
- Remove usage of packr while serving files.

## v0.0.1 - Running version (again)
After the previous maintainer has fixed the errors so it compiled, five years of inactivity led to new incompatible source code with new versions of Go.

Now all errors are fixed and the source code is partly restructured and cleaned up.

A GitHub action will now create executables for each OS if a new tag is pushed.
