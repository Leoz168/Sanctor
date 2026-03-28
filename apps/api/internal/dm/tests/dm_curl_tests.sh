#!/usr/bin/env bash
set -euo pipefail

API_BASE_URL="${API_BASE_URL:-http://localhost:8080}"

ALICE_EMAIL="alice.dm@test.com"
ALICE_USERNAME="alice_dm"
ALICE_PASSWORD="alicepass123"

BOB_EMAIL="bob.dm@test.com"
BOB_USERNAME="bob_dm"
BOB_PASSWORD="bobpass123"

echo "== Health =="
curl -s "$API_BASE_URL/api/health"
echo
echo

echo "== Register Alice =="
curl -s -X POST "$API_BASE_URL/api/auth/register" \
  -H "Content-Type: application/json" \
  -d '{"email":"'$ALICE_EMAIL'","username":"'$ALICE_USERNAME'","password":"'$ALICE_PASSWORD'","firstName":"Alice","lastName":"DM"}' || true
echo
echo

echo "== Register Bob =="
curl -s -X POST "$API_BASE_URL/api/auth/register" \
  -H "Content-Type: application/json" \
  -d '{"email":"'$BOB_EMAIL'","username":"'$BOB_USERNAME'","password":"'$BOB_PASSWORD'","firstName":"Bob","lastName":"DM"}' || true
echo
echo

echo "== Login Alice =="
ALICE_LOGIN="$(curl -s -X POST "$API_BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"'$ALICE_EMAIL'","password":"'$ALICE_PASSWORD'"}')"
printf '%s\n\n' "$ALICE_LOGIN"

echo "== Login Bob =="
BOB_LOGIN="$(curl -s -X POST "$API_BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"email":"'$BOB_EMAIL'","password":"'$BOB_PASSWORD'"}')"
printf '%s\n\n' "$BOB_LOGIN"

ALICE_TOKEN="$(printf '%s' "$ALICE_LOGIN" | node -e 'const fs=require("fs"); const input=fs.readFileSync(0,"utf8"); const data=JSON.parse(input); process.stdout.write(data.token || "")')"
BOB_TOKEN="$(printf '%s' "$BOB_LOGIN" | node -e 'const fs=require("fs"); const input=fs.readFileSync(0,"utf8"); const data=JSON.parse(input); process.stdout.write(data.token || "")')"

ALICE_ID="$(TOKEN="$ALICE_TOKEN" node -e 'const payload=JSON.parse(Buffer.from(process.env.TOKEN.split(".")[1],"base64url").toString()); process.stdout.write(payload.userId)')"
BOB_ID="$(TOKEN="$BOB_TOKEN" node -e 'const payload=JSON.parse(Buffer.from(process.env.TOKEN.split(".")[1],"base64url").toString()); process.stdout.write(payload.userId)')"

echo "ALICE_ID=$ALICE_ID"
echo "BOB_ID=$BOB_ID"
echo

echo "== Create / Get Direct Group =="
GROUP_JSON="$(curl -s -X POST "$API_BASE_URL/api/dm/groups/direct" \
  -H "Content-Type: application/json" \
  -d '{"userId":"'$ALICE_ID'","peerUserId":"'$BOB_ID'"}')"
printf '%s\n\n' "$GROUP_JSON"

GROUP_ID="$(printf '%s' "$GROUP_JSON" | node -e 'const fs=require("fs"); const input=fs.readFileSync(0,"utf8"); const data=JSON.parse(input); process.stdout.write(data.id || "")')"
echo "GROUP_ID=$GROUP_ID"
echo

echo "== Alice Groups =="
curl -s "$API_BASE_URL/api/dm/groups?userId=$ALICE_ID"
echo
echo

echo "== Bob Groups =="
curl -s "$API_BASE_URL/api/dm/groups?userId=$BOB_ID"
echo
echo

echo "== Alice Sends =="
curl -s -X POST "$API_BASE_URL/api/dm/messages/send" \
  -H "Content-Type: application/json" \
  -d '{"groupId":"'$GROUP_ID'","userId":"'$ALICE_ID'","content":"hello from alice via curl"}'
echo
echo

echo "== Bob Sends =="
curl -s -X POST "$API_BASE_URL/api/dm/messages/send" \
  -H "Content-Type: application/json" \
  -d '{"groupId":"'$GROUP_ID'","userId":"'$BOB_ID'","content":"hi alice, bob replying via curl"}'
echo
echo

echo "== Read Messages As Alice =="
curl -s "$API_BASE_URL/api/dm/messages?groupId=$GROUP_ID&userId=$ALICE_ID&limit=20"
echo
echo

echo "== WebSocket URLs =="
echo "ws://localhost:8080/api/dm/ws?groupId=$GROUP_ID&token=$ALICE_TOKEN"
echo "ws://localhost:8080/api/dm/ws?groupId=$GROUP_ID&token=$BOB_TOKEN"