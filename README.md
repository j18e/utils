# Utils
Utils is a collection of Go functions and types that tend to get reimplemented
in every new project. These include:
- an HTTP server which shuts down when a context is cancelled
- an os.Signal listener which returns an error
- a job runner which runs multiple jobs with a context which is cancelled upon
  receiving of the first error

One of the goals of utils is import minimalism: we strive to achieve all these
functions using the standard library and nothing else, if possible.
