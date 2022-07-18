function isValid()
{
  if [[ $1 =~ ^[0-9]{3}-[0-9]{3}-[0-9]{4}$ ]]; then
    return $(true)
  fi
  if [[ $1 =~ ^[(][0-9]{3}[)][[:space:]][0-9]{3}[-][0-9]{4}$ ]]; then
    return $(true)
  fi
  return $(false)
}

IFS=$'\n'
for i in $(cat file.txt); do
  if isValid $i; then
    echo $i
  fi
done
