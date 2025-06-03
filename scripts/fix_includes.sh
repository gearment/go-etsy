#!/bin/bash
# fix_includes.sh

input_file="$1"
output_file="$2"

jq '
.paths |= with_entries(
  .value |= with_entries(
    if (.value.parameters? // null) | type == "array" then
      .value.parameters |= map(
        if .name == "includes" and (.schema.type == "array") and (.schema.default == null) then
          .schema.default = []
        else
          .
        end
      )
    else
      .
    end
  )
)
' "$input_file" > "$output_file"

