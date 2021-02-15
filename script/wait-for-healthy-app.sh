EXPECTED_STATUS="200"
while true; do
  STATUS="$(curl http://localhost:8080/admin/health -s -w "%{http_code}")"

  if [ "$STATUS" == "$EXPECTED_STATUS" ]; then
    echo "********** application is healthy **********"
  break
  else
  echo "********** Waiting for application to be healthy **********"
  sleep 1
  fi
done