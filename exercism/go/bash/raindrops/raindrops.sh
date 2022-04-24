#!/usr/bin/env bash
# set -x
main() {
  echo "$1"

  topics=[input,input2,input3]
  for topic_name in $topics; do
    echo "process \"$topic_name\"..."
  done
  if [[ $2 = "aa" ]]; then
    echo good
  else
    echo bad
  fi

}

main "$@"
