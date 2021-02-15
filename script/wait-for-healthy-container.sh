EXPECTED_STATUS="healthy"
while true; do
  STATUS="$(docker inspect --format='{{ .State.Health.Status }}' "${CONTAINER}")"

  if [ "$STATUS" == "$EXPECTED_STATUS" ]; then
    echo "********** ${CONTAINER} is healthy **********"
  break
  else
  echo "********** Waiting for ${CONTAINER} to be healthy **********"
  sleep 1
  fi
done