// Code generated by stringdata. DO NOT EDIT.

package schema

// ActionSchemaJSON is the content of the file "actions.schema.json".
const ActionSchemaJSON = `{
  "$schema": "http://json-schema.org/draft-07/schema#",
  "$id": "actions.schema.json#",
  "title": "Action Definition",
  "description": "Describes an action with its scope and steps to perform.",
  "allowComments": true,
  "type": "object",
  "additionalProperties": false,
  "required": [
    "scopeQuery",
    "steps"
  ],
  "properties": {
    "$schema": {
      "description": "URL of the JSON Schema for an Action Definition.",
      "type": "string",
      "minLength": 1
    },
    "scopeQuery": {
      "description": "A Sourcegraph search query to generate a list of repositories over which to run the action. Use 'src actions scope-query' to see which repositories are matched by the query.",
      "type": "string",
      "minLength": 1
    },
    "steps": {
      "description": "A list of action steps to execute in each repository.",
      "type": "array",
      "minItems": 1,
      "items": {
        "type": "object",
        "required": ["type"],
        "additionalProperties": false,
        "properties": {
          "type": {
            "description": "Can be either \"command\", which executes the step in the native environment (OS) of the machine where 'src actions exec' is executed, or \"docker\" which runs a container with the repository contents mounted in at ` + "`" + `/work` + "`" + `. Note that local images (not from a Docker registry) must be built manually prior to executing.",
            "type": "string",
            "enum": ["command", "docker"]
          },
          "args": {
            "description": "The command and its argument to execute if \"type\" is \"command\", or a list of arguments to be passed to the Docker container if \"type\" is \"docker\".",
            "type": "array",
            "minItems": 1,
            "items": {
              "type": "string"
            }
          },
          "image": {
            "description": "The Docker image handle for running the container executing this step. Just like when running ` + "`" + `docker run` + "`" + `, ` + "`" + `args` + "`" + ` here override the default ` + "`" + `CMD` + "`" + ` to be executed.",
            "type": "string",
            "minLength": 1
          },
          "cacheDirs": {
            "description": "Names of directories to create for in a temporary location and mount into each \"docker\" step container under the specified name.",
            "type": "array",
            "minItems": 1,
            "items": {
              "type": "string"
            }
          }
        },
        "oneOf": [
          {
            "required": ["args"],
            "properties": {
              "type": { "const": "command" },
              "image": { "type": "null" }
            }
          },
          {
            "required": ["image"],
            "properties": {
              "type": { "const": "docker" }
            }
          }
        ]
      }
    }
  }
}
`
