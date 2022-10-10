function main() {
  echo "main"
}


message='name'

if [[ $message == name ]]; then
  echo message
  main
else
  echo "else"
fi

