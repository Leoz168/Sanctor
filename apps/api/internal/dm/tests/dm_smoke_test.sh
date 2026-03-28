#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
SESSION_FILE="$SCRIPT_DIR/dm_session.env"

if [[ ! -f "$SESSION_FILE" ]]; then
  echo "Missing $SESSION_FILE"
  echo "Run: ./generate_dm_session.sh"
  exit 1
fi

# shellcheck disable=SC1090
source "$SESSION_FILE"

if [[ -z "${ALICE_USER_ID:-}" || -z "${BOB_USER_ID:-}" ]]; then
  echo "Session file is missing user IDs. Re-run ./generate_dm_session.sh"
  exit 1
fi

echo "== Create or get direct DM group =="
GROUP_JSON="$(curl -s -X POST "$API_BASE_URL/api/dm/groups/direct" \
  -H "Content-Type: application/json" \
  -d "{\"userId\":\"$ALICE_USER_ID\",\"peerUserId\":\"$BOB_USER_ID\"}")"
echo "$GROUP_JSON"

GROUP_ID="$(printf '%s' "$GROUP_JSON" | node -e 'const fs=require("fs"); const input=fs.readFileSync(0,"utf8"); const data=JSON.parse(input); if(!data.id){ console.error(input); process.exit(1);} process.stdout.write(data.id)')"

echo
echo "== Alice sends message =="
curl -s -X POST "$API_BASE_URL/api/dm/messages/send" \
  -H "Content-Type: application/json" \
  -d "{\"groupId\":\"$GROUP_ID\",\"userId\":\"$ALICE_USER_ID\",\"content\":\"hello from alice\"}" \
  && echo

echo
echo "== Bob sends reply =="
curl -s -X POST "$API_BASE_URL/api/dm/messages/send" \
  -H "Content-Type: application/json" \
  -d "{\"groupId\":\"$GROUP_ID\",\"userId\":\"$BOB_USER_ID\",\"content\":\"hello from bob\"}" \
  && echo

echo
echo "== Alice group list =="
curl -s "$API_BASE_URL/api/dm/groups?userId=$ALICE_USER_ID" && echo

echo
echo "== Bob group list =="
curl -s "$API_BASE_URL/api/dm/groups?userId=$BOB_USER_ID" && echo

echo
echo "== Read DM history as Alice =="
curl -s "$API_BASE_URL/api/dm/messages?groupId=$GROUP_ID&userId=$ALICE_USER_ID&limit=20" && echo

echo
echo "== WebSocket URL for manual live test =="
printf '%s\n' "ws://localhost:8080/api/dm/ws?groupId=$GROUP_ID&token=$ALICE_TOKEN"
printf '%s\n' "ws://localhost:8080/api/dm/ws?groupId=$GROUP_ID&token=$BOB_TOKEN"
