# SM CLI Utilities

  sm-util [command] [options]

## command: mustache

  mustache will read the given template file, and using the given data render
  it to the specified output file location.

### options

  -data="{key1=value1}:~{key2=value2}": key=value data pair, overrides json file data
  -json="{string}": json data string
  -output="{file}": path to output file
  -template="{file}": path to template file

  -data|--data "{key1=value1:~key2=value2}

    Assign a value to a given key for the template rendering

  -json|--json "{string}"

    To be implemented

  -output|--output "{file}"

    specify the output file to write the rendered template to

  -template|--template "{file}"

    specify the template file to use


## command: json

  json reads a uri (file/url) and returns the value found at the given --path

### options

  -path="{path}"        path to retrieve value from
  -path|--path "{path}" path to retrieve value from
  -uri="{uri}"          json uri (url or path to file)
  -uri|--uri "{uri}"    json uri (url or path to file)

