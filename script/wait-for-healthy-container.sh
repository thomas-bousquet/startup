EXPECTED_STATUS="healthy"
while true; do
  STATUS="$(docker inspect --format='{{ .State.Health.Status }}' "${CONTAINER}")"

  if [ "$STATUS" == "$EXPECTED_STATUS" ]; then
    echo "********** ""${CONTAINER}"" container is healthy **********"
  break
  else
  echo "********** Waiting for ""${CONTAINER}"" container to be healthy **********"
  sleep 1
  fi
done